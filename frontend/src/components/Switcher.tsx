'use client'

import React, { SetStateAction, useState } from 'react'

const tabs = ['Users', 'Admins']

export const Switcher = () => {
  const [activeTab, setActiveTab] = useState('Users')

  const handleTabClick = (tabName: SetStateAction<string>) => {
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

      <div className='mt-4 text-center text-lg font-semibold text-white'>
        {activeTab === 'Users' && <p>Users</p>}
        {activeTab === 'Admins' && <p>Admins</p>}
      </div>
    </div>
  )
}
