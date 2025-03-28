import { NextResponse } from 'next/server';
import { SERVER_API } from '@/utils/apiUrl';
import errorHandler from '@/utils/errorHandler';

export async function POST(request: Request): Promise<NextResponse> {
    try {
        // Получаем данные формы (multipart/form-data)
        const formData = await request.formData();

        // Проверяем, что файл действительно присутствует в formData
        const file = formData.get('file');
        if (!file || !(file instanceof Blob)) {
            return new NextResponse(
                JSON.stringify({ error: 'No file provided or file is invalid' }),
                { status: 400 }
            );
        }

        // Получаем токен авторизации из заголовков запроса
        const token = request.headers.get('authorization');
        if (!token) {
            return new NextResponse(
                JSON.stringify({ error: 'No authorization token provided' }),
                { status: 401 }
            );
        }

        // Прокси-запрос к вашему бекенду для загрузки файла
        const res = await fetch(`${SERVER_API}/api/upload`, {
            method: 'POST',
            headers: {
                Authorization: token,
            },
            body: formData,
        });

        const data = await res.json();

        if (!res.ok) {
            return errorHandler({ message: data.error || 'Upload failed', status: res.status });
        }

        // Return a valid NextResponse object
        return NextResponse.json(data, { status: 200 });
    } catch (error) {
        // Make sure errorHandler returns a valid response object
        return errorHandler(error); // Ensure errorHandler returns a NextResponse object
    }
}
