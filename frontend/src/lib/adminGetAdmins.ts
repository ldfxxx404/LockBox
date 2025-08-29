'use client'

export async function adminGetAdmins() {
  try {
    const res = await fetch('/api/admin/admins', {
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
    console.error('Error in adminGetAdmins:', error)
    throw error
  }
}
