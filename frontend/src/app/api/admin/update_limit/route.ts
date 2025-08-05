import { NextResponse } from 'next/server'
import { UPDATE_LIMIT_URL } from '@/app/constants/api'
import { ErrorResponse } from '@/app/types/api'
import { UpdateStoragePayload } from '@/app/types/client'
import { clogger } from '@/utils/ColorLogger'

export async function PUT(req: Request) {
  try {
    const authHeader = req.headers.get('authorization')
    const token = authHeader?.split(' ')[1]
    const body = (await req.json()) as UpdateStoragePayload

    const res = await fetch(UPDATE_LIMIT_URL, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(body),
    })

    if (!res.ok || !token) {
      const error: ErrorResponse = {
        code: 401,
        message: 'Authtentication required',
        detail:
          'Missing or invalid authorization token. Please log in and try again.',
      }
      clogger.error(
        'Missing or invalid authorization token. Please log in and try again.'
      )
      return NextResponse.json(error, { status: error.code })
    } else {
      clogger.info('New limit set successfully')
    }
    const data = await res.json()
    return NextResponse.json(data)
  } catch (err) {
    console.error(err)
    return NextResponse.json(
      { error: 'Internal server error' },
      { status: 500 }
    )
  }
}
