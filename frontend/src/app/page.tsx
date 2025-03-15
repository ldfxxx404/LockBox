// TODO: add /upload,  /sgin in button's
import Link from 'next/link';

export default function Home() {
  return (
    <div style={styles.container}>
      <h1>Home Server Frontend v1.1</h1>
      <Link href="/login" passHref>
        <button style={styles.button}>Login</button>

      </Link>
      <Link href="/home" passHref>
        <button style={styles.button}>Home</button>
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

      <Link href="/sign_up" passHref>
        <button style={styles.button}>Sign Up</button>
      </Link>
    </div>
  );
}

// Пример стилей для компонента
const styles = {
  container: {
    textAlign: 'center' as 'center',  // Явное указание типа
    marginTop: '50px',
  } as React.CSSProperties,  // Указываем, что это объект стилей с типами из React
  button: {
    padding: '10px 20px',
    fontSize: '16px',
    cursor: 'pointer',
    border: 'none',
    borderRadius: '5px',
    backgroundColor: '#007bff',
    color: '#fff',
    transition: 'background-color 0.3s ease',
    margin: '10px',
  } as React.CSSProperties,
};
