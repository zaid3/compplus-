# ISOpilot production release checklist

This is the minimum operational checklist for a public, multi-company ISOpilot deployment. It complements the application controls; it is not a certification checklist for any customer.

## 1. Public URL and transport security

Before public launch, set:

```text
ISOPILOT_PUBLIC_URL=https://<final-app-domain>
ISOPILOT_CORS_ORIGINS=https://<final-app-domain>
ISOPILOT_COOKIE_SECURE=true
```

Keep `ISOPILOT_COOKIE_DOMAIN` empty unless cross-subdomain cookie sharing is genuinely required. A host-only session cookie reduces unnecessary exposure.

Terminate TLS at the trusted Coolify proxy, keep HTTP-to-HTTPS redirection enabled, and do not expose PostgreSQL, SeaweedFS or the headless Chrome service directly to the public network.

## 2. Authentication and account recovery

Self-signup is enabled by default because ISOpilot is intended for multiple independent companies. If a private invitation-only deployment is required, set:

```text
ISOPILOT_DISABLE_SIGNUP=true
```

Configure a real SMTP service before launch so email confirmation, invitations, magic links and password-reset flows can work reliably:

```text
ISOPILOT_SMTP_ADDR=
ISOPILOT_SMTP_USER=
ISOPILOT_SMTP_PASSWORD=
ISOPILOT_SMTP_TLS_REQUIRED=true
ISOPILOT_MAILER_SENDER_EMAIL=no-reply@isopilot.co.uk
```

Google and Microsoft OIDC can be enabled later with the corresponding ISOPILOT client ID and secret variables. Empty values keep those providers disabled.

Smoke-test at least these journeys before opening public signup:

- new company A signs up and creates its organisation;
- new company B signs up and creates a different organisation;
- a company A member cannot read, search, mutate or import into company B resources;
- owner/member role restrictions work as expected;
- sign-out invalidates the session;
- password reset and email confirmation links expire as configured;
- invitations cannot be reused after acceptance or expiry.

## 3. Secrets and database privileges

Coolify-generated `SERVICE_PASSWORD_*` and `SERVICE_REALBASE64_*` values are used for cryptographic keys, PostgreSQL and object storage. Do not replace them with repository-known passwords.

The application database role is deliberately non-superuser. Do not grant `SUPERUSER`, `CREATEDB`, `CREATEROLE` or replication privileges to the runtime account.

The production Compose file contains a one-time migration path for the earlier pilot database credentials. After the hardened deployment succeeds, verify the legacy passwords no longer authenticate.

## 4. Backups and recovery

Back up at minimum:

- PostgreSQL data;
- object-storage data/evidence;
- persistent OAuth signing key volume;
- Coolify environment/secrets using an approved secure method.

Encrypt backups, restrict access, define retention, and test a complete restore before relying on the backup process. A backup that has never been restored is not proven recoverable.

## 5. Multi-company isolation

ISOpilot keeps Probo's organisation/tenant authorisation model. The Fast Start installer additionally resolves framework references inside the selected organisation.

Before every major release, run tenant-isolation regression tests against framework imports, documents, evidence, risks, tasks, people, vendors, audit data, compliance pages and any newly added GraphQL or MCP operations.

Treat any cross-organisation data exposure as a release blocker.

## 6. Framework and legal-content maintenance

Framework content has two layers:

1. the underlying Probo framework engine and upstream framework definitions;
2. ISOpilot Fast Start working packs for Core, ISO/IEC 27001, ISO 9001, ISO 14001, ISO/IEC 42001 and UK GDPR.

Do not silently relabel a draft or expected standard as published. As of the 2026-08-07 release baseline, ISO 9001 remains the published 2015 edition with Amendment 1:2024; update the Fast Start pack only after the new edition is officially published and its changes have been reviewed.

When standards, legislation or regulator guidance change:

- record the authoritative source and effective/publication date;
- review affected framework mappings, prompts, template wording and evidence expectations;
- version the changed pack;
- avoid overwriting a customer's approved document without review;
- explain to users what changed and what they need to reassess.

ISOpilot templates use original practical wording. Do not copy protected standards text into the repository unless licensing expressly permits it.

## 7. Compliance-document reality check

A template is a starting point, not evidence of conformity. Before a document is approved, the organisation should confirm that:

- its scope, owners, systems, services, suppliers and locations are correct;
- stated controls actually operate;
- legal and contractual obligations are applicable and current;
- risks reflect the real organisation;
- records and evidence prove the described activities;
- exceptions and non-applicable requirements are justified;
- document review and approval are recorded.

The exact set of documents varies by organisation and applicability. ISOpilot should guide the user to sufficient documented information without implying that every template title is universally mandatory.

## 8. Application and dependency release gate

A release must not be promoted when the full-source build fails.

The GitHub Actions release check verifies the reviewed upstream Probo pin, key ISOpilot frameworks/guidance, removal of known static production credentials, frontend compilation, generated GraphQL/Relay code, Go generation and final runtime binaries.

For each upstream Probo update, review security/authentication changes and any schema or API changes before moving the pin. Preserve the overlay model unless there is a strong reason to alter the architecture.

## 9. Branding and accessibility

Customer-facing screens must use ISOpilot rather than Probo/Comp Plus+ except where technical attribution or upstream package naming is intentionally retained.

Keep the brand tokens defined in the application stylesheet, use Lucide-style line icons, visible keyboard focus, readable contrast, and text/icon labels in addition to status colour.

## 10. Launch acceptance

Do not call a release production-ready until all of the following are true:

- full-source CI build passes on the exact release commit;
- final HTTPS domain and secure cookie settings are active;
- SMTP account recovery/invitation journeys pass;
- two-company isolation smoke tests pass;
- backup and restore test passes;
- no known critical/high dependency or application vulnerability is accepted without a documented treatment;
- the current framework/version labels have been checked against authoritative sources;
- key Fast Start packs can be installed into a clean organisation without duplicate or cross-organisation data;
- policy/document approval, evidence upload, risk, task, audit and export/PDF journeys work;
- monitoring/logging is available for production troubleshooting.
