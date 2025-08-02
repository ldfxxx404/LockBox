import { REGISTER_URL } from '@/app/constants/api'
import { NextResponse } from 'next/server'
import { ErrorResponse } from '@/app/types/api'

export async function POST(req: Request) {
  try {
    const body = await req.json()

    const res = await fetch(REGISTER_URL, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(body),
    })

    if (!res.ok) {
      const error: ErrorResponse = {
        message: 'Registration failed!',
        detail:
          'This user already exists, please use other details to register.',
        code: 409,
      }
      console.log('dfa')
      return NextResponse.json(error, { status: error.code })
    }
    const responseData = await res.json()
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
