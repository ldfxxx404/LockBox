'use client'

import { ErrorResponse } from '@/app/types/api'

interface Payload {
  user_id: number
}

export async function adminMakeAdmin(param: Payload) {
  try {
    const res = await fetch(`/api/admin/make_admin/${param.user_id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        authorization: `Bearer ${sessionStorage.getItem('token')}`,
      },
    })

    if (!res.ok) {
      const error: ErrorResponse = {
        code: res.status,
        message: 'Authtentication required',
        detail:
          'Missing or invalid authorization token. Please log in and try again. ',
      }
      console.trace(error, { status: error.code })
      return null
    }
    const data = await res.json()
    return data
  } catch (err) {
    const error: ErrorResponse = {
      code: 500,
      message: 'Eternal server error',
    }
    console.error(error, { status: error.code })
    return null
  }
}
