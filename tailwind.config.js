/** @type {import('tailwindcss').Config} */
const plugin1 = require('@tailwindcss/line-clamp');
const plugin2 = require('@tailwindcss/typography');
module.exports = {
  content: ["./views/**/*.{html,js,django}"],
  theme: {
    extend: {},
  },
  plugins: [
    plugin1,
    plugin2,
  ],
}

