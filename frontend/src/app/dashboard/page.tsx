'use client'

import { useRedirect } from '@/app/hooks/useRedirect'
import Forbidden from '@/app/preloader/page'
import { adminGetUsers } from '@/app/lib/adminGetUsers'
import { useState, useEffect } from 'react'
import { UserInput } from '../components/InputForm'
import { adminMakeAdmin } from '../lib/adminMakeAdmin'
import { Button } from '../components/ActionButton'
import { adminRevokeAdmin } from '../lib/adminRevokeAdmin'

interface User {
  id: number
  name: string
  email: string
  is_admin: boolean
  storage_limit: number
}

export default function Admin() {
  const [users, setUsers] = useState<User[]>([])
  const [error, setError] = useState<string | null>(null)
  const [isLoading, setIsLoading] = useState(false)
  const { hasToken, isChecking } = useRedirect()
  const [searchTerm, setSearchTerm] = useState('')

  const makeAdmin = async (userId: number) => {
    //TODO: move userId: number to types.ts
    const res = await adminMakeAdmin({ user_id: userId })
    if (!res) {
      console.log('failed make admin')
    } else {
      console.log('success')
    }
  }
  const revokeAdmin = async (userId: number) => {
    //TODO: move userId: number to types.ts
    const res = await adminRevokeAdmin({ user_id: userId })
    if (!res) {
      console.log('failed make admin')
    } else {
      console.log('success')
    }
  }

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
        setError(err instanceof Error ? err.message : 'Unknown error')
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
        <h1 className='text-2xl font-bold mb-6'>Loading...</h1>
      </div>
    )
  }

  if (error) {
    return (
      <div className='flex flex-col items-center py-20'>
        <h1 className='text-2xl font-bold mb-6'>Error</h1>
        <p className='text-red-500'>{error}</p>
      </div>
    )
  }

  const filteredUsers = users.filter(
    user =>
      user.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      user.email.toLowerCase().includes(searchTerm.toLowerCase())
  )

  return (
    <div className='flex flex-col items-center py-20'>
      <h1 className='text-2xl font-bold mb-6'>Admin Panel</h1>

      <UserInput
        placeholder='Search'
        type='text'
        onChange={e => setSearchTerm(e.target.value)}
      />

      {filteredUsers.length === 0 ? (
        <p>Users not found</p>
      ) : (
        <ul className='bg-[#2d2f44] mt-8 px-8 py-6 rounded-xl shadow-lg w-full max-w-2xl space-y-3 h-[40rem] overflow-auto scrollbar-hidden'>
          {filteredUsers.map((user: User) => (
            <li
              key={user.id}
              className='border p-4 rounded-lg relative transition-colors'
            >
              <div className='absolute right-1.5 mt-2 mr-4'>
                <Button
                  onClick={() =>
                    user.is_admin ? revokeAdmin(user.id) : makeAdmin(user.id)
                  }
                  label={user.is_admin ? 'Revoke admin' : 'Make admin'}
                  type='submit'
                  className={user.is_admin ? '' : 'pr-7'}
                />
              </div>
              <p>
                <strong>User name:</strong> {user.name}
              </p>
              <p>
                <strong>Email:</strong> {user.email}
              </p>
              <p>
                <strong>Admin:</strong> {user.is_admin ? 'Yes' : 'No'}
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
