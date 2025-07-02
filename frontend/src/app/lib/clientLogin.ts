interface loginUserPayload {
  email: string
  password: string
}

export async function UserLogin(data: loginUserPayload) {
  const response = await fetch('/api/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data),
  })
  if (!response.ok) {
    throw new Error('Login failed')
  }
  return response.json()
}
