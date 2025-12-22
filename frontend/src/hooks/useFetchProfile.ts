import { getProfile } from '@/lib/clientProfile'
import { useEffect, useState } from 'react'
import { useSortFiles } from '@/hooks/useSortFiles'

export const useFetchProfile = () => {
  const [loading, setLoading] = useState<boolean>(true)
  const [error, setError] = useState<string | null>(null)
  const [limit, setLimit] = useState<number>(0)
  const [used, setUsed] = useState<number>(0)
  const [userName, setUserName] = useState<string>('')
  const { sortOrder, setFiles, sortFiles, files } = useSortFiles()

  useEffect(() => {
    async function fetchProfile() {
      setLoading(true)
      setError(null)
      try {
        const data = await getProfile()
        setFiles(data.files || [])
        setLimit(data.storage?.limit || 0)
        setUsed(data.storage?.used || 0)
        setUserName(data.user?.name || '')
      } catch (error) {
        const err = error as Error
        setError(err?.message || 'Profile Error')
      } finally {
        setLoading(false)
      }
    }
    fetchProfile()
  }, [setFiles])
  return {
    loading,
    error,
    limit,
    used,
    userName,
    files,
    setFiles,
    sortOrder,
    sortFiles,
  }
}
