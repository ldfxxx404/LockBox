import { DELETE_URL } from '@/constants/api'
import { NextResponse } from 'next/server'
import { ErrorResponse } from '@/types/errorResponse'
import { clogger } from '@/utils/ColorLogger'

export async function DELETE(req: Request) {
  const authHeader = req.headers.get('authorization')
  const token = authHeader?.split(' ')[1]
  const { searchParams } = new URL(req.url)
  const filename = searchParams.get('filename')

  try {
    if (!token || !filename) {
      const error: ErrorResponse = {
        message: !token ? 'Unauthorized' : 'Filename required',
        code: !token ? 401 : 400,
      }
      clogger.error(
        'Authorization token is missing or the required filename parameter is not provided.'
      )
      return NextResponse.json(error, { status: error.code })
    }

    const res = await fetch(
      DELETE_URL.replace('{filename}', encodeURIComponent(filename)),
      {
        method: 'DELETE',
        headers: {
          Authorization: `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
      }
    )

    if (!res.ok) {
      const text = await res.text()
      const error: ErrorResponse = {
        message: 'Delete file error',
        detail: text || 'File not found or token invalid!',
        code: res.status,
      }
      clogger.error(`Failed to delete "${filename}": ${error.detail}`)
      return NextResponse.json(error, { status: error.code })
    }

    const data = await res.json()
    clogger.info(`File "${filename}" successfully deleted`)
    return NextResponse.json(data, { status: 200 })
  } catch (err) {
    const error = err as Error
    const response: ErrorResponse = {
      message: 'Internal server error',
      detail: error.message,
      code: 500,
    }
    console.error('Error:', error)
    return NextResponse.json(response, { status: response.code })
  }
}
