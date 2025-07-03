interface loginPayload {
  email: string
  password: string
}

export async function UserLogin(data: loginPayload) {
  const res = await fetch('/api/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data),
  })

  const json = await res.json().catch(() => ({}))

  if (!res.ok) {
    throw new Error(json.message || 'Login failed')
  }
  return json
}
