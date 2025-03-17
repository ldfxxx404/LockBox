'use client';
import styles from '@/styles/Signup.module.css';

export default function Signup() {
  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Sign Up</h1>
      <form>
        <input
          type="text"
          name="username"
          placeholder="Username"
          className={styles.input}
        />
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
          Register
        </button>
      </form>
    </div>
  );
}