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

# Keep auth mail branding aligned with the product.
auth_service="pkg/iam/auth_service.go"
old_magic_sender='magicLinkDefaultSenderName = "Probo"'
new_magic_sender='magicLinkDefaultSenderName = "ISO Pilot"'
if grep -Fq "${old_magic_sender}" "${auth_service}"; then
  sed -i "s/${old_magic_sender}/${new_magic_sender}/" "${auth_service}"
fi
grep -Fq "${new_magic_sender}" "${auth_service}"

# Preserve compatibility with existing passwords at login while requiring a
# stronger baseline whenever a password is newly created, changed, or reset.
sed -i 's/v.Check(req.Password, "password", PasswordValidator())/v.Check(req.Password, "password", StrongPasswordValidator())/g' "${auth_service}"
sed -i 's/v.Check(req.NewPassword, "newPassword", PasswordValidator())/v.Check(req.NewPassword, "newPassword", StrongPasswordValidator())/' "${auth_service}"

grep -Fq 'v.Check(password, "password", PasswordValidator())' "${auth_service}"
test "$(grep -Fc 'v.Check(req.Password, "password", StrongPasswordValidator())' "${auth_service}")" -ge 2
grep -Fq 'v.Check(req.NewPassword, "newPassword", StrongPasswordValidator())' "${auth_service}"
grep -Fq 'func StrongPasswordValidator()' pkg/iam/validators.go

# Match the frontend to the server's 12-character policy for all new passwords.
for password_page_file in \
  apps/console/src/pages/iam/auth/SignUpPage.tsx \
  apps/console/src/pages/iam/auth/ResetPasswordPage.tsx \
  apps/console/src/pages/iam/auth/CreatePasswordPage.tsx; do
  sed -i 's/z.string().min(8)/z.string().min(12)/g' "${password_page_file}"
  grep -Fq 'z.string().min(12)' "${password_page_file}"
done

# Email-backed magic links are available again, but the public send endpoint is
# always wrapped in the server-side limiter from the ISO Pilot overlay.
connect_resolver="pkg/server/api/connect/v1/resolver.go"
old_magic_link_route='r.Post("/magic-link/send", magicLinkHandler.SendHandler)'
limited_magic_link_route='r.With(magicLinkRateLimitMiddleware).Post("/magic-link/send", magicLinkHandler.SendHandler)'
if grep -Fq "${old_magic_link_route}" "${connect_resolver}"; then
  sed -i 's#r.Post("/magic-link/send", magicLinkHandler.SendHandler)#r.With(magicLinkRateLimitMiddleware).Post("/magic-link/send", magicLinkHandler.SendHandler)#' "${connect_resolver}"
fi
grep -Fq "${limited_magic_link_route}" "${connect_resolver}"
grep -Fq 'r.Get("/magic-link/verify", magicLinkHandler.VerifyHandler)' "${connect_resolver}"

# Production auth UI/runtime guardrails. SSO remains dormant until an operator
# configures an identity provider; email/password, verification, recovery and
# magic-link flows are the supported public entry points.
sign_in_page="apps/console/src/pages/iam/auth/sign-in/SignInPage.tsx"
password_page="apps/console/src/pages/iam/auth/sign-in/PasswordSignInPage.tsx"
web_server="pkg/server/web/web.go"

grep -Fq 'MagicLinkForm' "${sign_in_page}"
grep -Fq '/auth/register' "${sign_in_page}"
grep -Fq '/auth/register' "${password_page}"
grep -Fq '/auth/forgot-password' "${password_page}"
! grep -Fq '/auth/sso-login' "${sign_in_page}"
grep -Fq 'EMAIL_NOT_VERIFIED' "${password_page}"

test -f pkg/server/api/connect/v1/public_auth_rate_limiter.go
grep -Fq 'AroundOperations(publicAuthRateLimitOperations)' pkg/server/api/connect/v1/graphql_handler.go
grep -Fq 'magicLinkRateLimitMiddleware' pkg/server/api/connect/v1/public_auth_rate_limiter.go

grep -Fq 'PROBOD_AUTH_COOKIE_SECURE="true"' entrypoint.sh
grep -Fq 'PROBOD_AUTH_COOKIE_SAMESITE="lax"' entrypoint.sh
grep -Fq 'https://app.isopilot.co.uk' entrypoint.sh
grep -Fq 'SMTP configuration is incomplete' entrypoint.sh
grep -Fq 'PROBOD_SMTP_TLS_REQUIRED="true"' entrypoint.sh
grep -Fq 'Strict-Transport-Security' "${web_server}"

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
