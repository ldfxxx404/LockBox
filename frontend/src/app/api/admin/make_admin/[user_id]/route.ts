import { NextRequest, NextResponse } from 'next/server'
import { MAKE_ADMIN_URL } from '@/constants/api'
import { ErrorResponse } from '@/types/errorResponse'
import { clogger } from '@/utils/ColorLogger'

export async function PUT(req: NextRequest) {
  try {
    const segments = req.nextUrl.pathname.split('/')
    const user_id_str = segments[segments.length - 1]
    const user_id = Number(user_id_str)
    const authHeader = req.headers.get('authorization')
    const token = authHeader?.split(' ')[1]

    if (!token) {
      const error: ErrorResponse = {
        message: 'Authentication required',
        code: 401,
      }
      clogger.error(
        'Missing or invalid authorization token. Please log in and try again.'
      )
      return NextResponse.json(error, { status: error.code })
    }

    const res = await fetch(`${MAKE_ADMIN_URL}/${user_id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        authorization: `Bearer ${token}`,
      },
    })
    const data = await res.json()

    if (!res.ok) {
      clogger.error(
        `API error: ${res.status} - ${data?.message || res.statusText}`
      )
      return NextResponse.json(data, { status: res.status })
    } else {
      clogger.info(
        'The user has been successfully granted administrator rights.'
      )
    }
    return NextResponse.json(data)
  } catch (err) {
    console.error(err)
    return NextResponse.json(
      { error: 'Internal server error' },
      { status: 500 }
    )
  }
}
