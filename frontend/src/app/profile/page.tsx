'use client'

import { useLogout } from '@/hooks/useLogout'
import { useRedirect } from '@/hooks/useRedirect'
import { useUpload } from '@/hooks/useUpload'
import { DeleteButton } from '@/components/DeleteButton'
import { Upload } from '@/components/UploadFile'
import { Button } from '@/components/ActionButton'
import { Sort } from '@/components/SortButton'
import { UserInput } from '@/components/InputForm'
import { Toaster } from 'react-hot-toast'
import { useRouter } from 'next/navigation'
import { PreviewButton } from '@/components/PreviewButton'
import { isAllowed } from '@/utils/checkExtension'
import { useIsAdmin } from '@/hooks/useIsAdmin'
import { StorageUsage } from '@/components/StorageUsage'
import { useFetchProfile } from '@/hooks/useFetchProfile'
import { useHandleDownload } from '@/hooks/useHandleDownload'
import { useFilterFiles } from '@/hooks/useFilterFiles'
import Forbidden from '@/app/preloader/page'

export default function UserProfile() {
  const {
    limit,
    used,
    userName,
    loading,
    error,
    files,
    setFiles,
    sortOrder,
    sortFiles,
  } = useFetchProfile()
  const { setSearchTerm, filteredFiles } = useFilterFiles(files)
  const handleLogout = useLogout()
  const router = useRouter()
  const isAdmin = useIsAdmin()
  const { hasToken, isChecking } = useRedirect()
  const { handleChange } = useUpload()
  const { handleDownload } = useHandleDownload()

  if (isChecking || !hasToken) {
    return <Forbidden />
  }

  return (
    <div className='min-h-screen bg-background flex flex-col items-center py-10'>
      <div className='bg-[var(--dracula-comment)] mt-20 px-8 py-6 rounded-xl shadow-lg w-full max-w-2xl max-sm:max-w-full max-sm:mt-0'>
        <h3 className='dracula-green text-lg font-semibold mb-4'>
          User storage information
        </h3>
        <div className='mb-2 text-[var(--dracula-white)]'>
          <span className='font-semibold'>User:</span> {userName}
        </div>
        <div className='mb-4'>
          <StorageUsage used={used} limit={limit} />
        </div>
        <div className='overflow-y-auto h-96 scrollbar-hidden max-sm:h-[60vh]'>
          <div
            className='flex justify-between items-center sticky top-0 bg-[var(--dracula-comment)] 
    max-sm:flex-col max-sm:items-stretch max-sm:gap-2 max-sm:p-2'
          >
            <h2 className='text-lg font-semibold max-sm:text-lg'>Files:</h2>

            <div className='mb-0.5 max-sm:w-full'>
              <UserInput
                placeholder='Search'
                type='text'
                onChange={e => setSearchTerm(e.target.value)}
                className='mt-0.5 w-full px-4 py-2 rounded-lg bg-[var(--dracula-comment)] text-[var(--dracula-white)] 
       placeholder-[var(--dracula-selection)] focus:outline-none focus:ring-2 
       focus:ring-[var(--dracula-purple-hover)] ring-2 ring-[var(--dracula-bar)] transition'
              />
            </div>
            <Sort
              onClick={sortFiles}
              label={`Sort ${sortOrder === 'asc' ? 'A-Z' : 'Z-A'}`}
            />
          </div>
          {loading && <div>Loading...</div>}
          {error && (
            <div className='text-[var(--dracula-red)] mb-4'>{error}</div>
          )}
          {!loading && !error && (
            <>
              {files.length === 0 ? (
                <div className='text-[var(--dracula-selection)]'>No files</div>
              ) : (
                <ul className='list-disc pl-5'>
                  {filteredFiles.map((file, idx) => (
                    <li
                      key={idx}
                      className='break-all flex justify-between items-baseline'
                    >
                      <a
                        href='#'
                        className='text-[var(--dracula-indigo)] hover:underline'
                        onClick={e => {
                          e.preventDefault()
                          handleDownload(file)
                        }}
                      >
                        {file}
                      </a>
                      <div className='max-sm: flex'>
                        {isAllowed(file) ? (
                          <PreviewButton filename={file} />
                        ) : null}
                        <DeleteButton
                          filename={file}
                          onDelete={() =>
                            setFiles(files.filter(f => f !== file))
                          }
                        />
                      </div>
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
