import type { StorybookConfig } from "@storybook/react-vite";

export default {
  framework: "@storybook/react-vite",
  stories: ["../src/**/*.mdx", "../src/**/*.stories.tsx"],
  addons: ["@storybook/addon-docs", "@storybook/addon-a11y", "@storybook/addon-themes"],
  typescript: { reactDocgen: "react-docgen" },
} satisfies StorybookConfig;
