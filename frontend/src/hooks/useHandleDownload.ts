import { FileDownload } from '@/lib/clientDownload'

export const useHandleDownload = () => {
  const handleDownload = async (filename: string) => {
    const result = await FileDownload(filename)
    if (result?.error) {
      alert(
        'Missing or invalid authorization token. Cannot download file. Please log in again.'
      )
      window.location.href = '/login'
    }
  }
  return { handleDownload }
}
