import { NextResponse } from "next/server";
import axios from "axios";

export async function POST(req: Request) {
  try {
    const formData = await req.json();

    if (!formData.email || !formData.password) {
      return NextResponse.json(
        { message: "Email and password are required" },
        { status: 400 }
      );
    }

    const response = await axios.post('http://localhost:5000/api/login', formData, {
      headers: { "Content-Type": "application/json" },
    });

    return NextResponse.json(response.data, { status: 200 });
  } catch (err) {
    if (axios.isAxiosError(err)) {
      return NextResponse.json(
        { message: err.response?.data?.message ?? "Login failed" },
        { status: err.response?.status ?? 500 }
      );
    }

    return NextResponse.json(
      { message: "Internal server error" },
      { status: 500 }
    );
  }
}
