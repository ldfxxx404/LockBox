'use client';
import { useEffect, useState } from 'react';
import axios from 'axios';
import styles from '@/styles/Storage.module.css';
import { API_URL } from '@/utils/apiUrl';

export default function Storage() {
    const [files, setFiles] = useState<string[]>([]);  // Массив файлов
    const [storageInfo, setStorageInfo] = useState({ limit: 10, used: 0 });  // Информация о хранилище
    const [error, setError] = useState<string | null>(null);  // Ошибка
    const [loading, setLoading] = useState(false);  // Статус загрузки
    const [selectedFile, setSelectedFile] = useState<File | null>(null);  // Выбранный файл для загрузки

    useEffect(() => {
        fetchFiles();  // Загружаем файлы при инициализации компонента
    }, []);

    const formatSize = (sizeMB: number): string => {
        return sizeMB >= 1024 ? `${(sizeMB / 1024).toFixed(2)} GB` : `${sizeMB} MB`;
    };

    const fetchFiles = async () => {
        setLoading(true);
        setError(null);  // Сбрасываем предыдущие ошибки

        try {
            const token = localStorage.getItem('token');
            const { data } = await axios.get(`${API_URL}/api/storage`, {
                headers: { Authorization: `Bearer ${token}` },
            });

            setFiles(data.files);  // Устанавливаем файлы
            setStorageInfo(data.storage);  // Устанавливаем информацию о хранилище
        } catch (err) {
            console.error(err);
            setError("Failed to fetch storage data. Please try again.");
            setFiles([]);  // Очищаем список файлов
            setStorageInfo({ limit: 10, used: 0 });  // Устанавливаем значения по умолчанию для хранилища
        } finally {
            setLoading(false);
        }
    };

    const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        if (e.target.files) {
            setSelectedFile(e.target.files[0]); // Обновляем выбранный файл
        }
    };

    const uploadFile = async (e: React.FormEvent) => {
        e.preventDefault();
        if (!selectedFile) return;
        setLoading(true);
        setError(null);  // Сбрасываем ошибки

        const formData = new FormData();
        formData.append('file', selectedFile);

        try {
            const token = localStorage.getItem('token');
            await axios.post(`${API_URL}/api/upload`, formData, {
                headers: {
                    Authorization: `Bearer ${token}`,
                    'Content-Type': 'multipart/form-data',
                },
            });

            await fetchFiles();  // Повторно загружаем файлы после загрузки
            setSelectedFile(null);  // Очищаем выбранный файл
        } catch (err) {
            console.error(err);
            setError("Failed to upload file. Please try again.");
        } finally {
            setLoading(false);
        }
    };

    const downloadFile = async (fileName: string) => {
        setLoading(true);
        setError(null);  // Сбрасываем ошибки

        try {
            const token = localStorage.getItem('token');
            const encodedFileName = encodeURIComponent(fileName);

            const response = await axios.get(`/api/proxy-download?fileName=${encodedFileName}`, {
                headers: { Authorization: `Bearer ${token}` },
                responseType: 'blob',
            });

            const url = window.URL.createObjectURL(new Blob([response.data]));
            const link = document.createElement('a');
            link.href = url;
            document.body.appendChild(link);
            link.click();
            link.remove();

            setTimeout(() => {
                window.open(url, '_blank');
                window.URL.revokeObjectURL(url);
            }, 500);
        } catch (err) {
            console.error(err);
            setError("Failed to download file. Please try again.");
        } finally {
            setLoading(false);
        }
    };

    const deleteFile = async (fileName: string) => {
        setLoading(true);
        setError(null);  // Сбрасываем ошибки

        try {
            const token = localStorage.getItem('token');
            const encodedFileName = encodeURIComponent(fileName);

            await axios.delete(`/api/storage?fileName=${encodedFileName}`, {
                headers: { Authorization: `Bearer ${token}` },
            });

            await fetchFiles();  // Повторно загружаем файлы после удаления
        } catch (err) {
            console.error(err);
            setError("Failed to delete file. Please try again.");
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className={styles.container}>
            <div className={styles.card}>
                <h1 className={styles.title}>Your Storage</h1>
                {error && <p className={styles.error}>{error}</p>} {/* Показываем ошибку если есть */}

                <div className={styles.storageInfo}>
                    <div className={styles.infoItem}>
                        <span>Storage Limit</span>
                        <strong>{formatSize(storageInfo.limit)}</strong>
                    </div>
                    <div className={styles.infoItem}>
                        <span>Used Space</span>
                        <strong>{formatSize(storageInfo.used)}</strong>
                    </div>
                </div>

                <form onSubmit={uploadFile} className={styles.uploadForm}>
                    <label className={styles.fileInputLabel}>
                        <input type="file" onChange={handleFileChange} className={styles.fileInput} />
                        Choose File
                    </label>
                    <button type="submit" className={styles.uploadButton} disabled={loading || !selectedFile}>
                        {loading ? 'Uploading...' : 'Upload'}
                    </button>
                </form>

                <div className={styles.filesSection}>
                    <h2 className={styles.sectionTitle}>Files</h2>
                    {files.length > 0 ? (
                        <ul className={styles.filesList}>
                            {files.map((file, index) => (
                                <li key={index} className={styles.fileItem}>
                                    <span className={styles.fileIcon}>📄</span>
                                    <span className={styles.fileName}>{file}</span>
                                    <button
                                        className={styles.deleteButton}
                                        onClick={() => deleteFile(file)}
                                        disabled={loading}
                                    >
                                        🗑️
                                    </button>
                                    <button
                                        className={styles.downloadButton}
                                        onClick={() => downloadFile(file)}
                                        disabled={loading}
                                    >
                                        Download
                                    </button>
                                </li>
                            ))}
                        </ul>
                    ) : (
                        <p className={styles.noFiles}>No files found. Upload your first file!</p>
                    )}
                </div>
            </div>
        </div>
    );
}
