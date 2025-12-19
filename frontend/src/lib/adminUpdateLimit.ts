'use client'

import { UpdateStoragePayload } from '@/types/apiTypes'
import { getToken } from '@/utils/getToken'

export async function UpdateLimit(userId: number, newLimit: number) {
  try {
    const payload: UpdateStoragePayload = {
      user_id: userId,
      new_limit: newLimit,
    }

    const res = await fetch('/api/admin/update_limit', {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        authorization: `Bearer ${getToken()}`,
      },
      body: JSON.stringify(payload),
    })
    const data = await res.json()

    if (!res.ok) {
      return { error: "Can't update storage" }
    }
    return data
  } catch (error) {
    console.error('Error: ', error)
    return { error: 'Network error' }
  }
}
