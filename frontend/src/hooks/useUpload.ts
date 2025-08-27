import { ChangeEvent, useState } from 'react'
import { FileUploader } from '@/lib/clientUpload'
import toast from 'react-hot-toast'

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
      alert('No file selected. Please choose a file to upload.')
      return
    }

    try {
      if (!sessionStorage.getItem('token')) {
        toast.error('Error uploading file, unauthorized')
      } else {
        await Promise.all(files.map(file => FileUploader(file)))
        setSelectedFiles([])
        const input = document.getElementById('file_upload') as HTMLInputElement
        if (input) input.value = ''

        setTimeout(() => {
          window.location.reload()
        }, 2000)
      }
    } catch (error) {
      console.error('Error uploading file: ', error)
      toast.error('Error uploading file')
    }
  }

  return { handleChange, handleSubmit }
}
