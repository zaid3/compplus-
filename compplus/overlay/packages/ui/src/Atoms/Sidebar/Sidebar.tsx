// Copyright (c) 2025-2026 Probo Inc <hello@probo.com>.
// MIT licensed upstream component, modified for the ISOPilot interface.

import { clsx } from "clsx";
import {
  createContext,
  type PropsWithChildren,
  useContext,
  useState,
} from "react";

import { Button } from "../Button/Button";
import { IconCollapse, IconExpand } from "../Icons";

const sidebarContext = createContext({ open: true });

function useSidebarState() {
  const [open, setOpenState] = useState<boolean>(() => {
    const stored = localStorage.getItem("sidebar-open");
    return stored !== null ? !!JSON.parse(stored) : true;
  });

  const setOpen = (value: boolean) => {
    setOpenState(value);
    localStorage.setItem("sidebar-open", JSON.stringify(value));
  };

  return [open, setOpen] as const;
}

export function Sidebar({ children }: PropsWithChildren) {
  const [open, setOpen] = useSidebarState();

  return (
    <sidebarContext.Provider value={{ open }}>
      <aside
        className={clsx(
          "mt-14 flex min-h-[calc(100vh-56px)] flex-none flex-col border-r border-border-mid bg-level-1 shadow-[1px_0_0_rgba(15,23,42,0.02)]",
          open ? "w-[288px]" : "w-[56px]",
        )}
      >
        <div className={clsx("flex-1 pb-3 pt-4", open ? "px-4" : "px-2")}>
          {children}
        </div>
        <div
          className={clsx(
            "sticky bottom-0 flex-none border-t border-border-mid bg-level-1 py-2",
            open ? "px-4" : "px-2",
          )}
        >
          <Button
            variant="tertiary"
            icon={open ? IconCollapse : IconExpand}
            onClick={() => setOpen(!open)}
            aria-label={open ? "Collapse sidebar" : "Expand sidebar"}
          />
        </div>
      </aside>
    </sidebarContext.Provider>
  );
}

export function useSidebarCollapsed(): boolean {
  return !useContext(sidebarContext).open;
}
