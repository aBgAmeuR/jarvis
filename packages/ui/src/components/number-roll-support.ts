// Rolling needs CSS mod() and a registered custom property; otherwise the plain text shows.

let rollSupport: boolean | undefined;

function detectRollSupport(): boolean {
  if (typeof CSS === "undefined" || typeof CSS.registerProperty !== "function") return false;
  const hasMod = CSS.supports(
    "transform",
    "translateY(clamp(-1lh, calc((mod(7.5, 10) - 5) * 1lh), 1lh))",
  );
  if (!hasMod) return false;
  try {
    CSS.registerProperty({
      name: "--number-roll-pos",
      syntax: "<number>",
      inherits: true,
      initialValue: "0",
    });
  } catch {
    // Already registered by another copy of this module.
  }
  return true;
}

export function supportsRoll(): boolean {
  rollSupport ??= detectRollSupport();
  return rollSupport;
}

export const subscribeToNothing = () => () => {};
export const noRollOnServer = () => false;
