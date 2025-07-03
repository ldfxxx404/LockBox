import { LOGIN_URL } from '@/app/constants/api'
import { NextResponse } from 'next/server'

console.log(`${LOGIN_URL}`)

export async function POST(req: Request) {
  try {
    const body = await req.json()

    const apiRes = await fetch(LOGIN_URL, {
      method: 'POST',
      headers: { 'Content-Type': 'Application/json' },
      body: JSON.stringify(body),
    })

    const json = await apiRes.json()
    console.log('Backend response: ', json)
    return NextResponse.json(json, { status: apiRes.status })

  } catch (err: any) {
    console.error(`Error in API route ${LOGIN_URL}`, err)
    return NextResponse.json(
      { message: 'Internal server error', detail: err.message },
      { status: 500 }
    )
  }
}
