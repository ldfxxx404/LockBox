'use client'

import { FileDelete } from '@/lib/clientDelete'
import { MouseEvent } from 'react'
import toast from 'react-hot-toast'

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
      alert(
        'Missing or invalid authorization token, can not delte file. Relogin'
      )
    } else {
      onDelete?.()
      toast.success(`${filename} deleted`)
    }
  }

  return (
    <button
      onClick={handleClick}
      className='text-red-500 hover:text-red-700 rounded-lg border mt-1.5 whitespace-nowrap ml-4 mr-4 pr-2 pl-2 cursor-pointer'
    >
      Delete
    </button>
  )
}
