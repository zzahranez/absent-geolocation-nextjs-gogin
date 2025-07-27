"use client"
import { useState } from "react";
import { Heart } from "lucide-react";

const Like = () => {
    const [liked, setLiked] = useState(false)

    function HandleLike() {
        setLiked(!liked)
    }
    return (
        <>
            <button onClick={HandleLike}>
                <Heart
                    className={`h-8 w-8 transition-colors ${liked ? "text-red-500 fill-red-500" : "text-gray-500"
                        }`}
                />
            </button>
        </>
    );
}

export default Like