import { ChangeEvent, FormEvent, useState } from 'react'
import { FileUploader } from '@/lib/clientUpload'
import toast from 'react-hot-toast'

export const useUpload = () => {
  const [selectedFiles, setSelectedFiles] = useState<File[]>([])

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const files = e.target.files
    if (files) {
      setSelectedFiles(Array.from(files))
    }
  }

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    if (selectedFiles.length === 0) {
      toast.error('File not chosen')
      return
    }

    try {
      if (!sessionStorage.getItem('token')) {
        alert('Error uploading file, unauthorized')
      } else {
        await Promise.all(selectedFiles.map(file => FileUploader(file)))
        setTimeout(() => {
          window.location.reload()
        }, 2000)
      }
      return selectedFiles.forEach(file => {
        toast.success(`File ${file.name} uploaded succesfully`)
      })
    } catch (error) {
      console.error('Error uploading file: ', error)
      alert('Error uploading file')
    }
  }
  return { handleChange, handleSubmit }
}
