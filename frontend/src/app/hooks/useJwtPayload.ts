import { jwtDecode, JwtPayload } from 'jwt-decode'
import { useMemo } from 'react'

interface MyTokenPayload extends JwtPayload {
user_id: string
}
export const useDecodedPayload = (): MyTokenPayload | null => {

  return useMemo(() => {
    const tokenPayload = localStorage.getItem('token',)
    if (!tokenPayload) return null

    try {
      return jwtDecode<MyTokenPayload>(tokenPayload)
    } catch (e) {
      console.error('Cannot decode JWT:', e)
      return null
    }
  }, [])
}
