"use client";

import { useState } from "react";
import axios from "axios";
import { useRouter } from "next/navigation";
import { toast, ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

export default function LoginPage() {
  const router = useRouter();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [isLoading, setIsLoading] = useState(false);

  const SpinnerDualRing = () => (
    <div className="relative w-5 h-5 mr-2">
      <div className="absolute inset-0 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
      <div className="absolute inset-1 border-2 border-white border-b-transparent rounded-full animate-spin animate-reverse"></div>
    </div>
  );

  const handleLogin = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setIsLoading(true);

    try {
      const res = await axios.post("http://localhost:8080/auth/login", {
        email,
        password,
      });

      if (!res.data.status) {
        toast.error("Login gagal! Email atau password salah");
      } else {
        toast.success("Login berhasil!");
        await new Promise((resolve) => setTimeout(resolve, 2000)); // tahan spinner
        router.push("/home");
      }
    } catch (err) {
      if (axios.isAxiosError(err)) {
        toast.error(err.response?.data?.message || "Server error");
      } else {
        toast.error("Error tidak diketahui.");
      }
    } finally {
      setIsLoading(false);
    }
  };


  return (
    <>
      <ToastContainer position="top-right" autoClose={2000} theme="colored" />
      <div className="min-h-screen flex items-center justify-center bg-gray-100 dark:bg-gray-900 px-4">
        <div className="w-full max-w-md p-8 bg-white dark:bg-gray-800 rounded-2xl shadow-lg">
          <h2 className="text-2xl font-bold text-center mb-6 dark:text-white">
            Login
          </h2>
          
        

          <form onSubmit={handleLogin} className="space-y-6">
            {/* Email Field */}
            <div>
              <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                Email
              </label>
              <input
                type="email"
                required
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                disabled={isLoading}
                className="w-full px-4 py-3 border border-gray-300 dark:border-gray-600 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-white disabled:opacity-50 disabled:cursor-not-allowed"
                placeholder="you@example.com"
              />
            </div>

            {/* Password Field */}
            <div>
              <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">
                Password
              </label>
              <input
                type="password"
                required
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                disabled={isLoading}
                className="w-full px-4 py-3 border border-gray-300 dark:border-gray-600 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-gray-700 dark:text-white disabled:opacity-50 disabled:cursor-not-allowed"
                placeholder="••••••••"
              />
            </div>

            {/* Submit Button */}
            <div>
              <button
                type="submit"
                disabled={isLoading}
                className={`w-full flex justify-center items-center gap-2 px-4 py-3 rounded-lg text-white font-semibold transition duration-300 transform hover:scale-105 ${isLoading
                    ? "bg-indigo-400 cursor-not-allowed scale-100"
                    : "bg-indigo-600 hover:bg-indigo-700 shadow-lg hover:shadow-xl"
                  }`}
              >
                {isLoading ? (
                  <>
                    <SpinnerDualRing />
                    Sedang login...
                  </>
                ) : (
                  "Login"
                )}
              </button>
            </div>
          </form>

          {/* Overlay Loading */}
          {isLoading && (
            <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
              <div className="bg-white dark:bg-gray-800 p-6 rounded-xl shadow-2xl text-center">
                <SpinnerDualRing />
                <p className="mt-4 text-gray-700 dark:text-gray-300 font-medium">
                  Sedang memproses login Anda...
                </p>
              </div>
            </div>
          )}
        </div>
      </div>
    </>
  );
}
