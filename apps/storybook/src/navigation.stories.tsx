import {
  Activity01Icon,
  DashboardSquare01Icon,
  Folder01Icon,
  Icon,
  ServerStack01Icon,
} from "@jarvis/icons";
import {
  Badge,
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbPage,
  BreadcrumbSeparator,
  Link,
  Sidebar,
  SidebarContent,
  SidebarGroupLabel,
  SidebarItem,
} from "@jarvis/ui";
import type { Meta, StoryObj } from "@storybook/react-vite";

const meta = {
  title: "Components/Navigation",
  parameters: { docs: { description: { component: "Sidebar, breadcrumb and link." } } },
} satisfies Meta;

export default meta;
type Story = StoryObj<typeof meta>;

export const BreadcrumbStory: Story = {
  parameters: {
    docs: {
      description: { story: "Path of the current location. The last item is the current page." },
    },
  },
  name: "Breadcrumb",
  render: () => (
    <Breadcrumb>
      <BreadcrumbItem>
        <BreadcrumbLink href="#supervision">Supervision</BreadcrumbLink>
      </BreadcrumbItem>
      <BreadcrumbSeparator />
      <BreadcrumbItem>
        <BreadcrumbLink href="#guests">Guests</BreadcrumbLink>
      </BreadcrumbItem>
      <BreadcrumbSeparator />
      <BreadcrumbItem>
        <BreadcrumbPage>
          <span className="text-muted-foreground font-mono text-xs">ct-102</span> adguard
        </BreadcrumbPage>
      </BreadcrumbItem>
    </Breadcrumb>
  ),
};

export const SidebarStory: Story = {
  parameters: { docs: { description: { story: "Primary navigation between modules." } } },
  name: "Sidebar",
  render: () => (
    <Sidebar className="h-64">
      <SidebarContent>
        <SidebarGroupLabel>Supervision</SidebarGroupLabel>
        <SidebarItem
          href="#overview"
          aria-current="page"
          count={
            <Badge tone="critical" dot>
              1
            </Badge>
          }
        >
          <Icon icon={DashboardSquare01Icon} />
          Overview
        </SidebarItem>
        <SidebarItem href="#guests" count={8}>
          <Icon icon={ServerStack01Icon} />
          Guests
        </SidebarItem>
        <SidebarItem href="#probes" count={21}>
          <Icon icon={Activity01Icon} />
          Probes
        </SidebarItem>
        <SidebarItem aria-disabled count="v2">
          <Icon icon={Folder01Icon} />
          Files
        </SidebarItem>
      </SidebarContent>
    </Sidebar>
  ),
};

export const LinkStory: Story = {
  parameters: { docs: { description: { story: "Internal and external links." } } },
  name: "Link",
  render: () => <Link href="#explore">Open in Grafana Explore</Link>,
};
