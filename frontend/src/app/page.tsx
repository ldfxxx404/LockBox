'use client'

import React, { useState } from 'react'
import axios from 'axios'
import Link from 'next/link'
import { useRouter } from 'next/navigation'

export default function HomePage() {
	const router = useRouter()
	const [selectedFile, setSelectedFile] = useState<File | null>(null)
	const [loading, setLoading] = useState(false)
	const [error, setError] = useState<string | null>(null)

	const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		if (e.target.files && e.target.files.length > 0) {
			setSelectedFile(e.target.files[0])
		}
	}

	const handleFileUpload = async () => {
		if (!selectedFile) return

		setLoading(true)
		setError(null)

		const formData = new FormData()
		formData.append('file', selectedFile)

		try {
			const token = localStorage.getItem('token')
			await axios.post(`/api/upload`, formData, {
				headers: {
					Authorization: `Bearer ${token}`,
					'Content-Type': 'multipart/form-data',
				},
			})

			router.push('/storage')
		} catch (err) {
			setError('Failed to upload file. Please try again.')
			console.error(err)
		} finally {
			setLoading(false)
		}
	}

	return (
		<div style={styles.container}>
			<h1>Welcome to LockBox</h1>
			<Link href='/login'>
				<button style={styles.button}>Login</button>
			</Link>
			<Link href='/storage'>
				<button style={styles.button}>Storage</button>
			</Link>
			<Link href='/profile'>
				<button style={styles.button}>Profile</button>
			</Link>
			<Link href='/register'>
				<button style={styles.button}>Sign Up</button>
			</Link>

			<div style={{ marginTop: '20px' }}>
				<input
					type='file'
					id='file-input'
					style={{ display: 'none' }}
					onChange={handleFileChange}
				/>

				<label htmlFor='file-input' style={styles.button}>
					{loading ? 'Uploading...' : 'Upload a File'}
				</label>

				{selectedFile && (
					<div style={{ marginTop: '20px' }}>
						<p>Selected file: {selectedFile.name}</p>
						<button
							onClick={handleFileUpload}
							style={styles.button}
							disabled={loading}
						>
							{loading ? 'Uploading...' : 'Confirm Upload'}
						</button>
					</div>
				)}

				{error && <p style={{ color: 'red' }}>{error}</p>}
			</div>
		</div>
	)
}

const styles = {
	container: {
		textAlign: 'center' as const,
		marginTop: '50px',
	},
	button: {
		padding: '10px 20px',
		fontSize: '16px',
		cursor: 'pointer',
		border: 'none',
		borderRadius: '5px',
		backgroundColor: '#5fbf22',
		color: '#fff',
		transition: 'background-color 0.3s ease',
		margin: '5px',
		display: 'inline-block',
		textAlign: 'center' as const,
	},
}
