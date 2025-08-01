import { NextResponse } from 'next/server'
import { MAKE_ADMIN_URL } from '@/app/constants/api'

export async function PUT(req: Request) {
  try {
    const authHeader = req.headers.get('authorization')
    const token = authHeader?.split(' ')[1]

    if (!token) {
      return NextResponse.json(
        { error: 'Unauthorized: Token is missing' },
        { status: 401 }
      )
    }

    const res = await fetch(MAKE_ADMIN_URL, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        authorization: `Bearer ${token}`,
      },
    })

    const data = await res.json()
    
    if (!res.ok) {
      console.error('Upstream error:', data)
      return NextResponse.json(
        { error: 'Failed to make admin' },
        { status: res.status }
      )
    }
    return NextResponse.json(data)

  } catch (err) {
    console.error(err)
    return NextResponse.json(
      { error: 'Internal server error' },
      { status: 500 }
    )
  }
}
