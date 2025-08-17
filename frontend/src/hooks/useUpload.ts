import { ChangeEvent, useState } from 'react'
import { FileUploader } from '@/lib/clientUpload'

export const useUpload = () => {
  const [selectedFiles, setSelectedFiles] = useState<File[]>([])

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const files = e.target.files
    if (files) {
      const filesArr = Array.from(files)
      setSelectedFiles(filesArr)
      handleSubmit(filesArr)
    }
  }

  const handleSubmit = async (filesToUpload?: File[]) => {
    const files = filesToUpload ?? selectedFiles
    if (files.length === 0) {
      alert('Files not chosen')
      return
    }

    try {
      if (!sessionStorage.getItem('token')) {
        alert('Error uploading file, unauthorized')
      } else {
        await Promise.all(files.map(file => FileUploader(file)))
        setSelectedFiles([])
        const input = document.getElementById('file_upload') as HTMLInputElement
        if (input) input.value = ''

        setTimeout(() => {
          window.location.reload()
        }, 1000)
      }
    } catch (error) {
      console.error('Error uploading file: ', error)
      alert('Error uploading file')
    }
  }

  return { handleChange, handleSubmit }
}
