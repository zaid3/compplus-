#!/bin/bash
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

set -euo pipefail

# This image is the production ISO Pilot runtime. Keep authentication bound to
# the canonical HTTPS origin even if an old hosting environment still contains
# legacy sslip.io variables from the initial deployment.
ISOPILOT_CANONICAL_URL="${ISOPILOT_CANONICAL_URL:-https://app.isopilot.co.uk}"

case "${ISOPILOT_CANONICAL_URL}" in
  https://*) ;;
  *)
    echo "Error: ISOPILOT_CANONICAL_URL must use https:// in production" >&2
    exit 1
    ;;
esac

export PROBOD_BASE_URL="${ISOPILOT_CANONICAL_URL}"
export PROBOD_API_CORS_ALLOWED_ORIGINS="${ISOPILOT_CANONICAL_URL}"
export PROBOD_AUTH_COOKIE_SECURE="true"
export PROBOD_AUTH_COOKIE_SAMESITE="lax"
export PROBOD_MAILER_SENDER_NAME="ISO Pilot"

# Public self-registration is disabled for the production launch. It can be
# deliberately enabled later only after outbound email/verification is ready.
if [ "${ISOPILOT_ENABLE_PUBLIC_SIGNUP:-false}" = "true" ]; then
  export PROBOD_AUTH_DISABLE_SIGNUP="false"
else
  export PROBOD_AUTH_DISABLE_SIGNUP="true"
fi

# Configuration file path
CONFIG_FILE="${CONFIG_FILE:-/etc/probod/config.yml}"

# When PROBOD_ENCRYPTION_KEY is set, always (re)generate the config from env vars.
# This includes literal values and aws:// / awssm:// / awsps:// AWS references.
# probod-bootstrap reads every PROBOD_* var; the entrypoint only checks this one
# to decide whether to run it, including when a stale config file exists on a
# persistent volume. When it is unset, fall back to an existing config file
# (e.g., mounted from a ConfigMap).
if [ -n "${PROBOD_ENCRYPTION_KEY:-}" ]; then
  echo "Generating configuration file from environment variables at: $CONFIG_FILE"
  probod-bootstrap -output "$CONFIG_FILE"
elif [ -f "$CONFIG_FILE" ]; then
  echo "Using existing configuration file at: $CONFIG_FILE"
else
  echo "Error: PROBOD_ENCRYPTION_KEY is unset and no config file found at $CONFIG_FILE" >&2
  exit 1
fi

# Execute probod with the generated config
exec probod -cfg-file "$CONFIG_FILE" "$@"
