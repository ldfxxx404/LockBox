'use client'
import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import Forbidden from '@/app/forbidden/page'

export default function UserProfile() {
  const [isCheckingAuth, setIsCheckingAuth] = useState(true)
  const router = useRouter()

  const handleLogout = () => {
    sessionStorage.removeItem('token')
    router.push('/')
  }

  useEffect(() => {
    const token = sessionStorage.getItem('token')
    setTimeout(() => {
      if (!token) {
        router.push('/404')
      } else {
        setIsCheckingAuth(false)
      }
    }, 5000)
  }, [router])

  if (isCheckingAuth) {
    return <Forbidden /> // отдельная страница для проверки на авторизацию (preload page)
  }

  return (
    <div className='min-h-screen bg-[#232536] flex flex-col items-center py-10'>
      <div className='bg-[#2d2f44] mt-8 px-8 py-6 rounded-xl shadow-lg w-full max-w-2xl'>
        <h3 className='text-indigo-400 text-lg font-semibold mb-4'>
          User storage information
        </h3>
        <ul className='space-y-3 text-white'>
          <li>
            <span className='text-[#b0b3c6]'>Storage usage:</span> 24 MiB{' '}
          </li>
          <li>
            <span className='text-[#b0b3c6]'>Free space:</span> 1000 MiB{' '}
          </li>
        </ul>
        <div className='justify-center flex gap-5 mt-8 '>
          <button className='bg-indigo-500 hover:bg-indigo-600 transition text-white font-semibold py-2 px-6 rounded-lg'>
            Upload
          </button>
          <button className='bg-indigo-500 hover:bg-indigo-600 transition text-white font-semibold py-2 px-6 rounded-lg'>
            Download
          </button>
          <button
            className='bg-indigo-500 hover:bg-indigo-600 transition text-white font-semibold py-2 px-6 rounded-lg'
            onClick={handleLogout}
          >
            Logout
          </button>
        </div>
      </div>
    </div>
  )
}
