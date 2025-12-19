import { jwtDecode } from 'jwt-decode'

export const getToken = () => {
  const token = sessionStorage.getItem('token')

  if (token) {
    jwtDecode(token)
  }

  return token
}
