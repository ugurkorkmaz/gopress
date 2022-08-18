const config = {
  jit: true,
  content: ["./src/**/*.{html,js,svelte,css}"],
  theme: {
    extend: {},
  },
  plugins: [],
  experimental: {
    applyComplexClasses: true,
  },
};

module.exports = config;
