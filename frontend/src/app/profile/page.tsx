'use client';
import { useState, FormEvent, ChangeEvent } from 'react';
import styles from '@/styles/Profile.module.css';

interface User {
  username: string;
  email: string;
}

export default function Profile() {
  // Начальные данные пользователя
  const [user, setUser] = useState<User>({
    username: 'John Doe',
    email: 'johndoe@example.com',
  });

  // Данные для редактирования профиля
  const [formData, setFormData] = useState<User>({
    username: user.username,
    email: user.email,
  });

  // Ошибка
  const [error, setError] = useState<string>('');

  // Обработчик изменения данных
  const handleInputChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  // Обработчик отправки формы
  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();

    const { username, email } = formData;

    if (!username || !email) {
      setError('All fields are required');
      return;
    }

    try {
      // Отправка данных на сервер
      const response = await fetch('/api/profile', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData),
      });

      if (!response.ok) {
        throw new Error('Failed to update profile');
      }

      const updatedUser = await response.json();
      setUser(updatedUser); // Обновление данных пользователя
      setError('');
      alert('Profile updated successfully!');
    } catch (err) {
      setError('An error occurred while updating the profile.');
    }
  };

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>User Profile</h1>

      {/* Информация о пользователе */}
      <div className={styles.profileInfo}>
        <h2>Your Information</h2>
        <p><strong>Username:</strong> {user.username}</p>
        <p><strong>Email:</strong> {user.email}</p>
      </div>

      {/* Форма редактирования */}
      <h2>Edit Profile</h2>
      {error && <div className={styles.error}>{error}</div>}
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          name="username"
          value={formData.username}
          onChange={handleInputChange}
          placeholder="Username"
          className={styles.input}
        />
        <input
          type="email"
          name="email"
          value={formData.email}
          onChange={handleInputChange}
          placeholder="Email"
          className={styles.input}
        />
        <button type="submit" className={styles.button}>Update Profile</button>
      </form>
    </div>
  );
}