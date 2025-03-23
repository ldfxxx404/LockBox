'use client';
import { useState, FormEvent, ChangeEvent } from 'react';
import styles from '@/styles/Profile.module.css';

interface User {
  username: string;
  email: string;
}

export default function Profile() {

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>User Profile</h1>

      {/* Информация о пользователе */}
      <div className={styles.profileInfo}>
        <h2>Your Information</h2>
        <p><strong>Username:</strong> { }</p>
        <p><strong>Email:</strong> { }</p>
      </div>

      {/* Форма редактирования */}
      <h2>Edit Profile</h2>
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
        <button type="submit" className={styles.button}>Update Profile</button>
      </form>
    </div>
  );
}