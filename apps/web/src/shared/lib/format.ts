/** Also the options of `NumberRoll`, so a rolling figure reads like a static one. */
export const numberOptions = { maximumFractionDigits: 1 } satisfies Intl.NumberFormatOptions;
export const integerOptions = { maximumFractionDigits: 0 } satisfies Intl.NumberFormatOptions;

const number = new Intl.NumberFormat(undefined, numberOptions);
const integer = new Intl.NumberFormat(undefined, integerOptions);
const rate = new Intl.NumberFormat(undefined, { maximumFractionDigits: 2 });
const dateTime = new Intl.DateTimeFormat(undefined, { dateStyle: "medium", timeStyle: "short" });

const utcDate = new Intl.DateTimeFormat(undefined, { dateStyle: "medium", timeZone: "UTC" });

const byteUnits = ["B", "KiB", "MiB", "GiB", "TiB", "PiB"] as const;

/** 1 536 → { value: 1.5, unit: "KiB" }. */
export function scaleBytes(bytes: number): { value: number; unit: string } {
  let value = bytes;
  let unit = 0;
  while (Math.abs(value) >= 1024 && unit < byteUnits.length - 1) {
    value /= 1024;
    unit++;
  }
  return { value, unit: byteUnits[unit] ?? "B" };
}

/** 1 536 → "1.5 KiB". */
export function formatBytes(bytes: number): string {
  const { value, unit } = scaleBytes(bytes);
  return `${number.format(value)} ${unit}`;
}

/** 0.4213 → "42 %". */
export function formatRatio(ratio: number): string {
  return `${integer.format(ratio * 100)} %`;
}

export function formatNumber(value: number): string {
  return number.format(value);
}

export function formatInteger(value: number): string {
  return integer.format(value);
}

/** 0.256 → "0.26 /s". Two digits: a home resolver idles under one query per second. */
export function formatPerSecond(value: number): string {
  return `${rate.format(value)} /s`;
}

/** 93 784 s → "1d 2h", 3 700 s → "1h 1m", 40 s → "40s". */
export function formatDuration(seconds: number): string {
  const days = Math.floor(seconds / 86_400);
  const hours = Math.floor((seconds % 86_400) / 3_600);
  const minutes = Math.floor((seconds % 3_600) / 60);
  if (days > 0) return `${days}d ${hours}h`;
  if (hours > 0) return `${hours}h ${minutes}m`;
  if (minutes > 0) return `${minutes}m`;
  return `${Math.floor(seconds)}s`;
}

export function formatDateTime(value: Date | string | number): string {
  return dateTime.format(new Date(value));
}

/** "3d 4h ago" for a past instant. */
export function formatAge(value: Date | string | number, now = Date.now()): string {
  const seconds = Math.max(0, (now - new Date(value).getTime()) / 1000);
  return `${formatDuration(seconds)} ago`;
}

/** A UTC day, such as the start of an availability bucket. */
export function formatUtcDate(value: Date | string | number): string {
  return utcDate.format(new Date(value));
}
