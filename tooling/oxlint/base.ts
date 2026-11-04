import { defineConfig } from "oxlint";

export default defineConfig({
  plugins: ["typescript", "unicorn", "oxc", "import"],
  categories: {
    correctness: "error",
    suspicious: "warn",
  },
  rules: {
    "typescript/no-explicit-any": "error",
    "typescript/no-non-null-assertion": "error",
    "typescript/ban-ts-comment": "error",
    "typescript/consistent-type-imports": "error",
    "import/no-cycle": "error",
    "import/no-unassigned-import": ["warn", { allow: ["**/*.css"] }],
    "no-console": "error",
    // Size and complexity.
    "max-lines": ["error", { max: 200, skipBlankLines: true, skipComments: true }],
    "max-lines-per-function": ["error", { max: 80, skipBlankLines: true, skipComments: true }],
    complexity: ["error", 15],
    "max-depth": ["error", 4],
    "max-params": ["error", 4],
  },
  ignorePatterns: ["dist", "storybook-static", "**/gen/**", "routeTree.gen.ts"],
});
