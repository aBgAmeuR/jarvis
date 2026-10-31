import { CloudUploadIcon, Icon } from "@jarvis/icons";
import { cn } from "cn";
import { type ComponentProps, type DragEvent, type ReactNode, useRef, useState } from "react";

import { Button } from "./button.tsx";

export type DropZoneProps = Omit<ComponentProps<"div">, "onDrop" | "children"> & {
  /** Called with the dropped or picked files. Uploading is the caller's job. */
  onFiles: (files: File[]) => void;
  /** Shown at rest, such as "Drop files to upload to /shared/photos". */
  label: ReactNode;
  /** Shown while files are dragged over, such as "Release to upload". */
  activeLabel?: ReactNode;
  multiple?: boolean;
};

/** Takes files by drag and drop or by the Browse button (the only way on mobile). */
export function DropZone({
  onFiles,
  label,
  activeLabel = "Release to upload",
  multiple = true,
  className,
  ...props
}: DropZoneProps) {
  const input = useRef<HTMLInputElement>(null);
  const [active, setActive] = useState(false);

  function onDragOver(event: DragEvent) {
    event.preventDefault();
    setActive(true);
  }

  function onDrop(event: DragEvent) {
    event.preventDefault();
    setActive(false);
    onFiles([...event.dataTransfer.files]);
  }

  return (
    <div
      onDragOver={onDragOver}
      onDragLeave={() => setActive(false)}
      onDrop={onDrop}
      className={cn(
        "flex flex-col items-center gap-1.5 rounded-md border border-dashed border-border-strong p-5 text-center text-muted-foreground",
        "transition-colors duration-100 ease-out",
        active && "border-solid border-foreground bg-muted text-foreground",
        className,
      )}
      {...props}
    >
      {active ? (
        <div className="font-medium">{activeLabel}</div>
      ) : (
        <>
          <Icon icon={CloudUploadIcon} className="size-4" />
          <div>{label}</div>
          <Button size="sm" onClick={() => input.current?.click()}>
            Browse…
          </Button>
          <input
            ref={input}
            type="file"
            multiple={multiple}
            hidden
            onChange={(event) => {
              onFiles([...(event.target.files ?? [])]);
              event.target.value = "";
            }}
          />
        </>
      )}
    </div>
  );
}
