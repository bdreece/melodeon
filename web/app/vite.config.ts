import autoprefixer from 'autoprefixer';
import cssnano from 'cssnano';
import daisyui from 'daisyui';
import tailwindcss from 'tailwindcss';

import { defineConfig } from 'vite';

export default defineConfig({
    build: {
        assetsDir: '',
        lib: {
            entry: 'src/index.ts',
            formats: ['umd', 'es'],
            name: 'Melodeon',
            fileName: 'melodeon',
        },
    },
    css: {
        postcss: {
            plugins: [
                autoprefixer(),
                cssnano(),
                tailwindcss({
                    content: [
                        '../templates/*.gotmpl',
                        './src/**/*.ts',
                    ],
                    theme: {
                        extend: {
                            fontFamily: {
                                sans: [
                                    '"Rubik"',
                                    'system-ui',
                                    'Helvetica',
                                    'sans-serif',
                                ],
                                display: [
                                    '"Pacifico"',
                                    'cursive',
                                ],
                            }
                        },
                    },
                    plugins: [daisyui],
                    daisyui: {
                        themes: [
                            {
                                custom: {
                                    'primary': '#773344',
                                    'secondary': '#E3B5A4',
                                    'accent': '#D44D5C',
                                    'neutral': '#0B0014',
                                    'base-100': '#F5E9E2'
                                }
                            },
                        ]
                    }
                }),
            ],
        },
    },
});
