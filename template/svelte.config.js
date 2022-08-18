import preprocess from "svelte-preprocess";
import adapter from "@sveltejs/adapter-static";

/** @type {import('@sveltejs/kit').Config} */
const config = {
  kit: {
    adapter: adapter({
      pages: "dist",
      assets: "dist",
      fallback: 'index.html',
      precompress: false,
    }),
    prerender: {
      default: true,
    },
    appDir: "app",
    outDir: ".kit",
    files: {
      assets: "src/assets",
    }
  },
  preprocess: [
    preprocess({
      postcss: true,
    }),
  ],
};

export default config;