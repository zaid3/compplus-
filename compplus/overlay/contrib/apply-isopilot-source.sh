#!/usr/bin/env bash
# Copyright (c) 2026 Probo Inc <hello@probo.com>.
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in
# all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

set -Eeuo pipefail

# Apply the small source-level adaptations that must be made to pinned upstream
# Probo after the ISO Pilot overlay is copied in. Each replacement is guarded so
# an upstream change fails loudly instead of silently producing a weaker build.

measure_file="pkg/probo/measure_service.go"
old_framework_lookup='framework.LoadByReferenceID(ctx, tx, scope, standard.Framework)'
new_framework_lookup='framework.LoadByOrganizationIDAndReferenceID(ctx, tx, scope, organization.ID, standard.Framework)'

if grep -Fq "${old_framework_lookup}" "${measure_file}"; then
  sed -i "s#${old_framework_lookup}#${new_framework_lookup}#" "${measure_file}"
fi

grep -Fq "${new_framework_lookup}" "${measure_file}"
grep -Eq 'MeasureID:[[:space:]]*measure.ID' "${measure_file}"

iso14001_file="compplus/templates/pack_iso14001.go"
biodiversity='Biodiversity, ecosystems, habitats and land-use impacts where relevant'
if ! grep -Fq "${biodiversity}" "${iso14001_file}"; then
  sed -i \
    's/"Water use where relevant", /"Water use where relevant", "Biodiversity, ecosystems, habitats and land-use impacts where relevant", /' \
    "${iso14001_file}"
fi

grep -Fq "${biodiversity}" "${iso14001_file}"

# Rebrand customer-facing locale VALUES only. Never mutate JSON keys: React i18n
# lookups depend on stable keys such as signInPage.newToProbo.
python3 contrib/rebrand-locales.py

grep -Fq '"newToProbo"' apps/console/src/_locales/en-US.json
grep -Fq 'ISO Pilot' apps/console/src/_locales/en-US.json

# Keep internal/source identifiers using the compact ISOPilot spelling where
# changing them to a spaced product name could make code invalid.
find compplus -type f \( -name '*.go' -o -name '*.ts' -o -name '*.tsx' -o -name '*.json' -o -name '*.md' \) \
  -exec sed -i 's/Comp Plus+/ISOPilot/g; s/ISOpilot/ISOPilot/g' {} +
find pkg/probo pkg/server/api/console -type f \( -name '*.go' -o -name '*.graphql' \) \
  -exec sed -i 's/Comp Plus+/ISOPilot/g; s/ISOpilot/ISOPilot/g' {} +

# Keep future auth mail branding aligned with the product even when SMTP is
# enabled later.
auth_service="pkg/iam/auth_service.go"
old_magic_sender='magicLinkDefaultSenderName = "Probo"'
new_magic_sender='magicLinkDefaultSenderName = "ISO Pilot"'
if grep -Fq "${old_magic_sender}" "${auth_service}"; then
  sed -i "s/${old_magic_sender}/${new_magic_sender}/" "${auth_service}"
fi
grep -Fq "${new_magic_sender}" "${auth_service}"

# Magic-link delivery requires an operator-provided SMTP service. The current
# production deployment has no SMTP credentials, so do not expose a public send
# endpoint that can only queue undeliverable email. The verification endpoint is
# retained so previously issued links would remain valid if one existed.
connect_resolver="pkg/server/api/connect/v1/resolver.go"
magic_link_send='r.Post("/magic-link/send", magicLinkHandler.SendHandler)'
if grep -Fq "${magic_link_send}" "${connect_resolver}"; then
  sed -i '\#r.Post("/magic-link/send", magicLinkHandler.SendHandler)#d' "${connect_resolver}"
fi
! grep -Fq "${magic_link_send}" "${connect_resolver}"
grep -Fq 'r.Get("/magic-link/verify", magicLinkHandler.VerifyHandler)' "${connect_resolver}"

# Append ISO Pilot visual tokens after upstream theme declarations so the
# existing component system keeps working while the product uses the blue brand
# palette in both light and dark mode.
theme_file="packages/ui/src/theme.css"
theme_overrides="packages/ui/src/isopilot-theme.css"
theme_marker="ISO Pilot production theme overrides"

test -f "${theme_overrides}"
if ! grep -Fq "${theme_marker}" "${theme_file}"; then
  printf '\n' >> "${theme_file}"
  cat "${theme_overrides}" >> "${theme_file}"
fi

grep -Fq "${theme_marker}" "${theme_file}"

echo "ISO Pilot source customizations applied successfully"
