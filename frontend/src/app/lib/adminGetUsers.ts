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

    if (!res.ok) {
      const errorData = await res.json()
      throw new Error(errorData.error || 'Failed to fetch users')
    }

    const data = await res.json()
    return data
  } catch (error) {
    console.error('Error in adminGetUsers:', error)
    throw error
  }
}
