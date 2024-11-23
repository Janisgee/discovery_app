import path from "node:path";
import { fileURLToPath } from "node:url";
import js from "@eslint/js";
import { FlatCompat } from "@eslint/eslintrc";
import prettierPlugin from "eslint-plugin-prettier/recommended";
import tailwindCSS from "eslint-plugin-tailwindcss";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);
const compat = new FlatCompat({
  baseDirectory: __dirname,
  recommendedConfig: js.configs.recommended,
  allConfig: js.configs.all,
});

const config = [
  ...compat.extends("next/core-web-vitals"),
  ...tailwindCSS.configs["flat/recommended"],
  {
    rules: {
      "tailwindcss/no-custom-classname": "off",
    },
  },
  prettierPlugin,
  { ignores: [".next/*"] },
];

export default config;
