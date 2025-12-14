import { UserRegister } from '@/lib/clientRegister'
import { useRouter } from 'next/navigation'
import { FormEvent, useState } from 'react'
import toast from 'react-hot-toast'

export const useHandleSubmitRegistration = () => {
  const router = useRouter()
  const [email, setEmail] = useState('')
  const [name, setName] = useState('')
  const [password, setPassword] = useState('')

  const handleSubmitRegistration = async (ev: FormEvent) => {
    ev.preventDefault()
    try {
      await UserRegister({ email, name, password })
      router.push('/login')
    } catch {
      toast.error('User with this email already exists')
    }
  }
  return {
    setEmail,
    setName,
    setPassword,
    handleSubmitRegistration,
  }
}
