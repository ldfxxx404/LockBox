import { ButtonHTMLAttributes } from 'react'

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  label: string
  type: 'button' | 'submit'
}

export const Button = ({ label, type }: ButtonProps) => {
  return (
    <button
      type={type}
      className='bg-[#6272a4] text-white border-none rounded-lg py-[0.7rem] text-lg mt-2 cursor-pointer transition-colors duration-200 hover:bg-[#5761a0]'
    >
      {label}
    </button>
  )
}
