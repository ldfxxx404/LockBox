'use client'

export async function storage() {
  try {
    const response = await fetch('/api/storage', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${sessionStorage.getItem('token')}`,
      },
    })

    const data = await response.json()

    if (!response.ok) {
      throw {
        message: data.message || 'Unknown error',
        code: data.code || response.status,
      }
    }

    return data
  } catch (error) {
    console.error('Client fetch error:', error)
    throw error
  }
}
