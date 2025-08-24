import { NextResponse } from 'next/server'
import { ErrorResponse } from '@/types/errorResponse'
import { clogger } from '@/utils/ColorLogger'
import { GetData } from '@/utils/GetUserData'

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

    const user = await GetData(req)
    if (user && typeof user === 'object' && 'email' in user) {
      clogger.info(`Validation successful — user "${user.email}" logged out`)
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
