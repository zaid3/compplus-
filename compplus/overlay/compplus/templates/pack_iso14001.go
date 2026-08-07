package templates

func iso14001Pack() Pack {
  templates := []TemplateDefinition{
    tmpl("ems-scope", "EMS-4.3", "Environmental Management System Scope", "governance document", "EMS Governance", "GOVERNANCE", "roles.environment_owner",
      "Define the organisational boundaries, activities, services, locations and authority covered by the environmental management system.",
      []string{
        "The EMS covers activities, products/services and supporting operations under the control or influence of {{organization.legal_name}} for {{organization.services}} across {{organization.locations}}.",
        "Scope decisions consider organisational context, interested-party requirements, environmental conditions and relevant upstream/downstream lifecycle relationships.",
        "The approved scope is maintained as controlled information and reviewed after significant operational, location, service, legal or environmental change.",
      },
      []string{"Approved EMS scope", "Current location/activity information"},
      []string{"ISO 14001:2026 EMS scope and context requirements"},
      []string{"activities/sites included", "material outsourced or influenced activities"}, nil),

    tmpl("environmental-policy", "EMS-5.2", "Environmental Policy", "policy", "EMS Governance", "POLICY", "roles.executive_owner",
      "Set management commitments for environmental protection, applicable obligations, objectives and continual improvement of environmental performance.",
      []string{
        "{{organization.legal_name}} commits to protecting the environment in ways relevant to its activities, products/services and environmental context.",
        "Applicable environmental legal, regulatory, contractual and other compliance obligations are identified and addressed.",
        "Environmental objectives are set, monitored and supported with resources and ownership.",
        "Pollution, waste, inefficient resource use and other significant adverse impacts are prevented, reduced or controlled where reasonably practicable within organisational control or influence.",
        "The EMS and environmental performance are continually improved using objectives, monitoring, audits, incidents, compliance evaluation and management review.",
      },
      []string{"Approved environmental policy", "Evidence policy was communicated/made available"},
      []string{"ISO 14001:2026 environmental policy and leadership requirements"},
      []string{"organisation-specific environmental commitments"}, nil),

    confidential(tmpl("aspects-impacts", "EMS-6.1.2", "Environmental Aspects and Impacts Register", "register", "Environmental Planning", "REGISTER", "roles.environment_owner",
      "Identify environmental aspects, resulting impacts and significance using a repeatable lifecycle-aware method.",
      []string{
        "Consider normal operations, abnormal conditions and reasonably foreseeable emergency situations for activities, services, facilities, travel, procurement, waste, energy, water, materials, emissions and other relevant interactions with the environment.",
        "Use a documented significance method considering scale/severity, frequency/likelihood, legal/stakeholder concern, control/influence and lifecycle relevance.",
        "Significant aspects receive operational controls, objectives or other action proportionate to risk and opportunity.",
        "The register is reviewed after material change and at planned intervals.",
      },
      []string{"Current aspects/impacts register", "Evidence supporting significance ratings"},
      []string{"ISO 14001:2026 environmental aspects and lifecycle perspective"},
      []string{"significance criteria and threshold", "site/service-specific aspects"},
      []string{"Energy use and associated environmental impact", "Business travel/transport", "Purchased goods/services and supplier influence", "Waste generation and disposal", "Water use where relevant", "Emissions/discharges/noise where relevant", "Emergency spill/fire/pollution scenarios where relevant"})),

    confidential(tmpl("compliance-obligations", "EMS-6.1.3", "Environmental Compliance Obligations Register", "register", "Environmental Compliance", "REGISTER", "roles.environment_owner",
      "Maintain current environmental legal, regulatory, permit, contractual and voluntary obligations relevant to activities and significant aspects.",
      []string{
        "Record the obligation source, applicable requirement, affected activity/site/aspect, responsible owner, compliance method, evidence and review date.",
        "Use authoritative sources and competent advice where interpretation is uncertain; do not rely on generic checklists as evidence that every obligation applies.",
        "Changes in law, permits, services, locations or environmental aspects trigger review and updates to controls/objectives where required.",
      },
      []string{"Environmental obligations register", "Relevant permits/licences/contracts where applicable"},
      []string{"ISO 14001:2026 compliance obligations"},
      []string{"jurisdictions/sites and applicable permits or sector obligations"}, nil)),

    tmpl("compliance-evaluation", "EMS-9.1.2", "Environmental Compliance Evaluation Procedure and Record", "procedure", "Environmental Compliance", "PROCEDURE", "roles.environment_owner",
      "Evaluate fulfilment of applicable environmental compliance obligations and retain evidence of status and corrective action.",
      []string{
        "Define what obligations are evaluated, method/evidence, responsible evaluator and frequency proportionate to risk and legal importance.",
        "Record compliant, noncompliant and uncertain results with supporting evidence; uncertainty is escalated for clarification rather than marked compliant by assumption.",
        "Noncompliance triggers containment/correction, root-cause/corrective action where needed and appropriate authority/customer notification when required.",
        "Maintain knowledge of compliance status and report material results to management review.",
      },
      []string{"Completed compliance evaluation", "Evidence/actions for identified gaps"},
      []string{"ISO 14001:2026 compliance evaluation"},
      []string{"evaluation cycle and competent evaluator(s)"}, nil),

    tmpl("environmental-objectives", "EMS-6.2", "Environmental Objectives and Action Programme", "register", "Environmental Performance", "REGISTER", "roles.environment_owner",
      "Turn significant aspects, obligations and environmental priorities into measurable objectives and action plans.",
      []string{
        "Objectives consider significant aspects, compliance obligations, risks/opportunities, technology, operational feasibility, resources and interested-party expectations.",
        "Each objective identifies target, baseline where useful, measure, actions, resources, owner, due date and evaluation method.",
        "Progress is monitored and actions are adjusted when performance is off track or circumstances change.",
      },
      []string{"Environmental objectives programme", "Performance review evidence"},
      []string{"ISO 14001:2026 objectives and planning"},
      []string{"targets, baselines and owners"},
      []string{"Reduce avoidable energy use", "Reduce waste and improve reuse/recycling", "Reduce unnecessary business travel/emissions where relevant", "Close environmental compliance actions on time", "Improve supplier environmental performance for material categories"}),

    tmpl("operational-control", "EMS-8.1", "Environmental Operational Control Procedure", "procedure", "Environmental Operations", "PROCEDURE", "roles.environment_owner",
      "Define operating criteria and controls for significant environmental aspects, compliance obligations and environmental risks across the lifecycle where the organisation controls or influences them.",
      []string{
        "Operational criteria are documented where their absence could lead to significant environmental impact or noncompliance.",
        "Controls cover relevant internal operations, maintenance, procurement, outsourced processes, contractors and change management.",
        "Environmental requirements are communicated to relevant external providers and lifecycle stages when the organisation has meaningful control or influence.",
        "Planned changes are reviewed for environmental impacts before implementation where practical.",
      },
      []string{"Operational control records", "Supplier/contractor requirements where applicable"},
      []string{"ISO 14001:2026 operational planning and control"},
      []string{"significant aspects requiring formal operating criteria"}, nil),

    tmpl("waste", "EMS-WASTE", "Waste and Resource Management Register", "register", "Resource and Waste", "REGISTER", "roles.environment_owner",
      "Track important waste streams, handling, contractors, legal duties and improvement opportunities.",
      []string{
        "Record waste stream/category, source, quantity or reasonable measure, storage/handling, disposal/recovery route, contractor, evidence and responsible owner where material.",
        "Apply the waste hierarchy or other relevant legal/organisational principles where applicable and seek prevention/reduction before disposal.",
        "Verify relevant contractor/carrier/disposal credentials and records where required by law or contract.",
      },
      []string{"Waste register", "Transfer/disposal/contractor evidence where applicable"},
      []string{"Supports ISO 14001:2026 aspects, operational control, compliance and performance"},
      []string{"material waste streams and measurement method"}, nil),

    tmpl("energy-water", "EMS-RESOURCE", "Energy, Water and Resource Monitoring Register", "register", "Environmental Performance", "REGISTER", "roles.environment_owner",
      "Monitor material resource use and detect trends, abnormal use and improvement opportunities.",
      []string{
        "Identify material resource streams such as electricity, gas/fuel, water, paper/materials or other sector-specific resources.",
        "Record consumption using available reliable data and normalise performance where a meaningful business activity factor improves interpretation.",
        "Investigate material unexpected changes and use trends to inform objectives, maintenance, purchasing and management review.",
      },
      []string{"Resource-consumption records", "Trend/performance review"},
      []string{"Supports ISO 14001:2026 environmental performance monitoring"},
      []string{"material resources, data source, baseline and reporting frequency"}, nil),

    confidential(tmpl("emissions", "EMS-EMISSIONS", "Environmental Emissions and Discharges Register", "register", "Environmental Performance", "REGISTER", "roles.environment_owner",
      "Record material emissions, discharges and related monitoring or permit requirements where relevant to the organisation.",
      []string{
        "Identify relevant air emissions, greenhouse-gas sources, water/land discharges, noise or other releases based on actual activities and legal/environmental context.",
        "Record source, measure/estimate method, applicable limit/target/obligation, monitoring frequency, responsible owner and evidence.",
        "Unexpected or noncompliant releases are handled through environmental incident and corrective-action processes.",
      },
      []string{"Emissions/discharge records where applicable", "Monitoring/permit evidence"},
      []string{"Supports ISO 14001:2026 aspects, compliance and performance monitoring"},
      []string{"which emission/discharge categories are material/applicable"}, nil)),

    confidential(tmpl("emergency", "EMS-8.2", "Environmental Emergency Preparedness and Response Plan", "plan", "Emergency Preparedness", "PLAN", "roles.environment_owner",
      "Prepare for environmental emergency situations, reduce adverse impact and test response arrangements.",
      []string{
        "Identify reasonably foreseeable emergency situations linked to environmental aspects, facilities, stored materials, utilities, suppliers and local conditions.",
        "Define immediate protective/containment actions, emergency contacts, authorities/landlord/supplier interfaces, communications and recovery responsibilities.",
        "Provide appropriate training/resources and test relevant response arrangements periodically where practicable.",
        "Review the plan and actual/tested responses after exercises, incidents and significant changes.",
      },
      []string{"Emergency plan", "Exercise/test record", "Emergency equipment/inspection evidence where relevant"},
      []string{"ISO 14001:2026 emergency preparedness and response"},
      []string{"credible environmental emergencies and emergency contacts"}, nil)),

    confidential(tmpl("incident-register", "EMS-INCIDENT", "Environmental Incident and Near-Miss Register", "register", "Environmental Incidents", "REGISTER", "roles.environment_owner",
      "Record environmental incidents, near misses, impacts, notifications, causes and corrective actions.",
      []string{
        "Record event date/location, aspect/impact, immediate action, actual/potential environmental harm and applicable notification/permit considerations.",
        "Material events receive cause analysis and corrective action, with effectiveness checked before closure.",
        "Trends are reviewed to improve operational controls, emergency plans, training and risk/aspect evaluations.",
      },
      []string{"Environmental incident register", "Completed event/action record"},
      []string{"Supports ISO 14001:2026 emergency, nonconformity and improvement"},
      []string{"incident escalation/notification authority"}, nil)),

    tmpl("contractor", "EMS-CONTRACTOR", "Supplier and Contractor Environmental Assessment", "assessment", "Third Parties", "TEMPLATE", "roles.environment_owner",
      "Apply proportionate environmental requirements to suppliers and contractors where procurement or outsourced work can create significant impact or compliance risk.",
      []string{
        "Assess environmental relevance using service/material type, onsite activity, waste, transport, hazardous materials, energy/resource intensity, legal duties and lifecycle influence.",
        "Communicate applicable controls, competence, emergency, waste, reporting and legal requirements before work starts where relevant.",
        "Monitor material provider performance and address nonconformity or improvement opportunities through supplier management.",
      },
      []string{"Completed supplier/contractor environmental assessment", "Relevant contractual/work instructions"},
      []string{"Supports ISO 14001:2026 lifecycle, procurement and externally provided process controls"},
      []string{"high-environmental-impact suppliers/contractors"}, nil),

    tmpl("lifecycle", "EMS-LIFECYCLE", "Lifecycle and Sustainable Procurement Assessment", "assessment", "Lifecycle", "TEMPLATE", "roles.environment_owner",
      "Consider material environmental impacts upstream and downstream when designing services, purchasing goods/services and influencing customer/end-of-life outcomes.",
      []string{
        "Consider acquisition/extraction, design, transport, use/operation and end-of-life stages only to the degree relevant to actual organisational control or influence.",
        "Use procurement specifications, supplier selection, design choices, customer information or disposal/reuse arrangements to influence significant lifecycle impacts where practical.",
        "Record important trade-offs and assumptions rather than claiming lifecycle benefits without evidence.",
      },
      []string{"Lifecycle/procurement assessment", "Supporting supplier/product information"},
      []string{"ISO 14001:2026 lifecycle perspective and operational control"},
      []string{"material purchased categories/services and lifecycle influence"}, nil),

    tmpl("monitoring", "EMS-9.1", "Environmental Monitoring and Measurement Plan", "plan", "Environmental Performance", "PLAN", "roles.environment_owner",
      "Define what environmental performance is monitored, methods, criteria, frequency, calibration/verification needs and evaluation responsibilities.",
      []string{
        "Monitoring focuses on significant aspects, operational controls, environmental objectives, compliance obligations and other indicators needed to evaluate EMS performance.",
        "Methods and equipment are suitable for intended use; calibration/verification is maintained where necessary for valid results.",
        "Results are analysed for trends, abnormal conditions, objective progress and need for corrective/improvement action.",
      },
      []string{"Monitoring plan", "Recent environmental performance report", "Calibration/verification where relevant"},
      []string{"ISO 14001:2026 monitoring, measurement, analysis and evaluation"},
      []string{"indicators, methods, frequency and owners"}, nil),

    tmpl("environmental-communication", "EMS-7.4", "Environmental Communication Plan", "plan", "EMS Governance", "PLAN", "roles.environment_owner",
      "Define reliable internal and external communication for relevant environmental information and obligations.",
      []string{
        "Determine what environmental information must or should be communicated, when, to whom, by whom and through which method.",
        "External claims are accurate, consistent with evidence and reviewed before publication where environmental performance or compliance is represented.",
        "Required communications with authorities, customers, landlords, suppliers, emergency responders or communities are assigned to named owners.",
      },
      []string{"Environmental communication plan", "Examples of relevant communications"},
      []string{"ISO 14001:2026 communication requirements"},
      []string{"external reporting/notification obligations"}, nil),

    tmpl("ems-audit", "EMS-9.2", "EMS Internal Audit Checklist and Report", "audit document", "Assurance", "REPORT", "roles.environment_owner",
      "Provide an audit-ready structure for testing EMS implementation, environmental controls, compliance evaluation and performance.",
      []string{
        "Audit scope and criteria consider EMS requirements, internal arrangements, significant aspects, compliance obligations, operational controls and previous results.",
        "Use objective evidence and competent, sufficiently independent auditors; record conformity, nonconformity and improvement observations.",
        "Communicate results and track corrective actions to verified closure.",
      },
      []string{"Completed EMS audit report", "Audit evidence/findings"},
      []string{"ISO 14001:2026 internal audit"},
      []string{"audit scope, dates and auditor(s)"},
      []string{"Context/scope/policy", "Aspects and significance", "Compliance obligations/evaluation", "Objectives", "Operational controls", "Emergency preparedness", "Monitoring/performance", "Supplier/lifecycle controls", "Incidents/nonconformities", "Improvement"}),

    confidential(tmpl("ems-review", "EMS-9.3", "EMS Management Review Agenda and Minutes", "record", "Governance", "RECORD", "roles.executive_owner",
      "Provide senior management with a ready agenda for environmental performance, changing context, compliance and improvement decisions.",
      []string{
        "Review previous actions, context/interested-party changes, significant aspects, compliance status, objectives, environmental performance trends, communications/complaints, audit results, incidents/nonconformities, resources, risks/opportunities and improvement options.",
        "Record decisions on continuing suitability/adequacy/effectiveness, changes, resources, objectives and improvement actions with owners/dates.",
      },
      []string{"Approved EMS management-review minutes", "Supporting environmental performance/compliance evidence"},
      []string{"ISO 14001:2026 management review"},
      []string{"meeting date, attendees, decisions and actions"}, nil)),

    tmpl("ems-corrective", "EMS-10", "Environmental Nonconformity and Corrective Action Form", "form", "Improvement", "TEMPLATE", "roles.environment_owner",
      "Provide a consistent record for environmental nonconformity, correction, cause, action and effectiveness review.",
      []string{
        "Describe the nonconformity/incident and environmental impact or compliance consequence.",
        "Take immediate action to control/correct the issue and address consequences where needed.",
        "Assess cause, recurrence/similar risks and select corrective action proportionate to significance.",
        "Verify effectiveness and update aspects, controls, obligations, objectives or documented information where necessary.",
      },
      []string{"Completed environmental corrective-action record"},
      []string{"ISO 14001:2026 nonconformity, corrective action and improvement"},
      []string{"closure authority for significant environmental actions"}, nil),
  }

  return makePack(
    "iso14001",
    "Comp Plus+ ISO 14001 Fast Start",
    "2026.08.07",
    "ISO 14001:2026",
    "Ready-made EMS policy, aspects/impacts, compliance obligations, lifecycle, operational control, emergency, monitoring, audit and management-review templates for the current 2026 edition.",
    templates,
    false,
    nil,
  )
}
