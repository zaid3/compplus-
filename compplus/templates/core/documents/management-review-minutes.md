# Integrated Management Review Minutes

**Organisation:** {{ organization.legal_name }}  
**Chair:** {{ roles.executive_owner }}  
**Meeting date:** {{ meeting.date | confirm }}  
**Next review:** {{ document.next_review_date }}

## Attendees

- {{ roles.executive_owner }}
- {{ roles.system_manager }}
- {{ roles.security_owner }}
- {{ roles.privacy_owner }}
- {{ roles.quality_owner }}
- {{ roles.environment_owner }}
- {{ roles.ai_owner }}
- {{ meeting.other_attendees | confirm }}

## Inputs, discussion and decisions

| Topic | CompPlus pre-filled input | Discussion and decision | Action owner | Due date |
|---|---|---|---|---|
| Previous actions | Open and overdue management-review actions | {{ meeting.previous_actions_decision | confirm }} | {{ meeting.previous_actions_owner | confirm }} | {{ meeting.previous_actions_due | confirm }} |
| Changes in context and interested parties | Changes recorded since the previous review | {{ meeting.context_decision | confirm }} | {{ meeting.context_owner | confirm }} | {{ meeting.context_due | confirm }} |
| Legal, regulatory, contractual and standards updates | Legal register changes and template update alerts | {{ meeting.legal_decision | confirm }} | {{ meeting.legal_owner | confirm }} | {{ meeting.legal_due | confirm }} |
| Objectives and performance | KPI results and missed targets | {{ meeting.objectives_decision | confirm }} | {{ meeting.objectives_owner | confirm }} | {{ meeting.objectives_due | confirm }} |
| Risks and opportunities | Highest residual risks and overdue treatments | {{ meeting.risks_decision | confirm }} | {{ meeting.risks_owner | confirm }} | {{ meeting.risks_due | confirm }} |
| Incidents, complaints and nonconformities | Trend summary and open corrective actions | {{ meeting.incidents_decision | confirm }} | {{ meeting.incidents_owner | confirm }} | {{ meeting.incidents_due | confirm }} |
| Audit and assessment results | Open findings and completed audit results | {{ meeting.audit_decision | confirm }} | {{ meeting.audit_owner | confirm }} | {{ meeting.audit_due | confirm }} |
| Supplier performance | Critical supplier reviews and issues | {{ meeting.suppliers_decision | confirm }} | {{ meeting.suppliers_owner | confirm }} | {{ meeting.suppliers_due | confirm }} |
| Resources and competence | Training completion, workload and resource constraints | {{ meeting.resources_decision | confirm }} | {{ meeting.resources_owner | confirm }} | {{ meeting.resources_due | confirm }} |
| Improvement opportunities | Suggestions, lessons and recurring issues | {{ meeting.improvement_decision | confirm }} | {{ meeting.improvement_owner | confirm }} | {{ meeting.improvement_due | confirm }} |

## Overall conclusion

**Is the management system suitable, adequate and effective?** {{ meeting.effectiveness | confirm }}

**Changes required:** {{ meeting.required_changes | confirm }}

**Resources approved:** {{ meeting.resources_approved | confirm }}

**Chair approval:** {{ meeting.approval_status | confirm }}  
**Approval date:** {{ meeting.approval_date | confirm }}