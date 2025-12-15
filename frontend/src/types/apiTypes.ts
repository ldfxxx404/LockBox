export interface UserPayload {
  id: number
  name: string
  email: string
  is_admin: boolean
  storage_limit: number
}

export interface LoginPayload {
  email: string
  password: string
}

export interface RegistrationPayload {
  email: string
  name: string
  password: string
}

export interface UpdateStoragePayload {
  new_limit: number
  user_id: number
}

export interface AdminActionPayload {
  user_id: number
}

export interface JwtPayload {
  is_admin: boolean
}
