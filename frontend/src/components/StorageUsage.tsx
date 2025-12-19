interface StorageUsageProps {
  used: number
  limit: number
}

export const StorageUsage = ({ used, limit }: StorageUsageProps) => {
  return (
    <>
      <div>
        <span className='font-semibold'>Usage:</span> {used} / {limit} MiB
      </div>
      <div className='w-full bg-[var(--background)] rounded h-3 mt-2'>
        <div
          className='bg-[var(--dracula-bar)] h-3 rounded'
          style={{ width: limit ? `${(used / limit) * 100}%` : '0%' }}
        ></div>
      </div>
    </>
  )
}
