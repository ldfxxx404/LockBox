'use client';
import { useState } from 'react';
import styles from '../Login.module.css';

export default function Login() {
  const [formData, setFormData] = useState({
    email: '',
    password: '',
  });

  const [error, setError] = useState('');

  // Тип для события изменения в input (React.ChangeEvent<HTMLInputElement>)
  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  // Тип для события отправки формы (React.FormEvent<HTMLFormElement>)
  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const { email, password } = formData;

    if (!email || !password) {
      setError('Both fields are required');
      return;
    }

    // Логика для входа (например, запрос на сервер)
    setTimeout(() => {
      setError('');
      alert('Logged in successfully!');
      // Тут может быть редирект или что-то другое после успешного входа
    }, 1000);
  };

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Login</h1>
      {error && <div className={styles.error}>{error}</div>}
      <form onSubmit={handleSubmit}>
        <input
          type="email"
          name="email"
          value={formData.email}
          onChange={handleInputChange}
          placeholder="Email"
          className={styles.input}
        />
        <input
          type="password"
          name="password"
          value={formData.password}
          onChange={handleInputChange}
          placeholder="Password"
          className={styles.input}
        />
        <button type="submit" className={styles.button}>
          Log In
        </button>
      </form>
      <a href="/sign_up" className={styles.link}>
        Don't have an account? Register
      </a>
    </div>
  );
}
