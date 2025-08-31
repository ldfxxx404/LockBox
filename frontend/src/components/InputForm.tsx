import { ChangeEvent, InputHTMLAttributes } from 'react'

interface InputFormProps extends InputHTMLAttributes<HTMLInputElement> {
  value?: string
  type?: 'email' | 'password' | 'name' | 'text'
  placeholder?: 'Email' | 'Password' | 'Name' | 'Search'
  required?: boolean
  onChange?: (ev: ChangeEvent<HTMLInputElement>) => void
}

export const UserInput = ({
  value,
  type,
  placeholder,
  required,
  onChange,
  ...props
}: InputFormProps) => {
  return (
    <input
      onChange={onChange}
      value={value}
      type={type}
      required={required}
      placeholder={placeholder}
      className='bg-[var(--dracula-selection-light)] text-foreground border-none rounded-lg py-3 px-4 text-base outline-none'
      {...props}
    />
  )
}
