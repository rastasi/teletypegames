import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  envPrefix: ['VITE_', 'WEBAPP_', 'YOUTUBE_', 'DISCORD_'],
  server: {
    host: true,
    allowedHosts: ['teletypegames.org'],
    proxy: {
      '/proxy/wiki': {
        target: 'https://wiki.teletype.hu',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/proxy\/wiki/, '')
      },
      '/proxy/git': {
        target: 'https://git.teletype.hu',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/proxy\/git/, '')
      }
    }
  }
})
