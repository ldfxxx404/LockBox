'use client'

export async function FileUploader(file: File) {
  const token = sessionStorage.getItem('token')
  const uploadForm = new FormData()
  uploadForm.append('file', file)

  try {
    const response = await fetch('/api/upload', {
      method: 'POST',
      body: uploadForm,
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
    if (!response.ok) {
      return { error: 'error' }
    }
  } catch (err) {
    return console.error('Error: ', err)
  }
}
