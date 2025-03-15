'use client';
import { useState } from 'react';
import styles from '../Upload.module.css';

export default function Upload() {
  const [file, setFile] = useState<File | null>(null);
  const [error, setError] = useState<string>('');
  const [uploading, setUploading] = useState<boolean>(false);
  const [uploadProgress, setUploadProgress] = useState<number>(0);

  // Обработчик выбора файла
  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const selectedFile = e.target.files ? e.target.files[0] : null;
    if (selectedFile) {
      setFile(selectedFile);
      setError('');
    }
  };

  // Форматирование размера файла
  const formatSize = (size: number) => {
    if (size < 1024) return `${size} B`;
    if (size < 1048576) return `${(size / 1024).toFixed(2)} KB`;
    return `${(size / 1048576).toFixed(2)} MB`;
  };

  // Имитация процесса загрузки
  const handleUpload = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!file) {
      setError('Please select a file to upload.');
      return;
    }

    setUploading(true);
    setUploadProgress(0);

    // Имитация процесса загрузки (здесь можно использовать реальный запрос API)
    const simulateUpload = new Promise<void>((resolve) => {
      let progress = 0;
      const interval = setInterval(() => {
        progress += 10;
        setUploadProgress(progress);

        if (progress === 100) {
          clearInterval(interval);
          resolve();
        }
      }, 500);
    });

    try {
      await simulateUpload;
      setUploading(false);
      alert('File uploaded successfully!');
      setFile(null); // Сброс файла после успешной загрузки
      setUploadProgress(0);
    } catch (err) {
      setUploading(false);
      setError('An error occurred during file upload.');
    }
  };

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Upload File</h1>

      {error && <div className={styles.error}>{error}</div>}

      {/* Форма загрузки */}
      <form onSubmit={handleUpload} className={styles.form}>
        <input
          type="file"
          onChange={handleFileChange}
          className={styles.input}
        />
        {file && (
          <div className={styles.fileInfo}>
            <p><strong>File:</strong> {file.name}</p>
            <p><strong>Size:</strong> {formatSize(file.size)}</p>
          </div>
        )}

        {uploading ? (
          <div className={styles.progressBar}>
            <div
              className={styles.progress}
              style={{ width: `${uploadProgress}%` }}
            ></div>
          </div>
        ) : (
          <button type="submit" className={styles.button}>
            Upload
          </button>
        )}
      </form>
    </div>
  );
}
