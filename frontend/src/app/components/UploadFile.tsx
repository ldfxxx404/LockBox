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
          onChange={onChange}
          id='file_upload'
          className='hidden'
          {...props}
        />

        <label
          htmlFor='file_upload'
          className='bg-indigo-500 hover:bg-indigo-600 transition text-white font-semibold py-2 px-6 rounded-lg'
        >
          <button type='submit'>Upload</button>
        </label>
      </form>
    </>
  )
}
