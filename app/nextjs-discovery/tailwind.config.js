/** @type {import('tailwindcss').Config} */

module.exports = {
  content: ['./src/**/*.{html,js}'],
  theme: {
    extend: {
      colors: {
        color: {
          yellow: '#FFFCA0',
          violet: ' #7C6AE2',
          light_grey: '#EBEBEB',
          dark_grey: '#676363',
        },
      },
      fontFamily: {
        kaisei_tokumin: ['var(--font-kaisei_tokumin)', 'sans-serif'],
        space_mono: ['var(--font-space_mono)', 'sans-serif'],
        inter: ['var(--font-inter)', 'sans-serif'],
      },
    },
  },
  plugins: [],
};
