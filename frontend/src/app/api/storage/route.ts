import { STORAGE_URL } from '@/app/constants/api'
import { NextResponse } from 'next/server'
import { ErrorResponse } from '@/app/types/api'

export async function GET(req: Request) {
  try {
    const authHeader = req.headers.get('authorization')
    const token = authHeader?.split(' ')[1]
    if (!token) {
      return NextResponse.json({ message: 'Unauthorized' }, { status: 401 })
    }

    const res = await fetch(STORAGE_URL, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
    })

    if (!res.ok) {
      const errData = await res.json()

      const error: ErrorResponse = {
        message: errData.message || 'Unknown error',
        code: errData.status || res.status,
      }
      return NextResponse.json(error, { status: res.status })
    }
    const data = await res.json()
    return NextResponse.json(data)
  } catch (error) {
    console.error('Storage API error:', error)
    const response: ErrorResponse = {
      message: 'Internal server error',
      code: 500,
    }
    return NextResponse.json(response, { status: 500 })
  }
}
