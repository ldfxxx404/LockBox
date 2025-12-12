import { ButtonHTMLAttributes } from 'react'
import { ReactNode } from 'react'

interface AuthSwitchButton extends ButtonHTMLAttributes<HTMLButtonElement> {
  onClick?: () => void
  children?: ReactNode
  placeholder?: string
}

export const AuthSwitchButton = ({
  onClick,
  placeholder,
  children,
  ...props
}: AuthSwitchButton) => {
  return (
    <button
      className='mt-[18px] text-[13px] text-[#bcbcbc] text-center cursor-pointer'
      onClick={onClick}
      {...props}
    >
      {children || placeholder}
    </button>
  )
}
