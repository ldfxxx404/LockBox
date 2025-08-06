import { PROFILE_URL } from '@/app/constants/api'
import { NextResponse } from 'next/server'
import { ErrorResponse } from '@/app/types/api'
import { clogger } from '@/utils/ColorLogger'

export async function GET(req: Request) {
  try {
    const authHeader = req.headers.get('authorization')
    const token = authHeader?.split(' ')[1]
    if (!token) {
      const error: ErrorResponse = {
        message: 'Unauthorized',
        code: 401,
      }
      clogger.error('Authorization token is missing for /profile request.')
      return NextResponse.json(error, { status: error.code })
    }

    const res = await fetch(PROFILE_URL, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
    })

    const data = await res.json()

    if (!res.ok) {
      const error: ErrorResponse = {
        message: data.message || 'Unknown error',
        code: data.status || res.status,
      }
      clogger.error(
        'Cannot get /profile. Missing or invalid authorization token. Please log in and try again.'
      )
      return NextResponse.json(error, { status: res.status })
    }

    clogger.info('Profile fetched successfully.')
    return NextResponse.json(data)
  } catch (error) {
    console.error('Profile API error:', error)
    const response: ErrorResponse = {
      message: 'Internal server error',
      code: 500,
    }
    return NextResponse.json(response, { status: 500 })
  }
}
