import { loginPayload } from '@/app/types/client'

export async function UserLogin(data: loginPayload) {
  try {
    const res = await fetch('/api/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    })

    const json = await res.json().catch(() => ({}))

    if (!res.ok) {
      const message =
        typeof json.message === 'string'
          ? json.message
          : `Login failed (${res.status})`
      throw new Error(message || 'Login failed')
    }
    try {
      const token = json.token
      sessionStorage.setItem('token', token)
    } catch (error) {
      console.log('Error get token', error)
    }

    return json
  } catch (err) {
    const error = err as Error
    throw new Error(error?.message || 'Network Error')
  }
}
