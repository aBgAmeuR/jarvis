import {
  Activity01Icon,
  Alert02Icon,
  Certificate01Icon,
  DashboardSquare01Icon,
  Database01Icon,
  GitCompareIcon,
  Globe02Icon,
  HardDriveIcon,
  Layers01Icon,
  NetworkIcon,
  Package01Icon,
  ServerStack01Icon,
  WifiIcon,
  type IconSvgElement,
} from "@jarvis/icons";
import type { LinkProps } from "@tanstack/react-router";

export type NavItem = {
  label: string;
  icon: IconSvgElement;
  to: LinkProps["to"];
};

export type NavGroup = {
  label: string;
  items: NavItem[];
};

export const navGroups: NavGroup[] = [
  {
    label: "Supervision",
    items: [
      { label: "Overview", icon: DashboardSquare01Icon, to: "/" },
      { label: "Node", icon: HardDriveIcon, to: "/node" },
      { label: "Guests", icon: ServerStack01Icon, to: "/guests" },
      { label: "Services", icon: Layers01Icon, to: "/services" },
    ],
  },
  {
    label: "Availability",
    items: [
      { label: "Probes", icon: Activity01Icon, to: "/probes" },
      { label: "Certificates", icon: Certificate01Icon, to: "/certificates" },
      { label: "Alerts", icon: Alert02Icon, to: "/alerts" },
    ],
  },
  {
    label: "Network",
    items: [
      { label: "Clients", icon: NetworkIcon, to: "/clients" },
      { label: "DNS", icon: Globe02Icon, to: "/dns" },
      { label: "VPN", icon: WifiIcon, to: "/vpn" },
    ],
  },
  {
    label: "Maintenance",
    items: [
      { label: "Updates", icon: Package01Icon, to: "/updates" },
      { label: "Backups", icon: Database01Icon, to: "/backups" },
      { label: "Drift", icon: GitCompareIcon, to: "/drift" },
    ],
  },
];
