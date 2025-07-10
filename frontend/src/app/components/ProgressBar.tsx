'use client'

import { ProgressHTMLAttributes } from 'react'
import { useState, useEffect } from 'react'

interface BarProps extends ProgressHTMLAttributes<HTMLProgressElement> {}

export const Bar = ({}: BarProps) => {
  const [progress, setProgress] = useState(0)

  useEffect(() => {
    const interval = setInterval(() => {
      setProgress(prev => {
        if (prev >= 100) {
          clearInterval(interval)
          return 100
        }
        return prev + 1
      })
    }, 50)

    return () => clearInterval(interval)
  }, [])

  return (
    <div className='mt-10 w-64 h-4 bg-gray-300 rounded overflow-hidden'>
      <div
        className='h-full bg-[#bd93f9] transition-all duration-200 ease-linear'
        style={{ width: `${progress}%` }}
      />
    </div>
  )
}
