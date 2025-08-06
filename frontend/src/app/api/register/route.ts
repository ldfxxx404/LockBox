import { REGISTER_URL } from '@/app/constants/api'
import { NextResponse } from 'next/server'
import { ErrorResponse } from '@/app/types/api'
import { clogger } from '@/utils/ColorLogger'

export async function POST(req: Request) {
  try {
    const body = await req.json()

    const res = await fetch(REGISTER_URL, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(body),
    })

    const responseData = await res.json()

    if (!res.ok) {
      const error: ErrorResponse = {
        message: 'Registration failed!',
        code: res.status,
      }
      clogger.warning(
        'This user already exists, please use other details to register.'
      )
      return NextResponse.json(error, { status: error.code })
    }

    clogger.info('User has registered successfully')
    return NextResponse.json(responseData)
  } catch (error) {
    console.error('Registration error', error)

    const err = error as Error
    const serverError: ErrorResponse = {
      message: 'Internal server error',
      detail: err.message,
      code: 500,
    }
    return NextResponse.json(serverError, { status: 500 })
  }
}
