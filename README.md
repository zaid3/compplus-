# ISOpilot

**Your compliance co-pilot.**

ISOpilot is a guided governance, risk and compliance platform that helps organisations turn compliance requirements into practical documents, controls, risks, tasks, evidence and audit work.

It keeps the proven Probo architecture and adds an ISOpilot experience for organisations that want one workspace for information security, quality, privacy, environmental and AI governance.

## Fast Start coverage

ISOpilot includes guided working packs for:

- ISO/IEC 27001:2022, including the 2024 climate-action amendment
- ISO 9001:2015 + Amendment 1:2024, the currently published ISO 9001 requirements
- ISO 14001:2026
- ISO/IEC 42001:2023
- UK GDPR and current UK data-protection changes, including the Data (Use and Access) Act 2025 provisions in force in 2026
- a shared Core management-system pack used across standards so common work is completed once

The underlying Probo framework engine also supports additional frameworks such as SOC 2, EU GDPR, ISO/IEC 27701, HIPAA, CCPA/CPRA, NIS2, DORA, FERPA, PCI DSS, HDS and 21 CFR Part 11.

## What a new organisation does

1. Create an account and organisation.
2. Describe the organisation in plain English: services, systems, team, processes and customers.
3. Choose the standards or regulations that apply.
4. Install the relevant Fast Start packs.
5. Review the prepared policies, procedures, registers, plans and forms. Confirm organisation-specific decisions instead of blindly accepting template text.
6. Review risks, framework requirements and measures; assign accountable owners.
7. Attach real evidence such as approvals, training records, supplier reviews, screenshots, logs and completed forms.
8. Run internal audits, record findings and complete corrective actions.
9. Hold management review, approve controlled documents and keep evidence current.

ISOpilot helps organise and demonstrate a management system. It does not automatically certify an organisation, replace competent legal or certification advice, or make a template true simply because it exists.

## How the workspace fits together

| Area | Plain-English meaning |
| --- | --- |
| Frameworks | The requirements you are working against. |
| Documents | Policies, procedures, registers, plans, forms and records describing or recording how the organisation works. |
| Risks | Things that could stop the organisation meeting its objectives or obligations, plus the actions used to reduce them. |
| Measures | Practical controls or activities used to meet requirements and reduce risk. |
| Tasks | Work that a named person needs to complete. |
| Evidence | Proof that a policy, control or process actually operates in real life. |
| Audits | Independent checks that requirements and the organisation's own arrangements are being followed and are effective. |
| Management review | Senior-management review of performance, risks, changes, resources, findings and improvement decisions. |

## Architecture

ISOpilot deliberately preserves Probo's Go, PostgreSQL, GraphQL, React/Relay and GRC domain architecture. ISOpilot-specific changes are maintained as an overlay and build layer so upstream security and product improvements can continue to be reviewed and adopted without rebuilding the platform from scratch.

The production build is pinned to a reviewed Probo commit for reproducibility. GitHub Actions builds the complete source image, while Coolify downloads the prebuilt runtime to keep deployment resource usage predictable.

## Multi-company use

The platform is designed for multiple independent organisations to create accounts and maintain separate workspaces. Organisation-scoped authorisation remains part of the Probo architecture, and ISOpilot's Fast Start framework installer resolves imported frameworks inside the selected organisation.

For a public SaaS deployment, use HTTPS, secure host-only cookies, exact CORS origins, SMTP for account recovery and invitations, generated persistent secrets, backups and restore testing. See `docs/ISOPILOT_PRODUCTION_RELEASE.md`.

## Branding

Customer-facing screens use the ISOpilot brand: Signal Blue actions, deep-ink typography, Inter/Inter Tight/IBM Plex Mono, consistent status semantics, keyboard focus states and the guiding co-pilot voice. Internal upstream package names and required MIT copyright notices are intentionally preserved for compatibility and attribution.

## Upstream and licence

ISOpilot is built on the open-source [Probo](https://github.com/getprobo/probo) project and retains the upstream MIT licence and copyright notices as required. Upstream source remains attributable to its original authors; ISOpilot-specific branding, templates, onboarding and deployment customisations live in this repository.

See [LICENSE](LICENSE) for licence terms.
