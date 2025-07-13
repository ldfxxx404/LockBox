interface LogoutResponse {
  message: string
}

export default async function logout(): Promise<LogoutResponse | null> {
  try {
    const token = sessionStorage.getItem('token')

    const res = await fetch('/api/logout', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
    })

      if (res.ok) {
      return { message: 'Logout succesfully' }
    }

    sessionStorage.removeItem('token')

    const data = (await res.json()) as LogoutResponse
    return data
  } catch (err) {
    console.error('Logout error:', err)
    return null
  }
}
