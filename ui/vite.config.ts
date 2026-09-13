import path from "node:path";
import react from "@vitejs/plugin-react";
import { tanstackRouter } from "@tanstack/router-plugin/vite";
import tailwindcss from "@tailwindcss/vite";
import { defineConfig } from "vite";
import { compression } from "vite-plugin-compression2";

export default defineConfig({
  plugins: [
    tanstackRouter(),
    react(),
    tailwindcss(),
    compression({ algorithms: ["brotliCompress", "gzip"], threshold: 1024 }),
  ],
  resolve: {
    alias: {
      "@": path.resolve(import.meta.dirname, "./src"),
    },
  },
  server: {
    proxy: {
      "/api": {
        target: "http://localhost:8080",
        changeOrigin: true,
      },
    },
  },
});
