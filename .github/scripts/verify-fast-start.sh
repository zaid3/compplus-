#!/usr/bin/env bash
set -Eeuo pipefail

API_BASE="${ISOPILOT_API_BASE:-http://127.0.0.1:8090}"
COOKIE_JAR="$(mktemp /tmp/isopilot-fast-start-cookie.XXXXXX)"
CURRENT_PACK="bootstrap"
CURRENT_PHASE="startup"
LAST_RESPONSE=""

cleanup() {
  rm -f "${COOKIE_JAR}"
}

diagnostics() {
  local exit_code=$?
  echo
  echo "=== FAST START VERIFICATION FAILED ==="
  echo "phase=${CURRENT_PHASE}"
  echo "pack=${CURRENT_PACK}"
  echo "exit_code=${exit_code}"
  if [[ -n "${LAST_RESPONSE}" ]]; then
    echo "=== LAST GRAPHQL RESPONSE ==="
    jq . <<<"${LAST_RESPONSE}" 2>/dev/null || printf '%s\n' "${LAST_RESPONSE}"
  fi
  if [[ -f compose.source-smoke.yaml ]]; then
    echo "=== SOURCE STACK STATE ==="
    docker compose -f compose.source-smoke.yaml ps -a || true
    echo "=== APPLICATION LOGS ==="
    docker compose -f compose.source-smoke.yaml logs --no-color --tail=400 probo || true
  fi
  exit "${exit_code}"
}

trap cleanup EXIT
trap diagnostics ERR

graphql() {
  local endpoint="$1"
  local payload="$2"
  curl -fsS --retry 4 --retry-all-errors --retry-delay 2 --max-time 45 \
    -c "${COOKIE_JAR}" -b "${COOKIE_JAR}" \
    -H 'Content-Type: application/json' \
    --data "${payload}" \
    "${API_BASE}${endpoint}"
}

assert_no_graphql_errors() {
  local response="$1"
  jq -e '(.errors // null) == null' <<<"${response}" >/dev/null
}

assert_nonnegative_counts() {
  local response="$1"
  jq -e '
    .data.installTemplatePack as $p
    | ["documentsCreated", "measuresCreated", "tasksCreated", "evidenceRequestsCreated"]
    | all(. as $field |
        ($p[$field] | type) == "number"
        and ($p[$field] | floor) == $p[$field]
        and $p[$field] >= 0)
  ' <<<"${response}" >/dev/null
}

assert_install_result() {
  local response="$1"
  local pack="$2"
  local expected_already_installed="$3"

  jq -e \
    --arg pack "${pack}" \
    --argjson expected "${expected_already_installed}" \
    '
      (.errors // null) == null
      and (.data.installTemplatePack != null)
      and (.data.installTemplatePack.packId == $pack)
      and (.data.installTemplatePack.framework.id != null)
      and ((.data.installTemplatePack.framework.name | type) == "string")
      and (.data.installTemplatePack.alreadyInstalled == $expected)
    ' <<<"${response}" >/dev/null

  assert_nonnegative_counts "${response}"

  if [[ "${pack}" == "iso27001" && "${expected_already_installed}" == "false" ]]; then
    jq -e '.data.installTemplatePack.statementOfApplicabilityCreated == true' <<<"${response}" >/dev/null
  fi
}

CURRENT_PHASE="wait-for-http"
ready=0
for attempt in $(seq 1 36); do
  code="$(curl -sS -o /tmp/isopilot-source-body -w '%{http_code}' --max-time 5 "${API_BASE}/" || true)"
  if [[ "${code}" =~ ^[0-9]+$ ]] && (( code >= 200 && code < 400 )); then
    ready=1
    break
  fi
  echo "Waiting for source-built app (${attempt}/36), HTTP=${code:-none}"
  sleep 5
done
[[ "${ready}" == "1" ]]

CURRENT_PHASE="signup"
email="ci-${GITHUB_RUN_ID:-local}-${GITHUB_RUN_ATTEMPT:-0}-$(date +%s)-${RANDOM}@example.com"
password='ISOPilot-CI-Test-123!'
signup_query='mutation SignUpPageMutation($input: SignUpInput!) { signUp(input: $input) { identity { id } } }'
signup_payload="$(jq -nc --arg q "${signup_query}" --arg email "${email}" --arg password "${password}" '{operationName:"SignUpPageMutation",query:$q,variables:{input:{email:$email,password:$password,fullName:"ISOPilot CI"}}}')"
LAST_RESPONSE="$(graphql '/api/connect/v1/graphql' "${signup_payload}")"
assert_no_graphql_errors "${LAST_RESPONSE}"
jq -e '.data.signUp.identity.id != null' <<<"${LAST_RESPONSE}" >/dev/null

CURRENT_PHASE="create-organization"
create_query='mutation NewOrganizationPageMutation($input: CreateOrganizationInput!) { createOrganization(input: $input) { organization { id name } } }'
create_payload="$(jq -nc --arg q "${create_query}" '{operationName:"NewOrganizationPageMutation",query:$q,variables:{input:{name:"ISOPilot CI Organisation"}}}')"
LAST_RESPONSE="$(graphql '/api/connect/v1/graphql' "${create_payload}")"
assert_no_graphql_errors "${LAST_RESPONSE}"
organization_id="$(jq -er '.data.createOrganization.organization.id' <<<"${LAST_RESPONSE}")"

CURRENT_PHASE="assume-organization"
assume_query='mutation AssumePageMutation($input: AssumeOrganizationSessionInput!) { assumeOrganizationSession(input: $input) { result { __typename ... on OrganizationSessionCreated { session { id } membership { id } } ... on PasswordRequired { reason } ... on SAMLAuthenticationRequired { reason } } } }'
assume_continue="/organizations/${organization_id}"
assume_payload="$(jq -nc --arg q "${assume_query}" --arg oid "${organization_id}" --arg continue_url "${assume_continue}" '{operationName:"AssumePageMutation",query:$q,variables:{input:{organizationId:$oid,"continue":$continue_url}}}')"
LAST_RESPONSE="$(graphql '/api/connect/v1/graphql' "${assume_payload}")"
assert_no_graphql_errors "${LAST_RESPONSE}"
jq -e '.data.assumeOrganizationSession.result.__typename == "OrganizationSessionCreated"' <<<"${LAST_RESPONSE}" >/dev/null

install_query='mutation TemplateLibraryDialogInstallMutation($input: InstallTemplatePackInput!) { installTemplatePack(input: $input) { packId version documentsCreated measuresCreated tasksCreated evidenceRequestsCreated statementOfApplicabilityCreated alreadyInstalled framework { id name } } }'
packs=(core iso27001 iso9001 uk-gdpr iso14001 iso42001)

make_install_payload() {
  local pack="$1"
  jq -nc \
    --arg q "${install_query}" \
    --arg oid "${organization_id}" \
    --arg pack "${pack}" \
    '{operationName:"TemplateLibraryDialogInstallMutation",query:$q,variables:{input:{organizationId:$oid,packId:$pack,answers:{legalName:"ISOPilot CI Organisation",services:"Compliance testing",locations:"United Kingdom",executiveOwner:"CI Executive",systemManager:"CI Compliance",securityOwner:"CI Security",privacyOwner:"CI Privacy",qualityOwner:"CI Quality",environmentalOwner:"CI Environment",aiOwner:"CI AI"}}}}'
}

for pack in "${packs[@]}"; do
  CURRENT_PACK="${pack}"
  CURRENT_PHASE="first-install"
  echo "=== FIRST INSTALL: ${pack} ==="
  payload="$(make_install_payload "${pack}")"
  LAST_RESPONSE="$(graphql '/api/console/v1/graphql' "${payload}")"
  echo "${LAST_RESPONSE}" | jq .
  assert_install_result "${LAST_RESPONSE}" "${pack}" false
  echo "PASS first-install ${pack}"
done

for pack in "${packs[@]}"; do
  CURRENT_PACK="${pack}"
  CURRENT_PHASE="idempotent-retry"
  echo "=== IDEMPOTENT RETRY: ${pack} ==="
  payload="$(make_install_payload "${pack}")"
  LAST_RESPONSE="$(graphql '/api/console/v1/graphql' "${payload}")"
  echo "${LAST_RESPONSE}" | jq .
  assert_install_result "${LAST_RESPONSE}" "${pack}" true
  echo "PASS retry ${pack}"
done

CURRENT_PACK="all"
CURRENT_PHASE="complete"
echo "Fast Start verification passed for: ${packs[*]}"
