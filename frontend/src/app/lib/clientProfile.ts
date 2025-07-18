'use client'

export async function getProfile() {
  const response = await fetch('/api/profile', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${sessionStorage.getItem('token')}`,
    },
  })
  const data = await response.json()
  if (!response.ok) {
    throw new Error(data.message || 'Profile fetch error')
  }
  return data
}