/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./web/layouts/**/*.html", "./web/pages/**/*.html"],
  theme: {
    colors: {
      "white": "#f8eee1",
      "yellow": "#ffbe0b",
      "orange": "#fb5607",
      "pink": "#ff006e",
      "purple": "#8338ec",
      "blue": "#3a86ff",
    },
    fontFamily: {
      body: ['"League Spartan"', "sans-serif"]
    }
  },
  plugins: [],
}


