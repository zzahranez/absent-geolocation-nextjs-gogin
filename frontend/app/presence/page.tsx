import Head from "next/head"
import Navbar from "../components/global/navbar"
import {
    Sunrise,
    Sunset,
    UserCheck,
    UserX,
    Thermometer,
    AlarmClock
} from "lucide-react"


const Presence = () => {
    return (
        <>
            <Head>
                <title>Absen Page</title>
            </Head>

            <div className="bg-gray-100 text-black dark:bg-black dark:text-gray-100 min-h-screen">
                <Navbar />

                {/* --- Section 1: Jadwal Absen --- */}
                <div className="container mx-5 pt-28">
                    <div className="grid grid-cols-1 md:grid-cols-2 gap-4 p-4">
                        {/* Morning Absent */}
                        <div className="bg-white dark:bg-gray-800 shadow-md rounded-bl-3xl rounded-tr-3xl p-6 flex items-center gap-4 h-48 me-14">
                            <Sunrise className="h-10 w-10 text-yellow-500 dark:text-yellow-300" />
                            <div>
                                <h2 className="text-xl font-semibold">Morning Absent</h2>
                                <p className="mt-2 text-gray-600 dark:text-gray-300">
                                    Absen pagi dimulai dari pukul 07.00 - 09.00 WIB.
                                </p>
                            </div>
                        </div>

                        {/* Afternoon Absent */}
                        <div className="bg-white dark:bg-gray-800 shadow-md rounded-bl-3xl rounded-tr-3xl p-6 flex items-center gap-4 h-48 me-14">
                            <Sunset className="h-10 w-10 text-orange-500 dark:text-orange-300" />
                            <div>
                                <h2 className="text-xl font-semibold">Afternoon Absent</h2>
                                <p className="mt-2 text-gray-600 dark:text-gray-300">
                                    Absen siang dibuka pukul 13.00 - 15.00 WIB.
                                </p>
                            </div>
                        </div>
                    </div>
                </div>

                {/* --- Section 2: Statistik Kehadiran --- */}
                <div className="w-full bg-gray-200 text-gray-100 dark:bg-gray-900 dark:text-black py-8 px-44 h">
                    <h1 className="text-xl text-black dark:text-white mb-5 font-semibold">Monthly Presence</h1>
                    <div className="grid grid-cols-1 md:grid-cols-2 gap-6">

                        {/* Total Hadir */}
                        <div className="bg-white dark:bg-gray-800 shadow-md rounded-bl-3xl rounded-tr-3xl p-6 flex items-center gap-4 h-32">
                            <UserCheck className="h-10 w-10 text-green-500 dark:text-green-300" />
                            <div>
                                <h2 className="text-xl font-semibold text-black dark:text-white">Total Hadir</h2>
                                <p className="mt-2 text-gray-600 dark:text-gray-300">
                                    Jumlah kehadiran seluruh siswa.
                                </p>
                            </div>
                        </div>

                        {/* Total Izin */}
                        <div className="bg-white dark:bg-gray-800 shadow-md rounded-bl-3xl rounded-tr-3xl p-6 flex items-center gap-4 h-32">
                            <UserX className="h-10 w-10 text-yellow-500 dark:text-yellow-300" />
                            <div>
                                <h2 className="text-xl font-semibold text-black dark:text-white">Total Izin</h2>
                                <p className="mt-2 text-gray-600 dark:text-gray-300">
                                    Jumlah siswa yang izin resmi.
                                </p>
                            </div>
                        </div>

                        {/* Total Sakit */}
                        <div className="bg-white dark:bg-gray-800 shadow-md rounded-bl-3xl rounded-tr-3xl p-6 flex items-center gap-4 h-32">
                            <Thermometer className="h-10 w-10 text-blue-500 dark:text-blue-300" />
                            <div>
                                <h2 className="text-xl font-semibold text-black dark:text-white">Total Sakit</h2>
                                <p className="mt-2 text-gray-600 dark:text-gray-300">
                                    Jumlah siswa yang sakit.
                                </p>
                            </div>
                        </div>

                        {/* Total Terlambat */}
                        <div className="bg-white dark:bg-gray-800 shadow-md rounded-bl-3xl rounded-tr-3xl p-6 flex items-center gap-4 h-32">
                            <AlarmClock className="h-10 w-10 text-red-500 dark:text-red-300" />
                            <div>
                                <h2 className="text-xl font-semibold text-black dark:text-white">Total Terlambat</h2>
                                <p className="mt-2 text-gray-600 dark:text-gray-300">
                                    Jumlah keterlambatan siswa.
                                </p>
                            </div>
                        </div>
                    </div>
                </div>


                {/* --- Section 3 --- */}
                <div className="w-full bg-gray-100 text-black dark:bg-black dark:text-gray-100 py-6 px-4">
                    <h1 className="text-xl font-semibold text-black dark:text-white mb-4 text-center">
                        Last Week Presence
                    </h1>

                    <div className="flex flex-col gap-4">
                        {/* Card 1 */}
                        <div className="bg-white dark:bg-gray-800 shadow-lg rounded-xl p-4 space-y-2">
                            {/* Tanggal dan Status */}
                            <div className="flex items-center justify-between">
                                <p className="text-sm text-gray-700 dark:text-gray-300">15 Juli 2025</p>
                                <span className="px-3 py-1 text-sm font-medium rounded-full bg-green-200 text-green-800 dark:bg-green-700 dark:text-white">
                                    Hadir
                                </span>
                            </div>

                            {/* Time In & Time Out */}
                            <div className="mt-4">
                                <p className="text-sm font-semibold text-gray-800 dark:text-gray-200">Time In</p>
                                <p className="text-sm text-gray-600 dark:text-gray-400">08:00</p>

                                <p className="text-sm font-semibold text-gray-800 dark:text-gray-200 mt-2">Time Out</p>
                                <p className="text-sm text-gray-600 dark:text-gray-400">17:00</p>
                            </div>
                        </div>



                    </div>

                    
                </div>
            </div>
        </>
    )
}

export default Presence