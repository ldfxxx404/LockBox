import styles from '@/styles/Error.module.css';

export default function NotFoundPage() {
  return (
    <div className={styles.errorContainer}>
      <div className={styles.errorContent}>
        <h1 className={styles.errorTitle}>404</h1>
        <p className={styles.errorText}>Страница не найдена</p>
        <p className={styles.errorDescription}>
        Упс! Кажется, эта страница сбежала от нас.</p>
      </div>
    </div>
  );
}
