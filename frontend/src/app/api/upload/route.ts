import { UPLOAD_URL } from '@/app/constants/api'
import { NextResponse } from 'next/server'
import { ErrorResponse } from '@/app/types/api'
import { clogger } from '@/utils/ColorLogger'

export async function POST(req: Request) {
  const formData = await req.formData()
  const file = formData.get('file')
  const authHeader = req.headers.get('authorization')
  const token = authHeader?.split(' ')[1]

  try {
    if (!token || !file) {
      const error: ErrorResponse = {
        message: !token
          ? 'Unauthorized'
          : !file
            ? 'File not upload'
            : 'Invalid request',
        code: !token ? 403 : 400,
      }
      clogger.error(
        'Error uploading file. Missing or invalid authorization token. Please log in and try again.'
      )
      return NextResponse.json(error, { status: error.code })
    } else {
      clogger.info('File uploaded successfully')
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
      const error: ErrorResponse = {
        message: 'Upload file error',
        detail: 'Token not found!',
        code: res.status,
      }
      const text = await res.text()
      console.error('UPLOAD_URL error:', text)
      return NextResponse.json(error, { status: error.code })
    }
    const data = await res.json()

    return NextResponse.json(data)
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
