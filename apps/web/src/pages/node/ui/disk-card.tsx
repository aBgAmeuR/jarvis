import type { DiskBody } from "@jarvis/api-client";
import {
  Badge,
  Card,
  CardContent,
  CardHeader,
  CardTitle,
  DescriptionDetails,
  DescriptionList,
  DescriptionTerm,
  Skeleton,
} from "@jarvis/ui";
import type { ReactNode } from "react";

import { formatDuration, formatInteger } from "@/shared/lib/format.ts";

import { diskKind } from "../model/disk-kind.ts";

function Row({ term, children }: { term: string; children: ReactNode }) {
  return (
    <>
      <DescriptionTerm>{term}</DescriptionTerm>
      <DescriptionDetails>{children}</DescriptionDetails>
    </>
  );
}

export function DiskCard({ disk }: { disk: DiskBody }) {
  return (
    <Card>
      <CardHeader>
        <CardTitle>
          <span className="text-muted-foreground font-mono text-xs">{disk.device}</span>{" "}
          {disk.model}
        </CardTitle>
        <Badge tone={disk.healthPassed ? "ok" : "critical"} dot>
          SMART {disk.healthPassed ? "passed" : "failed"}
        </Badge>
      </CardHeader>
      <CardContent>
        <DescriptionList>
          <Row term="Type">
            {diskKind(disk)} · {disk.protocol}
          </Row>
          <Row term="Serial">
            <span className="font-mono text-xs">{disk.serial}</span>
          </Row>
          <Row term="Temperature">
            {disk.temperatureCelsius === undefined
              ? "-"
              : `${formatInteger(disk.temperatureCelsius)} °C`}
          </Row>
          <Row term="Power on">
            {disk.powerOnHours === undefined ? "-" : formatDuration(disk.powerOnHours * 3_600)}
          </Row>
          <Row term="Power state">{disk.standby ? "Standby" : "Active"}</Row>
          {disk.nvme && (
            <>
              <Row term="Wear">{formatInteger(disk.nvme.percentageUsed)} %</Row>
              <Row term="Spare left">{formatInteger(disk.nvme.availableSparePercent)} %</Row>
              <Row term="Media errors">{formatInteger(disk.nvme.mediaErrors)}</Row>
              <Row term="Critical warning">
                {disk.nvme.criticalWarning === 0
                  ? "None"
                  : `0x${disk.nvme.criticalWarning.toString(16)}`}
              </Row>
            </>
          )}
          {disk.ata && (
            <>
              <Row term="Reallocated sectors">{formatInteger(disk.ata.reallocatedSectors)}</Row>
              <Row term="Pending sectors">{formatInteger(disk.ata.pendingSectors)}</Row>
              <Row term="Uncorrectable sectors">{formatInteger(disk.ata.uncorrectableSectors)}</Row>
              <Row term="UDMA CRC errors">{formatInteger(disk.ata.udmaCrcErrors)}</Row>
            </>
          )}
        </DescriptionList>
      </CardContent>
    </Card>
  );
}

const skeletonTerms = [
  "Type",
  "Serial",
  "Temperature",
  "Power on",
  "Power state",
  "Wear",
  "Spare left",
  "Media errors",
  "Critical warning",
];

/** Nine rows, as a disk with the four SMART fields of its protocol. */
export function DiskCardSkeleton() {
  return (
    <Card aria-hidden className="motion-safe:animate-none">
      <CardHeader>
        <CardTitle>
          <Skeleton className="h-lh w-40" />
        </CardTitle>
      </CardHeader>
      <CardContent>
        <DescriptionList>
          {skeletonTerms.map((term) => (
            <Row key={term} term={term}>
              <Skeleton className="h-lh w-24" />
            </Row>
          ))}
        </DescriptionList>
      </CardContent>
    </Card>
  );
}
