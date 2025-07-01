import { REGISTER_URL } from '@/app/constants/api'
import { NextResponse } from 'next/server'

export async function POST(req: Request) {
  try {
    const body = await req.json()
    console.log('API route got body:', body)

    const apiRes = await fetch(REGISTER_URL, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(body),
    })

    const json = await apiRes.json()
    console.log('Backend response:', json)

    return NextResponse.json(json, { status: apiRes.status })
  } catch (err: any) {
    console.error('Error in API route /api/register:', err)
    return NextResponse.json(
      { message: 'Internal server error', detail: err.message },
      { status: 500 }
    )
  }
}
