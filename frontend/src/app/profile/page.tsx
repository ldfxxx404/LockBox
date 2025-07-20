'use client'

import { useLogout } from '@/app/hooks/useLogout'
import Forbidden from '@/app/preloader/page'
import { useRedirect } from '@/app/hooks/useRedirect'
import { useUpload } from '../hooks/useUpload'
import { useEffect, useState } from 'react'
import { getProfile } from '../lib/clientProfile'
import { FileDownload } from '../lib/clientDownload'

export default function UserProfile() {
  const handleLogout = useLogout()
  const { hasToken, isChecking } = useRedirect()
  const { handleChange, handleSubmit } = useUpload()

  const [files, setFiles] = useState<string[]>([])
  const [limit, setLimit] = useState<number>(0)
  const [used, setUsed] = useState<number>(0)
  const [userName, setUserName] = useState<string>('')
  const [loading, setLoading] = useState<boolean>(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    async function fetchProfile() {
      setLoading(true)
      setError(null)
      try {
        const data = await getProfile()
        setFiles(data.files || [])
        setLimit(data.storage?.limit || 0)
        setUsed(data.storage?.used || 0)
        setUserName(data.user?.name || '')
      } catch (error) {
        const err = error as Error
        setError(err?.message || 'Profile Error')
      } finally {
        setLoading(false)
      }
    }
    fetchProfile()
  }, [])

  // Можно добавить обработчик для скачивания
  const handleDownload = async (filename: string) => {
    const result = await FileDownload(filename)
    if (result?.error) {
      alert('Ошибка скачивания файла')
    }
  }

  if (isChecking || !hasToken) {
    return <Forbidden />
  }

  return (
    <div className='min-h-screen bg-[#232536] flex flex-col items-center py-10'>
      <div className='bg-[#2d2f44] mt-8 px-8 py-6 rounded-xl shadow-lg w-full max-w-2xl'>
        <h3 className='text-indigo-400 text-lg font-semibold mb-4'>
          User storage information
        </h3>
        <div className='mb-2 text-white'>
          <span className='font-semibold'>User:</span> {userName}
        </div>

        <div className='mb-4'>
          <div>
            <span className='font-semibold'>Usage:</span> {used} / {limit} MiB
          </div>
          <div className='w-full bg-[#282a36] rounded h-3 mt-2'>
            <div
              className='bg-[#bd93f9] h-3 rounded'
              style={{ width: limit ? `${(used / limit) * 100}%` : '0%' }}
            ></div>
          </div>
        </div>

        <div>
          <h2 className='text-lg font-semibold mb-2'>Files:</h2>
          {loading && <div>Loading...</div>}
          {error && <div className='text-red-500 mb-4'>{error}</div>}
          {!loading && !error && (
            <>
              {files.length === 0 ? (
                <div className='text-gray-500'>No files</div>
              ) : (
                <ul className='list-disc pl-5'>
                  {files.map((file, idx) => (
                    <li key={idx} className='break-all'>
                      <a
                        href='#'
                        className='text-indigo-400 hover:underline'
                        onClick={e => {
                          e.preventDefault()
                          handleDownload(file)
                        }}
                      >
                        {file}
                      </a>
                    </li>
                  ))}
                </ul>
              )}
            </>
          )}
        </div>

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
