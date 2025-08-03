import { RegistrationPayload } from '@/app/types/client'

export async function UserRegister(RegistrationPayload: RegistrationPayload) {
  const res = await fetch('/api/register', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(RegistrationPayload),
  })

  const json = await res.json().catch(() => ({}))

  if (!res.ok) {
    throw new Error(json.message || 'Registration failed')
  } else return json
}
