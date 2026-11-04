import { createRootRoute, HeadContent, Outlet, Scripts } from "@tanstack/react-router";
import type { PropsWithChildren } from "react";

import { AppShell } from "@/app/shell/app-shell.tsx";
import { themeScript } from "@/app/theme/theme.ts";

import appCss from "../styles.css?url";

export const Route = createRootRoute({
  head: () => ({
    meta: [
      { charSet: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      { title: "Jarvis" },
    ],
    scripts: [{ children: themeScript }],
    links: [
      { rel: "icon", type: "image/svg+xml", href: "/favicon.svg" },
      { rel: "stylesheet", href: appCss },
    ],
  }),
  component: Root,
});

function Root() {
  return (
    <Document>
      <AppShell>
        <Outlet />
      </AppShell>
    </Document>
  );
}

function Document({ children }: PropsWithChildren) {
  return (
    <html lang="en" suppressHydrationWarning>
      <head>
        <HeadContent />
      </head>
      <body>
        {children}
        <Scripts />
      </body>
    </html>
  );
}
