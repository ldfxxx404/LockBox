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
      className='bg-[var(--dracula-button-bright)] text-[var(--dracula-white)] rounded-md px-6 py-2 text-base font-medium hover:bg-[var(--dracula-button-bright-hover)] transition cursor-pointer'
      onClick={onClick}
    >
      {label}
    </button>
  )
}
