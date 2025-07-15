'use client'

import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import Forbidden from '@/app/preloader/page'
import logout from '../lib/clientLogout'
import { FileUploader } from '../lib/clientUpload'

export default function UserProfile() {
  const [hasToken, setHasToken] = useState<boolean | null>(null)
  const router = useRouter()
  const [selectedFile, setSelectedFile] = useState<File | null>(null)

  useEffect(() => {
    if (typeof window !== 'undefined') {
      const token = sessionStorage.getItem('token')

      if (!token) {
        setHasToken(false)
        const timer = setTimeout(() => {
          router.push('/404')
        }, 5000)
        return () => clearTimeout(timer)
      } else {
        setHasToken(true)
        if (window.location.pathname !== '/profile') {
          router.push('/profile')
        }
      }
    }
  }, [router])

  const handleLogout = async () => {
    const result = await logout()
    if (result) {
      console.log(result)
      router.push('/')
    } else {
      console.error('Logout failed')
    }
  }

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0]
    if (file) {
      setSelectedFile(file)
    }
  }

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    if (!selectedFile) {
      alert('Файл не выбран')
      return
    }

    try {
      const token = sessionStorage.getItem('token')
      if (!token) {
        alert('File upload error')
      } else {
        await FileUploader(selectedFile)
        alert('Файл загружен успешно')
      }
    } catch (error) {
      console.error('Ошибка при загрузке файла:', error)
      alert('Ошибка загрузки файла')
    }
  }

  if (hasToken === false) {
    return <Forbidden />
  }

  if (hasToken === null) {
    return null
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
          <form onSubmit={handleSubmit}>
            <input
              onChange={handleChange}
              type='file'
              className='bg-indigo-500 hover:bg-indigo-600 transition text-white font-semibold py-2 px-6 rounded-lg'
            />
            <button
              type='submit'
              className='bg-indigo-500 hover:bg-indigo-600 transition text-white font-semibold py-2 px-6 rounded-lg'
            >
              Upload
            </button>
          </form>
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
