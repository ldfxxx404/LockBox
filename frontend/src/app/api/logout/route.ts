import { NextResponse } from 'next/server'
import { ErrorResponse } from '@/app/types/api'
import { clogger } from '@/utils/ColorLogger'

export async function POST(req: Request) {
  try {
    const authHeader = req.headers.get('authorization')
    const token = authHeader?.split(' ')[1]

    if (!token) {
      const error: ErrorResponse = {
        message: 'Unauthorized',
        code: 401,
      }
      clogger.error(
        'Missing or invalid authorization token. Please log in and try again.'
      )
      return NextResponse.json(error, { status: error.code })
    }

    clogger.info('Logout successfully')

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
