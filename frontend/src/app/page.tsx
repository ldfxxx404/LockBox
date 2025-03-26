import Link from 'next/link';
import React from 'react';

export default function HomePage() {
  return (
    <div style={styles.container}>
      <h1>Welcome to LockBox</h1>
      <Link href="/login" passHref>
        <button style={styles.button}>Login</button>
      </Link>

      <Link href="/storage" passHref>
        <button style={styles.button}>Storage</button>
      </Link>
      <Link href="/profile" passHref>
        <button style={styles.button}>Profile</button>
      </Link>

      <Link href="/upload" passHref>
        <button style={styles.button}>Upload file</button>
      </Link>

      <Link href="/register" passHref>
        <button style={styles.button}>Sign Up</button>
      </Link>
    </div>
  );
}

const styles = {
  container: {
    textAlign: 'center',  
    marginTop: '50px',
  } as React.CSSProperties,  
  button: {
    padding: '10px 20px',
    fontSize: '16px',
    cursor: 'pointer',
    border: 'none',
    borderRadius: '5px',
    backgroundColor: '#5fbf22',
    color: '#fff',
    transition: 'background-color 0.3s ease',
    margin: '5px',
  } as React.CSSProperties,
};