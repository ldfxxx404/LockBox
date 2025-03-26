import { NextResponse } from 'next/server';

type ServerError = {
  message?: string;
  status?: number;
};

export const serverErrorHandler = (
  error: unknown,
  defaultMessage: string = "An error occurred"
): NextResponse => {
  console.error("Server error:", error);

  if (error instanceof Error) {
    const err = error as ServerError & { response?: Response };
    
    if (err.response) {
      return NextResponse.json(
        { message: err.message || defaultMessage },
        { status: err.status || 500 }
      );
    }
    
    return NextResponse.json(
      { message: err.message || defaultMessage },
      { status: err.status || 500 }
    );
  }

  return NextResponse.json(
    { message: defaultMessage },
    { status: 500 }
  );
};
