import type { DiskBody } from "@jarvis/api-client";

export function diskKind(disk: DiskBody): "NVMe" | "SSD" | "HDD" {
  if (disk.protocol.toLowerCase() === "nvme") return "NVMe";
  return disk.rotational ? "HDD" : "SSD";
}
