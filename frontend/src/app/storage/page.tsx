'use client'

import { useEffect, useState } from 'react'
import { storage } from '../lib/clientStorage'

export default function Storage() {
  const [files, setFiles] = useState<string[]>([])
  const [limit, setLimit] = useState<number>(0)
  const [used, setUsed] = useState<number>(0)
  const [loading, setLoading] = useState<boolean>(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    async function fetchStorage() {
      setLoading(true)
      setError(null)
      try {
        const data = await storage()
        setFiles(data.files || [])
        setLimit(data.storage?.limit || 0)
        setUsed(data.storage?.used || 0)
      } catch (err: any) {
        setError(err?.message || 'Upload Error')
      } finally {
        setLoading(false)
      }
    }
    fetchStorage()
  }, [])

  return (
    <div className='max-w-2xl mx-auto mt-10 p-6 bg-[#6272a4] rounded shadow'>
      <h1 className='text-2xl font-bold mb-4'>Your storage</h1>
      {loading && <div>Loading...</div>}
      {error && <div className='text-red-500 mb-4'>{error}</div>}
      {!loading && !error && (
        <>
          <div className='mb-4'>
            <div>
              <span className='font-semibold'>Uasge:</span> {used} / {limit} MiB
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
            {files.length === 0 ? (
              <div className='text-gray-500'>No files</div>
            ) : (
              <ul className='list-disc pl-5'>
                {files.map((file, idx) => (
                  <li key={idx} className='break-all'>
                    {file}
                  </li>
                ))}
              </ul>
            )}
          </div>
        </>
      )}
    </div>
  )
}
