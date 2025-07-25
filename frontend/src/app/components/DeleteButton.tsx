'use client'

import { FileDelete } from '@/app/lib/clientDelete'
import { MouseEvent } from 'react'

interface DeleteButtonProps {
  filename: string
  onDelete?: () => void
  className?: string
}

export const DeleteButton = ({ filename, onDelete }: DeleteButtonProps) => {
  const handleClick = async (e: MouseEvent) => {
    e.preventDefault()
    const result = await FileDelete(filename)
    if (result?.error) {
      alert('Delete file error')
    } else {
      onDelete?.()
    }
  }

  return (
    <button
      onClick={handleClick}
      className='text-red-500 hover:text-red-700 rounded-lg border mt-1.5 whitespace-nowrap ml-4 mr-4 pr-2 pl-2 '
    >
      Delete
    </button>
  )
}
