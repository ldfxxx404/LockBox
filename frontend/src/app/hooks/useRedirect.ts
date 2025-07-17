'use client'

import { useRouter } from 'next/navigation'
import { useEffect } from 'react'

export const useRedirect = () => {
  const router = useRouter()
  const token: Boolean = false

  useEffect(() => {
    if (typeof window !== 'undefined' && !sessionStorage.getItem('token')) {
      !!token
      const timer = setTimeout(() => {
        router.push('/404')
      }, 5000)
      return () => clearTimeout(timer)
    } else {
      !token
      if (window.location.pathname !== '/profile') {
        router.push('/profile')
      }
    }
  }, [router])
  return token
}
