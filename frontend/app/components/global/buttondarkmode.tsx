"use client"

import { useTheme } from "next-themes"
import { useEffect, useState } from "react"
import { Moon, Sun } from "lucide-react"

const ButtonDarkMode = () => {
    const { theme, setTheme } = useTheme()
    const [mounted, setMounted] = useState(false)

    // Ini agar icon tidak flash saat render SSR
    useEffect(() => {
        setMounted(true)
    }, [])

    if (!mounted) return null

    const isDark = theme === "dark"

    return (
        <button
            onClick={() => setTheme(isDark ? "light" : "dark")}
            className={`relative inline-flex items-center justify-center w-8 h-8 rounded-full transition-all duration-500 transform hover:scale-110 active:scale-95 shadow-lg hover:shadow-xl
            ${isDark ? 'bg-gradient-to-r from-purple-600 to-blue-600 shadow-purple-500/25' : 'bg-gradient-to-r from-yellow-400 to-orange-500 shadow-orange-500/25'}
            border-4 border-white/20 backdrop-blur-sm group`}
        >
            <div className="absolute inset-0 rounded-full bg-white/10 opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>

            <div className={`transition-all duration-500 transform ${isDark ? 'rotate-180 scale-110' : 'rotate-0 scale-100'}`}>
                {isDark ? <Moon className="w-7 h-7 text-white drop-shadow-lg" /> : <Sun className="w-7 h-7 text-white drop-shadow-lg" />}
            </div>

            <div
                className={`absolute inset-0 rounded-full border-2 border-white/30 transition-all duration-500 ${
                    isDark ? "animate-pulse" : "animate-spin"
                }`}
                style={{
                    animation: isDark ? "pulse 2s infinite" : "spin 8s linear infinite",
                }}
            ></div>

            <div
                className={`absolute inset-0 rounded-full blur-xl opacity-30 transition-all duration-500 ${
                    isDark
                        ? "bg-gradient-to-r from-purple-400 to-blue-400"
                        : "bg-gradient-to-r from-yellow-300 to-orange-400"
                }`}
            ></div>
        </button>
    )
}

export default ButtonDarkMode
