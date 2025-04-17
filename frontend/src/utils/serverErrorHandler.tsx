import { NextResponse } from 'next/server'

interface ServerError {
  message?: string | null
  status?: number
}

export const serverErrorHandler = (
  error: unknown,
  defaultMessage: string = 'Произошла ошибка'
): NextResponse => {
  let message: string = defaultMessage
  let status: number = 500

  if (error instanceof Response) {
    status = error.status
  } else if (error instanceof Error) {
    message = error.message || defaultMessage
  } else if (typeof error === 'object' && error !== null) {
    const err = error as ServerError
    message = err.message || defaultMessage
    status = err.status || 500
  }

  if (message === null) {
    message = defaultMessage
  }

  console.error(`❌ [API ERROR] Status: ${status}, Message: ${message}`)
  return NextResponse.json({ message }, { status })
}
