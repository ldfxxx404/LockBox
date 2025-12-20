'use client'

import { Button } from '@/components/ActionButton'
import { UserInput } from '@/components/InputForm'
import { AuthSwitchButton } from '@/components/AuthSwitchButton'
import { useHandleSubmitLogin } from '@/hooks/useHandleSubmitLogin'

export default function LoginPage() {
  const { handleSubmitLogin, email, setEmail, password, setPassword, router } =
    useHandleSubmitLogin()

  return (
    <main className='min-h-screen flex items-center justify-center bg-background'>
      <div className='bg-[var(--dracula-comment)] px-10 py-8 rounded-xl shadow-[0_2px_16px_0_#0002] min-w-[400px] max-w-[90vw] flex flex-col items-center max-sm:min-w-0'>
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
        <AuthSwitchButton
          onClick={() => router.push('/register')}
          placeholder='Don’t have an account? Register'
        />
      </div>
    </main>
  )
}
