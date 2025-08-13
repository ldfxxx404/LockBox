export default function Main() {
  return (
    <div className='flex items-center justify-center min-h-screen bg-[var(--background)] max-sm:pl-6 pr-6'>
      <div className='bg-[#343746] px-10 py-12 rounded-xl shadow-lg flex flex-col items-center w-full max-w-md'>
        <h1 className='text-3xl text-[var(--foreground)] mb-6 font-semibold max-sm:text-2xl'>
          Welcome to LockBox
        </h1>
        <p className='text-[var(--foreground)] text-center mb-8 opacity-80'>
          Your secure file storage. Fast. Private. Reliable.
        </p>
        <div className='flex gap-4 w-full justify-center'>
          <a
            href='/login'
            className='bg-[#6272a4] text-white rounded-md px-6 py-2 text-base font-medium hover:bg-[#5861a0] transition'
          >
            Sign in
          </a>
          <a
            href='/register'
            className='bg-[#44475a] text-[var(--foreground)] rounded-md px-6 py-2 text-base font-medium hover:bg-[#51536a] transition'
          >
            Sign up
          </a>
        </div>
      </div>
    </div>
  )
}
