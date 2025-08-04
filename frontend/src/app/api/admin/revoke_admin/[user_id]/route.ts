import { REVOKE_ADMIN_URL } from '@/app/constants/api'
import { NextResponse, NextRequest } from 'next/server'
import { clogger } from '@/utils/ColorLogger'
import { ErrorResponse } from '@/app/types/api'

export async function PUT(req: NextRequest) {
  try {
    const segments = req.nextUrl.pathname.split('/')
    const user_id_str = segments[segments.length - 1]
    const user_id = Number(user_id_str)
    const authHeader = req.headers.get('authorization')
    const token = authHeader?.split(' ')[1]

    const res = await fetch(`${REVOKE_ADMIN_URL}/${user_id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        authorization: `Bearer ${token}`,
      },
    })
    const data = await res.json()

    if (!res.ok || !token) {
      const error: ErrorResponse = {
        code: 401,
        message: 'Authtentication required',
        detail:
          'Missing or invalid authorization token. Please log in and try again.',
      }
      clogger.error(
        'Missing or invalid authorization token. Please log in and try again.'
      )
      return NextResponse.json(error, { status: error.code })
    } else {
      clogger.info(
        "You've been demoted from the privileged caste, please go fuck yourself"
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
