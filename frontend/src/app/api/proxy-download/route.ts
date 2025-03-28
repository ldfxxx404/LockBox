import { NextResponse } from 'next/server';
import { SERVER_API } from '@/utils/apiUrl';

export async function GET(req: Request) {
    try {
        const { searchParams } = new URL(req.url);
        const fileName = searchParams.get('fileName');
        if (!fileName) {
            return NextResponse.json({ error: 'File name is required' }, { status: 400 });
        }

        // Декодируем имя файла на сервере для корректной обработки
        const decodedFileName = decodeURIComponent(fileName);

        // Далее мы выполняем работу с файлом на сервере (например, получаем его с другого API)
        const token = req.headers.get('authorization') || '';

        const backendResponse = await fetch(`${SERVER_API}/api/storage/${decodedFileName}`, {
            method: 'GET',
            headers: { Authorization: token },
        });

        if (!backendResponse.ok) {
            console.error('Error fetching file from backend:', backendResponse.statusText);
            return NextResponse.json({ error: 'Failed to fetch file' }, { status: backendResponse.status });
        }

        // Бинарное содержимое файла
        const fileBlob = await backendResponse.blob();

        // Кодируем имя файла обратно в URL-совместимую форму
        const encodedFileName = encodeURIComponent(decodedFileName);

        // Возвращаем файл с правильными заголовками
        return new Response(fileBlob, {
            headers: {
                'Content-Type': backendResponse.headers.get('Content-Type') || 'application/octet-stream',
                'Content-Disposition': `attachment; filename="${encodedFileName}"`, // Возвращаем обратно закодированное имя
            },
        });
    } catch (error) {
        console.error('Internal error:', error);
        return NextResponse.json({ error: 'Internal Server Error' }, { status: 500 });
    }
}
