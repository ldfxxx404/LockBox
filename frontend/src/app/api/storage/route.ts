import { PROFILE_URL } from '@/constants/api'
import { NextResponse } from 'next/server'
import { ErrorResponse } from '@/types/errorResponse'
import { clogger } from '@/utils/ColorLogger'

export async function GET(req: Request) {
  try {
    const authHeader = req.headers.get('authorization')
    const token = authHeader?.split(' ')[1]
    if (!token) {
      clogger.error('Unauthorized request to /profile: missing token')
      return NextResponse.json({ message: 'Unauthorized' }, { status: 401 })
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
      clogger.error(
        `Failed to fetch profile: ${data.message || 'Unknown error'}`
      )
      const error: ErrorResponse = {
        message: data.message || 'Unknown error',
        code: data.status || res.status,
      }
      return NextResponse.json(error, { status: res.status })
    }

    clogger.info('Profile fetched successfully')
    return NextResponse.json(data)
  } catch (error) {
    clogger.error(`Storage API error: ${(error as Error).message}`)
    const response: ErrorResponse = {
      message: 'Internal server error',
      code: 500,
    }
    return NextResponse.json(response, { status: 500 })
  }
}
