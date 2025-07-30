import { ButtonHTMLAttributes } from 'react'

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  label: string
  type: 'button' | 'submit'
  onClick?: () => void
}

export const Button = ({ label, type, onClick }: ButtonProps) => {
  return (
    <button
      type={type}
      className='bg-[#6272a4] text-white rounded-md px-6 py-2 text-base font-medium hover:bg-[#5861a0] transition'
      onClick={onClick}
    >
      {label}
    </button>
  )
}
