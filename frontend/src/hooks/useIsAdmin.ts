import { jwtDecode } from 'jwt-decode'
import { JwtPayload } from '@/types/apiTypes'
import { useEffect, useState } from 'react'
import { getToken } from '@/utils/getToken'

export const useIsAdmin = () => {
  const [isAdmin, setIsAdmin] = useState(false)

  useEffect(() => {
    const token = getToken()
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
