import type { StorybookConfig } from "@storybook/react-vite";

export default {
  framework: "@storybook/react-vite",
  stories: ["../src/**/*.mdx", "../src/**/*.stories.tsx"],
  addons: [
    "@storybook/addon-docs",
    "@storybook/addon-a11y",
    "@storybook/addon-themes",
    "@storybook/addon-vitest",
  ],
  typescript: { reactDocgen: "react-docgen" },
  viteFinal: (config) => ({
    ...config,
    build: {
      ...config.build,
      rolldownOptions: {
        ...config.build?.rolldownOptions,
        onLog(level, log, defaultHandler) {
          if (log.code === "MODULE_LEVEL_DIRECTIVE") return;
          defaultHandler(level, log);
        },
      },
    },
  }),
} satisfies StorybookConfig;
