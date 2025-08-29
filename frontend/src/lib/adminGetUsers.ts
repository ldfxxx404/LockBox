'use client'

export async function adminGetUsers() {
  try {
    const res = await fetch('/api/admin/users', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        authorization: `Bearer ${sessionStorage.getItem('token')}`,
      },
    })
    const data = await res.json()

    if (!res.ok) {
      throw new Error(data.error || 'Failed to fetch users')
    }

    return data
  } catch (error) {
    console.error('Error in adminGetUsers:', error)
    throw error
  }
}
