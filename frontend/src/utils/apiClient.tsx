import {NextResponse} from 'next/server';
import {SERVER_API} from './apiUrl';
import {serverErrorHandler} from './serverErrorHandler';

interface ApiResponse<T = unknown> {
    data?: T;
    message?: string;
}

export async function callApi<T>(
    endpoint: string,
    method: string,
    body: unknown,
    successStatus: number = 200
): Promise<NextResponse> {
    try {
        const res = await fetch(`${SERVER_API}${endpoint}`, {
            method,
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(body),
        });

        let responseData: ApiResponse<T>;

        try {
            responseData = await res.json();
        } catch {
            return serverErrorHandler({
                message: "Invalid response format from the server",
                status: res.status,
            });
        }

        if (!res.ok) {
            return serverErrorHandler({
                message: responseData.message || "Error",
                status: res.status,
            });
        }

        return NextResponse.json(responseData.data || responseData, {
            status: successStatus,
        });
    } catch (error: unknown) {
        return serverErrorHandler(error);
    }
}
