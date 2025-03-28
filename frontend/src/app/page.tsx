'use client';

import React, {useState} from 'react';
import axios from 'axios';
import {API_URL} from '@/utils/apiUrl';
import Link from "next/link";

export default function HomePage() {
    const [selectedFile, setSelectedFile] = useState<File | null>(null);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);

    const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        if (e.target.files) {
            setSelectedFile(e.target.files[0]); // Обновляем выбранный файл
        }
    };

    const handleFileUpload = async () => {
        if (!selectedFile) return;

        setLoading(true);
        setError(null);

        const formData = new FormData();
        formData.append('file', selectedFile);

        try {
            const token = localStorage.getItem('token');
            const response = await axios.post(`${API_URL}/api/upload`, formData, {
                headers: {
                    Authorization: `Bearer ${token}`,
                    'Content-Type': 'multipart/form-data',
                },
            });
            console.log('File uploaded successfully:', response.data);
        } catch (err) {
            setError('Failed to upload file. Please try again.');
            console.error(err);
        } finally {
            setLoading(false);
        }
    };

    return (
        <div style={styles.container}>
            <h1>Welcome to LockBox</h1>
            <Link href="/login" passHref>
                <button style={styles.button}>Login</button>
            </Link>
            <Link href="/storage" passHref>
                <button style={styles.button}>Storage</button>
            </Link>
            <Link href="/profile" passHref>
                <button style={styles.button}>Profile</button>
            </Link>
            <Link href="/register" passHref>
                <button style={styles.button}>Sign Up</button>
            </Link>

            <div style={{marginTop: '20px'}}>
                <input
                    type="file"
                    style={{display: 'none'}}
                    id="file-input"
                    onChange={handleFileChange}
                />
                <label htmlFor="file-input">
                    <button
                        style={styles.button}
                        disabled={loading}
                        onClick={() => document.getElementById('file-input')?.click()}
                    >
                        {loading ? 'Uploading...' : 'Upload a File'}
                    </button>
                </label>

                {selectedFile && (
                    <div style={{marginTop: '20px'}}>
                        <p>Selected file: {selectedFile.name}</p>
                        <button onClick={handleFileUpload} style={styles.button}>
                            Confirm Upload
                        </button>
                    </div>
                )}

                {error && <p style={{color: 'red'}}>{error}</p>}
            </div>
        </div>
    );
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
    },
};
