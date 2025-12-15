import { useState } from 'react'

export const useSortFiles = () => {
  const [sortOrder, setSortOrder] = useState<'asc' | 'desc'>('asc')
  const [files, setFiles] = useState<string[]>([])

  const sortFiles = () => {
    const newOrder = sortOrder === 'asc' ? 'desc' : 'asc'
    setSortOrder(newOrder)

    setFiles(prevFiles => {
      const sorted = [...prevFiles].sort((a, b) => {
        return newOrder === 'asc' ? a.localeCompare(b) : b.localeCompare(a)
      })
      return sorted
    })
  }
  return {
    setFiles,
    files,
    sortFiles,
    sortOrder,
  }
}
