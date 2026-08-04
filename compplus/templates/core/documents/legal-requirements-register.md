# Legal, Regulatory and Contractual Requirements Register

**Organisation:** {{ organization.legal_name }}  
**Register owner:** {{ roles.system_manager }}  
**Next review:** {{ document.next_review_date }}

CompPlus has added starter entries from the selected packs. Confirm whether each entry applies and add any sector, customer, licence or contract obligations that are specific to the organisation.

| Requirement | Source | Why applicable | Owner | Actions or controls | Evidence | Review frequency | Last reviewed | Next review |
|---|---|---|---|---|---|---|---|---|
| UK data protection law | UK GDPR and Data Protection Act 2018, as amended | {{ organization.processes_personal_data | conditional }} | {{ roles.privacy_owner }} | Maintain processing records, lawful bases, rights handling, contracts, security, DPIAs and breach procedures. | Privacy records in CompPlus | At least annually and after legal change | {{ register.last_reviewed | confirm }} | {{ document.next_review_date }} |
| Customer and supplier contracts | Signed agreements and service terms | Contractual commitments affect service, security, privacy, quality or environmental performance. | {{ roles.system_manager }} | Record obligations, owners, evidence, renewal dates and required controls. | Contract register and linked documents | At renewal and annually | {{ register.last_reviewed | confirm }} | {{ document.next_review_date }} |
| Selected certification standards | Standards lawfully accessed by the organisation | {{ organization.selected_standards | confirm }} | {{ roles.system_manager }} | Maintain mapped frameworks, measures, evidence, audits and management reviews. | CompPlus framework and audit records | After a standards update and annually | {{ register.last_reviewed | confirm }} | {{ document.next_review_date }} |
| {{ legal.custom_requirement | confirm }} | {{ legal.custom_source | confirm }} | {{ legal.custom_applicability | confirm }} | {{ legal.custom_owner | confirm }} | {{ legal.custom_actions | confirm }} | {{ legal.custom_evidence | confirm }} | {{ legal.custom_frequency | confirm }} | {{ register.last_reviewed | confirm }} | {{ document.next_review_date }} |

## Review conclusion

**Changes identified:** {{ legal.review_changes | confirm }}  
**Actions created:** {{ legal.review_actions | confirm }}  
**Approved by:** {{ roles.executive_owner }}  
**Approval date:** {{ approval.date | confirm }}