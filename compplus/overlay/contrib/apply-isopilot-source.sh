#!/usr/bin/env bash
set -Eeuo pipefail

# Apply the small source-level adaptations that must be made to pinned upstream
# Probo after the ISOPilot overlay is copied in. Each replacement is guarded so
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

# Rebrand customer-facing source strings only. Internal package names, API names,
# database identifiers and upstream copyright notices remain unchanged.
find apps -type f -path '*/_locales/*.json' \
  -exec sed -i 's/Probo/ISOPilot/g; s/Comp Plus+/ISOPilot/g; s/ISOpilot/ISOPilot/g' {} +
find compplus -type f \( -name '*.go' -o -name '*.ts' -o -name '*.tsx' -o -name '*.json' -o -name '*.md' \) \
  -exec sed -i 's/Comp Plus+/ISOPilot/g; s/ISOpilot/ISOPilot/g' {} +
find pkg/probo pkg/server/api/console -type f \( -name '*.go' -o -name '*.graphql' \) \
  -exec sed -i 's/Comp Plus+/ISOPilot/g; s/ISOpilot/ISOPilot/g' {} +

echo "ISOPilot source customizations applied successfully"
