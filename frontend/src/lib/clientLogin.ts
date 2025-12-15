'use client'

import { LoginPayload } from '@/types/apiTypes'
import { ErrorResponse } from '@/types/errorResponse'

export async function UserLogin(loginPayload: LoginPayload) {
  try {
    const res = await fetch('/api/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(loginPayload),
    })
    const data = await res.json()

    if (!res.ok) {
      const error: ErrorResponse = {
        code: 401,
        message: 'Login error!',
        detail: 'Check your credentials',
      }
      return console.error(error, { status: error.code })
    }
    try {
      const token = data.token
      sessionStorage.setItem('token', token)
    } catch (error) {
      console.error('Error get token', error)
    }
    return data
  } catch (err) {
    const error: ErrorResponse = {
      code: 500,
      message: 'Eternal server error',
    }
    return console.error(error, { status: error.code }, err)
  }
}
