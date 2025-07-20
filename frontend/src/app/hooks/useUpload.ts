import { ChangeEvent, FormEvent, useState } from 'react'
import { FileUploader } from '../lib/clientUpload'

export const useUpload = () => {
  const [selectedFile, setSelectedFile] = useState<File | null>(null)

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0]
    if (file) {
      setSelectedFile(file)
    }
  }

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    if (!selectedFile) {
      alert('File not chosen')
      return
    }

    try {
      if (!sessionStorage.getItem('token')) {
        alert('File upload error')
      } else {
        await FileUploader(selectedFile)
        alert('File successfully uploaded')
      }
    } catch (error) {
      console.error('File upload error:', error)
      alert('File upload error')
    }
  }
  return { handleChange, handleSubmit }
}
