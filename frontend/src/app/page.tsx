// TODO: add /home, /login, /storage, /profile, /upload, /login /sgin in button's
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
  } as React.CSSProperties,
};
