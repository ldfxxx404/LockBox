import { callApi } from '@/utils/apiClient'

export async function POST(request: Request) {
  const formData = await request.json()
  return callApi('/api/register', 'POST', formData, 200)
}
