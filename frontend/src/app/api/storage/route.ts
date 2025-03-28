import { NextResponse } from 'next/server';
import { SERVER_API } from '@/utils/apiUrl';

export async function GET(request: Request) {
    try {
        // Извлекаем JWT-токен из заголовка авторизации
        const authorizationHeader = request.headers.get('authorization');
        if (!authorizationHeader || !authorizationHeader.startsWith('Bearer ')) {
            return new NextResponse(
                JSON.stringify({ error: 'No valid authorization token provided' }),
                { status: 401 }
            );
        }

        const token = authorizationHeader.split(' ')[1]; // Extract token after "Bearer"

        // Прокси-запрос к бэкенду
        const res = await fetch(`${SERVER_API}/api/storage`, {
            method: 'GET',
            headers: { Authorization: `Bearer ${token}` },
        });

        const data = await res.json();

        // Если сервер вернул ошибку, возвращаем ошибку с данным статусом
        if (!res.ok) {
            return new NextResponse(
                JSON.stringify({ error: data.error || 'Failed to fetch storage' }),
                { status: res.status }
            );
        }

        // Возвращаем успешный ответ с данными
        return NextResponse.json(data, { status: 200 });

    } catch (error: unknown) {
        // Ловим все остальные ошибки и возвращаем ошибку сервера
        console.error(error);  // Log the error for debugging
        return new NextResponse(
            JSON.stringify({ error: 'Internal server error' }),
            { status: 500 }
        );
    }
}
