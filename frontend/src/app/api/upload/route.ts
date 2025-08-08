import { UPLOAD_URL } from '@/constants/api'
import { NextResponse } from 'next/server'
import { ErrorResponse } from '@/types/errorResponse'
import { clogger } from '@/utils/ColorLogger'

export async function POST(req: Request) {
  try {
    const formData = await req.formData()
    const file = formData.get('file')
    const authHeader = req.headers.get('authorization')
    const token = authHeader?.split(' ')[1]

    if (!token || !file) {
      const error: ErrorResponse = {
        message: !token ? 'Unauthorized' : 'File not uploaded',
        code: !token ? 403 : 400,
      }
      clogger.error(
        !token
          ? 'Missing or invalid authorization token. Please log in and try again.'
          : 'File not provided in the request.'
      )
      return NextResponse.json(error, { status: error.code })
    }

    const uploadForm = new FormData()
    uploadForm.append('file', file)

    const res = await fetch(UPLOAD_URL, {
      method: 'POST',
      body: uploadForm,
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })

    if (!res.ok) {
      const errorText = await res.text()
      clogger.error(`Upload failed: ${errorText}`)
      const error: ErrorResponse = {
        message: 'Upload file error',
        detail: errorText || 'Unknown error',
        code: res.status,
      }
      return NextResponse.json(error, { status: error.code })
    }

    clogger.info('File uploaded successfully')
    const data = await res.json()
    return NextResponse.json(data)
  } catch (err) {
    const error = err as Error
    clogger.error(`Internal server error during file upload: ${error.message}`)
    const response: ErrorResponse = {
      message: 'Internal server error',
      detail: error.message,
      code: 500,
    }
    return NextResponse.json(response, { status: response.code })
  }
}
