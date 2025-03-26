
import { NextResponse } from 'next/server';

export async function POST(request: Request): Promise<NextResponse> {
    try {
      const formData = await request.json();
  
      const res = await fetch("http://localhost:5000/api/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(formData),
      });
  
      if (!res.ok) {
        const errorData = await res.json();
        return NextResponse.json(
          { message: errorData.message || "Login failed" },
          { status: res.status }
        );
      }
  
      const data = await res.json();
      return NextResponse.json(data, { status: 200 });
    } catch (error) {
      return NextResponse.json(
        { message: "An unexpected error occurred" },
        { status: 500 }
      );
    }
  }
  

