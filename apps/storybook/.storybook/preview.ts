import { withThemeByClassName } from "@storybook/addon-themes";
import type { Preview } from "@storybook/react-vite";

import "./preview.css";

export default {
  tags: ["autodocs"],
  decorators: [
    withThemeByClassName({
      themes: { light: "", dark: "dark" },
      defaultTheme: "light",
      parentSelector: "html",
    }),
  ],
  parameters: {
    layout: "centered",
    controls: {
      expanded: true,
      sort: "requiredFirst",
      matchers: { color: /(background|color)$/i, date: /Date$/i },
    },
    options: {
      storySort: {
        order: [
          "Introduction",
          "Foundations",
          ["Colors", "Typography"],
          "Components",
          ["Button", "Badge", "Input", "Select", "Tabs", "Card", "Data", "Table", "List", "*"],
          "Charts",
        ],
      },
    },
    docs: {
      toc: { headingSelector: "h2, h3", title: "On this page" },
      story: { inline: true },
    },
    a11y: { test: "error" },
  },
} satisfies Preview;
