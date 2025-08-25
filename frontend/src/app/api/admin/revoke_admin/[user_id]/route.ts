import { REVOKE_ADMIN_URL } from '@/constants/api'
import { NextResponse, NextRequest } from 'next/server'
import { clogger } from '@/utils/ColorLogger'
import { ErrorResponse } from '@/types/errorResponse'
import { GetData } from '@/utils/GetUserData'

export async function PUT(req: NextRequest) {
  try {
    const segments = req.nextUrl.pathname.split('/')
    const user_id_str = segments[segments.length - 1]
    const user_id = Number(user_id_str)

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

    const res = await fetch(`${REVOKE_ADMIN_URL}/${user_id}`, {
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
      const user = await GetData(req)
      if (
        user &&
        typeof user === 'object' &&
        'email' in user &&
        'name' in user
      ) {
        clogger.info(
          `Admin rights successfully revoked by user: "${user.email}" from the user with UID: ${user_id}.`
        )
      }
    }

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
