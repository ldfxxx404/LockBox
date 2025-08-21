import type { Metadata } from 'next'
import localFont from 'next/font/local'
import './globals.css'

const geistSans = localFont({
  src: [
    {
      path: '../../public/fonts/Geist/webfonts/Geist-Regular.woff2',
      weight: '400',
      style: 'normal',
    },
    {
      path: '../../public/fonts/Geist/webfonts/Geist-Bold.woff2',
      weight: '700',
      style: 'normal',
    },
    {
      path: '../../public/fonts/Geist/webfonts/Geist-Italic[wght].woff2',
      style: 'italic',
    },
  ],
  variable: '--font-geist-sans',
  display: 'swap',
})

export const metadata: Metadata = {
  title: 'LockBox',
  description: 'Made by lnkssr & ldfxxx',
}

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <html lang='en'>
      <body className={`${geistSans.variable} antialiased`}>{children}</body>
    </html>
  )
}
