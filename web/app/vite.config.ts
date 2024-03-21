import type { Config as DaisyUIConfig } from 'daisyui'
import type { Config as TailwindConfig } from 'tailwindcss';
import type { BuildOptions, CSSOptions, PluginOption } from 'vite';

import autoprefixer from 'autoprefixer';
import cssnano from 'cssnano';
import daisyui from 'daisyui';
import tailwindcss from 'tailwindcss';
import { defineConfig } from 'vite';
import { svelte, vitePreprocess } from '@sveltejs/vite-plugin-svelte';

const build: BuildOptions = {
    assetsDir: '',
    lib: {
        entry: 'src/index.ts',
        formats: ['umd', 'es'],
        name: 'Melodeon',
        fileName: 'melodeon',
    },
}

const daisyUIConfig: DaisyUIConfig = {
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

const tailwindConfig: TailwindConfig = {
    content: [
        '../templates/*.gotmpl',
        './src/**/*.{ts,svelte}',
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
    daisyui: daisyUIConfig,
};

const css: CSSOptions = {
    postcss: {
        plugins: [
            autoprefixer(),
            cssnano(),
            tailwindcss(tailwindConfig),
        ],
    }
}

const plugins: PluginOption[] = [
    svelte({
        preprocess: vitePreprocess(),
    }),
];

export default defineConfig({ plugins, build, css });
