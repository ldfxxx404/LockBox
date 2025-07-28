'use client'

import { useRouter } from 'next/navigation'
import { useEffect, useState } from 'react'

export const useRedirect = () => {
  const router = useRouter()
  const [isChecking, setIsChecking] = useState(true)
  const [hasToken, setHasToken] = useState(false)

  useEffect(() => {
    if (typeof window === 'undefined') return
    const token = sessionStorage.getItem('token')
    setHasToken(!!token)
    setIsChecking(false)

    if (!token) {
      const timer = setTimeout(() => {
        router.push('/404')
      }, 5000)
      return () => clearTimeout(timer)
    } else if (window.location.pathname !== '/profile') {
      router.push('/profile')
    }
    if (window.location.pathname === '/admin') {
      router.push('/admin')
    }
  }, [router])

  return { hasToken, isChecking }
}
