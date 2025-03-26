import { NextResponse } from 'next/server';
import { SERVER_API } from '@/utils/apiUrl';

export async function POST(request: Request): Promise<NextResponse> {
  try {
    const formData = await request.json();
    
    const res = await fetch(`${SERVER_API}/api/register`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(formData),
    });

    if (!res.ok) {
      const errorData = await res.json();
      return NextResponse.json(
        { 
          message: errorData.message || "Registration failed" 
        },
        { status: res.status }
      );
    }

    const data = await res.json();
    return NextResponse.json(data, { status: 201 });

  } catch (error) {
    console.error("Registration error:", error);
    return NextResponse.json(
      { 
        message: "Server connection failed. Please try again later." 
      },
      { status: 500 }
    );
  }
}
