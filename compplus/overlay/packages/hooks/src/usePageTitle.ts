import { useEffect } from "react";

export function usePageTitle(title: string) {
  useEffect(() => {
    document.title = title ? `${title} - ISOpilot` : "ISOpilot";
    return () => {
      document.title = "ISOpilot";
    };
  }, [title]);
}
