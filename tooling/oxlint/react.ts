import { defineConfig } from "oxlint";

import base from "./base.ts";

export default defineConfig({
  extends: [base],
  ignorePatterns: base.ignorePatterns,
  jsPlugins: ["@shadcn/lint"],
  plugins: ["typescript", "unicorn", "oxc", "import", "react", "jsx-a11y", "react-perf"],
  settings: {
    shadcn: { themeImports: ["^@jarvis/"] },
  },
  rules: {
    "react/rules-of-hooks": "error",
    "react/exhaustive-deps": "error",
    "react/jsx-key": "error",
    "react/react-in-jsx-scope": "off",
    "shadcn/no-raw-colors": "error",
    "shadcn/no-unknown-classes": ["error", { allow: ["animate-*", "*:animate-*"] }],
    "shadcn/no-inline-styles": "error",
    "react-perf/jsx-no-new-object-as-prop": "off",
    "react-perf/jsx-no-new-array-as-prop": "off",
    "react-perf/jsx-no-new-function-as-prop": "off",
  },
});
