import { addIconSelectors } from '@iconify/tailwind';
import daisyui from 'daisyui';

/** @type {import('tailwindcss').Config} */
export default {
    content: {
        files: ["../**/*.cshtml", "./src/**/*.{js,ts,svelte}"],
        relative: false,
    },
    theme: {
        extend: {
            fontSize: {
                sm: '0.750rem',
                base: '1rem',
                xl: '1.333rem',
                '2xl': '1.777rem',
                '3xl': '2.369rem',
                '4xl': '3.158rem',
                '5xl': '4.210rem',
            },
            fontFamily: {
                display: 'Rubik Doodle Shadow',
                sans: 'Rubik',
            },
            fontWeight: {
                normal: '400',
                bold: '700',
            },
        },
    },
    plugins: [
        addIconSelectors(['tabler']),
        daisyui,
    ],
    daisyui: {
        themes: [
            {
                light: {
                    "primary": "#1B66A1",
                    "secondary": "#B7A495",
                    "accent": "#FF9E73",
                    "neutral": "#172135",
                    "base-100": "#DFE5F1",
                },
                dark: {
                    "primary": "#5eaae4",
                    "secondary": "#6a5748",
                    "accent": "#8a2900",
                    "neutral": "#172135",
                    "base-100": "#0e1420",
                },
            },
        ],
    },
}

