import react from "@jarvis/oxlint-config/react";
import { defineConfig } from "oxlint";

const above = (...layers: string[]) => [
  "error",
  {
    patterns: [
      {
        group: layers.map((layer) => `@/${layer}/**`),
        message: "A layer imports only from the layers below it.",
      },
    ],
  },
];

export default defineConfig({
  extends: [react],
  overrides: [
    {
      files: ["src/shared/**"],
      rules: { "no-restricted-imports": above("app", "pages", "entities", "routes") },
    },
    {
      files: ["src/app/**"],
      rules: { "no-restricted-imports": above("pages", "entities", "routes") },
    },
    {
      files: ["src/pages/**"],
      rules: { "no-restricted-imports": above("app", "pages", "routes") },
    },
  ],
});
