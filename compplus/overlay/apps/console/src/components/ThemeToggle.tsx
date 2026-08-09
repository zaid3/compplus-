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

import { MoonIcon, SunIcon } from "@phosphor-icons/react";

import { useDisplayMode } from "#/lib/displayMode/useDisplayMode";

export function ThemeToggle() {
  const { displayMode, toggleDisplayMode } = useDisplayMode();
  const isDark = displayMode === "dark";

  return (
    <button
      type="button"
      className="isopilot-theme-toggle group inline-flex h-9 items-center gap-2 rounded-xl px-2.5 text-sm font-medium transition-all duration-200"
      onClick={toggleDisplayMode}
      aria-label={isDark ? "Switch to light mode" : "Switch to dark mode"}
      aria-pressed={isDark}
      title={isDark ? "Switch to light mode" : "Switch to dark mode"}
    >
      {isDark ? <MoonIcon size={16} /> : <SunIcon size={16} />}
      <span className="hidden sm:inline">{isDark ? "Dark" : "Light"}</span>
      <span className="isopilot-theme-toggle-track relative inline-flex h-5 w-9 rounded-full transition-colors duration-200">
        <span
          className={`isopilot-theme-toggle-thumb absolute top-0.5 size-4 rounded-full transition-transform duration-200 ${isDark ? "translate-x-[18px]" : "translate-x-0.5"}`}
        />
      </span>
    </button>
  );
}
