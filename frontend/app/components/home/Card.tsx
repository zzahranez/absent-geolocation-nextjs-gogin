interface PropCardHome {
    name: string;
    title: string;
    description: string;
}
function CardHome({ name, title, description }: PropCardHome) {
    return (
        <div className="group relative rounded-3xl border border-zinc-200 dark:border-zinc-700/50 bg-white/90 dark:bg-zinc-900/60 backdrop-blur-xl shadow-sm hover:shadow-2xl transition-all duration-500 hover:scale-[1.03] will-change-transform overflow-hidden">

            {/* Subtle radial gradient background on hover */}
            <div
                className="absolute inset-0 opacity-0 group-hover:opacity-100 transition-opacity duration-500 rounded-3xl pointer-events-none"
                style={{
                    background: 'radial-gradient(circle at 50% 50%, rgba(59, 130, 246, 0.1) 0%, transparent 60%)',
                }}
            />

            {/* Gradient overlay on hover (left to right) */}
           <div className="absolute inset-0 rounded-3xl bg-gradient-to-r from-transparent via-orange-400/30 to-transparent dark:via-blue-900/10 opacity-0 group-hover:opacity-100 transition-opacity duration-500 pointer-events-none" />



            {/* Inner subtle border glow on hover */}
            <div className="absolute inset-0 rounded-3xl border border-transparent group-hover:border-blue-200 dark:group-hover:border-blue-800/50 transition-colors duration-500 pointer-events-none" />

            {/* Content */}
            <div className="relative z-10 p-6">
                <h2 className="text-xl md:text-2xl font-semibold text-zinc-800 dark:text-zinc-50 flex items-center gap-2.5 transition-colors duration-300 group-hover:text-black dark:group-hover:text-blue-300">
                    <span className="inline-block text-lg transform transition-transform duration-500 group-hover:rotate-12 group-hover:scale-125">
                        🚀
                    </span>
                    <span className="tracking-tight">{title}</span>
                </h2>

                <p className="text-sm font-medium text-zinc-500 dark:text-zinc-400 mt-1">
                    By{' '}
                    <span className="text-zinc-700 dark:text-zinc-200 font-semibold transition-colors duration-300">
                        {name}
                    </span>
                </p>

                <p className="mt-4 text-base leading-relaxed text-zinc-700 dark:text-zinc-300 line-clamp-3 group-hover:line-clamp-none transition-all duration-300">
                    {description}
                </p>
            </div>

            {/* Optional: Floating action indicator (bottom right) */}
            <div className="absolute bottom-4 right-6 opacity-0 group-hover:opacity-100 transform translate-y-2 group-hover:translate-y-0 transition-all duration-300">
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="h-4 w-4 text-blue-500 dark:text-blue-400"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                >
                    <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        strokeWidth={2}
                        d="M13 7l5 5m0 0l-5 5m5-5H6"
                    />
                </svg>
            </div>
        </div>
    );
}


export default CardHome;