import { jwtDecode } from 'jwt-decode'

export const getToken = () => {
  const token = sessionStorage.getItem('token')

  if (token) {
    try {
      jwtDecode(token)
    } catch (error) {
      console.error('Invalid token:', error)
      return null
    }
  }

  return token
}
