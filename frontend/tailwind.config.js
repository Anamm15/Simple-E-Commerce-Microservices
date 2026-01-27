/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        "primary-1": "#4A90E2",
        "primary-2": "#5CD65C",
        "secondary-1": "#2C3E50",
        "secondary-2": "#A0EACD",
        "tertary-1": "#B3D9FF",
        "tertary-2": "#B8F2D8",
      },
    },
  },
  plugins: [],
}