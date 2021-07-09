import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import styleImport from 'vite-plugin-style-import'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        vue(),
        styleImport({
            libs: [
                {
                    libraryName: 'element-plus',
                    esModule: true,
                    ensureStyleFile: true,
                    resolveStyle: (name) => {
                        name = name.slice(3)
                        return `element-plus/packages/theme-chalk/src/${name}.scss`;
                    },
                    resolveComponent: (name) => {
                        return `element-plus/lib/${name}`;
                    },
                },
                // {
                //     libraryName: 'element-plus',
                //     esModule: true,
                //     ensureStyleFile: true,
                //     resolveStyle: (name) => {
                //         return `element-plus/lib/theme-chalk/${name}.css`;
                //     },
                //     resolveComponent: (name) => {
                //         return `element-plus/lib/${name}`;
                //     },
                // }
            ]
        })
    ],
    resolve: {
        // alias: [
        //     {find: 'src', replacement: path.resolve(__dirname, '.')},
        //     {find: '@components', replacement: path.resolve(__dirname, './src/components')}
        // ]
        alias: {
            'src': path.resolve(__dirname, './src'),
            '@components': path.resolve(__dirname, './src/components')
        }
    },
    server: {
        port: 3000,
    },
    optimizeDeps: {
        include: [],
        exclude: []
    },
})
