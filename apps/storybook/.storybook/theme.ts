import { create } from "storybook/theming";

const base = {
  brandTitle: "Jarvis UI",
  fontBase: '"Geist Variable", ui-sans-serif, system-ui, sans-serif',
  fontCode: '"Geist Mono Variable", ui-monospace, monospace',
  barSelectedColor: "#171717",
} as const;

export const theme = create({
  ...base,
  base: "light",
  colorPrimary: "#171717",
  colorSecondary: "#1c64d7",
  appBg: "#fafafa",
  appContentBg: "#ffffff",
  appPreviewBg: "#ffffff",
  appBorderColor: "#e5e5e5",
  appBorderRadius: 8,
  textColor: "#171717",
  textMutedColor: "#666666",
  barBg: "#ffffff",
  barTextColor: "#666666",
  inputBorder: "#e5e5e5",
  inputBorderRadius: 6,
});
