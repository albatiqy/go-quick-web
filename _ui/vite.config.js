import { resolve } from 'path'
import { defineConfig } from 'vite'

export default defineConfig({
    build: {
        outDir: '../_embed/ui',
        emptyOutDir: true,
        target: 'esnext',
        rollupOptions: {
            input: {
              main: resolve(__dirname, 'index.html'),
              nested: resolve(__dirname, 'blank.html')
            }
          }
    }
})