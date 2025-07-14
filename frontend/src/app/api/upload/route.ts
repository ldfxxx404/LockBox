import { UPLOAD_URL } from '@/app/constants/api'

export async function POST(req: Request) {
  const formData = await req.formData()
  const file = formData.get('file')
  const authHeader = req.headers.get('authorization')
  const token = authHeader?.split(' ')[1]

  try {
    if (!token) {
      return Response.json({ error: 'Unauthorized' }, { status: 403 })
    }
    if (!file) {
      return Response.json({ error: 'File not upload' }, { status: 400 })
    }
    const uploadForm = new FormData()
    uploadForm.append('file', file)

    const response = await fetch(UPLOAD_URL, {
      method: 'POST',
      body: uploadForm,
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })

    if (!response.ok) {
      const text = await response.text()
      console.error('UPLOAD_URL error:', text)
      return Response.json(
        { error: 'Upload failed', details: text },
        { status: 520 }
      )
    }
    const data = await response.json()

    return Response.json(data)
  } catch (err) {
    console.error('Error:', err)
    Response.json({ error: 'Internal server error' }, { status: 500 })
  }
}
