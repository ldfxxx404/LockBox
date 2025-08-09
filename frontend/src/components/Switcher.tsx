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
      <div className='flex space-x-4 bg-[#2d2f44] p-2 rounded-xl shadow-md'>
        {tabs.map(tab => (
          <button
            key={tab}
            onClick={() => handleTabClick(tab)}
            className={`px-6 py-2 rounded-lg font-medium transition-colors duration-300 
              ${
                activeTab === tab
                  ? 'bg-purple-600 text-white shadow-md'
                  : 'bg-[#1f2133] text-gray-300 hover:bg-[#3c3f5e]'
              }`}
          >
            {tab}
          </button>
        ))}
      </div>
    </div>
  )
}
