import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  build :{
    assetsDir:""
  },
  plugins: [vue()],
  server:{port:5000,proxy:{
    '/api':{
        target: 'http://localhost:3000',
        changeOrigin: true

    },
  },}
  
})
