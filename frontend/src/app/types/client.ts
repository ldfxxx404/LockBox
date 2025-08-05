// export interface User {
//   id: string
//   name: string
//   role?: 'admin' | 'user'
// }

export interface loginPayload {
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
