import { NextResponse } from 'next/server'
import { ErrorResponse } from '@/app/types/api'

export async function POST(req: Request) {
  try {
    const authHeader = req.headers.get('Authorization')
    const token = authHeader?.split(' ')[1]

    if (!token) {
      const error: ErrorResponse = {
        message: 'Unauthorized',
        code: 401,
      }
      return NextResponse.json(error, { status: error.code })
    }

    return NextResponse.json({ success: true, message: 'Logout successfully' })
  } catch (error) {
    const response: ErrorResponse = {
      message: 'Logout failed',
      code: 500,
    }
    console.error('Logout error:', error)
    return NextResponse.json(response, { status: response.code })
  }
}
