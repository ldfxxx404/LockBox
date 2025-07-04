'use client'

import { useDecodedPayload } from '../hooks/useJwtPayload'

export default function UserProfile() {
  const user = useDecodedPayload()
  const id = user?.user_id

  return (
    <div className='min-h-screen bg-[#232536] flex flex-col items-center py-10'>
      <div className='flex items-center gap-6 bg-[#2d2f44] px-8 py-6 rounded-xl shadow-lg w-full max-w-2xl border-b border-[#35364a]'>
        <div className='flex-shrink-0'>
          <div className='w-20 h-20 rounded-full bg-[#35364a] border-4 border-indigo-400 flex items-center justify-center text-3xl font-bold text-indigo-300'>
            {id} {/*имплементировать получение имени пользователя*/}
          </div>
        </div>
        <div>
          <h2 className='text-2xl font-semibold text-white mb-1'>A A</h2>
          <p className='text-[#b0b3c6]'>a@a.a</p>{' '}
          {/* Реализовать получение почты текущего пользователя*/}
        </div>
      </div>
      <div className='bg-[#2d2f44] mt-8 px-8 py-6 rounded-xl shadow-lg w-full max-w-2xl'>
        <h3 className='text-indigo-400 text-lg font-semibold mb-4'>
          User storage information
        </h3>
        <ul className='space-y-3 text-white'>
          <li>
            <span className='text-[#b0b3c6]'>Storage usage:</span> 24 MiB{' '}
            {/* Имплементировать получение кол-ва занятой памяти*/}
          </li>
          <li>
            <span className='text-[#b0b3c6]'>Free space:</span> 1000 MiB{' '}
            {/* Имплементировать получение кол-ва свободной памяти*/}
          </li>
        </ul>
        <div className='justify-center flex gap-5 mt-8 '>
          <button className='bg-indigo-500 hover:bg-indigo-600 transition text-white font-semibold py-2 px-6 rounded-lg'>
            Upload
          </button>
          <button className='bg-indigo-500 hover:bg-indigo-600 transition text-white font-semibold py-2 px-6 rounded-lg'>
            Download
          </button>
          <button className='bg-indigo-500 hover:bg-indigo-600 transition text-white font-semibold py-2 px-6 rounded-lg'>
            Logout
          </button>
        </div>
      </div>
    </div>
  )
}
