import { NextResponse } from 'next/server'
import { UPDATE_LIMIT_URL } from '@/app/constants/api'
import { ErrorResponse } from '@/app/types/api'
import { UpdateStoragePayload } from '@/app/types/client'
import { clogger } from '@/utils/ColorLogger'

export async function PUT(req: Request) {
  try {
    const authHeader = req.headers.get('authorization')
    const token = authHeader?.split(' ')[1]

    if (!token) {
      const error: ErrorResponse = {
        code: 401,
        message: 'Authentication required',
      }
      clogger.error(
        'Missing or invalid authorization token. Please log in and try again.'
      )
      return NextResponse.json(error, { status: error.code })
    }

    const body = (await req.json()) as UpdateStoragePayload

    const res = await fetch(UPDATE_LIMIT_URL, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(body),
    })

    const data = await res.json()

    if (!res.ok) {
      clogger.error(
        `API error: ${res.status} - ${data?.message || res.statusText}`
      )
      return NextResponse.json(data, { status: res.status })
    }

    clogger.info('New limit set successfully')
    return NextResponse.json(data)
  } catch (err) {
    console.error(err)
    const error: ErrorResponse = {
      code: 500,
      message: 'Internal server error',
      detail: (err as Error).message,
    }
    return NextResponse.json(error, { status: error.code })
  }
}
