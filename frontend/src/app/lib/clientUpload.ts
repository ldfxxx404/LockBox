'use client'

export async function FileUploader(file: File) {
  const uploadForm = new FormData()
  uploadForm.append('file', file)

  try {
    const response = await fetch('/api/upload', {
      method: 'POST',
      body: uploadForm,
      headers: {
        Authorization: `Bearer ${sessionStorage.getItem('token')}`,
      },
    })
    if (!response.ok) {
      return { error: 'error' }
    }
  } catch (err) {
    return console.error('Error: ', err)
  }
}
