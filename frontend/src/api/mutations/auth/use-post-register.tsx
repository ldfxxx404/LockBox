import { API_ENDPOINTS } from '@/api/enums/endpoints'
import { RegisterFormBody, RegisterResponse } from '@/api/types/auth'
import { api } from '@/api/utils/api'
import {
  useMutation,
  UseMutationOptions,
} from '@tanstack/react-query'

export const postRegisterOptions = (
  body: RegisterFormBody,
  options?: UseMutationOptions<RegisterResponse>
): UseMutationOptions<RegisterResponse> => ({
  mutationFn: async () =>
    await api.post(API_ENDPOINTS.Register(), { data: body }),
  ...options,
})

export const usePostRegister = (
  body: RegisterFormBody,
  options?: UseMutationOptions<RegisterResponse>
) => {
  return useMutation(postRegisterOptions(body, options))
}
