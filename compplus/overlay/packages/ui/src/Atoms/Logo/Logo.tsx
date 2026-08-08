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

type Props = {
  className?: string;
  withPicto?: boolean;
};

function PilotMark({ size = 28 }: { size?: number }) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 32 32"
      width={size}
      height={size}
      role="img"
      aria-label="ISOPilot"
      fill="none"
      className="shrink-0"
    >
      <rect width="32" height="32" rx="8" fill="#2F6BFF" />
      <path
        d="M8.4 17.05 24.9 8.55c.83-.43 1.69.43 1.27 1.26l-8.5 16.49c-.43.84-1.67.7-1.9-.21l-1.4-5.6-5.59-1.4c-.91-.23-1.05-1.47-.38-2.04Z"
        fill="#FFFFFF"
      />
      <path d="m14.37 20.49 4.1-4.1" stroke="#DCE7FF" strokeWidth="1.5" strokeLinecap="round" />
    </svg>
  );
}

export function Logo({ className, withPicto }: Props) {
  if (withPicto) {
    return (
      <span className={clsx("inline-flex items-center gap-2.5 text-txt-primary", className)} aria-label="ISOPilot">
        <PilotMark />
        <span
          className="whitespace-nowrap text-[19px] font-bold leading-none tracking-[-0.55px]"
          style={{ fontFamily: '"Inter Tight", "Inter", ui-sans-serif, system-ui, sans-serif' }}
        >
          ISOPilot
        </span>
      </span>
    );
  }

  return (
    <span
      className={clsx("inline-flex whitespace-nowrap text-[17px] font-bold leading-none tracking-[-0.45px] text-txt-primary", className)}
      style={{ fontFamily: '"Inter Tight", "Inter", ui-sans-serif, system-ui, sans-serif' }}
      aria-label="ISOPilot"
    >
      ISOPilot
    </span>
  );
}
