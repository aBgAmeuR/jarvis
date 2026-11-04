import type { ComponentProps } from "react";

export function Logo(props: ComponentProps<"svg">) {
  return (
    <svg viewBox="0 0 48 48" fill="none" aria-hidden {...props}>
      <circle
        cx="24"
        cy="24"
        r="19"
        stroke="currentColor"
        strokeWidth="2.5"
        strokeLinecap="round"
        strokeDasharray="84 35.4"
      />
      <circle cx="24" cy="24" r="6" fill="currentColor" />
      <circle cx="37.4" cy="10.6" r="5" fill="#D4537E" />
    </svg>
  );
}
