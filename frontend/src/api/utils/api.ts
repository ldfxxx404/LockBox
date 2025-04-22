import { Env } from '@/config/env'
import axios, { AxiosRequestConfig } from 'axios'

export const $api = axios.create({
  baseURL: Env.NEXT_PUBLIC_API_URL,
  validateStatus: status => status >= 200 && status < 300,
  withCredentials: true,
})

export type RequestConfig = Partial<AxiosRequestConfig>

const fetchApi = async <T>(url: string, options: RequestConfig): Promise<T> => {
  const response = $api
    .request<T>({
      url,
      ...options,
    })
    .catch(reason => reason)

  return (await response)!.data
}

export const api = {
  get: async <T>(url: string, options?: RequestConfig) =>
    await fetchApi<T>(url, { method: 'GET', ...options }),
  post: async <T>(url: string, options?: Partial<RequestConfig>) =>
    await fetchApi<T>(url, { method: 'POST', ...options }),
  put: async <T>(url: string, options?: Partial<RequestConfig>) =>
    await fetchApi<T>(url, { method: 'PUT', ...options }),
  patch: async <T>(url: string, options?: Partial<RequestConfig>) =>
    await fetchApi<T>(url, { method: 'PATCH', ...options }),
  delete: async <T>(url: string, options?: Partial<RequestConfig>) =>
    await fetchApi<T>(url, { method: 'DELETE', ...options }),
}
