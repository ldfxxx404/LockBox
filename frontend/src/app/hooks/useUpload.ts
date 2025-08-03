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
        alert('Error uploading file, unauthorized')
      } else {
        await FileUploader(selectedFile)
        setTimeout(() => {
          window.location.reload()
        }, 1000)
      }
    } catch (error) {
      console.error('Error uploading file: ', error)
      alert('Error uploading file, unauthorized')
    }
  }
  return { handleChange, handleSubmit }
}
