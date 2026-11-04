// Splits a formatted number into keyed parts, so the digits that stay keep their place.

type DigitPart = { type: "digit"; key: string; digit: number };
type SymbolPart = { type: "symbol"; key: string; value: string };
export type Part = DigitPart | SymbolPart;
export type RenderedPart = Part & { exiting?: boolean; entered?: boolean };

type Atom =
  | { kind: "integer"; digit: number }
  | { kind: "group"; value: string }
  | { kind: "fraction"; digit: number }
  | { kind: "symbol"; type: string; value: string };

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

export function toParts(
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
export function merge(prev: RenderedPart[], next: Part[]): RenderedPart[] {
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
export function rollDelta(from: number, to: number, dir: number): number {
  const up = (((to - from) % 10) + 10) % 10;
  if (dir > 0) return up;
  if (dir < 0) return up - 10;
  return up > 5 ? up - 10 : up;
}
