import { Icon, Moon02Icon, Sun03Icon } from "@jarvis/icons";
import { Button, cn, Tooltip } from "@jarvis/ui";

import { useTheme } from "./use-theme.ts";

const iconMotion =
  "absolute inset-0 transition-[opacity,rotate,scale] duration-200 ease-out motion-reduce:transition-none";

export function ThemeToggle() {
  const { theme, toggle } = useTheme();
  const dark = theme === "dark";
  const label = dark ? "Switch to light theme" : "Switch to dark theme";

  return (
    <Tooltip content={label} side="bottom">
      <Button variant="ghost" size="icon" aria-label={label} onClick={toggle}>
        <span className="relative size-3.5">
          <Icon
            icon={Sun03Icon}
            className={cn(
              iconMotion,
              dark ? "scale-100 rotate-0 opacity-100" : "scale-50 -rotate-90 opacity-0",
            )}
          />
          <Icon
            icon={Moon02Icon}
            className={cn(
              iconMotion,
              dark ? "scale-50 rotate-90 opacity-0" : "scale-100 rotate-0 opacity-100",
            )}
          />
        </span>
      </Button>
    </Tooltip>
  );
}
