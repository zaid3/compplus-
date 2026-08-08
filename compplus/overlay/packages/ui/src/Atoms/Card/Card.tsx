// Copyright (c) 2025-2026 Probo Inc <hello@probo.com>.
// MIT licensed upstream component, modified for the ISOPilot interface.

import type { CSSProperties, PropsWithChildren } from "react";
import { tv } from "tailwind-variants";

import { Slot } from "../Slot";

type Props = PropsWithChildren<{
  padded?: boolean;
  className?: string;
  style?: CSSProperties;
  asChild?: boolean;
}>;

const card = tv({
  base: "rounded-2xl border border-border-mid bg-level-1 shadow-sm",
  variants: {
    padded: {
      true: "p-5",
    },
  },
});

export function Card({
  padded = false,
  children,
  className,
  asChild,
  style,
}: Props) {
  const Component = asChild ? Slot : "div";
  return (
    <Component style={style} className={card({ padded, className })}>
      {children}
    </Component>
  );
}
