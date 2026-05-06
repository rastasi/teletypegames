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
        target: 'https://wiki.teletypegames.org',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/proxy\/wiki/, '')
      },
      '/_assets': {
        target: 'https://wiki.teletypegames.org',
        changeOrigin: true,
      },
      '/_error': {
        target: 'https://wiki.teletypegames.org',
        changeOrigin: true,
      },
      '/proxy/git': {
        target: 'https://git.teletypegames.org',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/proxy\/git/, '')
      }
    }
  }
})
