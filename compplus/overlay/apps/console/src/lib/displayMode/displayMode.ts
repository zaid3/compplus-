// Copyright (c) 2026 Probo Inc <hello@probo.com>.
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

export type DisplayMode = "light" | "dark";

type Listener = () => void;

const STORAGE_KEY = "isopilot-display-mode";
const listeners = new Set<Listener>();
let override: DisplayMode | null = null;

function readStoredMode(): DisplayMode | null {
  if (typeof window === "undefined") {
    return null;
  }

  try {
    const stored = window.localStorage.getItem(STORAGE_KEY);
    return stored === "light" || stored === "dark" ? stored : null;
  } catch {
    return null;
  }
}

function getSystemDisplayMode(): DisplayMode {
  if (typeof window === "undefined" || !window.matchMedia) {
    return "light";
  }

  return window.matchMedia("(prefers-color-scheme: dark)").matches ? "dark" : "light";
}

export function getDisplayMode(): DisplayMode {
  return override ?? getSystemDisplayMode();
}

function applyDisplayMode(mode: DisplayMode): void {
  if (typeof document === "undefined") {
    return;
  }

  document.documentElement.classList.toggle("dark", mode === "dark");
  document.documentElement.dataset.displayMode = mode;
}

function notify(): void {
  for (const listener of listeners) {
    listener();
  }
}

export function setDisplayMode(mode: DisplayMode): void {
  override = mode;

  if (typeof window !== "undefined") {
    try {
      window.localStorage.setItem(STORAGE_KEY, mode);
    } catch {
      // Browsers with blocked storage can still use the mode for this session.
    }
  }

  applyDisplayMode(mode);
  notify();
}

export function toggleDisplayMode(): void {
  setDisplayMode(getDisplayMode() === "dark" ? "light" : "dark");
}

export function subscribeDisplayMode(listener: Listener): () => void {
  listeners.add(listener);
  return () => {
    listeners.delete(listener);
  };
}

function initDisplayMode(): void {
  if (typeof window === "undefined") {
    return;
  }

  override = readStoredMode();
  applyDisplayMode(getDisplayMode());

  if (!window.matchMedia) {
    return;
  }

  const mediaQuery = window.matchMedia("(prefers-color-scheme: dark)");
  mediaQuery.addEventListener("change", () => {
    if (override !== null) {
      return;
    }

    applyDisplayMode(getDisplayMode());
    notify();
  });
}

initDisplayMode();
