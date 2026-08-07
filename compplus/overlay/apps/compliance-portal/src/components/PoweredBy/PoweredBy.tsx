import { ProboLogo } from "@probo/ui/src/v2/ProboLogo/ProboLogo";
import { Text } from "@probo/ui/src/v2/typography/Text";
import type { ReactNode } from "react";

import { dotPatternStyle } from "#/components/MediaTile/variants";

import { poweredBy } from "./variants";

export interface PoweredByProps {
  label?: ReactNode;
  href?: string;
}

export function PoweredBy({ label = "Powered by", href = "/" }: PoweredByProps) {
  const slots = poweredBy();

  return (
    <footer className={slots.root()}>
      <div className={slots.backdrop()} style={dotPatternStyle} />
      <div className={slots.backdropFade()} />
      <div className={slots.content()}>
        <Text size={1} color="neutral">{label}</Text>
        <a href={href} aria-label="Comp Plus+">
          <ProboLogo className={slots.logo()} aria-hidden />
        </a>
      </div>
    </footer>
  );
}
