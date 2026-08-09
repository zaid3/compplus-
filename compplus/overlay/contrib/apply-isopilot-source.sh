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

# Rebrand customer-facing locale strings only. Internal package names, API names,
# database identifiers and upstream copyright notices remain unchanged.
find apps -type f -path '*/_locales/*.json' \
  -exec sed -i 's/Probo/ISO Pilot/g; s/Comp Plus+/ISO Pilot/g; s/ISOpilot/ISO Pilot/g; s/ISOPilot/ISO Pilot/g' {} +

# Keep internal/source identifiers using the compact ISOPilot spelling where
# changing them to a spaced product name could make code invalid.
find compplus -type f \( -name '*.go' -o -name '*.ts' -o -name '*.tsx' -o -name '*.json' -o -name '*.md' \) \
  -exec sed -i 's/Comp Plus+/ISOPilot/g; s/ISOpilot/ISOPilot/g' {} +
find pkg/probo pkg/server/api/console -type f \( -name '*.go' -o -name '*.graphql' \) \
  -exec sed -i 's/Comp Plus+/ISOPilot/g; s/ISOpilot/ISOPilot/g' {} +

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
