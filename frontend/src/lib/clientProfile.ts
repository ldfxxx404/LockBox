'use client'

import { getToken } from '@/utils/getToken'

export async function getProfile() {
  const response = await fetch('/api/profile', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${getToken()}`,
    },
  })
  const data = await response.json()
  if (!response.ok) {
    throw new Error(data.message || 'Profile fetch error')
  }
  return data
}
