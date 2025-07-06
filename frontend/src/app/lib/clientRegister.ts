interface RegistrationPayload {
  email: string
  name: string
  password: string
}

export async function UserRegister(data: RegistrationPayload) {
  const res = await fetch('/api/register', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  })

  const json = await res.json().catch(() => ({}))

  if (res.ok) {
    localStorage.setItem('email', JSON.stringify(data.email))
    localStorage.setItem('name', JSON.stringify(data.name))
  } else {
    throw new Error(json.message || 'Registration failed')
  }

  return json
}
