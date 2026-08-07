import type { ComponentProps } from "react";

export type ProboLogoProps = ComponentProps<"svg">;

export function ProboLogo(props: ProboLogoProps) {
  return (
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 142 32" role="img" aria-label="ISOPilot" fill="none" {...props}>
      <rect x="0" y="0" width="32" height="32" rx="8" fill="#2F6BFF" />
      <path
        d="M8.4 17.05 24.9 8.55c.83-.43 1.69.43 1.27 1.26l-8.5 16.49c-.43.84-1.67.7-1.9-.21l-1.4-5.6-5.59-1.4c-.91-.23-1.05-1.47-.38-2.04Z"
        fill="#FFFFFF"
      />
      <path d="m14.37 20.49 4.1-4.1" stroke="#DCE7FF" strokeWidth="1.5" strokeLinecap="round" />
      <text
        x="42"
        y="21.2"
        fill="currentColor"
        fontFamily="Inter Tight, Inter, ui-sans-serif, system-ui, sans-serif"
        fontSize="19"
        fontWeight="700"
        letterSpacing="-0.55"
      >
        ISOPilot
      </text>
    </svg>
  );
}
