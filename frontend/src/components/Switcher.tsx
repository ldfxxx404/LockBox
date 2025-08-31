'use client'

import React, { Dispatch, SetStateAction } from 'react'

const tabs = ['Users', 'Admins'] as const

interface SwitcherProps {
  activeTab: string
  setActiveTab: Dispatch<SetStateAction<string>>
}

export const Switcher = ({ activeTab, setActiveTab }: SwitcherProps) => {
  const handleTabClick = (tabName: string) => {
    setActiveTab(tabName)
  }

  return (
    <div className='mb-6'>
      <div className='flex space-x-4 bg-[var(--dracula-comment)] p-2 rounded-xl shadow-md'>
        {tabs.map(tab => (
          <button
            key={tab}
            onClick={() => handleTabClick(tab)}
            className={`px-6 py-2 rounded-lg font-medium transition-colors duration-300 cursor-pointer 
              ${
                activeTab === tab
                  ? 'bg-[var(--dracula-purple)] text-[var(--dracula-white)] shadow-md'
                  : 'bg-[var(--dracula-card)] text-[var(--dracula-selection)] hover:bg-[var(--dracula-selection-light)]'
              }`}
          >
            {tab}
          </button>
        ))}
      </div>
    </div>
  )
}
