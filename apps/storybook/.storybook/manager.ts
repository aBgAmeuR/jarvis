import { addons } from "storybook/manager-api";

import { theme } from "./theme.ts";

addons.setConfig({
  theme,
  sidebar: { showRoots: true },
  toolbar: { zoom: { hidden: false }, eject: { hidden: true } },
});
