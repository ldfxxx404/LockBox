import { PROFILE_URL } from '@/constants/apiEndpoints'
import { clogger } from './ColorLogger'
import { ErrorResponse } from '@/types/errorResponse'
import { NextResponse } from 'next/server'

export async function GetData(req: Request) {
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
        authorization: `Bearer ${token}`,
      },
    })

    if (!res.ok) {
      clogger.error('Error requesting profile')
      return null
    }

    const data = await res.json()
    const email = <string>data.user.email
    const name = <string>data.user.name
    const id = <number>data.user.id
    return { email, name, id }
  } catch (error) {
    console.error('Profile API error:', error)
    const response: ErrorResponse = {
      message: 'Internal server error',
      code: 500,
    }
    return NextResponse.json(response, { status: 500 })
  }
}
