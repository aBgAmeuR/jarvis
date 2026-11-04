import { z } from "zod";

const pair = /([\w.-]+)=("(?:[^"\\]|\\.)*"|\S*)/g;
const jsonObject = z.record(z.string(), z.unknown());

export type LogField = { key: string; value: string };

/** `pretty` is set for a JSON line. Fields are empty when the line has no structure. */
export function parseLine(line: string): { fields: LogField[]; pretty?: string } {
  const json = line.startsWith("{") ? jsonObject.safeParse(safeJson(line)) : undefined;
  if (json?.success) {
    return {
      fields: Object.entries(json.data).map(([key, value]) => ({
        key,
        value: typeof value === "string" ? value : JSON.stringify(value),
      })),
      pretty: JSON.stringify(json.data, null, 2),
    };
  }
  return {
    fields: [...line.matchAll(pair)].map(([, key = "", raw = ""]) => ({
      key,
      value: raw.startsWith('"') ? raw.slice(1, -1).replaceAll('\\"', '"') : raw,
    })),
  };
}

function safeJson(text: string): unknown {
  try {
    return JSON.parse(text);
  } catch {
    return undefined;
  }
}
