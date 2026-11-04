/** Refresh interval per data type, in ms. See docs/api.md, Polling. */
export const polling = {
  metrics: 10_000,
  inventory: 60_000,
  catalog: 300_000,
} as const;
