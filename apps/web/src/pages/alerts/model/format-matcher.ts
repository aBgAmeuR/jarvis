import type { MatcherBody } from "@jarvis/api-client";

/** Alertmanager syntax: `vmid=116`, `name!~Foo.*`. */
export function formatMatcher({ name, value, isEqual, isRegex }: MatcherBody): string {
  const operator = `${isEqual ? "=" : "!"}${isRegex ? "~" : isEqual ? "" : "="}`;
  return `${name}${operator}${value}`;
}
