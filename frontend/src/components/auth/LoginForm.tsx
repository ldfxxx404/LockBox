import { forwardRef, HTMLAttributes } from 'react'

export const LoginForm = forwardRef<
  HTMLDivElement,
  HTMLAttributes<HTMLDivElement>
>(({}, ref) => {
  return <div ref={ref}></div>
})

LoginForm.displayName = 'LoginForm'
