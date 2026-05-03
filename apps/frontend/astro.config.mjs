// @ts-check
import { defineConfig } from 'astro/config';
import tailwind from '@astrojs/tailwind';

export default defineConfig({
  devToolbar: {
    enabled: false
  },
  integrations: [tailwind()],
  server: {
    host: true,
    allowedHosts: ['teletypegames.org']
  }
});
