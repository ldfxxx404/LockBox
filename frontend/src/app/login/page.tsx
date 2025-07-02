'use client'

import { FormEvent, useState } from "react"
import { useRouter } from "next/navigation";
import { LOGIN_URL } from "../constants/api";

export default function LoginPage() {
  const router = useRouter();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmitLogin = async (ev: FormEvent) => {
    ev.preventDefault();
    
    const response = await fetch(LOGIN_URL, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password })
    })
    if (response.ok) {
      router.push('/')
    } else {
      console.log('error')
    }
  }

  return (
    <main className='min-h-screen flex items-center justify-center bg-background'>
      <div className='bg-[#343746] px-10 py-8 rounded-xl shadow-[0_2px_16px_0_#0002] min-w-[400px] max-w-[90vw] flex flex-col items-center'>
        <h1 className='text-[32px] mb-8 text-foreground'>Login</h1>
        <form className='w-full flex flex-col gap-[18px]' onSubmit={handleSubmitLogin}>
          <input
            value={email}
            onChange={ev => setEmail(ev.target.value)}
            type='email'
            placeholder='Email'
            required
            className='bg-[#44475a] text-foreground border-none rounded-lg py-3 px-4 text-base outline-none'
          />
          <input
            value={password}
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
            Sign in
          </button>
        </form>
        <div className='mt-[18px] text-[13px] text-[#bcbcbc] text-center'></div>
      </div>
    </main>
  )
}
