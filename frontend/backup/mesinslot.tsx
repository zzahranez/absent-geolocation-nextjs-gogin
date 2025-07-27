// "use client";

// import Navbar from "@/app/components/navbar";
// import React, { useState, useEffect } from "react";


// const slotItems = ["🍒", "🍋", "🔔", "⭐", "💎", "🍉", "🍇"];

// export default function UltimateSlotMachine() {
//     const [slots, setSlots] = useState(["🍒", "🍋", "🔔"]);
//     const [spinning, setSpinning] = useState(false);
//     const [credits, setCredits] = useState(10000);
//     const [bet, setBet] = useState(50);
//     const [winAmount, setWinAmount] = useState(0);
//     const [showWinModal, setShowWinModal] = useState(false);
//     const [jackpotGlow, setJackpotGlow] = useState(false);
//     const [totalWins, setTotalWins] = useState(0);
//     const [winStreak, setWinStreak] = useState(0);
//     const [reelAnimations, setReelAnimations] = useState([false, false, false]);

//     useEffect(() => {
//         if (showWinModal) {
//             const timer = setTimeout(() => setShowWinModal(false), 3000);
//             return () => clearTimeout(timer);
//         }
//     }, [showWinModal]);

//     const checkWin = (result: string[]): number => {
//         const [a, b, c] = result;
//         if (a === b && b === c) {
//             if (a === "💎") return bet * 150;
//             if (a === "⭐") return bet * 100;
//             if (a === "🔔") return bet * 75;
//             if (a === "🍒") return bet * 50;
//             return bet * 30;
//         }
//         if ((a === b || b === c || a === c)) {
//             if (result.includes("💎")) return bet * 8;
//             if (result.includes("⭐")) return bet * 6;
//             if (result.includes("🔔")) return bet * 4;
//             return bet * 2;
//         }
//         if (result.includes("💎") && result.includes("⭐")) return bet * 3;
//         return 0;
//     };

//     const spin = async () => {
//         if (spinning || credits < bet) return;
//         setSpinning(true);
//         setCredits(prev => prev - bet);
//         setWinAmount(0);
//         setShowWinModal(false);
//         setJackpotGlow(false);

//         const reelStopTimes = [1600, 2200, 2800];
//         const tempSlots = [...slots];
//         setReelAnimations([true, true, true]);

//         for (let i = 0; i < 3; i++) {
//             setTimeout(() => {
//                 const newSymbol = slotItems[Math.floor(Math.random() * slotItems.length)];
//                 tempSlots[i] = newSymbol;
//                 setSlots([...tempSlots]);

//                 setReelAnimations(prev => {
//                     const newAnim = [...prev];
//                     newAnim[i] = false;
//                     return newAnim;
//                 });

//                 if (i === 2) {
//                     const win = checkWin(tempSlots);
//                     if (win > 0) {
//                         setWinAmount(win);
//                         setCredits(prev => prev + win);
//                         setTotalWins(prev => prev + win);
//                         setWinStreak(prev => prev + 1);
//                         setShowWinModal(true);

//                         if (win >= bet * 50) {
//                             setJackpotGlow(true);
//                         }
//                     } else {
//                         setWinStreak(0);
//                     }
//                     setTimeout(() => setSpinning(false), 600);
//                 }
//             }, reelStopTimes[i]);
//         }
//     };

//     const maxBet = () => setBet(Math.min(500, credits));
//     const addCredits = () => setCredits(prev => prev + 1000);

//     return (
//         <>
//             <Navbar />
//             <div className="min-h-screen relative overflow-hidden bg-black">
//                 {/* === ULTIMATE QUANTUM BACKGROUND === */}
//                 <div className="absolute inset-0">
//                     {/* Animated Particle Field */}
//                     <Particles />

//                     {/* Quantum Tunnel Effect */}
//                     <div className="absolute inset-0 bg-gradient-radial from-cyan-500/10 via-transparent to-transparent animate-spin-slow blur-3xl pointer-events-none"></div>

//                     {/* Flowing Wave Lines */}
//                     <div className="absolute inset-0 opacity-20 dark:opacity-40">
//                         {[...Array(10)].map((_, i) => (
//                             <div
//                                 key={i}
//                                 className="absolute left-0 right-0 h-px bg-gradient-to-r from-transparent via-cyan-400 to-transparent animate-pulse"
//                                 style={{
//                                     top: `${i * 10 + 5}%`,
//                                     animationDelay: `${i * 0.5}s`,
//                                     animationDuration: `${3 + i}s`,
//                                 }}
//                             ></div>
//                         ))}
//                     </div>

//                     {/* Corner Glows */}
//                     <div className="absolute top-20 left-20 w-64 h-64 bg-cyan-500/20 rounded-full blur-3xl animate-pulse"></div>
//                     <div className="absolute bottom-20 right-20 w-64 h-64 bg-purple-500/20 rounded-full blur-3xl animate-pulse delay-1000"></div>
//                 </div>

//                 {/* === MAIN UI === */}
//                 <div className="relative z-10 flex flex-col items-center justify-center p-6 pt-16 md:pt-24">
//                     <div className="w-full max-w-4xl">

//                         {/* === MEGA HEADER === */}
//                         <div className="text-center mb-10">
//                             <h1
//                                 className={`text-7xl md:text-9xl font-black bg-clip-text text-transparent mb-4 drop-shadow-2xl leading-none ${jackpotGlow ? 'animate-pulse' : ''
//                                     }`}
//                                 style={{
//                                     background: 'linear-gradient(90deg, #f97316, #ec4899, #8b5cf6, #06b6d4)',
//                                     WebkitBackgroundClip: 'text',
//                                     backgroundClip: 'text',
//                                 }}
//                             >
//                                 ROYAL SLOTS
//                             </h1>
//                             <p className="text-xl md:text-2xl text-gray-400 font-medium tracking-wide">
//                                 Spin. Win. Repeat.
//                             </p>
//                             {winStreak > 2 && (
//                                 <div className="mt-4 inline-flex items-center gap-2 px-6 py-3 bg-gradient-to-r from-yellow-500 to-red-500 text-white font-black text-lg rounded-full animate-bounce">
//                                     🔥 {winStreak}x WIN STREAK! 🔥
//                                 </div>
//                             )}
//                         </div>

//                         {/* === STATS DASHBOARD === */}
//                         <div className="grid grid-cols-3 gap-6 mb-10">
//                             <Stat label="Balance" value={credits} color="orange" icon="💰" />
//                             <Stat label="Current Bet" value={bet} color="green" icon="🎯" />
//                             <Stat label="Total Wins" value={totalWins} color="blue" icon="🏆" />
//                         </div>

//                         {/* === SLOT MACHINE === */}
//                         <div className="relative">
//                             <div className="backdrop-blur-xl bg-white/20 dark:bg-gray-900/40 rounded-3xl p-8 shadow-2xl border border-white/30 dark:border-gray-700/50 relative overflow-hidden transition-all hover:scale-105">
//                                 {/* Quantum Corners */}
//                                 <QuantumCorners />

//                                 {/* Display */}
//                                 <div className="bg-gradient-to-br from-gray-900 to-black dark:from-gray-950 dark:to-black rounded-3xl p-8 mb-8 relative overflow-hidden shadow-inner">
//                                     {spinning && (
//                                         <div className="absolute top-0 left-0 w-full h-1 bg-gradient-to-r from-transparent via-cyan-400 to-transparent animate-pulse"></div>
//                                     )}
//                                     <div className="grid grid-cols-3 gap-6">
//                                         {slots.map((item, index) => (
//                                             <div
//                                                 key={index}
//                                                 className={`
//                                                     w-28 h-28 md:w-32 md:h-32 relative rounded-2xl shadow-2xl border-2 flex items-center justify-center text-5xl
//                                                     transition-all duration-300
//                                                     ${reelAnimations[index]
//                                                         ? 'animate-spin-slower scale-95 opacity-80'
//                                                         : 'hover:scale-105 active:scale-100'
//                                                     }
//                                                     ${showWinModal ? 'border-yellow-400 shadow-yellow-400/50' : 'border-gray-700'}
//                                                 `}
//                                                 style={{
//                                                     background: 'linear-gradient(145deg, #1f2937, #111827)',
//                                                 }}
//                                             >
//                                                 <span className={reelAnimations[index] ? 'blur-sm' : 'drop-shadow-2xl'}>
//                                                     {item}
//                                                 </span>
//                                             </div>
//                                         ))}
//                                     </div>
//                                 </div>

//                                 {/* === WIN MODAL === */}
//                                 {showWinModal && (
//                                     <div className="absolute inset-0 flex items-center justify-center z-30 pointer-events-none">
//                                         <div className={`
//                                             bg-gradient-to-r from-yellow-400 via-orange-500 to-red-500 dark:from-yellow-300 dark:via-orange-400 dark:to-red-400
//                                             text-black font-black text-3xl md:text-5xl px-8 py-6 rounded-3xl shadow-2xl animate-pop
//                                             transform scale-110 opacity-100
//                                         `}>
//                                             {winAmount >= bet * 50 ? (
//                                                 <div className="text-center">
//                                                     🎰 <span className="block">MEGA JACKPOT!</span>
//                                                     <span className="text-2xl md:text-4xl">+{winAmount.toLocaleString()}</span>
//                                                 </div>
//                                             ) : winAmount >= bet * 20 ? (
//                                                 <div className="text-center">
//                                                     💎 <span className="block">BIG WIN!</span>
//                                                     <span className="text-2xl">+{winAmount.toLocaleString()}</span>
//                                                 </div>
//                                             ) : (
//                                                 <div className="text-center">
//                                                     🎉 <span className="block">WIN +{winAmount}</span>
//                                                 </div>
//                                             )}
//                                         </div>
//                                     </div>
//                                 )}

//                                 {/* Controls */}
//                                 <div className="grid grid-cols-2 gap-6 mb-6">
//                                     <div>
//                                         <label className="block text-gray-200 font-bold mb-3">Bet Amount</label>
//                                         <select
//                                             value={bet}
//                                             onChange={(e) => setBet(Number(e.target.value))}
//                                             disabled={spinning}
//                                             className="w-full bg-white/70 dark:bg-gray-800/70 text-gray-900 dark:text-white rounded-xl p-4 border border-gray-300 focus:border-orange-500 focus:ring-2 focus:ring-orange-300 outline-none transition"
//                                         >
//                                             {[10, 25, 50, 100, 250, 500].map(b => (
//                                                 <option key={b} value={b}>{b} Credits</option>
//                                             ))}
//                                         </select>
//                                     </div>
//                                     <div className="flex flex-col gap-3 justify-end">
//                                         <button
//                                             onClick={maxBet}
//                                             disabled={spinning}
//                                             className="w-full bg-gradient-to-r from-blue-600 to-purple-700 text-white font-bold py-3 rounded-xl hover:scale-105 active:scale-95 transition"
//                                         >
//                                             MAX BET
//                                         </button>
//                                         <button
//                                             onClick={addCredits}
//                                             className="w-full bg-gradient-to-r from-green-600 to-emerald-600 text-white font-bold py-3 rounded-xl hover:scale-105 active:scale-95 transition"
//                                         >
//                                             +1,000 CREDITS
//                                         </button>
//                                     </div>
//                                 </div>

//                                 {/* Spin Button */}
//                                 <button
//                                     onClick={spin}
//                                     disabled={spinning || credits < bet}
//                                     className={`
//                                         w-full py-6 rounded-2xl font-black text-2xl relative overflow-hidden group
//                                         transition-all duration-300
//                                         ${spinning || credits < bet
//                                             ? 'bg-gray-500 cursor-not-allowed'
//                                             : 'bg-gradient-to-r from-orange-500 via-red-500 to-pink-500 hover:scale-105 active:scale-95 shadow-xl hover:shadow-3xl'
//                                         }
//                                     `}
//                                 >
//                                     <span className="relative z-10 flex items-center justify-center gap-3 text-white">
//                                         {spinning ? (
//                                             <>
//                                                 <div className="w-8 h-8 border-3 border-white border-t-transparent rounded-full animate-spin"></div>
//                                                 <span>SPINNING...</span>
//                                             </>
//                                         ) : credits < bet ? (
//                                             <span>INSUFFICIENT CREDITS</span>
//                                         ) : (
//                                             <>
//                                                 <span className="text-3xl">🎰</span>
//                                                 <span>SPIN TO WIN</span>
//                                                 <span className="text-3xl">💫</span>
//                                             </>
//                                         )}
//                                     </span>
//                                 </button>
//                             </div>
//                         </div>

//                         {/* === PAYTABLE === */}
//                         <PayTable bet={bet} />

//                     </div>
//                 </div>
//             </div>

//             {/* === GLOBAL ANIMATIONS === */}
//             <style jsx global>{`
//                 @keyframes pop {
//                     0% { transform: scale(0.8); opacity: 0; }
//                     50% { transform: scale(1.1); opacity: 1; }
//                     100% { transform: scale(1); opacity: 1; }
//                 }
//                 .animate-pop {
//                     animation: pop 0.6s ease-out;
//                 }

//                 @keyframes spin-slow {
//                     from { transform: rotate(0deg); }
//                     to { transform: rotate(360deg); }
//                 }
//                 .animate-spin-slow { animation: spin-slow 20s linear infinite; }

//                 .animate-spin-slower {
//                     animation: spin 2s linear infinite;
//                 }

//                 .bg-gradient-radial {
//                     background: radial-gradient(circle, var(--tw-gradient-stops));
//                 }
//             `}</style>
//         </>
//     );
// }

// // === REUSABLE COMPONENTS ===

// function Particles() {
//     return (
//         <div className="absolute inset-0">
//             {[...Array(30)].map((_, i) => {
//                 const size = Math.random() * 4 + 1;
//                 const delay = Math.random() * 3;
//                 const duration = Math.random() * 3 + 2;
//                 return (
//                     <div
//                         key={i}
//                         className="absolute rounded-full animate-pulse opacity-60"
//                         style={{
//                             left: `${Math.random() * 100}%`,
//                             top: `${Math.random() * 100}%`,
//                             width: `${size}px`,
//                             height: `${size}px`,
//                             background: ['cyan', 'magenta', 'yellow', 'orange'][Math.floor(Math.random() * 4)],
//                             animationDelay: `${delay}s`,
//                             animationDuration: `${duration}s`,
//                         }}
//                     ></div>
//                 );
//             })}
//         </div>
//     );
// }

// function QuantumCorners() {
//     return (
//         <>
//             <div className="absolute top-3 left-3 w-5 h-5 rounded-full bg-gradient-to-br from-orange-400 to-red-500 shadow-lg shadow-orange-400/50"></div>
//             <div className="absolute top-3 right-3 w-5 h-5 rounded-full bg-gradient-to-br from-pink-400 to-purple-500 shadow-lg shadow-pink-400/50"></div>
//             <div className="absolute bottom-3 left-3 w-5 h-5 rounded-full bg-gradient-to-br from-blue-400 to-cyan-500 shadow-lg shadow-blue-400/50"></div>
//             <div className="absolute bottom-3 right-3 w-5 h-5 rounded-full bg-gradient-to-br from-green-400 to-emerald-500 shadow-lg shadow-green-400/50"></div>
//         </>
//     );
// }

// function Stat({ label, value, color, icon }: { label: string; value: number; color: string; icon: string }) {
//     const colorClasses = {
//         orange: 'from-orange-500 to-red-500',
//         green: 'from-green-500 to-emerald-500',
//         blue: 'from-blue-500 to-cyan-500',
//     };

//     return (
//         <div className="bg-white/20 dark:bg-gray-800/30 backdrop-blur-md rounded-2xl p-5 border border-white/20 shadow-lg hover:shadow-2xl transition-all">
//             <div className="flex items-center gap-2 mb-2">
//                 <span className="text-lg">{icon}</span>
//                 <p className="text-sm text-gray-300 font-semibold">{label}</p>
//             </div>
//             <p className={`text-3xl font-black bg-gradient-to-r ${colorClasses[color as keyof typeof colorClasses]} bg-clip-text text-transparent`}>
//                 {value.toLocaleString()}
//             </p>
//         </div>
//     );
// }

// function PayTable({ bet }: { bet: number }) {
//     return (
//         <div className="mt-10 bg-white/20 dark:bg-gray-800/30 backdrop-blur-md rounded-2xl p-6 border border-white/20 shadow-lg">
//             <h3 className="text-2xl font-black text-center text-white mb-5">💎 PREMIUM PAYTABLE</h3>
//             <div className="space-y-3">
//                 {[
//                     { sym: '💎💎💎', mult: 150, label: 'Diamond Jackpot', color: 'from-yellow-400 to-orange-500' },
//                     { sym: '⭐⭐⭐', mult: 100, label: 'Star Burst', color: 'from-purple-400 to-pink-500' },
//                     { sym: '🔔🔔🔔', mult: 75, label: 'Bell Ring', color: 'from-blue-400 to-cyan-500' },
//                     { sym: '🍒🍒🍒', mult: 50, label: 'Cherry Bomb', color: 'from-red-400 to-pink-500' },
//                 ].map((row, i) => (
//                     <div key={i} className="flex justify-between items-center p-3 rounded-xl bg-gradient-to-r to-transparent border border-gray-600/50">
//                         <div>
//                             <span className="font-black text-xl">{row.sym}</span>
//                             <span className="text-sm text-gray-300 ml-3">{row.label}</span>
//                         </div>
//                         <span className={`font-black text-lg bg-gradient-to-r ${row.color} bg-clip-text text-transparent`}>
//                             {Intl.NumberFormat().format(bet * row.mult)}
//                         </span>
//                     </div>
//                 ))}
//             </div>
//         </div>
//     );
// }