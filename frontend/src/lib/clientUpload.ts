'use client'

import toast from 'react-hot-toast'
import { getToken } from '@/utils/getToken'

export async function FileUploader(file: File) {
  const uploadForm = new FormData()
  uploadForm.append('file', file)

  try {
    const response = await fetch('/api/upload', {
      method: 'POST',
      body: uploadForm,
      headers: {
        Authorization: `Bearer ${getToken()}`,
      },
    })
    if (!response.ok) {
      toast.error(
        'Upload failed – not enough disk space or unsupported file format.'
      )
      return { error: true }
    }
    toast.success(`${file.name} uploaded successfully`)
    return await response.json()
  } catch (err) {
    return console.error('Error: ', err)
  }
}
