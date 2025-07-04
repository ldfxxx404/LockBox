'use client'

import { FormEvent, useState } from 'react'
import { UserRegister } from '../lib/clientRegister'
import { REGISTER_URL } from '../constants/api'
import { useRouter } from 'next/navigation'

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
    } catch (err) {
      console.log('Registration fault URL:', REGISTER_URL, err)
    }
  }

  return (
    <main className='min-h-screen flex items-center justify-center bg-background'>
      <div className='bg-[#343746] px-10 py-8 rounded-xl shadow-[0_2px_16px_0_#0002] min-w-[400px] max-w-[90vw] flex flex-col items-center'>
        <h1 className='text-foreground text-[32px] mb-8'>Registration</h1>
        <form
          className='w-full flex flex-col gap-[18px]'
          onSubmit={handleSubmitRegistration}
        >
          <input
            onChange={ev => setEmail(ev.target.value)}
            type='email'
            placeholder='Email'
            required
            className='bg-[#44475a] text-foreground border-none rounded-lg py-3 px-4 text-base outline-none'
          />
          <input
            onChange={ev => setName(ev.target.value)}
            type='text'
            placeholder='Name'
            required
            className='bg-[#44475a] text-foreground border-none rounded-lg py-3 px-4 text-base outline-none'
          />
          <input
            onChange={ev => setPassword(ev.target.value)}
            type='password'
            placeholder='Password'
            required
            className='bg-[#44475a] text-foreground border-none rounded-lg py-3 px-4 text-base outline-none'
          />
          <button
            type='submit'
            className='bg-[#6272a4] text-white border-none rounded-lg py-[0.7rem] text-lg mt-2 cursor-pointer transition-colors duration-200 hover:bg-[#5761a0]'
          >
            Sign up
          </button>
        </form>
      </div>
    </main>
  )
}
