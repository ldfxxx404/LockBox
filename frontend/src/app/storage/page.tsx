'use client';
import { useState, useEffect } from 'react';
import styles from '@/styles/Storage.module.css';

interface FileData {
  name: string;
  size: number; // Размер файла в байтах
}

export default function Storage() {
  const [files, setFiles] = useState<FileData[]>([]);
  const [totalSize, setTotalSize] = useState<number>(0); // Общий размер файлов
  const [error, setError] = useState<string>('');

  // Форматирование размера файла для отображения в удобном виде (например, KB, MB)
  const formatSize = (size: number) => {
    if (size < 1024) return `${size} B`;
    if (size < 1048576) return `${(size / 1024).toFixed(2)} KB`;
    return `${(size / 1048576).toFixed(2)} MB`;
  };

  // Загрузка данных о файлах с сервера
  useEffect(() => {
    const fetchFiles = async () => {
      try {
        const response = await fetch('/api/storage');
        if (!response.ok) {
          throw new Error('Failed to fetch files');
        }
        const data = await response.json();
        setFiles(data.files);
        setTotalSize(data.totalSize);
      } catch (err) {
        setError('An error occurred while fetching files.');
      }
    };

    fetchFiles();
  }, []);

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Storage</h1>

      {error && <div className={styles.error}>{error}</div>}

      <div className={styles.storageInfo}>
        <p><strong>Total files:</strong> {files.length}</p>
        <p><strong>Total size:</strong> {formatSize(totalSize)}</p>
      </div>

      {/* Список файлов */}
      <div className={styles.fileList}>
        {files.map((file, index) => (
          <div key={index} className={styles.fileItem}>
            <p><strong>File:</strong> {file.name}</p>
            <p><strong>Size:</strong> {formatSize(file.size)}</p>
          </div>
        ))}
      </div>
    </div>
  );
}