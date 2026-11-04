import { cn } from "cn";
import {
  useEffect,
  useMemo,
  useRef,
  useState,
  useSyncExternalStore,
  type ComponentProps,
} from "react";

import type { CssVars } from "../lib/styles.ts";

/*
 * Adapted from the assistant-ui NumberRoll. Digits roll through a registered
 * CSS custom property and `mod()`; no animation library is involved.
 */

const defaultDuration = 500;
const digitCells = Array.from({ length: 10 }, (_, cell) => cell);

type DigitPart = { type: "digit"; key: string; digit: number };
type SymbolPart = { type: "symbol"; key: string; value: string };
type Part = DigitPart | SymbolPart;
type RenderedPart = Part & { exiting?: boolean; entered?: boolean };

type Display = { value: number; formatted: string; rendered: RenderedPart[]; dir: number };

type Atom =
  | { kind: "integer"; digit: number }
  | { kind: "group"; value: string }
  | { kind: "fraction"; digit: number }
  | { kind: "symbol"; type: string; value: string };

let rollSupport: boolean | undefined;

function detectRollSupport(): boolean {
  if (typeof CSS === "undefined" || typeof CSS.registerProperty !== "function") return false;
  const hasMod = CSS.supports(
    "transform",
    "translateY(clamp(-1lh, calc((mod(7.5, 10) - 5) * 1lh), 1lh))",
  );
  if (!hasMod) return false;
  try {
    CSS.registerProperty({
      name: "--number-roll-pos",
      syntax: "<number>",
      inherits: true,
      initialValue: "0",
    });
  } catch {
    // Already registered by another copy of this module.
  }
  return true;
}

function supportsRoll(): boolean {
  rollSupport ??= detectRollSupport();
  return rollSupport;
}

const subscribeToNothing = () => () => {};
const noRollOnServer = () => false;

function toAtoms(
  value: number,
  formatter: Intl.NumberFormat,
  prefix: string | undefined,
  suffix: string | undefined,
): Atom[] {
  const atoms: Atom[] = [];
  if (prefix) atoms.push({ kind: "symbol", type: "prefix", value: prefix });
  for (const part of formatter.formatToParts(value)) {
    if (part.type === "integer" || part.type === "fraction") {
      for (const char of part.value) {
        const digit = char.charCodeAt(0) - 48;
        if (digit >= 0 && digit <= 9) {
          atoms.push({ kind: part.type, digit });
        } else {
          atoms.push({ kind: "symbol", type: part.type, value: char });
        }
      }
    } else if (part.type === "group") {
      atoms.push({ kind: "group", value: part.value });
    } else {
      const type = part.type === "minusSign" || part.type === "plusSign" ? "sign" : part.type;
      atoms.push({ kind: "symbol", type, value: part.value });
    }
  }
  if (suffix) atoms.push({ kind: "symbol", type: "suffix", value: suffix });
  return atoms;
}

function toParts(
  value: number,
  formatter: Intl.NumberFormat,
  prefix: string | undefined,
  suffix: string | undefined,
): Part[] {
  const atoms = toAtoms(value, formatter, prefix, suffix);
  const counts = new Map<string, number>();
  const nextKey = (type: string) => {
    const count = counts.get(type) ?? 0;
    counts.set(type, count + 1);
    return `${type}:${count}`;
  };

  // Integer digits and separators are keyed right to left, so the ones digit
  // is always int:0. 999 → 1,000 keeps the surviving places and only the new
  // leading parts enter.
  const placeKeys = new Map<number, string>();
  for (let index = atoms.length - 1; index >= 0; index--) {
    const atom = atoms[index];
    if (atom?.kind === "integer") placeKeys.set(index, nextKey("int"));
    else if (atom?.kind === "group") placeKeys.set(index, nextKey("group"));
  }

  return atoms.map((atom, index): Part => {
    switch (atom.kind) {
      case "integer":
        return { type: "digit", key: placeKeys.get(index) ?? nextKey("int"), digit: atom.digit };
      case "group":
        return {
          type: "symbol",
          key: placeKeys.get(index) ?? nextKey("group"),
          value: atom.value,
        };
      case "fraction":
        return { type: "digit", key: nextKey("fraction"), digit: atom.digit };
      case "symbol":
        return { type: "symbol", key: `${nextKey(atom.type)}:${atom.value}`, value: atom.value };
    }
  });
}

/** Keeps the parts that left on screen, flagged `exiting`, and flags the new ones `entered`. */
function merge(prev: RenderedPart[], next: Part[]): RenderedPart[] {
  const nextKeys = new Set(next.map((part) => part.key));
  const prevKeys = new Set(prev.map((part) => part.key));
  const out: RenderedPart[] = [];
  let index = 0;

  const emitExited = (until: string | undefined) => {
    for (let old = prev[index]; old && old.key !== until; old = prev[index]) {
      index++;
      if (!nextKeys.has(old.key)) out.push(old.exiting ? old : { ...old, exiting: true });
    }
  };

  for (const part of next) {
    if (prevKeys.has(part.key)) {
      emitExited(part.key);
      index++;
      out.push(part);
    } else {
      out.push({ ...part, entered: true });
    }
  }
  emitExited(undefined);
  return out;
}

/** Steps to turn from one digit to another: up, down, or the short way for `dir` 0. */
function rollDelta(from: number, to: number, dir: number): number {
  const up = (((to - from) % 10) + 10) % 10;
  if (dir > 0) return up;
  if (dir < 0) return up - 10;
  return up > 5 ? up - 10 : up;
}

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
      {digitCells.map((cell) => (
        <span
          key={cell}
          data-d={cell}
          className="absolute inset-0 text-center before:content-[attr(data-d)]"
          style={{
            transform: `translateY(clamp(-1lh, calc((mod(mod(${cell} - var(--number-roll-pos), 10) + 5, 10) - 5) * 1lh), 1lh))`,
          }}
        />
      ))}
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

  // One removal timer per exiting part, so a later batch does not extend the
  // life of parts already leaving.
  const exitTimers = useRef(new Map<string, ReturnType<typeof setTimeout>>());
  const lastDuration = useRef(duration);
  useEffect(() => {
    const timers = exitTimers.current;
    if (lastDuration.current !== duration) {
      lastDuration.current = duration;
      for (const timer of timers.values()) clearTimeout(timer);
      timers.clear();
    }
    const exiting = new Set(display.rendered.filter((part) => part.exiting).map((p) => p.key));
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
  }, [display.rendered, duration]);
  useEffect(() => {
    const timers = exitTimers.current;
    return () => {
      for (const timer of timers.values()) clearTimeout(timer);
      timers.clear();
    };
  }, []);

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
