export type Theme = "light" | "dark";

const storageKey = "theme";

/** Runs in `<head>` before first paint to avoid a theme flash. */
export const themeScript = `try{var t=localStorage.getItem("${storageKey}");var d=t?t==="dark":matchMedia("(prefers-color-scheme: dark)").matches;var r=document.documentElement;r.classList.toggle("dark",d);r.style.colorScheme=d?"dark":"light"}catch(e){}`;

export function readTheme(): Theme {
  return document.documentElement.classList.contains("dark") ? "dark" : "light";
}

export function applyTheme(theme: Theme) {
  const root = document.documentElement;
  root.classList.toggle("dark", theme === "dark");
  root.style.colorScheme = theme;
  try {
    localStorage.setItem(storageKey, theme);
  } catch {
    // Storage can be blocked; the choice then lasts until reload.
  }
}
