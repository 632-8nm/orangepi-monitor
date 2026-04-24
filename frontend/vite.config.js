import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import { resolve } from "node:path";

export default defineConfig({
  plugins: [vue()],
  build: {
    outDir: "../static/dist",
    emptyOutDir: true,
    lib: {
      entry: resolve(__dirname, "src/main.js"),
      name: "OrangePiMonitorApp",
      formats: ["iife"],
      fileName: () => "app.js",
    },
    rollupOptions: {
      output: {
        assetFileNames: "app.css",
      },
    },
  },
});
