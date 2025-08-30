'use client'

import { Button } from '@/components/ActionButton'
import { useRouter } from 'next/navigation'

export default function NotFound() {
  const router = useRouter()
  return (
    <div className='flex flex-col items-center justify-center min-h-[100dvh] w-screen overflow-hidden text-center p-4'>
      <h1 className='text-9xl max-sm:text-6xl text-[#ff5555]'>403</h1>
      <h2 className='text-4xl max-sm:text-2xl max-sm:mb-55 text-[#ffb86c] mb-8'>
        Access denied
      </h2>
      <Button label='Go back' type='submit' onClick={() => router.push('/profile')} />
    </div>
  )
}
