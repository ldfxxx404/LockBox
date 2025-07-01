import { REGISTER_URL } from '../constants/api'

interface RegistrationPayload {
  email: string
  name: string
  password: string
}

export async function UsrReg(data: RegistrationPayload) {
  const res = await fetch(REGISTER_URL, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  })
  if (!res.ok) {
    const error = await res.json()
    throw new Error(error.message || 'Registration fault nya register.ts')
  }
  return await res.json()
}
