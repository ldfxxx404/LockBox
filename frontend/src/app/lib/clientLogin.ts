interface loginPayload {
  email: string
  password: string
}

export async function UserLogin(data: loginPayload) {
  try {
    const res = await fetch('/api/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    })

    // Пытаемся распарсить JSON, если не получилось — json останется пустым объектом
    const json = await res.json().catch(() => ({}))

    if (!res.ok) {
      const message =
        typeof json.message === 'string'
          ? json.message
          : `Login failed (${res.status})`
      throw new Error(message || 'Login failed')
    }
    try {
      // Логика для парсинга токена
      const token = json.token
      console.log('token: ', token)
      localStorage.setItem('token', token)
    } catch (error) {
      console.log('Error get token', error)
    }

    return json
  } catch (err) {
    const error = err as Error
    throw new Error(error?.message || 'Network Error')
  }
}
