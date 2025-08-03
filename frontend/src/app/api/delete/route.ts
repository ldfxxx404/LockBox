import { DELETE_URL } from '@/app/constants/api'
import { NextResponse } from 'next/server'
import { ErrorResponse } from '@/app/types/api'
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
        code: !token ? 403 : 400,
      }
      clogger.error('Unauthorized! Token not found')
      return NextResponse.json(error, { status: error.code })
    } else {
      clogger.info('File successfully deleted')
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
      const error: ErrorResponse = {
        message: 'Delete file error',
        detail: 'File not found or token invalid!',
        code: res.status,
      }
      const text = await res.text()
      console.error('DELETE_URL error:', text)
      return NextResponse.json(error, { status: error.code })
    }

    const data = await res.json()
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
