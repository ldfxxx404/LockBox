export default function LoginPage() {
  return (
    <main
      style={{
        minHeight: '100vh',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        background: 'var(--background)',
      }}
    >
      <div
        style={{
          background: '#343746',
          padding: '2.5rem 2rem',
          borderRadius: 12,
          boxShadow: '0 2px 16px 0 #0002',
          minWidth: 400,
          maxWidth: '90vw',
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
        }}
      >
        <h1
          style={{ color: 'var(--foreground)', fontSize: 32, marginBottom: 32 }}
        >
          Login
        </h1>
        <form
          style={{
            width: '100%',
            display: 'flex',
            flexDirection: 'column',
            gap: 18,
          }}
        >
          <input
            type='email'
            placeholder='Email'
            required
            style={{
              background: '#44475a',
              color: 'var(--foreground)',
              border: 'none',
              borderRadius: 8,
              padding: '0.75rem 1rem',
              fontSize: 16,
              marginBottom: 0,
              outline: 'none',
            }}
          />
          <input
            type='password'
            placeholder='Password'
            required
            style={{
              background: '#44475a',
              color: 'var(--foreground)',
              border: 'none',
              borderRadius: 8,
              padding: '0.75rem 1rem',
              fontSize: 16,
              marginBottom: 0,
              outline: 'none',
            }}
          />
          <button
            type='submit'
            style={{
              background: '#6272a4',
              color: '#fff',
              border: 'none',
              borderRadius: 8,
              padding: '0.7rem 0',
              fontSize: 18,
              marginTop: 8,
              cursor: 'pointer',
              transition: 'background 0.2s',
            }}
          >
            Sign in
          </button>
        </form>
        <div
          style={{
            marginTop: 18,
            fontSize: 13,
            color: '#bcbcbc',
            textAlign: 'center',
          }}
        ></div>
      </div>
    </main>
  )
}
