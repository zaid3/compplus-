# Integrated Internal Audit Programme

**Organisation:** {{ organization.legal_name }}  
**Programme owner:** {{ roles.system_manager }}  
**Audit cycle:** {{ audit.cycle | confirm }}

CompPlus proposes audit areas from the selected packs and risk information. Confirm dates, scope and auditor independence. Findings and corrective actions should be recorded directly in CompPlus.

| Audit area | Standards or obligations | Risk basis | Scope | Auditor | Planned date | Status | Report |
|---|---|---|---|---|---|---|---|
| Management-system governance | All selected packs | Annual management-system review | Context, scope, leadership, roles, objectives, document control, audit, review and improvement | {{ audit.governance_auditor | confirm }} | {{ audit.governance_date | confirm }} | Planned | Add after completion |
| Operational and technical controls | Selected operational packs | Highest residual risks and changed processes | {{ audit.operational_scope | confirm }} | {{ audit.operational_auditor | confirm }} | {{ audit.operational_date | confirm }} | Planned | Add after completion |
| Legal, privacy and supplier obligations | UK GDPR and applicable contracts | Personal data, critical suppliers and legal change | Processing records, DPIAs, rights, breaches, contracts, supplier monitoring and retention | {{ audit.privacy_auditor | confirm }} | {{ audit.privacy_date | confirm }} | Planned | Add after completion |
| {{ audit.custom_area | confirm }} | {{ audit.custom_criteria | confirm }} | {{ audit.custom_risk_basis | confirm }} | {{ audit.custom_scope | confirm }} | {{ audit.custom_auditor | confirm }} | {{ audit.custom_date | confirm }} | Planned | Add after completion |

## Audit report template

### Audit objective and scope

{{ audit.report_objective_scope | confirm }}

### Criteria and documents reviewed

{{ audit.report_criteria | confirm }}

### People interviewed and samples tested

{{ audit.report_samples | confirm }}

### Evidence and observations

{{ audit.report_evidence | confirm }}

### Conformities and effective practices

{{ audit.report_conformities | confirm }}

### Major nonconformities

{{ audit.report_major_findings | confirm }}

### Minor nonconformities

{{ audit.report_minor_findings | confirm }}

### Observations and improvement opportunities

{{ audit.report_opportunities | confirm }}

### Conclusion

{{ audit.report_conclusion | confirm }}

### Corrective actions and deadlines

{{ audit.report_actions | confirm }}