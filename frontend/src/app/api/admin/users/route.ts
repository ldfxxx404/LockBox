import { GET_USERS_URL } from '@/app/constants/api'
import { NextResponse } from 'next/server'

export async function GET(req: Request) {
  try {
    const authHeader = req.headers.get('authorization')
    const token = authHeader?.split(' ')[1]
    if (!token) {
      return NextResponse.json(
        { error: 'Unauthorized: Token is missing' },
        { status: 401 }
      )
    }

    const res = await fetch(GET_USERS_URL, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        authorization: `Bearer ${token}`,
      },
    })

    const data = await res.json()

    if (!res.ok) {
      console.error('Error fetching users:', data)
      return NextResponse.json(
        { error: 'Failed to fetch users' },
        { status: res.status }
      )
    }

    return NextResponse.json(data)
  } catch (err) {
    console.error('Internal server error:', err)
    return NextResponse.json(
      { error: 'Internal server error' },
      { status: 500 }
    )
  }
}
