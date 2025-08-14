'use client'

import { FormEvent, useState } from 'react'
import { UserRegister } from '@/lib/clientRegister'
import { useRouter } from 'next/navigation'
import { Button } from '@/components/ActionButton'
import { UserInput } from '@/components/InputForm'
import toast, { Toaster } from 'react-hot-toast'

export default function RegisterPage() {
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

  return (
    <main className='min-h-screen flex items-center justify-center bg-background'>
      <div className='bg-[#343746] px-10 py-8 rounded-xl shadow-[0_2px_16px_0_#0002] min-w-[400px] max-w-[90vw] flex flex-col items-center max-sm:min-w-0'>
        <h1 className='text-foreground text-[32px] mb-8'>Registration</h1>
        <form
          className='w-full flex flex-col gap-[18px]'
          onSubmit={handleSubmitRegistration}
        >
          <UserInput
            onChange={ev => setEmail(ev.target.value)}
            type='email'
            placeholder='Email'
            required={true}
          />

          <UserInput
            onChange={ev => setName(ev.target.value)}
            type='name'
            placeholder='Name'
            required={true}
          />

          <UserInput
            onChange={ev => setPassword(ev.target.value)}
            type='password'
            placeholder='Password'
            required={true}
          />

          <Button label='Sign Up' type='submit' />
        </form>
      </div>
      <Toaster
        position='top-center'
        toastOptions={{
          style: {
            background: '#343746',
            color: '#fff',
          },
        }}
      />
    </main>
  )
}
