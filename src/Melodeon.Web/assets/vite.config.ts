import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'

// https://vitejs.dev/config/
export default defineConfig({
    publicDir: '../wwwroot',
    build: {
        copyPublicDir: false,
        emptyOutDir: true,
        assetsDir: '',
        lib: {
            entry: 'src/index.ts',
            name: 'melodeon',
            fileName: 'index',
            formats: ['es'],
        },
        rollupOptions: {
            external: [
                'htmx.org',
                'htmx-ext-head-support',
                'htmx-ext-sse',
            ],
            output: {
                paths: {
                    'htmx.org': 'https://unpkg.com/htmx.org@1.9.12/dist/htmx.min.js',
                    'htmx-ext-head-support': 'https://unpkg.com/htmx.org@1.9.12/dist/ext/head-support.js',
                    'htmx-ext-sse': 'https://unpkg.com/htmx.org@1.9.12/dist/ext/sse.js',
                }
            }
        }
    },
    plugins: [svelte()],
})
