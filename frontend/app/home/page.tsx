"use client"
import Head from "next/head";
import Navbar from "../components/global/navbar";
import Card from "../components/home/Card"
import Like from "../components/like";
import { useEffect, useState } from "react";
export default function Home() {
    const [message, setMessage] = useState("")
    useEffect(() => {
        const msg = localStorage.getItem('message')
        if (msg) {
            setMessage(msg)
            localStorage.removeItem('message')
        }
    }, [])


    setTimeout(() => {
        setMessage("")
    }, 3000);


    return (
        <>
            <Head>
                <title>Home</title>
            </Head>
            <div className=" text-black bg-gray-200 dark:bg-gray-900 dark:text-white min-h-screen">
                <Navbar />
                <div className="pt-28">
                    <div
                        className="relative bg-center bg-cover text-white dark:text-white"
                        style={{ backgroundImage: "url('/your-image.jpg')" }} // Gambar dari folder /public
                    >
                        {/* Overlay gelap */}
                        <div className="absolute inset-0 bg-black/40 dark:bg-gray-900/60" />

                        {/* Konten Tengah */}
                        <div className="relative z-10 flex flex-col items-center justify-center text-center gap-y-4 py-24 px-4">
                            <Like />
                            <h1 className="text-3xl md:text-4xl font-bold">Welcome to the Geolocations Absent</h1>
                            <h1 className="text-3xl md:text-4xl font-bold">Zeno Robo Absent</h1>
                            <h1 className="text-lg md:text-xl">Scroll For More Information</h1>
                            {message && <p>{message}</p>}
                        </div>
                    </div>

                    {/* Section 2 */}
                    <div className="w-full bg-gray-100 text-black dark:bg-gray-900 dark:text-white pb-7">
                        <div className=" mx-12 text-center mb-16 pt-12">
                            <h1 className=" font-bold text-xl">Service</h1>
                        </div>
                        {/* Card */}

                        <div className=" mx-4">
                            <div className="grid grid-cols-3 gap-10 mx-12 ">
                                <Card name="Jack" description="This is a geolocoation service absent" title="Absent" />
                                <Card name="Jack" description="This is a geolocoation service absent" title="Absent" />
                                <Card name="Jack" description="This is a geolocoation service absent" title="Absent" />
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </>
    )
}



