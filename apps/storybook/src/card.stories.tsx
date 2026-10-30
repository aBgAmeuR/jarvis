import type { Meta, StoryObj } from '@storybook/react-vite'

import { Badge, Button, Card, CardContent, CardFooter, CardHeader, CardTitle } from '@jarvis/ui'

const meta = {
  title: 'Components/Card',
  component: Card,
} satisfies Meta<typeof Card>

export default meta
type Story = StoryObj<typeof meta>

const alerts = [
  { name: 'ZfsPoolDegraded', object: 'pve / hdd', severity: 'critical', since: '12 min' },
  { name: 'CertificateExpiresSoon', object: 'grafana', severity: 'warning', since: '2 d' },
  { name: 'AptUpdatesPending', object: 'CT 116', severity: 'info', since: '5 h' },
] as const

export const ToHandle: Story = {
  render: (args) => (
    <Card {...args} className="w-96">
      <CardHeader>
        <CardTitle>To handle</CardTitle>
        <Badge tone="critical">3</Badge>
        <Button size="sm" variant="ghost">
          All alerts
        </Button>
      </CardHeader>
      <CardContent className="p-0">
        <ul className="divide-y divide-border">
          {alerts.map((alert) => (
            <li key={alert.name} className="flex items-center gap-2 px-3 py-1.5">
              <Badge tone={alert.severity} dot>
                {alert.severity}
              </Badge>
              <span className="flex-1 truncate">{alert.name}</span>
              <span className="text-fg-muted">{alert.object}</span>
              <span className="w-10 text-right text-fg-subtle">{alert.since}</span>
            </li>
          ))}
        </ul>
      </CardContent>
      <CardFooter className="text-xs">Refreshed every 10 s</CardFooter>
    </Card>
  ),
}

export const Simple: Story = {
  render: (args) => (
    <Card {...args} className="w-56">
      <CardHeader>
        <CardTitle>CPU</CardTitle>
      </CardHeader>
      <CardContent>
        <div className="text-2xl font-semibold">23.4 %</div>
        <div className="text-xs text-fg-muted">16 cores, load 1.82</div>
      </CardContent>
    </Card>
  ),
}
