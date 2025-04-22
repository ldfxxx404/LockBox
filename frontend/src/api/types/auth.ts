import { BaseSuccessResponse } from './base'
import { User } from './users'

export interface RegisterFormBody {
  email: string
  password: string
  name: string
}

export interface RegisterResponse
  extends BaseSuccessResponse,
    Omit<RegisterFormBody, 'name'> {
  user: User
}

export interface LogoutResponse extends BaseSuccessResponse {}
