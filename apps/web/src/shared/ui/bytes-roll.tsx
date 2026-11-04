import { NumberRoll } from "@jarvis/ui";

import { numberOptions, scaleBytes } from "../lib/format.ts";

/** `formatBytes` as a rolling figure: the unit slides when it changes. */
export function BytesRoll({ bytes }: { bytes: number }) {
  const { value, unit } = scaleBytes(bytes);
  return <NumberRoll value={value} format={numberOptions} suffix={` ${unit}`} />;
}
