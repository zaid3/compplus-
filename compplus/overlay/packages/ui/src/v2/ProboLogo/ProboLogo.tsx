import type { ComponentProps } from "react";

export type ProboLogoProps = ComponentProps<"svg">;

export function ProboLogo(props: ProboLogoProps) {
  return (
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 148 28" role="img" aria-label="Comp Plus+" fill="none" {...props}>
      <path d="M14 1.75 24.5 5.6v7.12c0 6.45-4.34 10.72-10.5 13.53C7.84 23.44 3.5 19.17 3.5 12.72V5.6L14 1.75Z" fill="currentColor" />
      <path d="M9.45 13.95 12.4 16.9l6.35-7" stroke="white" strokeWidth="2.35" strokeLinecap="round" strokeLinejoin="round" />
      <path d="M21.9 18.15h4.6m-2.3-2.3v4.6" stroke="currentColor" strokeWidth="2.1" strokeLinecap="round" />
      <text x="34" y="19.1" fill="currentColor" fontFamily="Inter, ui-sans-serif, system-ui, sans-serif" fontSize="17" fontWeight="700" letterSpacing="-0.45">Comp Plus+</text>
    </svg>
  );
}
