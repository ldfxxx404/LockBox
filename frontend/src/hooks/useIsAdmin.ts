import { jwtDecode } from 'jwt-decode'
import { JwtPayload } from '@/types/apiTypes'
import { useEffect, useState } from 'react'

export const useIsAdmin = () => {
  const [isAdmin, setIsAdmin] = useState(false)

  useEffect(() => {
    const token = sessionStorage.getItem('token')
    if (!token) return

    try {
      const payload = jwtDecode<JwtPayload>(token)
      setIsAdmin(!!payload.is_admin)
    } catch {
      setIsAdmin(false)
    }
  }, [])

  return isAdmin
}
