import { defineConfig } from 'vite'
const { resolve } = require('path')
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
// export default defineConfig({
//   build :{
//     assetsDir:""
//   },
//   plugins: [vue()],
//   server:{port:5000,proxy:{
//     '/api':{
//         target: 'http://localhost:3000',
//         changeOrigin: true

//     },
//   },}
  
  
// })

export default ({ command, mode }) => {
  if (command === 'serve') {
    return {
      // serve specific config
      plugins: [vue()],
      server:{port:5000,proxy:{
            '/api':{
                target: 'http://localhost:3000',
                changeOrigin: true
        
            },
          },}
    }
  } else {
    return {
      // build specific config
      // base:"/cp/",
      build:{
        assetsDir:"",
        rollupOptions: {
          input: {
            main: resolve(__dirname, 'index.html'),
            nested: resolve(__dirname, 'cp/index.html')
          },
        }
      },
      plugins: [vue()],
    }
  }
}