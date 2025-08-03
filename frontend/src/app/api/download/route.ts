import { DOWNLOAD_URL } from '@/app/constants/api'
import { NextResponse } from 'next/server'
import { ErrorResponse } from '@/app/types/api'
import { clogger } from '@/utils/ColorLogger'

export async function GET(req: Request) {
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
      return NextResponse.json(error, { status: error.code })
    }
    const res = await fetch(
      DOWNLOAD_URL.replace('{filename}', encodeURIComponent(filename)),
      {
        method: 'GET',
        headers: {
          Authorization: `Bearer ${token}`,
          Accept: 'application/octet-stream',
        },
      }
    )

    if (!res.ok) {
      const error: ErrorResponse = {
        message: 'Download file error',
        detail: 'File not found or token invalid!',
        code: res.status,
      }
      clogger.error(
        'File not found or invalid authorization token. Please log in and try again.'
      )
      const text = await res.text()
      console.error('DOWNLOAD_URL error:', text)
      return NextResponse.json(error, { status: error.code })
    } else {
      clogger.info('File downloaded successfully')
    }

    return new NextResponse(res.body)
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
