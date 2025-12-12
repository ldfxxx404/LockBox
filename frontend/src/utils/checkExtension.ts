import { allowedFileExtensions } from '@/constants/allowedFileExtensions'
import path from 'path'

export const isAllowed = (file: string) =>
  allowedFileExtensions.includes(path.extname(file).toLowerCase())
