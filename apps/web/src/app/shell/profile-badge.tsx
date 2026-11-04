import { useGetMe } from "@jarvis/api-client";
import { Icon, UserIcon } from "@jarvis/icons";
import { Badge, Skeleton, Tooltip } from "@jarvis/ui";

export function ProfileBadge() {
  const { data: account } = useGetMe({ query: { staleTime: Infinity, refetchInterval: false } });

  if (!account) return <Skeleton className="h-4.5 w-20 rounded-sm" />;

  return (
    <Tooltip content={account.isAdmin ? "Administrator" : "Member"}>
      <Badge className="min-w-0">
        <Icon icon={UserIcon} className="size-3 shrink-0" />
        <span className="truncate">{account.user}</span>
      </Badge>
    </Tooltip>
  );
}
