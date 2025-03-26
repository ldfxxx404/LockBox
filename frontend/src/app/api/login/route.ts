import { NextResponse } from 'next/server';
import { SERVER_API } from '@/utils/apiUrl';

export async function POST(request: Request): Promise<NextResponse> {
  try {
    const formData = await request.json();
    
    const res = await fetch(`${SERVER_API}/api/login`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(formData),
    });

    if (!res.ok) {
      const errorData = await res.json();
      return NextResponse.json(
        { 
          message: errorData.message || "Invalid email or password" 
        },
        { status: res.status }
      );
    }

    const data = await res.json();
    return NextResponse.json({
      user: data.user,
      token: data.token
    }, { status: 200 });

  } catch (error) {
    console.error("Login error:", error);
    return NextResponse.json(
      { 
        message: "Server connection failed. Please try again later." 
      },
      { status: 500 }
    );
  }
}
