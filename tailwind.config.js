/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./layouts/**/*.html", "./pages/**/*.html"],
  theme: {
    colors: {
      "white": "#f8eee1",
      "darkgreen": "#275657",
      "brown": {
        100: "#ECE1CD",
        200: "#42362B",
        500: "#AA7338",
        900: "#C99654",
      }
    },
    fontFamily: {
      display: ['"Playwrite HU"', "cursive"],
      body: ['"League Spartan"', "sans-serif"]
    }
  },
  plugins: [],
}


