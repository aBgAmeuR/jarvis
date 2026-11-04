/** "2026-10-30 14:39:10", local time. */
export function formatLogTime(value: string): string {
  return new Date(value).toLocaleString("sv-SE");
}
