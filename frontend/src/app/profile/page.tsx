'use client'

import { useState } from 'react'
import { useLogout } from '@/app/hooks/useLogout'
import { FileUploader } from '@/app/lib/clientUpload'
import Forbidden from '@/app/preloader/page'
import { useRedirect } from '@/app/hooks/useRedirect'

export default function UserProfile() {
  const [selectedFile, setSelectedFile] = useState<File | null>(null)
  const handleLogout = useLogout()
  const hasToken = useRedirect()

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
      if (!sessionStorage.getItem('token')) {
        alert('File upload error')
      } else {
        await FileUploader(selectedFile)
        alert('File successfully uploaded')
      }
    } catch (error) {
      console.error('File upload error:', error)
      alert('Ошибка загрузки файла')
    }
  }

  if (hasToken) {
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
  return <Forbidden />
}
