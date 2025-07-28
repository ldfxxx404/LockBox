'use client'

import { useRedirect } from '@/app/hooks/useRedirect'
import Forbidden from '@/app/preloader/page'
import { adminGetUsers } from '@/app/lib/adminGetUsers'
import { useState, useEffect } from 'react'

export default function Admin() {
  const [users, setUsers] = useState([])
  const [error, setError] = useState<string | null>(null)
  const [isLoading, setIsLoading] = useState(false)
  const { hasToken, isChecking } = useRedirect()

  useEffect(() => {
    const fetchUsers = async () => {
      if (!hasToken) return

      setIsLoading(true)
      setError(null)

      try {
        const usersData = await adminGetUsers()
        setUsers(usersData)
      } catch (err) {
        console.error('Failed to fetch users:', err)
        setError(err instanceof Error ? err.message : 'Неизвестная ошибка')
      } finally {
        setIsLoading(false)
      }
    }

    fetchUsers()
  }, [hasToken])

  if (isChecking) {
    return <Forbidden />
  }

  if (!hasToken) {
    return <Forbidden />
  }

  if (isLoading) {
    return (
      <div className='flex flex-col items-center py-20'>
        <h1 className='text-2xl font-bold mb-6'>Загрузка...</h1>
      </div>
    )
  }

  if (error) {
    return (
      <div className='flex flex-col items-center py-20'>
        <h1 className='text-2xl font-bold mb-6'>Ошибка</h1>
        <p className='text-red-500'>{error}</p>
      </div>
    )
  }

  return (
    <div className='flex flex-col items-center py-20'>
      <h1 className='text-2xl font-bold mb-6'>Admin Panel</h1>
      {users.length === 0 ? (
        <p>Нет пользователей</p>
      ) : (
        <ul className='w-full max-w-md space-y-3'>
          {users.map((user: any) => (
            <li key={user.id} className='border p-4 rounded shadow'>
              <p>
                <strong>Имя:</strong> {user.name}
              </p>
              <p>
                <strong>Email:</strong> {user.email}
              </p>
              <p>
                <strong>Admin:</strong> {user.is_admin ? 'Да' : 'Нет'}
              </p>
              <p>
                <strong>Limit:</strong> {user.storage_limit} MB
              </p>
            </li>
          ))}
        </ul>
      )}
    </div>
  )
}
