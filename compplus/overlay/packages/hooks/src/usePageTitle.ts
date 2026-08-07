import { useEffect } from "react";

export function usePageTitle(title: string) {
  useEffect(() => {
    document.title = title ? `${title} - ISOPilot` : "ISOPilot";
    return () => {
      document.title = "ISOPilot";
    };
  }, [title]);
}
