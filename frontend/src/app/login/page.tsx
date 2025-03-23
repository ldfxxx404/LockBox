'use client';
import styles from '@/styles/Login.module.css';

export default function Login() {
  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Login</h1>
      <form>
        <input
          type="email"
          name="email"
          placeholder="Email"
          className={styles.input}
        />
        <input
          type="password"
          name="password"
          placeholder="Password"
          className={styles.input}
        />
        <button type="submit" className={styles.button}>
          Log In
        </button>
      </form>
      <a href="/register" className={styles.link}>
        Don't have an account? Register
      </a>
    </div>
  );
}