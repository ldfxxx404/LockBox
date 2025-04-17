import { SERVER_API } from '@/utils/apiUrl'
import { NextResponse } from 'next/server'

export async function GET(req: Request) {
  try {
    const { searchParams } = new URL(req.url)
    const fileName = searchParams.get('fileName')
    if (!fileName) {
      return NextResponse.json(
        { error: 'File name is required' },
        { status: 400 }
      )
    }

    const decodedFileName = decodeURIComponent(fileName)

    const token = req.headers.get('authorization') || ''

    const backendResponse = await fetch(
      `${SERVER_API}/api/storage/${decodedFileName}`,
      {
        method: 'GET',
        headers: { Authorization: token },
      }
    )

    if (!backendResponse.ok) {
      console.error(
        'Error fetching file from backend:',
        backendResponse.statusText
      )
      return NextResponse.json(
        { error: 'Failed to fetch file' },
        { status: backendResponse.status }
      )
    }

    const fileBlob = await backendResponse.blob()

    const encodedFileName = encodeURIComponent(decodedFileName)

    return new Response(fileBlob, {
      headers: {
        'Content-Type':
          backendResponse.headers.get('Content-Type') ||
          'application/octet-stream',
        'Content-Disposition': `attachment; filename="${encodedFileName}"`,
      },
    })
  } catch (error) {
    console.error('Internal error:', error)
    return NextResponse.json(
      { error: 'Internal Server Error' },
      { status: 500 }
    )
  }
}
