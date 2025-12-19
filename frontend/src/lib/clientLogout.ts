import { getToken } from "@/utils/getToken"

export default async function logout() {
  try {
    const res = await fetch('/api/logout', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${getToken()}`,
      },
    })

    if (res.ok) {
      sessionStorage.clear()
      return { message: 'Logout succesfully' }
    }

    const data = await res.json()

    return data
  } catch (err) {
    console.error('Logout error:', err)
    return null
  }
}
