'use client';
import { useState } from 'react';
import styles from '../Home.module.css'; // Путь к стилям

export default function Home() {
  const [user, setUser] = useState<string | null>(null); // Например, если есть авторизация

  return (
    <div className={styles.container}>
      <header className={styles.header}>
        <h1 className={styles.title}>Welcome to Our App!</h1>
      </header>

      {/* Вставить информацию о пользователе, если он авторизован */}
      {user ? (
        <div className={styles.userInfo}>
          <p>Hello, {user}!</p>
          <button className={styles.button}>Go to Profile</button>
        </div>
      ) : (
        <div className={styles.guestInfo}>
          <p>Welcome, guest! Please log in or sign up to access more features.</p>
          <button className={styles.button}>Login</button>
          <button className={styles.button}>Sign Up</button>
        </div>
      )}

      {/* Секция с основным содержимым */}
      <section className={styles.content}>
        <h2>Explore Features</h2>
        <div className={styles.feature}>
          <h3>File Storage</h3>
          <p>Upload and manage your files easily.</p>
        </div>
        <div className={styles.feature}>
          <h3>Profile Management</h3>
          <p>Update and manage your profile settings.</p>
        </div>
        <div className={styles.feature}>
          <h3>Community</h3>
          <p>Connect with others and share your ideas.</p>
        </div>
      </section>

      <footer className={styles.footer}>
        <p>© 2025 Your App Name. All rights reserved.</p>
      </footer>
    </div>
  );
}
