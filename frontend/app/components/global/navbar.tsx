"use client"
import Link from "next/link"
import ButtonDarkMode from "./buttondarkmode"
import Image from "next/image"
import { usePathname } from "next/navigation"

const Navbar = () => {
    const pathname = usePathname()

    return (
        <>
            <nav className="bg-transparent py-4 fixed w-full z-50 px-24">
                <div className="flex items-center justify-between">
                    {/* Logo Section */}
                    <div className="flex items-center">
                        <Image
                            src="/Logo.png"
                            alt="My Site Logo"
                            width={45}
                            height={70}

                        />
                    </div>

                    {/* Navigation Menu */}
                    <div className="flex items-center space-x-8">
                        <ul className="flex items-center space-x-6">
                            {[
                                { href: "/home", label: "Home" },
                                { href: "/absent", label: "Absent" },
                                { href: "/presence", label: "Presence" },
                            ].map((item) => (
                                <li key={item.href}>
                                    <Link
                                        href={item.href}
                                        className={`relative group font-medium transition duration-200
                  ${pathname === item.href
                                                ? "text-orange-700 dark:text-purple-400 font-semibold"
                                                : "text-black dark:text-white hover:text-orange-700 dark:hover:text-purple-400"
                                            }
                `}
                                    >
                                        {item.label}
                                        <span
                                            className={`
                    absolute -bottom-1 left-0 h-0.5 transition-all duration-200
                    ${pathname === item.href
                                                    ? "w-full bg-orange-700 dark:bg-purple-400"
                                                    : "w-0 group-hover:w-full bg-orange-700 dark:bg-purple-400"
                                                }
                  `}
                                        ></span>
                                    </Link>
                                </li>
                            ))}
                            <li className="ml-4">
                                <ButtonDarkMode />
                            </li>
                        </ul>
                    </div>
                </div>
            </nav>
        </>
    )
}

export default Navbar