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

const Absent = () => {
    return (
        <>
            <Head>
                <title>Absen Page</title>
            </Head>

            <div className="bg-gray-100 text-black dark:bg-black dark:text-gray-100 min-h-screen">
                <Navbar />

                {/* --- Section 1: Jadwal Absen --- */}

            </div>
        </>
    )
}

export default Absent
