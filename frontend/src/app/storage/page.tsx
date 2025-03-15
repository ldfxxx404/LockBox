'use client';
import { useState, useEffect } from 'react';
import styles from '../Storage.module.css';

interface FileData {
  name: string;
  size: number;  // Размер файла в байтах
}

export default function Storage() {
  const [files, setFiles] = useState<FileData[]>([]);
  const [totalSize, setTotalSize] = useState<number>(0); // Общий размер файлов
  const [error, setError] = useState<string>('');

  // Имитация загрузки файлов (например, из API или localStorage)
  useEffect(() => {
    // Замените это на реальный запрос для получения списка файлов
    const loadedFiles: FileData[] = [
      { name: 'file1.txt', size: 2000 },
      { name: 'file2.png', size: 5000 },
      { name: 'file3.pdf', size: 8000 },
    ];

    setFiles(loadedFiles);
    setTotalSize(loadedFiles.reduce((total, file) => total + file.size, 0));
  }, []);

  // Форматирование размера файла для отображения в удобном виде (например, KB, MB)
  const formatSize = (size: number) => {
    if (size < 1024) return `${size} B`;
    if (size < 1048576) return `${(size / 1024).toFixed(2)} KB`;
    return `${(size / 1048576).toFixed(2)} MB`;
  };

  // Обработчик удаления файла
  const handleDelete = (fileName: string) => {
    setFiles((prevFiles) => prevFiles.filter(file => file.name !== fileName));

    // Обновляем общий размер после удаления
    setTotalSize((prevTotal) => prevTotal - (files.find(file => file.name === fileName)?.size || 0));

    alert(`File ${fileName} deleted successfully!`);
  };

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Storage</h1>

      {error && <div className={styles.error}>{error}</div>}

      <div className={styles.storageInfo}>
        <p><strong>Total files:</strong> {files.length}</p>
        <p><strong>Total size:</strong> {formatSize(totalSize)}</p>
      </div>

      <div className={styles.fileList}>
        <h2>Files</h2>
        {files.length === 0 ? (
          <p>No files available</p>
        ) : (
          files.map((file) => (
            <div key={file.name} className={styles.fileItem}>
              <p><strong>{file.name}</strong> - {formatSize(file.size)}</p>
              <button onClick={() => handleDelete(file.name)} className={styles.deleteButton}>Delete</button>
            </div>
          ))
        )}
      </div>
    </div>
  );
}
