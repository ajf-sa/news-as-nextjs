import { defineConfig } from 'vite'
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
      base:"/cp/",
      build:{
        assetsDir:""
      },
      plugins: [vue()],
    }
  }
}