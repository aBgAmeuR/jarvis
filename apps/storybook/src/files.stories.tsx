import { Button, Card, CardContent, DropZone, Progress } from "@jarvis/ui";
import type { Meta, StoryObj } from "@storybook/react-vite";
import { useState } from "react";

const meta = {
  title: "Components/Files",
  parameters: {
    docs: {
      description: {
        component: "Upload surface for the Files module, the only place Jarvis writes.",
      },
    },
  },
} satisfies Meta;

export default meta;
type Story = StoryObj<typeof meta>;

function Upload() {
  const [names, setNames] = useState<string[]>([]);
  return (
    <Card className="w-96">
      <CardContent className="flex flex-col gap-3">
        <DropZone
          label={
            <>
              Drop files to upload to <span className="font-mono text-xs">/shared/photos</span>
            </>
          }
          onFiles={(files) => setNames(files.map((file) => file.name))}
        />
        {names.length > 0 && (
          <div className="text-muted-foreground">Picked: {names.join(", ")}</div>
        )}
        <div className="flex flex-col gap-1">
          <div className="flex items-center gap-2">
            <span className="truncate font-mono text-xs">IMG_4821.HEIC</span>
            <span className="text-muted-foreground ml-auto">64 % · 12.4 MB/s</span>
          </div>
          <Progress value={64} aria-label="IMG_4821.HEIC upload" />
        </div>
        <div className="flex flex-col gap-1">
          <div className="flex items-center gap-2">
            <span className="truncate font-mono text-xs">walkthrough.mp4</span>
            <span className="text-critical ml-auto">Failed · quota</span>
            <Button size="sm" variant="ghost">
              Retry
            </Button>
          </div>
          <Progress value={38} tone="critical" aria-label="walkthrough.mp4 upload" />
        </div>
      </CardContent>
    </Card>
  );
}

export const UploadStory: Story = {
  parameters: {
    docs: { description: { story: "Drop zone with keyboard-accessible file selection." } },
  },
  name: "Drop zone and progress",
  render: () => <Upload />,
};
