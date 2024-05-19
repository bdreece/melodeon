import nodeResolve from '@rollup/plugin-node-resolve';
import typescript from '@rollup/plugin-typescript';
import terser from '@rollup/plugin-terser';
import postcss from 'rollup-plugin-postcss';
import autoprefixer from 'autoprefixer';
import cssnano from 'cssnano';

/** @type {import('rollup').RollupOptions} */
export default {
    input: 'src/index.js',
    output: {
        name: 'Melodeon',
        file: 'dist/melodeon.umd.cjs',
        format: 'umd',
    },
    plugins: [
        nodeResolve(),
        typescript(),
        terser(),
        postcss({
            extract: 'style.css',
            plugins: [autoprefixer(), cssnano()],
        }),
    ],
};
