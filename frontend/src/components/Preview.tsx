import { FilePreview } from '@/lib/clientFilePreview'

export const PreviewButton = ({ filename }: { filename: string }) => {
  const handleClick = async () => {
    const result = await FilePreview(filename)
    if (result?.error) {
      alert(
        'Missing or invalid authorization token. Cannot preview file. Please log in again.'
      )
      window.location.href = '/login'
    }
  }

  return (
    <button
      onClick={handleClick}
      className='text-[var(--dracula-green)] hover:text-[var(--dracula-green-hover)] text-md rounded-lg border mt-1.5 whitespace-nowrap ml-1 mr-1 pr-1 pl-1 cursor-pointer'
    >
      Preview
    </button>
  )
}
