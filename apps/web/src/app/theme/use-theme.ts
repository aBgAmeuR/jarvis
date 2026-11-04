import { useCallback, useSyncExternalStore } from "react";

import { applyTheme, readTheme, type Theme } from "./theme.ts";

function subscribe(onChange: () => void) {
  const observer = new MutationObserver(onChange);
  observer.observe(document.documentElement, { attributes: true, attributeFilter: ["class"] });
  return () => observer.disconnect();
}

export function useTheme(): { theme: Theme; toggle: () => void } {
  const theme = useSyncExternalStore<Theme>(subscribe, readTheme, () => "light");
  const toggle = useCallback(() => applyTheme(readTheme() === "dark" ? "light" : "dark"), []);
  return { theme, toggle };
}
