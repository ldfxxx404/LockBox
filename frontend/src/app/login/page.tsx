"use client";
import { useState } from "react";
import axios from "axios";
import styles from "@/styles/Signup.module.css";
import { useRouter } from "next/navigation";

export default function Signup() {
  const router = useRouter();
  const [formData, setFormData] = useState({ email: "", password: "" });
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState(false);
  const [loading, setLoading] = useState(false);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData((prev) => ({ ...prev, [e.target.name]: e.target.value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    setSuccess(false);
    setLoading(true);

    try {
      const response = await axios.post('/api/login', formData);
      console.log("User registered:", response.data);
      setSuccess(true);
      router.push("/profile");
    } catch (err) {
      if (axios.isAxiosError(err)) {
        setError(err.response?.data?.message || "Registration failed");
      } else {
        setError("An unexpected error occurred");
      }
    } finally {
      setLoading(false);
    }
  };
  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Login</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="email"
          name="email"
          placeholder="Email"
          value={formData.email}
          onChange={handleChange} 
          className={styles.input}
          required
        />
        <input
          type="password"
          name="password"
          placeholder="Password"
          value={formData.password}
          onChange={handleChange} 
          className={styles.input}
          required
        />
        {error && <p className={styles.error}>{error}</p>}
        {success && <p className={styles.success}>Login successful!</p>}
        <button type="submit" className={styles.button} disabled={loading}>
          {loading ? "Logging in..." : "Log In"}
        </button>
      </form>
      <a href="/register" className={styles.link}>
        Don't have an account? Register
      </a>
    </div>
  );
}
