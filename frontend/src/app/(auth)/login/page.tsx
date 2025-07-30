'use client'

import { FormEvent, useState } from 'react'
import { useRouter } from 'next/navigation'
import { LOGIN_URL } from '../../constants/api'
import { UserLogin } from '../../lib/clientLogin'
import { Button } from '../../components/ActionButton'
import { UserInput } from '../../components/InputForm'

export default function LoginPage() {
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

  return (
    <main className='min-h-screen flex items-center justify-center bg-background'>
      <div className='bg-[#343746] px-10 py-8 rounded-xl shadow-[0_2px_16px_0_#0002] min-w-[400px] max-w-[90vw] flex flex-col items-center'>
        <h1 className='text-[32px] mb-8 text-foreground'>Login</h1>
        <form
          className='w-full flex flex-col gap-[18px]'
          onSubmit={handleSubmitLogin}
        >
          <UserInput
            value={email}
            type='email'
            placeholder='Email'
            onChange={ev => setEmail(ev.target.value)}
            required={true}
          />
          <UserInput
            value={password}
            type='password'
            placeholder='Password'
            onChange={ev => setPassword(ev.target.value)}
            required={true}
          />

          <Button label='Sign In' type='submit' />
        </form>
        <div className='mt-[18px] text-[13px] text-[#bcbcbc] text-center'></div>
      </div>
    </main>
  )
}
