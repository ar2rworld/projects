/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  darkMode: 'class',
  theme: {
    colors: {
      black: "#000",
      white: "#fff",
      bg: {
        1: 'hsl(0, 0%, 100%)',
        2: 'hsl(206, 20%, 90%)',
        3: 'hsl(206, 20%, 80%)',
      },
      fg: {
        1: 'hsl(0, 0%, 13%)',
        2: 'hsl(0, 0%, 20%)',
        3: 'hsl(0, 0%, 30%)',
      },
      link: 'hsl(208, 77%, 47%)',
      hover: 'hsl(208, 77%, 55%)',
      active: 'hsl(208, 77%, 40%)',
      d: {
        bg: {
          0: 'rgb(16, 20, 30)',
          1: 'hsl(0, 0%, 18%)',
          2: 'hsl(0, 0%, 30%)',
          3: 'hsl(0, 0%, 40%)',
        },
        fg: {
          1: 'hsl(0, 0%, 90%)',
          2: 'hsl(0, 0%, 70%)',
          3: 'hsl(0, 0%, 60%)',
        },
        link: 'hsl(206, 96%, 72%)',
        hover: 'hsl(206, 96%, 78%)',
        active: 'hsl(206, 96%, 64%)',
      },
    },
    fontFamily: {
      sans: ['Courier New', 'Courier', 'monospace'],
      serif: ['Merriweather', 'serif'],
      mono: ['ui-monospace', 'Cascadia Code', 'Source Code Pro', 'Menlo', 'Consolas', 'DejaVu Sans Mono', 'monospace']
    }
  },
  plugins: [],
}

