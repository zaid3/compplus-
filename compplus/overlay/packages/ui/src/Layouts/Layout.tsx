// Copyright (c) 2025-2026 Probo Inc <hello@probo.com>.
// MIT licensed upstream component, modified for the ISOPilot interface.

import { clsx } from "clsx";
import {
  createContext,
  type PropsWithChildren,
  type ReactNode,
  useContext,
  useEffect,
  useMemo,
  useState,
} from "react";
import { createPortal } from "react-dom";
import { Link } from "react-router";

import { Logo } from "../Atoms/Logo/Logo";
import { Sidebar } from "../Atoms/Sidebar/Sidebar";
import { Toasts } from "../Atoms/Toasts/Toasts";
import { ConfirmDialog } from "../Molecules/Dialog/ConfirmDialog";

type Props = PropsWithChildren<{
  headerLeading?: ReactNode;
  headerTrailing: ReactNode;
  sidebar?: ReactNode;
}>;

const LayoutContext = createContext<{ setDrawer: (v: boolean) => void }>({
  setDrawer: () => {},
});

export function Layout({
  headerLeading,
  headerTrailing,
  sidebar,
  children,
}: Props) {
  const [hasDrawer, setDrawer] = useState(false);
  const layoutContext = useMemo(() => ({ setDrawer }), []);

  return (
    <LayoutContext value={layoutContext}>
      <div className="min-h-screen bg-level-0 text-txt-primary">
        <header className="fixed inset-x-0 top-0 z-20 flex h-14 items-center border-b border-border-mid bg-level-1/95 px-4 shadow-sm backdrop-blur supports-[backdrop-filter]:bg-level-1/90">
          <Link
            to="/"
            className="inline-flex shrink-0 items-center rounded-lg focus-visible:outline-none"
            aria-label="ISOPilot home"
          >
            <Logo withPicto className="h-8 w-auto" />
          </Link>
          {headerLeading && (
            <>
              <div className="mx-3 h-6 w-px bg-border-mid" aria-hidden="true" />
              {headerLeading}
            </>
          )}
          <div className="ml-auto">{headerTrailing}</div>
        </header>

        <div className="flex min-h-screen" id="main">
          {sidebar && <Sidebar>{sidebar}</Sidebar>}
          <main
            className={clsx(
              "mt-14 w-full min-w-0 bg-level-0 transition-all duration-300",
              hasDrawer && "pr-105",
            )}
          >
            <div className="mx-auto min-h-[calc(100vh-56px)] w-full max-w-[1280px] px-6 py-9 lg:px-9 lg:py-10">
              {children}
            </div>
          </main>
        </div>
        <Toasts />
        <ConfirmDialog />
      </div>
    </LayoutContext>
  );
}

export function Drawer({
  children,
  className,
}: PropsWithChildren<{ className?: string }>) {
  const { setDrawer } = useContext(LayoutContext);
  useEffect(() => {
    setDrawer(true);
    return () => setDrawer(false);
  }, [setDrawer]);

  return createPortal(
    <aside
      className={clsx(
        "fixed right-0 top-0 h-screen w-105 border-l border-border-mid bg-level-1 px-6 pb-8 pt-20 shadow-lg",
        className,
      )}
    >
      {children}
    </aside>,
    document.body,
  );
}
