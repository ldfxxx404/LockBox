import { LOGIN_URL } from '@/app/constants/api'
import { NextResponse } from 'next/server'
import { ErrorResponse } from '@/app/types/api'
import { clogger } from '../../../utils/ColorLogger'

export async function POST(req: Request) {
  try {
    const data = await req.json()
    const res = await fetch(LOGIN_URL, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    })

    const responseData = await res.json()

    if (!res.ok) {
      const error: ErrorResponse = {
        message: 'Login failed! Check your credentials',
        detail: responseData.message || 'Unknown error',
        code: res.status,
      }
      clogger.error('Login failed! Check your credentials')
      return NextResponse.json(error, { status: res.status })
    }

    clogger.info('Login completed successfully')
    return NextResponse.json(responseData)
  } catch (error) {
    console.error('Login error', error)

    const err = error as Error
    const serverError: ErrorResponse = {
      message: 'Internal server error',
      detail: err.message,
      code: 500,
    }
    return NextResponse.json(serverError, { status: serverError.code })
  }
}
