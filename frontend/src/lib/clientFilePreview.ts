'use client'

export async function FilePreview(filename: string) {
  try {
    const response = await fetch(
      `/api/download?filename=${encodeURIComponent(filename)}`,
      {
        method: 'GET',
        headers: {
          Authorization: `Bearer ${sessionStorage.getItem('token')}`,
        },
      }
    )

    if (!response.ok) {
      return { error: 'Preview error' }
    }

    const blob = await response.blob()
    const url = window.URL.createObjectURL(blob)

    window.open(url, '_blank')

    setTimeout(() => window.URL.revokeObjectURL(url), 5000)

    return { success: true }
  } catch (err) {
    console.error('Error: ', err)
    return { error: 'Network error' }
  }
}
