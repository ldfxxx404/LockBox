import { LOGIN_URL } from '@/app/constants/api'
import { NextResponse } from 'next/server'

export async function getLoginResponseHandler(req: Request) {
  try {
    const body = await req.json()

    const res = await fetch(LOGIN_URL, {
      method: 'POST',
      headers: { 'Content-Type': 'Application/json' },
      body: JSON.stringify(body),
    })

    const json = await res.json()
    console.log('Backend response: ', json)
  } catch (err: any) {
    console.error(`Error in API route ${LOGIN_URL}`, err)
    return NextResponse.json(
      { message: 'Internal server error', detail: err.message },
      { status: 500 }
    )
  }
}
