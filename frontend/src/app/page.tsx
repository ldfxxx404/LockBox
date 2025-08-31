export default function Main() {
  return (
    <div className='flex items-center justify-center min-h-screen bg-[var(--background)] max-sm:pl-6 pr-6'>
      <div className='bg-[var(--dracula-comment)] px-10 py-12 rounded-xl shadow-lg flex flex-col items-center w-full max-w-md'>
        <h1 className='text-3xl text-[var(--foreground)] mb-6 font-semibold max-sm:text-2xl'>
          Welcome to LockBox
        </h1>
        <p className='text-[var(--foreground)] text-center mb-8 opacity-80'>
          Your secure file storage. Fast. Private. Reliable.
        </p>
        <div className='flex gap-4 w-full justify-center'>
          <a
            href='/login'
            className='bg-[var(--dracula-button-bright)] text-[var(--dracula-white)] rounded-md px-6 py-2 text-base font-medium hover:bg-[var(--dracula-button-bright-hover)] transition'
          >
            Sign in
          </a>
          <a
            href='/register'
            className='bg-[var(--dracula-button-dark)] text-[var(--foreground)] rounded-md px-6 py-2 text-base font-medium hover:bg-[var(--dracula-button-dark-hover)] transition'
          >
            Sign up
          </a>
        </div>
      </div>
    </div>
  )
}
