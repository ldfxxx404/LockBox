import { ChangeEvent, FormEvent } from 'react'

interface UploadProps {
  onSubmit?: (e: FormEvent<HTMLFormElement>) => void
  onChange?: (e: ChangeEvent<HTMLInputElement>) => void
}

export const Upload = ({ onSubmit, onChange, ...props }: UploadProps) => {
  return (
    <>
      <form onSubmit={onSubmit} className='flex gap-4'>
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
          className='bg-[#6272a4] hover:bg-[#5861a0] transition text-white font-medium py-2 px-6 rounded-lg'
        >
          <button type='submit'>Upload</button>
        </label>
      </form>
    </>
  )
}
