'use client'

import { getToken } from "@/utils/getToken";

export async function FileDelete(
  filename: string
): Promise<{ success: boolean; error?: string }> {
  try {
    const token = getToken()
    if (!token) {
      return { success: false, error: 'Unauthorized' }
    }

    const response = await fetch(
      `/api/delete?filename=${encodeURIComponent(filename)}`,
      {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
      }
    )

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}))
      return {
        success: false,
        error: errorData.message || 'Delete failed',
      }
    }

    return { success: true }
  } catch (err) {
    console.error('Delete error:', err)
    return {
      success: false,
      error: err instanceof Error ? err.message : 'Network error',
    }
  }
}
