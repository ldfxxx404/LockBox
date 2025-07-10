import { Bar } from '@/app/components/ProgressBar'

export default function Forbidden() {
  return (
    <div className='flex flex-col items-center justify-center mt-72'>
      <h1 role='status' className='text-2xl font-bold'>
        Checking in progress, please wait...
      </h1>
      <Bar />
    </div>
  )
}
