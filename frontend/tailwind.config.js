/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      backgroundImage: {
        "space-infy": "url('./src/img/infy-starry-bg.svg')",
      },
    },
  },
  plugins: [],
};
