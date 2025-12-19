'use client'

import { getToken } from '@/utils/getToken'
import { useRouter } from 'next/navigation'
import { useEffect, useState } from 'react'

export const useRedirect = () => {
  const router = useRouter()
  const [isChecking, setIsChecking] = useState(true)
  const [hasToken, setHasToken] = useState(false)

  useEffect(() => {
    if (typeof window === 'undefined') return

    const token = getToken()
    setHasToken(!!token)
    setIsChecking(false)

    if (!token) {
      const timer = setTimeout(() => {
        router.push('/404')
      }, 5000)
      return () => clearTimeout(timer)
    }

    const currentPath = window.location.pathname
    if (currentPath !== '/profile' && currentPath !== '/dashboard') {
      router.push('/profile')
    }
  }, [router])

  return { hasToken, isChecking }
}
