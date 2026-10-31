import { defineConfig } from "oxlint";

import base from "./base.ts";

export default defineConfig({
  extends: [base],
  ignorePatterns: base.ignorePatterns,
  plugins: ["typescript", "unicorn", "oxc", "import", "react", "jsx-a11y", "react-perf"],
  rules: {
    "react/rules-of-hooks": "error",
    "react/exhaustive-deps": "error",
    "react/jsx-key": "error",
    "react/react-in-jsx-scope": "off",
    "react-perf/jsx-no-new-object-as-prop": "off",
    "react-perf/jsx-no-new-array-as-prop": "off",
    "react-perf/jsx-no-new-function-as-prop": "off",
  },
});
