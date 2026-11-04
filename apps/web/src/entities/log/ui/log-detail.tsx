import type { EntryBody } from "@jarvis/api-client";
import {
  Button,
  DescriptionDetails,
  DescriptionList,
  DescriptionTerm,
  Drawer,
  DrawerBody,
  DrawerContent,
  DrawerFooter,
  DrawerHeader,
  DrawerTitle,
} from "@jarvis/ui";

import { formatLogTime } from "../model/format-log-time.ts";
import { parseLine } from "../model/parse-line.ts";

type LogDetailProps = {
  entry: EntryBody | undefined;
  open: boolean;
  onOpenChange: (open: boolean) => void;
};

export function LogDetail({ entry, open, onOpenChange }: LogDetailProps) {
  const { fields, pretty } = parseLine(entry?.line ?? "");

  return (
    <Drawer open={open} onOpenChange={onOpenChange}>
      <DrawerContent>
        {entry && (
          <>
            <DrawerHeader>
              <DrawerTitle>Log line</DrawerTitle>
            </DrawerHeader>
            <DrawerBody>
              <DescriptionList>
                <DescriptionTerm>Time</DescriptionTerm>
                <DescriptionDetails>{formatLogTime(entry.timestamp)}</DescriptionDetails>
                <DescriptionTerm>Level</DescriptionTerm>
                <DescriptionDetails>{entry.level ?? "-"}</DescriptionDetails>
                <DescriptionTerm>Service</DescriptionTerm>
                <DescriptionDetails>{entry.service}</DescriptionDetails>
                {entry.instance && (
                  <>
                    <DescriptionTerm>Guest</DescriptionTerm>
                    <DescriptionDetails>
                      {entry.instance}
                      {entry.vmid !== undefined && ` (${entry.vmid})`}
                    </DescriptionDetails>
                  </>
                )}
              </DescriptionList>
              <pre className="bg-background rounded-sm border p-2 font-mono text-xs break-words whitespace-pre-wrap">
                {pretty ?? entry.line}
              </pre>
              {fields.length > 1 && (
                <DescriptionList>
                  {fields.map(({ key, value }, index) => (
                    <div key={`${key}-${index}`} className="contents">
                      <DescriptionTerm>{key}</DescriptionTerm>
                      <DescriptionDetails className="font-mono text-xs break-words">
                        {value}
                      </DescriptionDetails>
                    </div>
                  ))}
                </DescriptionList>
              )}
            </DrawerBody>
            <DrawerFooter>
              <Button size="sm" onClick={() => navigator.clipboard.writeText(entry.line)}>
                Copy line
              </Button>
            </DrawerFooter>
          </>
        )}
      </DrawerContent>
    </Drawer>
  );
}
