import { ButtonHTMLAttributes } from 'react'

interface SortButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  onClick: () => void
  label: string
}

export const Sort = ({ onClick, label, ...props }: SortButtonProps) => {
  return (
    <button
      onClick={onClick}
      className='bg-[var(--dracula-button-bright)] hover:bg-[var(--dracula-button-bright-hover)] text-white px-3 py-1 rounded text-sm mr-2 cursor-pointer'
      {...props}
    >
      {label}
    </button>
  )
}
