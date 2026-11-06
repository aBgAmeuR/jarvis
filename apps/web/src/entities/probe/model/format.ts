const percent = new Intl.NumberFormat(undefined, {
  minimumFractionDigits: 2,
  maximumFractionDigits: 2,
});

/** Never rounds up to 100 %: 0.99996 → "99.99 %". */
export function formatAvailability(ratio: number): string {
  return `${percent.format(ratio < 1 ? Math.min(ratio * 100, 99.99) : 100)} %`;
}
