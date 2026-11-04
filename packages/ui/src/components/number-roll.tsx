import { cn } from "cn";
import {
  useEffect,
  useMemo,
  useRef,
  useState,
  useSyncExternalStore,
  type ComponentProps,
  type Dispatch,
  type SetStateAction,
} from "react";

import type { CssVars } from "../lib/styles.ts";
import { merge, rollDelta, toParts, type RenderedPart } from "./number-roll-parts.ts";
import { noRollOnServer, subscribeToNothing, supportsRoll } from "./number-roll-support.ts";

/*
 * Adapted from the assistant-ui NumberRoll. Digits roll through a registered
 * CSS custom property and `mod()`; no animation library is involved.
 */

const defaultDuration = 500;
const digitCells = Array.from({ length: 10 }, (_, cell) => cell);

type Display = { value: number; formatted: string; rendered: RenderedPart[]; dir: number };

function NumberRollDigit({ digit, dir }: { digit: number; dir: number }) {
  const [state, setState] = useState({ digit, roll: digit });
  if (state.digit !== digit) {
    setState({ digit, roll: state.roll + rollDelta(state.digit, digit, dir) });
  }
  const style: CssVars = { "--number-roll-pos": state.roll };

  return (
    <span
      data-slot="number-roll-digit"
      className="relative inline-block overflow-clip [transition-property:--number-roll-pos] duration-(--number-roll-duration) ease-(--number-roll-ease) motion-reduce:transition-none"
      style={style}
    >
      {/* The glyphs come from ::before, so find-in-page and copy never see the strip. */}
      <span data-d={digit} className="invisible before:content-[attr(data-d)]" />
      {digitCells.map((cell) => {
        const cellStyle: CssVars = {
          "--number-roll-y": `clamp(-1lh, calc((mod(mod(${cell} - var(--number-roll-pos), 10) + 5, 10) - 5) * 1lh), 1lh)`,
        };
        return (
          <span
            key={cell}
            data-d={cell}
            className="absolute inset-0 translate-y-(--number-roll-y) text-center before:content-[attr(data-d)]"
            style={cellStyle}
          />
        );
      })}
    </span>
  );
}

function NumberRollPart({ part, dir }: { part: RenderedPart; dir: number }) {
  const style: CssVars = { "--number-roll-shift": dir === 0 ? "0%" : dir > 0 ? "35%" : "-35%" };

  return (
    <span
      data-slot="number-roll-part"
      className={cn(
        "inline-grid grid-cols-[1fr] transition-[grid-template-columns,opacity,translate] duration-(--number-roll-fade) ease-out motion-reduce:transition-none",
        part.entered &&
          "starting:translate-y-(--number-roll-shift) starting:grid-cols-[0fr] starting:opacity-0",
        part.exiting &&
          "pointer-events-none translate-y-[calc(var(--number-roll-shift)*-1)] grid-cols-[0fr] opacity-0",
      )}
      style={style}
    >
      <span className="min-w-0 overflow-hidden">
        {part.type === "digit" ? (
          <NumberRollDigit digit={part.digit} dir={dir} />
        ) : (
          <span data-slot="number-roll-symbol" className="whitespace-pre">
            {part.value}
          </span>
        )}
      </span>
    </span>
  );
}

/**
 * Removes each exiting part once its fade ends. One timer per part, so a later
 * batch does not extend the life of parts already leaving.
 */
function useExitTimers(
  rendered: RenderedPart[],
  duration: number,
  setDisplay: Dispatch<SetStateAction<Display>>,
) {
  const exitTimers = useRef(new Map<string, ReturnType<typeof setTimeout>>());
  const lastDuration = useRef(duration);
  useEffect(() => {
    const timers = exitTimers.current;
    if (lastDuration.current !== duration) {
      lastDuration.current = duration;
      for (const timer of timers.values()) clearTimeout(timer);
      timers.clear();
    }
    const exiting = new Set(rendered.filter((part) => part.exiting).map((p) => p.key));
    for (const [key, timer] of timers) {
      if (!exiting.has(key)) {
        clearTimeout(timer);
        timers.delete(key);
      }
    }
    for (const key of exiting) {
      if (timers.has(key)) continue;
      timers.set(
        key,
        setTimeout(() => {
          timers.delete(key);
          setDisplay((current) => ({
            ...current,
            rendered: current.rendered.filter((part) => !(part.exiting && part.key === key)),
          }));
        }, duration),
      );
    }
  }, [rendered, duration, setDisplay]);
  useEffect(() => {
    const timers = exitTimers.current;
    return () => {
      for (const timer of timers.values()) clearTimeout(timer);
      timers.clear();
    };
  }, []);
}

export type NumberRollProps = Omit<ComponentProps<"span">, "children" | "prefix"> & {
  value: number;
  /** Keep the object stable, such as a module constant: a new one rebuilds the formatter. */
  format?: Intl.NumberFormatOptions;
  locales?: Intl.LocalesArgument;
  prefix?: string;
  suffix?: string;
  /** Roll direction. `auto` follows the sign of the change. */
  trend?: "auto" | "up" | "down";
  /** Digit roll in ms. Characters entering and leaving fade in 60 % of it. */
  duration?: number;
};

/**
 * A number whose digits roll like an odometer when `value` changes. Entering
 * and leaving characters, such as the `K` of `1.1K`, slide and fade. It shows
 * the plain formatted text on the server, without CSS `mod()`, and under
 * reduced motion. Screen readers get the text, but a change is not announced.
 */
export function NumberRoll({
  value,
  format,
  locales,
  prefix,
  suffix,
  trend = "auto",
  duration = defaultDuration,
  className,
  style,
  ...props
}: NumberRollProps) {
  const enhanced = useSyncExternalStore(subscribeToNothing, supportsRoll, noRollOnServer);

  const formatter = useMemo(() => new Intl.NumberFormat(locales, format), [locales, format]);
  const parts = useMemo(
    () => toParts(value, formatter, prefix, suffix),
    [value, formatter, prefix, suffix],
  );
  const formatted = `${prefix ?? ""}${formatter.format(value)}${suffix ?? ""}`;

  const [display, setDisplay] = useState<Display>({ value, formatted, rendered: parts, dir: 0 });

  if (display.formatted !== formatted) {
    const directions = { up: 1, down: -1, auto: Math.sign(value - display.value) };
    setDisplay({
      value,
      formatted,
      rendered: merge(display.rendered, parts),
      dir: directions[trend],
    });
  }

  useExitTimers(display.rendered, duration, setDisplay);

  const rootStyle: CssVars = {
    "--number-roll-duration": `${duration}ms`,
    "--number-roll-fade": "calc(var(--number-roll-duration) * 0.6)",
    "--number-roll-ease": "cubic-bezier(0.23, 1, 0.32, 1)",
    ...style,
  };

  return (
    <span
      data-slot="number-roll"
      className={cn("inline-block whitespace-nowrap tabular-nums", className)}
      style={rootStyle}
      {...props}
    >
      <span className="sr-only">{formatted}</span>
      {enhanced ? (
        <span aria-hidden className="inline-block select-none">
          {display.rendered.map((part) => (
            <NumberRollPart key={part.key} part={part} dir={display.dir} />
          ))}
        </span>
      ) : (
        <span aria-hidden>{formatted}</span>
      )}
    </span>
  );
}
