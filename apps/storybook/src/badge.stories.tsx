import type { Meta, StoryObj } from '@storybook/react-vite'

import { Badge } from '@jarvis/ui'

const meta = {
  title: 'Components/Badge',
  component: Badge,
  args: { children: 'warning', tone: 'warning', dot: false },
  argTypes: {
    tone: { control: 'inline-radio', options: ['critical', 'warning', 'info', 'ok', 'neutral'] },
  },
} satisfies Meta<typeof Badge>

export default meta
type Story = StoryObj<typeof meta>

export const Default: Story = {}

export const Severities: Story = {
  render: (args) => (
    <div className="flex items-center gap-2">
      <Badge {...args} tone="critical">
        critical
      </Badge>
      <Badge {...args} tone="warning">
        warning
      </Badge>
      <Badge {...args} tone="info">
        info
      </Badge>
      <Badge {...args} tone="ok">
        ok
      </Badge>
      <Badge {...args} tone="neutral">
        lxc
      </Badge>
    </div>
  ),
}

export const WithDot: Story = {
  ...Severities,
  args: { dot: true },
}
