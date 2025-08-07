import { ButtonHTMLAttributes } from 'react'

interface SortButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  onClick: () => void
  label: string
}

export const Sort = ({ onClick, label, ...props }: SortButtonProps) => {
  return (
    <button
      onClick={onClick}
      className='bg-[#6272a4] hover:bg-[#5861a0] text-white px-3 py-1 rounded text-sm mr-2'
      {...props}
    >
      {label}
    </button>
  )
}
