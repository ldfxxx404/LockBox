'use client'

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
      const errorDetails = await res.json()
      console.log('Some error happens', res.status, errorDetails)
      return null
    }

    const data = await res.json()
    return data
  } catch (err) {
    console.log(err)
    return null
  }
}
