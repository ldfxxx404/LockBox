'use client'

import { Button } from '@/components/ActionButton'
import { useRouter } from 'next/navigation'

export default function NotFound() {
  const router = useRouter()
  return (
    <div className='items-center justify-center flex flex-col mt-56'>
      <h1 className='text-9xl text-[#ff5555]'>404</h1>
      <h2 className='text-4xl text-[#ffb86c]'>Page not found</h2>
      <div className='pt-96'>
        <Button
          label='Go back'
          type='submit'
          onClick={() => router.push('/')}
        />
      </div>
    </div>
  )
}
