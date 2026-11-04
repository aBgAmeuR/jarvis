/** Drops the leading logfmt time and level: the row already shows both. */
export function cleanLine(line: string): string {
  return line.replace(/^time=\S+\s+/, "").replace(/^level=\S+\s+/, "");
}
