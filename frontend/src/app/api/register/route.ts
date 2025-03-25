import { NextResponse } from 'next/server';
import axios, { AxiosError } from 'axios';
import { API_REGISTER_URL } from '@/app/api/constants';

export async function POST(request: Request): Promise<NextResponse> {
  try {
    const formData = await request.json();
    
    const response = await axios.post(
      API_REGISTER_URL,
      formData,
      { headers: { "Content-Type": "application/json" } }
    );
    
    return NextResponse.json(response.data, { status: 200 });
  } catch (err: unknown) {
    if (axios.isAxiosError(err)) {
      return NextResponse.json(
        { message: err.response?.data?.message ?? err.response?.data ?? "Registration failed" },
        { status: err.response?.status ?? 500 }
      );
    }
    
    return NextResponse.json(
      { message: "An unexpected error occurred" },
      { status: 500 }
    );
  }
}
