import { useRender } from "@base-ui/react/use-render";
import { Icon, SidebarLeft01Icon } from "@jarvis/icons";
import { cn } from "cn";
import {
  createContext,
  useCallback,
  useContext,
  useEffect,
  useMemo,
  useState,
  type ComponentProps,
  type ReactNode,
} from "react";

import { Button, type ButtonProps } from "./button.tsx";

type SidebarState = {
  /** Desktop: the sidebar takes its width, or is folded away. */
  open: boolean;
  /** Mobile: the sidebar slides over the page. */
  openMobile: boolean;
  toggle: () => void;
  closeMobile: () => void;
};

/** Without a provider the sidebar is always shown, as in a story. */
const fallbackState: SidebarState = {
  open: true,
  openMobile: true,
  toggle: () => {},
  closeMobile: () => {},
};

const SidebarContext = createContext<SidebarState>(fallbackState);

const desktopQuery = "(min-width: 768px)";

/**
 * Holds the open state of the sidebar and binds Ctrl/⌘ + B to it. Wrap the
 * `Sidebar` and the `SidebarInset` in it.
 */
export function SidebarProvider({
  defaultOpen = true,
  className,
  ...props
}: ComponentProps<"div"> & { defaultOpen?: boolean }) {
  const [open, setOpen] = useState(defaultOpen);
  const [openMobile, setOpenMobile] = useState(false);

  const toggle = useCallback(() => {
    if (window.matchMedia(desktopQuery).matches) {
      setOpen((value) => !value);
    } else {
      setOpenMobile((value) => !value);
    }
  }, []);
  const closeMobile = useCallback(() => setOpenMobile(false), []);

  useEffect(() => {
    const onKeyDown = (event: KeyboardEvent) => {
      if (event.key.toLowerCase() === "b" && (event.metaKey || event.ctrlKey)) {
        event.preventDefault();
        toggle();
      }
    };
    window.addEventListener("keydown", onKeyDown);
    return () => window.removeEventListener("keydown", onKeyDown);
  }, [toggle]);

  const state = useMemo(
    () => ({ open, openMobile, toggle, closeMobile }),
    [open, openMobile, toggle, closeMobile],
  );

  return (
    <SidebarContext value={state}>
      <div className={cn("flex h-svh w-full overflow-hidden", className)} {...props} />
    </SidebarContext>
  );
}

export function useSidebar(): SidebarState {
  return useContext(SidebarContext);
}

/**
 * The app navigation, 184 px wide, on the sidebar depth. Folds away on
 * desktop and slides over the page on mobile, both through `SidebarTrigger`.
 */
export function Sidebar({ className, children, ...props }: ComponentProps<"nav">) {
  const { open, openMobile, closeMobile } = useSidebar();
  return (
    <>
      {openMobile && (
        <div
          aria-hidden
          className="bg-foreground/30 fixed inset-0 z-30 md:hidden"
          onClick={closeMobile}
        />
      )}
      <nav
        inert={!open && !openMobile ? true : undefined}
        className={cn(
          "shrink-0 overflow-hidden border-sidebar-border bg-sidebar text-sidebar-foreground",
          "transition-[width,translate] duration-150 ease-out motion-reduce:transition-none",
          "max-md:fixed max-md:inset-y-0 max-md:left-0 max-md:z-40 max-md:w-46 max-md:border-r",
          !openMobile && "max-md:-translate-x-full",
          open ? "md:w-46 md:border-r" : "md:w-0",
          className,
        )}
        {...props}
      >
        <div className="flex h-full w-46 flex-col">{children}</div>
      </nav>
    </>
  );
}

/** The page next to the sidebar. */
export function SidebarInset({ className, ...props }: ComponentProps<"div">) {
  return <div className={cn("flex min-h-0 min-w-0 flex-1 flex-col", className)} {...props} />;
}

/** Folds or shows the sidebar. */
export function SidebarTrigger({ onClick, ...props }: ButtonProps) {
  const { toggle } = useSidebar();
  return (
    <Button
      variant="ghost"
      size="icon"
      aria-label="Toggle sidebar"
      onClick={(event) => {
        onClick?.(event);
        toggle();
      }}
      {...props}
    >
      <Icon icon={SidebarLeft01Icon} />
    </Button>
  );
}

/** The top of the sidebar: the app name, or a switcher. */
export function SidebarHeader({ className, ...props }: ComponentProps<"header">) {
  return (
    <header className={cn("flex h-9 shrink-0 items-center gap-2 px-3", className)} {...props} />
  );
}

/** The scrollable middle of the sidebar, holding the groups. */
export function SidebarContent({ className, ...props }: ComponentProps<"div">) {
  return (
    <div
      className={cn("flex min-h-0 flex-1 flex-col gap-px overflow-y-auto p-1.5 pt-0", className)}
      {...props}
    />
  );
}

/** The bottom of the sidebar, always visible. */
export function SidebarFooter({ className, ...props }: ComponentProps<"footer">) {
  return <footer className={cn("flex shrink-0 flex-col gap-px p-1.5", className)} {...props} />;
}

/** A module name above its pages, such as Supervision. */
export function SidebarGroupLabel({ className, ...props }: ComponentProps<"div">) {
  return (
    <div
      className={cn("px-2 pt-2 pb-0.5 text-2xs font-medium text-muted-foreground", className)}
      {...props}
    />
  );
}

export type SidebarItemProps = useRender.ComponentProps<"a"> & {
  /** Shown on the right: a count, or a severity badge. */
  count?: ReactNode;
};

/**
 * A page link. The current page is the one with `aria-current="page"`, which
 * the router link sets. Put an `Icon` first.
 */
export function SidebarItem({ render, className, count, children, ...props }: SidebarItemProps) {
  return useRender({
    defaultTagName: "a",
    render,
    props: {
      ...props,
      className: cn(
        "flex items-center gap-1.75 rounded-sm px-1.5 py-0.75 text-sm text-sidebar-foreground",
        "transition-colors duration-100 ease-out",
        "hover:bg-sidebar-accent hover:text-sidebar-accent-foreground",
        "aria-[current=page]:bg-sidebar-accent aria-[current=page]:text-sidebar-accent-foreground",
        "aria-disabled:pointer-events-none aria-disabled:opacity-45",
        "outline-hidden focus-visible:outline-2 focus-visible:outline-solid focus-visible:-outline-offset-2 focus-visible:outline-sidebar-ring",
        "[&_svg]:size-3.25 [&_svg]:shrink-0",
        className,
      ),
      children: (
        <>
          {children}
          {count !== undefined && (
            <span className="text-muted-foreground ml-auto text-xs font-normal">{count}</span>
          )}
        </>
      ),
    },
  });
}
