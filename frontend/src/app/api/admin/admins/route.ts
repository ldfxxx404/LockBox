import { GET_ADMINS_URL } from '@/constants/api'
import { clogger } from '@/utils/ColorLogger'
import { NextResponse } from 'next/server'
import { ErrorResponse } from '@/types/errorResponse'

export async function GET(req: Request) {
  try {
    const authHeader = req.headers.get('authorization')
    const token = authHeader?.split(' ')[1]

    if (!token) {
      const error: ErrorResponse = {
        code: 401,
        message: 'Unauthorized',
      }
      clogger.error('Missing or invalid authorization token.')
      return NextResponse.json(error, { status: error.code })
    }

    const res = await fetch(GET_ADMINS_URL, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        authorization: `Bearer ${token}`,
      },
    })

    const data = await res.json()

    if (!res.ok) {
      const message = data?.message || 'Failed to fetch users.'
      clogger.error(`Fetch failed: ${res.status} - ${message}`)
      return NextResponse.json(
        {
          code: res.status,
          message: 'Failed to fetch users',
          detail: message,
        },
        { status: res.status }
      )
    }

    clogger.info('Authorized. User list received.')
    return NextResponse.json(data)
  } catch (err) {
    clogger.error('Internal server error')
    const error: ErrorResponse = {
      code: 500,
      message: 'Internal server error',
      detail: (err as Error).message,
    }
    return NextResponse.json(error, { status: error.code })
  }
}
