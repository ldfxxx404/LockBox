'use client'

import { useLogout } from '@/hooks/useLogout'
import Forbidden from '@/app/preloader/page'
import { useRedirect } from '@/hooks/useRedirect'
import { useUpload } from '@/hooks/useUpload'
import { useEffect, useState } from 'react'
import { getProfile } from '@/lib/clientProfile'
import { FileDownload } from '@/lib/clientDownload'
import { DeleteButton } from '@/components/DeleteButton'
import { Upload } from '@/components/UploadFile'
import { Button } from '@/components/ActionButton'
import { Sort } from '@/components/SortButton'
import { UserInput } from '@/components/InputForm'
import { Toaster } from 'react-hot-toast'
import { jwtDecode } from 'jwt-decode'
import { useRouter } from 'next/navigation'

export default function UserProfile() {
  const handleLogout = useLogout()
  const { hasToken, isChecking } = useRedirect()
  const { handleChange } = useUpload()

  const [files, setFiles] = useState<string[]>([])
  const [limit, setLimit] = useState<number>(0)
  const [used, setUsed] = useState<number>(0)
  const [userName, setUserName] = useState<string>('')
  const [loading, setLoading] = useState<boolean>(true)
  const [error, setError] = useState<string | null>(null)
  const [sortOrder, setSortOrder] = useState<'asc' | 'desc'>('asc')
  const [searchTerm, setSearchTerm] = useState('')
  const [isAdmin, setIsAdmin] = useState(false)
  const router = useRouter()

  const sortFiles = () => {
    const newOrder = sortOrder === 'asc' ? 'desc' : 'asc'
    setSortOrder(newOrder)

    setFiles(prevFiles => {
      const sorted = [...prevFiles].sort((a, b) => {
        return newOrder === 'asc' ? a.localeCompare(b) : b.localeCompare(a)
      })
      return sorted
    })
  }

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

  useEffect(() => {
    const token = sessionStorage.getItem('token')
    if (token) {
      const payload = jwtDecode<jwtPayload>(token)
      setIsAdmin(!!payload.is_admin)
    }
  }, [])

  const handleDownload = async (filename: string) => {
    const result = await FileDownload(filename)
    if (result?.error) {
      alert(
        'Missing or invalid authorization token. Cannot download file. Please log in again.'
      )
      window.location.href = '/login'
    }
  }

  const filteredFiles = files.filter(filename =>
    filename.toLowerCase().includes(searchTerm.toLowerCase())
  )
  if (isChecking || !hasToken) {
    return <Forbidden />
  }

  interface jwtPayload {
    is_admin: boolean
  }

  const token = sessionStorage.getItem('token')
  if (token) {
    jwtDecode(token)
  }

  return (
    <div className='min-h-screen bg-background flex flex-col items-center py-10'>
      <div className='bg-[#343746] mt-20 px-8 py-6 rounded-xl shadow-lg w-full max-w-2xl max-sm:max-w-full max-sm:mt-0'>
        {/*TODO: change bg-color from bg-[#343746] to bg-[#2d2f44] same dashboard*/}{' '}
        <h3 className='dracula-green text-lg font-semibold mb-4'>
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
        <div className='overflow-y-auto h-96 scrollbar-hidden max-sm:h-[60vh]'>
          <div
            className='flex justify-between items-center sticky top-0 bg-[#343746] 
    max-sm:flex-col max-sm:items-stretch max-sm:gap-2 max-sm:p-2'
          >
            <h2 className='text-lg font-semibold max-sm:text-lg'>Files:</h2>

            <div className='mb-0.5 max-sm:w-full'>
              <UserInput
                placeholder='Search'
                type='text'
                onChange={e => setSearchTerm(e.target.value)}
                className='mt-0.5 w-full px-4 py-2 rounded-lg bg-[#2d2f44] text-white 
       placeholder-gray-400 focus:outline-none focus:ring-2 
       focus:ring-purple-500 transition'
              />
            </div>
            <Sort
              onClick={sortFiles}
              label={`Sort ${sortOrder === 'asc' ? 'A-Z' : 'Z-A'}`}
            />
          </div>
          {loading && <div>Loading...</div>}
          {error && <div className='text-red-500 mb-4'>{error}</div>}
          {!loading && !error && (
            <>
              {files.length === 0 ? (
                <div className='text-gray-500'>No files</div>
              ) : (
                <ul className='list-disc pl-5'>
                  {filteredFiles.map((file, idx) => (
                    <li
                      key={idx}
                      className='break-all flex justify-between items-baseline'
                    >
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
                      <DeleteButton
                        filename={file}
                        onDelete={() => setFiles(files.filter(f => f !== file))}
                      />
                    </li>
                  ))}
                </ul>
              )}
            </>
          )}
        </div>
        <div className='justify-center flex gap-4 mt-8'>
          <div className='flex gap-4 w-full justify-center'>
            <Upload onChange={handleChange} />
            <Button label='Logout' type='submit' onClick={handleLogout} />
            {isAdmin ? (
              <Button
                label='Dashboard'
                type='button'
                onClick={() => router.push('/dashboard')}
              />
            ) : null}
          </div>
        </div>
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
    </div>
  )
}
