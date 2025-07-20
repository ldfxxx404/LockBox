'use client'

export async function FileDownload(filename: string) {
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
      return { error: 'Download error' }
    }
    const blob = await response.blob()
    const url = window.URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = filename
    document.body.appendChild(a)
    a.click()
    a.remove()
    window.URL.revokeObjectURL(url)
    return { success: true }
  } catch (err) {
    console.error('Error: ', err)
    return { error: 'Network error' }
  }
}
