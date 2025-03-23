'use client';
import styles from '@/styles/Upload.module.css';

export default function Upload() {
  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Upload File</h1>

      {/* Форма загрузки */}
      <form className={styles.form}>
        <input
          type="file"
          className={styles.input}
        />
        <div className={styles.fileInfo}>
          <p><strong>File:</strong> [File Name]</p>
          <p><strong>Size:</strong> [File Size]</p>
        </div>

        <div className={styles.progressBar}>
          <div
            className={styles.progress}
            style={{ width: `0%` }}
          ></div>
        </div>

        <button type="submit" className={styles.button}>
          Upload
        </button>
      </form>
    </div>
  );
}