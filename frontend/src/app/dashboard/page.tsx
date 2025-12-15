'use client'

import { useRedirect } from '@/hooks/useRedirect'
import Forbidden from '@/app/preloader/page'
import { adminGetUsers } from '@/lib/adminGetUsers'
import { useState, useEffect } from 'react'
import { UserInput } from '@/components/InputForm'
import { adminMakeAdmin } from '@/lib/adminMakeAdmin'
import { Switcher } from '@/components/Switcher'
import { adminRevokeAdmin } from '@/lib/adminRevokeAdmin'
import { UpdateLimit } from '@/lib/adminUpdateLimit'
import { UserPayload } from '@/types/apiTypes'
import { adminGetAdmins } from '@/lib/adminGetAdmins'
import toast, { Toaster } from 'react-hot-toast'
import { Button } from '@/components/ActionButton'
import { useRouter } from 'next/navigation'

export default function Admin() {
  const router = useRouter()
  const [users, setUsers] = useState<UserPayload[]>([])
  const [error, setError] = useState<string | null>(null)
  const [searchTerm, setSearchTerm] = useState('')
  const [activeTab, setActiveTab] = useState('Users')
  const [admins, setAdmins] = useState<UserPayload[]>([])
  const { hasToken, isChecking } = useRedirect()

  const [limitInputs, setLimitInputs] = useState<Record<number, number | ''>>(
    {}
  )

  const handleLimitChange = (userId: number, value: number | '') => {
    setLimitInputs(prev => ({ ...prev, [userId]: value }))
  }

  const handleLimitUpdate = async (userId: number) => {
    const limit = limitInputs[userId]

    if (limit === '' || isNaN(Number(limit))) {
      toast.error('Please, insert limit value')
      return
    }

    const res = await UpdateLimit(userId, Number(limit))

    if (res?.error) {
      toast.error('Error updating limit. Incorrect value entered.')
    } else {
      toast.success('Limit successfuly updated')
    }
  }

  const makeAdmin = async (userId: number) => {
    const res = await adminMakeAdmin({ user_id: userId })
    if (!res) {
      toast.error('Failed to grant admin rights')
    } else {
      toast.success('Administrator privileges granted successfully')
    }
  }
  const revokeAdmin = async (userId: number) => {
    const res = await adminRevokeAdmin({ user_id: userId })
    if (!res) {
      toast.error('Failed to revoke admin rights')
    } else {
      toast.success('Admin rights revoked successfully')
    }
  }

  const filteredUsers = (activeTab === 'Users' ? users : admins).filter(
    user =>
      (activeTab === 'Users' ? !user.is_admin : user.is_admin) &&
      (user.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
        user.email.toLowerCase().includes(searchTerm.toLowerCase()))
  )

  useEffect(() => {
    const fetchData = async () => {
      if (!hasToken) return
      setError(null)
      try {
        if (activeTab === 'Users') {
          const usersData = await adminGetUsers()
          setUsers(usersData)
        } else {
          const adminsData = await adminGetAdmins()
          setAdmins(adminsData)
        }
      } catch (err) {
        console.error('Failed to fetch data:', err)
        setError(err instanceof Error ? err.message : 'Unknown error')
      }
    }
    fetchData()
  }, [hasToken, activeTab])

  if (isChecking || !hasToken || error) {
    if (error) {
      router.push('/403')
    }
    return <Forbidden />
  }

  return (
    <div className='flex flex-col items-center px-4 py-8 min-h-screen bg-[var(--background)] text-[var(--dracula-white)]'>
      <h1 className='text-4xl font-extrabold mb-8 tracking-wide'>
        Admin Panel
      </h1>

      <Switcher activeTab={activeTab} setActiveTab={setActiveTab} />
      <div className='mb-5'>
        <Button
          label='Back to profile'
          type='button'
          onClick={() => router.push('/profile')}
        />
      </div>

      <UserInput
        placeholder='Search'
        type='text'
        onChange={e => setSearchTerm(e.target.value)}
        className='mb-6 w-full max-w-md px-4 py-2 rounded-lg bg-[var(--dracula-comment)] text-[var(--dracula-white)] placeholder-[var(--dracula-selection)] focus:outline-none focus:ring-2 focus:ring-[var(--dracula-purple)] transition'
      />

      {filteredUsers.length === 0 ? (
        <p>Users not found</p>
      ) : (
        <ul className='bg-[var(--dracula-comment)] mt-4 px-6 py-6 rounded-2xl shadow-2xl w-full max-w-3xl space-y-4 h-[40rem] overflow-y-auto scrollbar-hidden'>
          {filteredUsers.map((user: UserPayload) => (
            <li
              key={user.id}
              className='bg-[var(--dracula-card)] p-6 rounded-xl shadow-md hover:shadow-lg transition-shadow duration-300 relative max-sm:p-4 max-sm:text-sm'
            >
              <div className='flex flex-col sm:flex-row sm:justify-between sm:items-start gap-4'>
                <div className='flex-1 space-y-1'>
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
                </div>

                <div className='w-full sm:w-60 space-y-2'>
                  <input
                    type='number'
                    value={limitInputs[user.id] ?? ''}
                    onChange={e =>
                      handleLimitChange(user.id, Number(e.target.value))
                    }
                    placeholder='New limit (MB)'
                    className='w-full bg-[var(--dracula-comment)] text-[var(--dracula-white)] border-none rounded-lg py-2 px-3 text-sm placeholder-[var(--dracula-selection)] focus:outline-none focus:ring-2 focus:ring-[var(--dracula-purple)] transition'
                  />

                  <button
                    onClick={() => handleLimitUpdate(user.id)}
                    type='button'
                    className='w-full bg-[var(--dracula-purple)] hover:bg-[var(--dracula-purple-hover)] text-[var(--dracula-white)] text-sm font-medium py-1.5 px-2 rounded-md transition duration-200'
                  >
                    Update Limit
                  </button>

                  <button
                    onClick={() =>
                      user.is_admin ? revokeAdmin(user.id) : makeAdmin(user.id)
                    }
                    type='button'
                    className={`w-full text-[var(--dracula-white)] text-sm font-medium py-1.5 px-2 rounded-md transition duration-200 ${
                      user.is_admin
                        ? 'bg-[var(--dracula-red)] hover:bg-[var(--dracula-red-hover)]'
                        : 'bg-[var(--dracula-green)] hover:bg-[var(--dracula-green-hover)]'
                    }`}
                  >
                    {user.is_admin ? 'Revoke admin' : 'Make admin'}
                  </button>
                </div>
              </div>
            </li>
          ))}
        </ul>
      )}
      <Toaster // // TODO: move this toast to anonther file
        position='top-center'
        toastOptions={{
          style: {
            background: '#343746',
            color: '#fff',
          },
        }}
      />
    </div>
  )
}
