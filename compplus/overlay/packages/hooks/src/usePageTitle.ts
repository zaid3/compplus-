import { useEffect } from "react";

export function usePageTitle(title: string) {
  useEffect(() => {
    document.title = title ? `${title} - Comp Plus+` : "Comp Plus+";
    return () => {
      document.title = "Comp Plus+";
    };
  }, [title]);
}
