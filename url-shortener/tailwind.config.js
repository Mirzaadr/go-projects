/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./static/templates/*.html"],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],
}