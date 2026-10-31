import { createRootRoute, HeadContent, Outlet, Scripts } from "@tanstack/react-router";
import type { PropsWithChildren } from "react";

import appCss from "../styles.css?url";

export const Route = createRootRoute({
  head: () => ({
    meta: [
      { charSet: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      { title: "Jarvis" },
    ],
    links: [{ rel: "stylesheet", href: appCss }],
  }),
  component: Root,
});

function Root() {
  return (
    <Document>
      <div className="mx-auto flex min-h-screen max-w-6xl flex-col gap-4 p-4">
        <header className="flex items-center gap-2">
          <h1 className="text-base font-semibold">Jarvis</h1>
        </header>
        <main>
          <Outlet />
        </main>
      </div>
    </Document>
  );
}

function Document({ children }: PropsWithChildren) {
  return (
    <html lang="en">
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
