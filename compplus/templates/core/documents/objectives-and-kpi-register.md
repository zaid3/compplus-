# Objectives and KPI Register

**Organisation:** {{ organization.legal_name }}  
**Register owner:** {{ roles.executive_owner }}  
**Reporting year:** {{ objectives.reporting_year | confirm }}

CompPlus proposes the starter objectives below. Confirm each target, owner and due date. Remove any objective that is not relevant and add organisation-specific objectives.

| Objective | Metric | Baseline | Target | Owner | Due date | Review frequency | Status | Evidence |
|---|---|---|---|---|---|---|---|---|
| Complete overdue compliance actions | Number of overdue high-priority actions | {{ metrics.overdue_high_priority }} | 0 | {{ roles.system_manager }} | {{ organization.certification_target | confirm }} | Monthly | Open | CompPlus task report |
| Complete scheduled reviews on time | Percentage of scheduled reviews completed by the due date | {{ metrics.review_completion_rate }} | 95% or higher | {{ roles.system_manager }} | {{ objectives.year_end | confirm }} | Quarterly | Open | Review and audit records |
| Maintain staff competence | Percentage of required training completed | {{ metrics.training_completion_rate }} | 100% | {{ roles.system_manager }} | {{ objectives.training_due_date | confirm }} | Monthly | Open | Training register and completion records |
| Close corrective actions effectively | Percentage closed by the due date and verified | {{ metrics.corrective_action_rate }} | 90% or higher | {{ roles.system_manager }} | {{ objectives.year_end | confirm }} | Monthly | Open | Corrective-action register |
| {{ objectives.custom_1 | confirm }} | {{ objectives.custom_1_metric | confirm }} | {{ objectives.custom_1_baseline | confirm }} | {{ objectives.custom_1_target | confirm }} | {{ objectives.custom_1_owner | confirm }} | {{ objectives.custom_1_due | confirm }} | {{ objectives.custom_1_frequency | confirm }} | Open | {{ objectives.custom_1_evidence | confirm }} |

## Review decision

**Progress summary:** {{ objectives.progress_summary | confirm }}  
**Changes or resources required:** {{ objectives.actions | confirm }}  
**Reviewed by:** {{ roles.executive_owner }}  
**Review date:** {{ approval.date | confirm }}