import { API_ENDPOINTS } from '@/api/enums/endpoints'
import { RegisterFormBody, RegisterResponse } from '@/api/types/auth'
import { api } from '@/api/utils/api'
import { useMutation, UseMutationOptions } from '@tanstack/react-query'

export const postLoginOptions = (
  body: RegisterFormBody,
  options?: UseMutationOptions<RegisterResponse>
): UseMutationOptions<RegisterResponse> => ({
  mutationFn: async () => await api.post(API_ENDPOINTS.Login(), { data: body }),
  ...options,
})

export const usePostLogin = (
  body: RegisterFormBody,
  options?: UseMutationOptions<RegisterResponse>
) => {
  return useMutation(postLoginOptions(body, options))
}
