import { LOGIN_URL } from '@/app/constants/api'
import { NextResponse } from 'next/server'

export async function getLoginResponseHandler(req: Request) {
  const body = await req.json()

  const res = await fetch(LOGIN_URL, {
    method: 'POST',
    headers: { 'Content-Type': 'Application/json' },
    body: JSON.stringify(body),
  })
  const data = await res.json()

  if (!res.ok) {
    return NextResponse.json({ error: 'Login failed' }, { status: 401 })
  }
  return NextResponse.json(data, { status: 200 })
}
