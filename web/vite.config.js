import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

// Prefix-aware assets: the platform injects NEXT_PUBLIC_BASE_PATH (fold)
// as the component's full path; Vite bakes it as the build base. Absolute
// '/' would break asset URLs under /{project}/{component} routing.
export default defineConfig({
  plugins: [react()],
  base: process.env.NEXT_PUBLIC_BASE_PATH || '/',
});
