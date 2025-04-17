import { SERVER_API } from '@/utils/apiUrl'
import { NextResponse } from 'next/server'

export async function GET(request: Request) {
  try {
    const authorizationHeader = request.headers.get('authorization')
    if (!authorizationHeader || !authorizationHeader.startsWith('Bearer ')) {
      return new NextResponse(
        JSON.stringify({ error: 'No valid authorization token provided' }),
        { status: 401 }
      )
    }

    const token = authorizationHeader.split(' ')[1]

    const res = await fetch(`${SERVER_API}/api/storage`, {
      method: 'GET',
      headers: { Authorization: `Bearer ${token}` },
    })

    const data = await res.json()

    if (!res.ok) {
      return new NextResponse(
        JSON.stringify({ error: data.error || 'Failed to fetch storage' }),
        { status: res.status }
      )
    }

    return NextResponse.json(data, { status: 200 })
  } catch (error: unknown) {
    console.error(error)
    return new NextResponse(
      JSON.stringify({ error: 'Internal server error' }),
      { status: 500 }
    )
  }
}
