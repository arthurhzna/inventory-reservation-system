import { defineConfig, loadEnv } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '')

  return {
    plugins: [react()],
    define: {
      __URL_API_INVENTORY__: JSON.stringify(env.URL_API_INVENTORY ?? '8000'),
    },
  }
})
