'use client'

import { useRedirect } from '@/app/hooks/useRedirect'
import Forbidden from '@/app/preloader/page'

export default function admin() {
  const { hasToken, isChecking } = useRedirect()

  if (isChecking || !hasToken) {
    return <Forbidden />
  }

  return (
  <div>
      31234
      <ul>
            <li>
                  a
                  d
                  v
            </li>
      </ul>
  </div>
  
  )
}
