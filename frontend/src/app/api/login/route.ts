import { LOGIN_URL } from '@/app/constants/api'
import { NextResponse } from 'next/server'
import { ErrorResponse } from '@/app/types/api'

export async function POST(request: Request) {
  try {
    const data = await request.json()
    const res = await fetch(LOGIN_URL, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    })

    if (!res.ok) {
      const errData = await res.json()
      const error: ErrorResponse = {
        message: 'Login failed! Check your credentials',
        detail: errData.message || 'Unknown error',
        code: res.status,
      }
      return NextResponse.json(error, { status: res.status })
    }
    const responseData = await res.json()
    return NextResponse.json(responseData)
  } catch (error) {
    console.error('Login error', error)

    const err = error as Error
    const serverError: ErrorResponse = {
      message: 'Internal server error',
      detail: err.message,
      code: 500,
    }
    return NextResponse.json(serverError, { status: 500 })
  }
}
