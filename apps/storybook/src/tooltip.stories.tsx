import type { Meta, StoryObj } from '@storybook/react-vite'

import { Badge, Button, Tooltip } from '@jarvis/ui'

const meta = {
  title: 'Components/Tooltip',
  component: Tooltip,
  args: {
    content: 'Open in Grafana Explore',
    side: 'top',
    children: <Button>Explore</Button>,
  },
  argTypes: {
    side: { control: 'inline-radio', options: ['top', 'right', 'bottom', 'left'] },
    align: { control: 'inline-radio', options: ['start', 'center', 'end'] },
    children: { control: false },
  },
} satisfies Meta<typeof Tooltip>

export default meta
type Story = StoryObj<typeof meta>

export const Default: Story = {}

export const Group: Story = {
  render: () => (
    <div className="flex gap-2">
      <Tooltip content="Grafana">
        <Button>Grafana</Button>
      </Tooltip>
      <Tooltip content="homelab-iac">
        <Button>Repo</Button>
      </Tooltip>
      <Tooltip content="Since 2026-10-30 09:12 UTC">
        <Badge tone="critical" tabIndex={0}>
          critical
        </Badge>
      </Tooltip>
    </div>
  ),
}
