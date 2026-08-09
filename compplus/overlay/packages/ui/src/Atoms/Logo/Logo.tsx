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
      aria-label="ISO Pilot"
      fill="none"
      className="shrink-0 drop-shadow-[0_4px_10px_rgba(47,107,255,0.22)]"
    >
      <defs>
        <linearGradient id="iso-pilot-mark" x1="5" y1="3" x2="28" y2="30" gradientUnits="userSpaceOnUse">
          <stop stopColor="#4D83FF" />
          <stop offset="1" stopColor="#245DE5" />
        </linearGradient>
      </defs>
      <rect width="32" height="32" rx="9" fill="url(#iso-pilot-mark)" />
      <path
        d="M8.4 17.05 24.9 8.55c.83-.43 1.69.43 1.27 1.26l-8.5 16.49c-.43.84-1.67.7-1.9-.21l-1.4-5.6-5.59-1.4c-.91-.23-1.05-1.47-.38-2.04Z"
        fill="#FFFFFF"
      />
      <path
        d="m14.37 20.49 4.1-4.1"
        stroke="#DCE7FF"
        strokeWidth="1.5"
        strokeLinecap="round"
      />
    </svg>
  );
}

export function Logo({ className, withPicto }: Props) {
  if (withPicto) {
    return (
      <span
        className={clsx("inline-flex items-center gap-2.5 text-txt-primary", className)}
        aria-label="ISO Pilot"
      >
        <PilotMark />
        <span
          className="whitespace-nowrap text-[19px] leading-none font-bold tracking-[-0.55px]"
          style={{ fontFamily: '"Inter Tight", "Inter", ui-sans-serif, system-ui, sans-serif' }}
        >
          ISO Pilot
        </span>
      </span>
    );
  }

  return (
    <span
      className={clsx(
        "inline-flex whitespace-nowrap text-txt-primary text-[17px] leading-none font-bold tracking-[-0.45px]",
        className,
      )}
      style={{ fontFamily: '"Inter Tight", "Inter", ui-sans-serif, system-ui, sans-serif' }}
      aria-label="ISO Pilot"
    >
      ISO Pilot
    </span>
  );
}
