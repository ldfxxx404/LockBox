'use client'

import { jwtDecode, JwtPayload } from 'jwt-decode'
import { useEffect, useState } from 'react'

interface MyTokenPayload extends JwtPayload {
  user_id: string
}

export const useDecodedPayload = (): MyTokenPayload | null => {
  const [payload, setPayload] = useState<MyTokenPayload | null>(null)

  useEffect(() => {
    const token = localStorage.getItem('token')
    if (!token) return

    try {
      const decoded = jwtDecode<MyTokenPayload>(token)
      setPayload(decoded)
    } catch (e) {
      console.error('Cannot decode JWT:', e)
    }
  }, [])

  return payload
}
