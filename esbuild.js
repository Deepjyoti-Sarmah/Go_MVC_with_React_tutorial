const esbuild = require("esbuild");

esbuild
  .build({
    entryPoints: ["frontend/Application.tsx", "frontend/style.scss"],
    outdir: "public/assets",
    bundle: true,
    watch: true,
    minify: true,
    plugins: [sassPlugin()],
  })
  .then(() => console.log("⚡ Build complete! ⚡"))
  .catch(() => process.exit(1));