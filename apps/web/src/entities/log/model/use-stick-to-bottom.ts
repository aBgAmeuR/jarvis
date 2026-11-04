import { useEffect, useRef } from "react";

export function useStickToBottom() {
  const ref = useRef<HTMLDivElement>(null);
  const stuck = useRef(true);

  useEffect(() => {
    const element = ref.current;
    if (element && stuck.current) element.scrollTop = element.scrollHeight;
  });

  const onScroll = () => {
    const element = ref.current;
    if (element)
      stuck.current = element.scrollHeight - element.scrollTop - element.clientHeight < 24;
  };

  return { ref, onScroll };
}
