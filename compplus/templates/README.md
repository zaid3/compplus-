# CompPlus Template Library

This directory contains the reusable, organisation-scoped template packs used by CompPlus.

## Goal

A customer should answer company questions once, then receive pre-filled policies, procedures, registers, tasks and evidence requests. The customer should mainly review highlighted decisions, attach evidence and approve.

## Design principles

1. **Ask once, reuse everywhere** — company profile answers are referenced by placeholders such as `{{ organization.name }}` and `{{ roles.security_owner }}`.
2. **One control, many standards** — a single implementation activity can map to ISO/IEC 27001, ISO 9001, ISO 14001, ISO/IEC 42001 and UK GDPR.
3. **Minimal user work** — every template identifies auto-filled fields, confirmation fields and evidence-required fields.
4. **Plain English first** — customer-facing wording is original CompPlus guidance, not copied ISO text.
5. **Auditor traceability** — mappings are metadata references only. Licensed standards wording must not be redistributed without permission.
6. **Versioned sources** — every pack records its standard edition, legal guidance date and source URLs.
7. **Organisation isolation** — installing a pack copies editable instances into one organisation; the master library remains read-only.

## Initial packs

- `core` — shared management-system templates used across all packs.
- `iso27001` — ISO/IEC 27001:2022 with Amendment 1:2024.
- `uk-gdpr` — current UK GDPR and ICO guidance, including Data (Use and Access) Act changes.
- `iso9001` — ISO 9001:2015 with Amendment 1:2024 until the 2026 edition is published.
- `iso14001` — ISO 14001:2026.
- `iso42001` — ISO/IEC 42001:2023.

## Template lifecycle

1. Customer selects standards.
2. CompPlus runs the organisation questionnaire.
3. Relevant templates are copied into the organisation.
4. Auto-fill placeholders are resolved.
5. Customer sees only unresolved confirmations and evidence requests.
6. Approved documents become controlled versions with owners and review dates.
7. Template-pack updates create migration suggestions; they never overwrite approved customer content automatically.

## Status

This branch establishes the library structure and content model. It does not alter the current production deployment. The next implementation step is a pack installer that creates Probo frameworks, controls, measures, documents, tasks and evidence requests from these files.