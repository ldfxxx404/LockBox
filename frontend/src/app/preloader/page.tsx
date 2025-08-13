'use client'

import { Bar } from '@/components/ProgressBar'

export default function Preloader() {
  return (
    <div className='flex flex-col items-center justify-center mt-72'>
      <h1 role='status' className='text-2xl font-bold max-sm:ml-6'>
        Checking in progress, please wait...
      </h1>
      <Bar />
    </div>
  )
}
