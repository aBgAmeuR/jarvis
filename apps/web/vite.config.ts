import tailwindcss from "@tailwindcss/vite";
import { tanstackStart } from "@tanstack/react-start/plugin/vite";
import react from "@vitejs/plugin-react";
import { defineConfig } from "vite";

export default defineConfig({
  resolve: { tsconfigPaths: true },
  plugins: [tanstackStart({ spa: { enabled: true } }), react(), tailwindcss()],
  preview: { host: "127.0.0.1" },
  server: {
    proxy: {
      "/api": {
        target: "http://127.0.0.1:8080",
        headers: { "Remote-User": "dev", "Remote-Groups": "admin" },
      },
    },
  },
});
