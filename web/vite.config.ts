import {defineConfig} from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig({
    plugins: [react()],
    server: {
        port: 9000,
        strictPort: true,
        open: true,
        proxy: {
            "^/api/": "http://localhost:9090/"
        }
    }
});
