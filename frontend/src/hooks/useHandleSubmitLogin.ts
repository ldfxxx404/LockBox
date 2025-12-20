import { useRouter } from 'next/navigation'
import { useState } from 'react'
import { FormEvent } from 'react'
import { UserLogin } from '@/lib/clientLogin'

export const useHandleSubmitLogin = () => {
  const router = useRouter()
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')

  const handleSubmitLogin = async (ev: FormEvent) => {
    ev.preventDefault()
    try {
      await UserLogin({ email, password })
      router.push('/profile')
    } catch (err) {
      console.log(err)
    }
  }
  return { handleSubmitLogin, email, setEmail, password, setPassword, router }
}
