import { useRouter } from 'next/navigation'
import logout from '@/lib/clientLogout'

export const useLogout = () => {
  const router = useRouter()
  const handleLogout = async () => {
  const result = await logout()

    if (result) {
      router.push('/')
    } else {
      console.error('Logout failed')
    }
  }
  return handleLogout
}
