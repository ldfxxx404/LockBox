import { useState } from 'react'

export const useFilterFiles = (files: string[]) => {
  const [searchTerm, setSearchTerm] = useState('')

  const filteredFiles = files.filter(filename =>
    filename.toLowerCase().includes(searchTerm.toLowerCase())
  )
  return { filteredFiles, setSearchTerm, searchTerm }
}
