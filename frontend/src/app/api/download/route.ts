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
        code: !token ? 401 : 400,
      }
      clogger.error(
        !token
          ? 'Authorization token is missing.'
          : 'Filename query parameter is required.'
      )
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
      const text = await res.text()
      const error: ErrorResponse = {
        message: 'Download file error',
        detail: text || 'File not found or token invalid!',
        code: res.status,
      }
      clogger.error(`Download failed for "${filename}": ${error.detail}`)
      console.error('DOWNLOAD_URL error:', text)
      return NextResponse.json(error, { status: error.code })
    }

    clogger.info(`File "${filename}" downloaded successfully`)
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
