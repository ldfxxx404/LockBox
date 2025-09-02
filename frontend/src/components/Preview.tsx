'use client'

import { MouseEvent } from 'react'
import toast from 'react-hot-toast'

interface DeleteButtonProps {
  filename: string
  onDelete?: () => void
  className?: string
}

export const Preview = () => {
  return (
    <button className='text-emerald-500 hover:text-emerald-600 text-md rounded-lg border mt-1.5 whitespace-nowrap ml-1 mr-1 pr-1 pl-1 cursor-pointer'>
      Preview
    </button>
  )
}
