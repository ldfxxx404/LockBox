import { ChangeEvent } from 'react'

interface UploadProps {
  onChange?: (e: ChangeEvent<HTMLInputElement>) => void
}

export const Upload = ({ onChange, ...props }: UploadProps) => {
  return (
    <>
      <input
        type='file'
        multiple
        onChange={onChange}
        id='file_upload'
        className='hidden'
        {...props}
      />
      <label
        htmlFor='file_upload'
        className='bg-[var(--dracula-button-bright)] hover:bg-[var(--dracula-button-bright-hover)] transition text-white font-medium py-2 px-6 rounded-lg cursor-pointer'
      >
        Upload
      </label>
    </>
  )
}
