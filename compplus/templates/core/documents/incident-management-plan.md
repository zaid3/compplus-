# Integrated Incident Management Plan

**Plan owner:** {{ roles.system_manager }}  
**Executive contact:** {{ roles.executive_owner }}  
**Review and test cycle:** At least annually

## 1. Purpose

This plan helps {{ organization.legal_name }} respond quickly and consistently to security, privacy, service-quality, environmental and AI-related incidents.

## 2. Report immediately

Workers must report suspected incidents through:

{{ incident.reporting_channels | confirm }}

No worker should delay reporting while trying to prove the incident is genuine.

## 3. Response roles

- Incident lead: {{ incident.lead | confirm }}
- Technical or operational lead: {{ incident.technical_lead | confirm }}
- Privacy lead: {{ roles.privacy_owner }}
- Communications lead: {{ incident.communications_lead | confirm }}
- Executive decision-maker: {{ roles.executive_owner }}
- Legal, insurance or specialist contacts: {{ incident.external_contacts | confirm }}

## 4. Severity and escalation

{{ incident.escalation | confirm }}

A critical incident must be escalated immediately where there is serious harm, major service interruption, significant legal or contractual exposure, loss of sensitive information, environmental harm or unsafe AI behaviour.

## 5. Response process

1. **Receive and record** — open an incident record and capture reporter, time, systems, people and initial facts.
2. **Triage** — classify type, severity, affected assets or people and immediate risk.
3. **Contain** — prevent further harm while preserving evidence.
4. **Investigate** — establish timeline, cause, scope and impact.
5. **Notify and communicate** — meet legal, regulatory, contractual and customer notification requirements.
6. **Recover** — restore safe service and confirm controls are operating.
7. **Review** — identify lessons, risks and corrective actions.
8. **Close** — obtain approval that actions and evidence are complete.

## 6. Required incident record

The register must contain facts, decisions, timestamps, contacts, affected information or services, containment, notifications, recovery, costs, evidence and corrective actions.

## 7. Playbooks

Select at least three relevant playbooks:

{{ incident.playbooks | confirm }}

Suggested starting playbooks: phishing or account compromise, ransomware or malware, personal-data breach, cloud outage, supplier failure, environmental spill, serious quality failure and harmful AI output.

## 8. Regulatory considerations

The incident lead must check applicable notification deadlines immediately. For a potentially reportable UK personal-data breach, the privacy owner starts the assessment and breach log without waiting for every fact to be known.

## 9. Testing

The plan is tested through a tabletop or simulation. Record participants, scenario, decisions, lessons and actions:

{{ incident.test_record | evidence }}
