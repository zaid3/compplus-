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
      aria-label="ISOpilot"
      fill="none"
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
      <span className={clsx("inline-flex items-center gap-2.5 text-[#0B1220]", className)} aria-label="ISOpilot">
        <PilotMark />
        <span
          className="text-[19px] leading-none font-bold tracking-[-0.55px]"
          style={{ fontFamily: '"Inter Tight", "Inter", ui-sans-serif, system-ui, sans-serif' }}
        >
          ISOpilot
        </span>
      </span>
    );
  }

  return (
    <span
      className={clsx("inline-flex text-[#0B1220] text-[17px] leading-none font-bold tracking-[-0.45px]", className)}
      style={{ fontFamily: '"Inter Tight", "Inter", ui-sans-serif, system-ui, sans-serif' }}
      aria-label="ISOpilot"
    >
      ISOpilot
    </span>
  );
}
