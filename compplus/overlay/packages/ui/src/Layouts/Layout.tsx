// Copyright (c) 2025-2026 Probo Inc <hello@probo.com>.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

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
import { Link, useLocation } from "react-router";

import { Logo } from "../Atoms/Logo/Logo";
import { Sidebar } from "../Atoms/Sidebar/Sidebar";
import { Toasts } from "../Atoms/Toasts/Toasts";
import { ConfirmDialog } from "../Molecules/Dialog/ConfirmDialog";

type Props = PropsWithChildren<{
  headerLeading?: ReactNode;
  headerTrailing: ReactNode;
  sidebar?: ReactNode;
}>;

const COMPACT_MEDIA_QUERY = "(max-width: 1023px)";

const LayoutContext = createContext<{ setDrawer: (v: boolean) => void }>({
  setDrawer: () => {},
});

function useCompactLayout(): boolean {
  const [compact, setCompact] = useState(() =>
    typeof window !== "undefined"
      ? window.matchMedia(COMPACT_MEDIA_QUERY).matches
      : false,
  );

  useEffect(() => {
    const media = window.matchMedia(COMPACT_MEDIA_QUERY);
    const update = () => setCompact(media.matches);

    update();
    media.addEventListener("change", update);
    return () => media.removeEventListener("change", update);
  }, []);

  return compact;
}

export function Layout({
  headerLeading,
  headerTrailing,
  sidebar,
  children,
}: Props) {
  const [hasDrawer, setDrawer] = useState(false);
  const [mobileNavOpen, setMobileNavOpen] = useState(false);
  const compact = useCompactLayout();
  const location = useLocation();
  const layoutContext = useMemo(
    () => ({
      setDrawer,
    }),
    [],
  );

  useEffect(() => {
    setMobileNavOpen(false);
  }, [location.pathname]);

  useEffect(() => {
    if (!compact) {
      setMobileNavOpen(false);
    }
  }, [compact]);

  useEffect(() => {
    if (!mobileNavOpen) return;

    const previousOverflow = document.body.style.overflow;
    document.body.style.overflow = "hidden";

    const onKeyDown = (event: KeyboardEvent) => {
      if (event.key === "Escape") {
        setMobileNavOpen(false);
      }
    };
    window.addEventListener("keydown", onKeyDown);

    return () => {
      document.body.style.overflow = previousOverflow;
      window.removeEventListener("keydown", onKeyDown);
    };
  }, [mobileNavOpen]);

  return (
    <LayoutContext value={layoutContext}>
      <div className="isopilot-shell text-txt-primary bg-level-0 min-h-screen min-w-0 overflow-x-hidden">
        <header className="isopilot-header fixed top-0 z-30 left-0 right-0 px-3 sm:px-4 lg:px-5 flex items-center border-b border-border-solid h-14 bg-level-1">
          {sidebar && compact && (
            <button
              type="button"
              className="isopilot-mobile-menu-button mr-1 inline-flex size-10 shrink-0 items-center justify-center rounded-xl text-txt-secondary transition-colors hover:bg-subtle-hover hover:text-txt-primary focus-visible:outline-none focus-visible:shadow-focus"
              onClick={() => setMobileNavOpen(true)}
              aria-label="Open navigation"
              aria-expanded={mobileNavOpen}
              aria-controls="isopilot-mobile-navigation"
            >
              <MenuIcon />
            </button>
          )}

          <Link
            to="/"
            className="inline-flex shrink-0 items-center rounded-xl px-1.5 py-1 transition-opacity hover:opacity-90"
            aria-label="ISO Pilot home"
          >
            <Logo withPicto className="isopilot-header-logo h-8" />
          </Link>

          {headerLeading && !compact && (
            <>
              <div className="mx-4 h-6 w-px bg-border-solid" aria-hidden="true" />
              {headerLeading}
            </>
          )}

          <div className="ml-auto flex min-w-0 items-center gap-1 sm:gap-2">
            {headerTrailing}
          </div>
        </header>

        <div className="flex min-h-screen min-w-0" id="main">
          {sidebar && !compact && <Sidebar>{sidebar}</Sidebar>}
          <main
            className={clsx(
              "w-full min-w-0 mt-14 transition-all duration-300",
              hasDrawer && "xl:pr-105",
            )}
          >
            <div className="isopilot-content-frame py-5 px-4 sm:py-7 sm:px-6 lg:py-10 lg:px-10 max-w-[1280px] w-full mx-auto min-w-0 min-h-[calc(100vh-56px)]">
              {children}
            </div>
          </main>
        </div>

        {sidebar && compact && mobileNavOpen && (
          <div className="fixed inset-0 z-40 lg:hidden">
            <button
              type="button"
              className="absolute inset-0 bg-dialog/45 backdrop-blur-[2px]"
              onClick={() => setMobileNavOpen(false)}
              aria-label="Close navigation"
            />
            <aside
              id="isopilot-mobile-navigation"
              role="dialog"
              aria-modal="true"
              aria-label="ISO Pilot navigation"
              className="isopilot-mobile-nav absolute inset-y-0 left-0 z-10 flex w-[min(88vw,340px)] max-w-full flex-col border-r border-border-solid bg-level-1 shadow-mid"
            >
              <div className="flex h-14 shrink-0 items-center justify-between border-b border-border-solid px-4">
                <Logo withPicto className="h-8" />
                <button
                  type="button"
                  className="inline-flex size-10 items-center justify-center rounded-xl text-txt-secondary transition-colors hover:bg-subtle-hover hover:text-txt-primary focus-visible:outline-none focus-visible:shadow-focus"
                  onClick={() => setMobileNavOpen(false)}
                  aria-label="Close navigation"
                >
                  <CloseIcon />
                </button>
              </div>

              {headerLeading && (
                <div className="shrink-0 border-b border-border-solid px-4 py-3">
                  {headerLeading}
                </div>
              )}

              <nav className="flex-1 overflow-y-auto overscroll-contain px-3 py-3">
                {sidebar}
              </nav>
            </aside>
          </div>
        )}

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
    return () => {
      setDrawer(false);
    };
  }, [setDrawer]);

  return createPortal(
    <aside
      className={clsx(
        "fixed top-14 right-0 bottom-0 z-30 w-full max-w-full overflow-y-auto border-l border-border-solid bg-level-1 px-4 pb-[max(2rem,env(safe-area-inset-bottom))] pt-6 shadow-mid sm:w-[420px] sm:px-6",
        className,
      )}
    >
      {children}
    </aside>,
    document.body,
  );
}

function MenuIcon() {
  return (
    <svg
      aria-hidden="true"
      viewBox="0 0 24 24"
      className="size-5"
      fill="none"
      stroke="currentColor"
      strokeWidth="1.8"
      strokeLinecap="round"
    >
      <path d="M4 7h16M4 12h16M4 17h16" />
    </svg>
  );
}

function CloseIcon() {
  return (
    <svg
      aria-hidden="true"
      viewBox="0 0 24 24"
      className="size-5"
      fill="none"
      stroke="currentColor"
      strokeWidth="1.8"
      strokeLinecap="round"
    >
      <path d="m6 6 12 12M18 6 6 18" />
    </svg>
  );
}
