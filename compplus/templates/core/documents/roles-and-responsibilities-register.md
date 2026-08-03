# Roles and Responsibilities Register

**Organisation:** {{ organization.legal_name }}  
**Register owner:** {{ roles.executive_owner }}  
**Next review:** {{ document.next_review_date }}

Use this register to confirm who owns each part of the management system. One person may hold several roles in a small organisation. Add a deputy where the role is critical.

| Role | Primary owner | Deputy | Main responsibilities | Authority | Review date |
|---|---|---|---|---|---|
| Executive owner | {{ roles.executive_owner }} | {{ roles.executive_deputy | confirm }} | Approve policy, resources, risk acceptance and management-review decisions. | Approve major management-system decisions and exceptions. | {{ document.next_review_date }} |
| Management-system manager | {{ roles.system_manager }} | {{ roles.system_manager_deputy | confirm }} | Coordinate templates, actions, evidence, audits, reviews and improvements. | Assign work and escalate overdue or high-risk issues. | {{ document.next_review_date }} |
| Information-security owner | {{ roles.security_owner }} | {{ roles.security_deputy | confirm }} | Coordinate security risks, controls, incidents and awareness. | Require treatment and escalate unacceptable security risk. | {{ document.next_review_date }} |
| Privacy owner | {{ roles.privacy_owner }} | {{ roles.privacy_deputy | confirm }} | Coordinate processing records, rights requests, DPIAs, breaches and privacy advice. | Stop or escalate processing with unmanaged privacy risk. | {{ document.next_review_date }} |
| Quality owner | {{ roles.quality_owner }} | {{ roles.quality_deputy | confirm }} | Coordinate process performance, complaints, nonconformities and improvement. | Require corrective action for quality failures. | {{ document.next_review_date }} |
| Environmental owner | {{ roles.environment_owner }} | {{ roles.environment_deputy | confirm }} | Coordinate environmental aspects, obligations, objectives, incidents and monitoring. | Escalate significant environmental risk or non-compliance. | {{ document.next_review_date }} |
| AI governance owner | {{ roles.ai_owner }} | {{ roles.ai_deputy | confirm }} | Coordinate AI inventory, impact assessments, risks, oversight and monitoring. | Pause or escalate AI use that exceeds approved risk. | {{ document.next_review_date }} |

## Approval

**Confirmed by:** {{ roles.executive_owner }}  
**Confirmation date:** {{ approval.date | confirm }}