/** Swaps API series names for the short labels the chart tooltip shows. A name without a label is kept. */
export function labelSeries<T extends { name: string }>(
  lines: readonly T[],
  labels: Readonly<Partial<Record<string, string>>>,
): T[] {
  return lines.map((line) => ({ ...line, name: labels[line.name] ?? line.name }));
}
