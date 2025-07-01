interface RegistrationPayload {
  email: string
  name: string
  password: string
}

export async function UsrReg(data: RegistrationPayload) {
  const res = await fetch('/api/register', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  })

  const json = await res.json().catch(() => ({}))

  if (!res.ok) {
    throw new Error(json.message || 'Registration failed')
  }

  return json
}
