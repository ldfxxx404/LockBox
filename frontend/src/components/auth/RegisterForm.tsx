'use client'

import { usePostRegister } from '@/api/mutations/auth/use-post-register'
import { useForm } from '@tanstack/react-form'
import { forwardRef, HTMLAttributes } from 'react'

export const RegisterForm = forwardRef<
  HTMLDivElement,
  HTMLAttributes<HTMLDivElement>
>(({}, ref) => {
  const form = useForm({
    defaultValues: {
      email: '',
      password: '',
      name: '',
    },
  })

  const {} = usePostRegister(form.state.values)

  return <div ref={ref}></div>
})

RegisterForm.displayName = 'LoginForm'
