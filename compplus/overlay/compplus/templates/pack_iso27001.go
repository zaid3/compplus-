package templates

import "fmt"

func iso27001Pack() Pack {
  templates := []TemplateDefinition{
    tmpl("isms-scope", "CLAUSE-4.3", "ISMS Scope Statement", "governance document", "ISMS Governance", "GOVERNANCE", "roles.security_owner",
      "Define the boundaries and applicability of the information security management system, including interfaces and dependencies.",
      []string{
        "The ISMS covers information, people, processes, technology and third parties used by {{organization.legal_name}} to deliver {{organization.services}} within {{organization.locations}}.",
        "External and internal context, interested-party requirements and interfaces with outsourced activities are considered when defining boundaries.",
        "The approved scope is maintained as controlled documented information and reviewed after significant organisational, service, technology or supplier change.",
      },
      []string{"Approved ISMS scope", "Current service/system/location inventory"},
      []string{"ISO/IEC 27001:2022 clause 4.3", "ISO/IEC 27001:2022/Amd 1:2024 context update"},
      []string{"information systems and cloud environments in scope", "interfaces/dependencies with third parties", "any scope exclusions and rationale"}, nil),

    tmpl("security-policy", "CLAUSE-5.2", "Information Security Policy", "policy", "ISMS Governance", "POLICY", "roles.executive_owner",
      "Set management direction and commitments for protecting information and continually improving the ISMS.",
      []string{
        "{{organization.legal_name}} protects the confidentiality, integrity and availability of information in line with business, legal, regulatory, contractual and stakeholder requirements.",
        "Information-security objectives are established and monitored, and sufficient resources are provided to operate and improve the ISMS.",
        "Risk is managed using the approved assessment and treatment method, and controls are selected according to risk, obligations and business need.",
        "Security responsibilities are assigned, communicated and supported by appropriate competence, awareness and escalation routes.",
        "Security incidents, nonconformities and control weaknesses are reported, investigated and used to drive continual improvement.",
      },
      []string{"Approved policy", "Evidence policy was communicated to relevant workers"},
      []string{"ISO/IEC 27001:2022 clause 5.2"},
      []string{"senior management approval", "how the policy is made available to relevant interested parties"}, nil),

    tmpl("security-roles", "CLAUSE-5.3", "Information Security Roles and Responsibility Matrix", "register", "ISMS Governance", "REGISTER", "roles.security_owner",
      "Assign information-security responsibilities, decision rights, escalation and reporting responsibilities.",
      []string{
        "Top management assigns responsibility for ISMS conformity and for reporting ISMS performance.",
        "Operational security duties have named owners and deputies, with segregation of conflicting duties where practical.",
        "Responsibilities are reviewed after organisational change and communicated to affected people.",
      },
      []string{"Approved security role matrix", "Relevant job descriptions or responsibility evidence"},
      []string{"ISO/IEC 27001:2022 clause 5.3", "Annex A organisational controls"},
      []string{"security incident decision owner", "risk acceptance authority", "access approval authority"}, nil),

    tmpl("risk-method", "CLAUSE-6.1.2", "Information Security Risk Assessment Methodology", "procedure", "ISMS Risk", "PROCEDURE", "roles.security_owner",
      "Define repeatable information-security risk criteria and a consistent assessment process focused on loss of confidentiality, integrity and availability.",
      []string{
        "Risk criteria include likelihood, impact, risk level, acceptance criteria and rules for when assessment must be repeated.",
        "Information-security risks identify the affected information/service, threat or event, vulnerability/cause, potential consequence and accountable risk owner.",
        "Impact considers confidentiality, integrity, availability, legal/regulatory, contractual, operational, financial and reputational consequences where relevant.",
        "Results are prioritised against approved risk criteria and retained as documented information.",
      },
      []string{"Approved risk methodology", "Completed risk assessment using the method"},
      []string{"ISO/IEC 27001:2022 clause 6.1.2"},
      []string{"risk acceptance criteria", "likelihood and impact definitions", "assessment/reassessment triggers"},
      []string{"Recommended starter method: likelihood 1-5 multiplied by impact 1-5", "Assess inherent risk before treatment and residual risk after treatment", "High or out-of-appetite residual risk requires explicit risk-owner approval"}),

    confidential(tmpl("risk-register", "CLAUSE-6.1.2-R", "Information Security Risk Register", "register", "ISMS Risk", "REGISTER", "roles.security_owner",
      "Provide a ready structure for recording security risks, owners, scoring, treatment and review status.",
      []string{
        "Each risk has a unique reference, affected asset/service/process, risk description, owner, inherent likelihood/impact and resulting level.",
        "Record selected treatment option, linked controls/actions, target date, residual likelihood/impact, acceptance decision and next review.",
      },
      []string{"Current risk register", "Evidence supporting high-risk treatment/acceptance decisions"},
      []string{"ISO/IEC 27001:2022 clauses 6.1.2 and 8.2"},
      []string{"initial priority risks", "risk review frequency"},
      []string{"Example starter risk: unauthorised account access causes disclosure or alteration of sensitive information — assess based on actual authentication/access controls", "Example starter risk: critical cloud/service outage prevents delivery of important services — assess resilience and recovery arrangements", "Example starter risk: supplier compromise exposes information or disrupts service — assess critical suppliers and contractual controls"})),

    confidential(tmpl("risk-treatment", "CLAUSE-6.1.3", "Information Security Risk Treatment Plan", "plan", "ISMS Risk", "PLAN", "roles.security_owner",
      "Translate risk decisions into owned control and treatment actions and record residual-risk approval.",
      []string{
        "For each unacceptable risk, select a treatment option and determine controls needed to reduce risk to an approved level.",
        "Compare selected controls with the Annex A reference set so necessary controls are not overlooked; additional controls may be used when justified by risk or obligations.",
        "Each treatment action records owner, resources, due date, status, expected risk reduction and evidence.",
        "Risk owners approve the treatment plan and accept residual risk after treatment.",
      },
      []string{"Approved risk treatment plan", "Risk-owner approval/acceptance evidence"},
      []string{"ISO/IEC 27001:2022 clause 6.1.3 and 8.3"},
      []string{"treatment owners and target dates", "risk-owner residual-risk approvals"}, nil)),

    tmpl("soa-guide", "CLAUSE-6.1.3-SOA", "Statement of Applicability Review Guide", "procedure", "ISMS Governance", "PROCEDURE", "roles.security_owner",
      "Give users a simple workflow for completing and maintaining the native Comp Plus+ Statement of Applicability.",
      []string{
        "Review every Annex A reference against assessed risks, applicable legal/regulatory/contractual requirements and business needs.",
        "For controls retained as applicable, record a meaningful inclusion rationale and link implementation evidence or measures.",
        "For controls marked not applicable, record a defensible exclusion rationale showing why the control is unnecessary in the organisation's context.",
        "Confirm implementation status through linked measures/evidence rather than assuming that a control is implemented because it is applicable.",
        "Review the SoA after material risk, scope, technology, supplier, obligation or control change and before external certification/assurance activity.",
      },
      []string{"Approved Statement of Applicability", "Evidence supporting applicability and implementation decisions"},
      []string{"ISO/IEC 27001:2022 clause 6.1.3"},
      []string{"final rationale for any excluded Annex A control"},
      []string{"Comp Plus+ creates all 93 Annex A references as provisional entries for review; users confirm applicability, rationale and implementation evidence"}),

    tmpl("security-objectives", "CLAUSE-6.2", "Information Security Objectives and Measurement Plan", "register", "ISMS Performance", "REGISTER", "roles.security_owner",
      "Maintain measurable security objectives, owners, resources, deadlines, monitoring and evaluation methods.",
      []string{
        "Objectives are consistent with the security policy and take account of applicable requirements, risk assessment and risk treatment results.",
        "Each objective identifies what will be done, required resources, responsible owner, completion date and evaluation method.",
        "Objectives are monitored, communicated and updated when priorities or risks change.",
      },
      []string{"Security objectives/KPI review"},
      []string{"ISO/IEC 27001:2022 clause 6.2"},
      []string{"security targets and reporting frequency"},
      []string{"Critical access reviews completed on schedule", "High-risk vulnerabilities remediated within approved service levels", "Security awareness completion and effectiveness", "Backup/recovery tests completed successfully", "High-risk treatment actions closed on time"}),

    tmpl("communication", "CLAUSE-7.4", "Information Security Communication Plan", "plan", "ISMS Governance", "PLAN", "roles.security_owner",
      "Define what security information must be communicated, when, to whom, by whom and through which channels.",
      []string{
        "Internal communications cover policies, responsibilities, threats, incidents, planned changes, training and security performance as relevant to the audience.",
        "External communications cover customers, suppliers, authorities, affected individuals and other stakeholders when required by contract, law, incident response or assurance commitments.",
        "Sensitive communications follow approved confidentiality, legal-review and incident-communication rules.",
      },
      []string{"Communication plan", "Examples of relevant security communications"},
      []string{"ISO/IEC 27001:2022 clause 7.4"},
      []string{"regulatory/customer incident communication contacts"}, nil),

    tmpl("asset-classification", "A.5.9", "Information Asset Management and Classification Policy", "policy", "Asset Security", "POLICY", "roles.security_owner",
      "Maintain an accountable inventory of important information and associated assets and protect them according to sensitivity and business importance.",
      []string{
        "Information and associated assets in scope are identified and assigned an accountable owner.",
        "Classification considers confidentiality, integrity, availability, legal/contractual duties and business impact, with handling rules proportionate to classification.",
        "Asset records are reviewed after acquisition, major change, transfer, disposal and at planned intervals.",
        "Information is labelled or otherwise clearly handled according to classification where this adds practical protection.",
      },
      []string{"Asset/information inventory", "Classification and handling evidence"},
      []string{"ISO/IEC 27001:2022 Annex A 5.9-5.13"},
      []string{"classification levels and handling rules"},
      []string{"Suggested starter levels: Public; Internal; Confidential; Highly Restricted — tailor to actual business needs"}),

    tmpl("acceptable-use", "A.5.10", "Acceptable Use of Information and Technology Policy", "policy", "People and Assets", "POLICY", "roles.security_owner",
      "Set simple, enforceable expectations for authorised use of information, devices, accounts, cloud services and communication tools.",
      []string{
        "Business information and systems are used only for authorised purposes and in line with law, contracts, policy and role responsibilities.",
        "Users protect credentials, devices and sensitive information; sharing accounts or bypassing security controls is prohibited unless explicitly authorised for a controlled purpose.",
        "Only approved software/services and storage locations are used for business information unless an exception is documented and approved.",
        "Suspected loss, misuse, compromise or policy breach is reported promptly through the incident route.",
      },
      []string{"Approved acceptable-use policy", "Worker acknowledgement or communication evidence"},
      []string{"ISO/IEC 27001:2022 Annex A 5.10"},
      []string{"rules for limited personal use, removable media and unapproved cloud/AI services"}, nil),

    tmpl("access-control", "A.5.15", "Identity and Access Control Policy", "policy", "Access Control", "POLICY", "roles.security_owner",
      "Ensure access to information and systems is authorised, least-privilege, attributable, periodically reviewed and removed promptly when no longer needed.",
      []string{
        "Every user receives a unique identity unless a documented technical exception requires a controlled shared/service account.",
        "Access requires approved business need and follows least privilege, need-to-know and segregation principles.",
        "Privileged access receives stronger approval, authentication, monitoring and review than standard access.",
        "Authentication strength is proportionate to risk; multi-factor authentication is used for privileged, remote and sensitive access where technically feasible.",
        "Access is reviewed periodically and promptly changed or removed after role change, termination, contract end or identified misuse.",
      },
      []string{"Access-control policy", "Recent access approval/review evidence", "Privileged account inventory"},
      []string{"ISO/IEC 27001:2022 Annex A 5.15-5.18 and relevant technological access controls"},
      []string{"access review frequency", "MFA exceptions, if any", "privileged-access approval authority"}, nil),

    tmpl("jml", "A.5.18", "Joiner, Mover and Leaver Access Procedure", "procedure", "Access Control", "PROCEDURE", "roles.security_owner",
      "Give managers and IT a repeatable checklist for granting, changing and removing access through the worker lifecycle.",
      []string{
        "Joiners receive only manager-approved access required for their role after required checks/agreements and induction steps are complete.",
        "Role or responsibility changes trigger prompt review of existing access, removal of unnecessary permissions and approval of new access.",
        "Leaver access is disabled or removed at the appropriate effective time, company assets are recovered, shared secrets are changed where needed and ongoing confidentiality obligations are reinforced.",
        "Exceptions and delayed removals are documented, risk-assessed and approved.",
      },
      []string{"Completed joiner/mover/leaver example", "Termination access-removal evidence"},
      []string{"ISO/IEC 27001:2022 Annex A access and people controls"},
      []string{"HR/manager notification route", "target time for high-risk leaver access removal"}, nil),

    confidential(tmpl("supplier-security", "A.5.19", "Supplier and ICT Supply Chain Security Procedure", "procedure", "Third Parties", "PROCEDURE", "roles.security_owner",
      "Assess and manage information-security risks created by suppliers, cloud services and technology supply chains.",
      []string{
        "Security due diligence is proportionate to supplier criticality, information access, system access, location, sub-processors/subcontractors and service dependency.",
        "Security requirements are included in agreements and cover incident notification, access, confidentiality, continuity, vulnerability/change management, audit/assurance and secure exit where relevant.",
        "Material changes to supplier services or security posture are risk-assessed before acceptance when practical.",
        "Critical supplier performance and assurance are reviewed at planned intervals and after incidents.",
      },
      []string{"Supplier security assessments", "Relevant security clauses", "Critical supplier monitoring evidence"},
      []string{"ISO/IEC 27001:2022 Annex A 5.19-5.22"},
      []string{"critical-supplier review frequency", "minimum assurance evidence by supplier tier"}, nil)),

    tmpl("cloud-security", "A.5.23", "Cloud Service Security Standard", "procedure", "Cloud Security", "PROCEDURE", "roles.security_owner",
      "Set minimum security requirements for acquiring, configuring, operating, changing and exiting cloud services.",
      []string{
        "Cloud services are approved based on business need, data sensitivity, legal/contractual requirements, supplier assurance and service resilience.",
        "Responsibilities under the shared-responsibility model are documented for security configuration, identity, logging, backup, vulnerability management, incident response and data protection.",
        "Secure baseline configurations, least-privilege administration, MFA, logging and encryption are used proportionate to risk.",
        "Cloud changes and exit plans address data portability, retention/deletion, access removal and evidence preservation.",
      },
      []string{"Cloud service inventory", "Cloud security review/configuration evidence"},
      []string{"ISO/IEC 27001:2022 Annex A 5.23"},
      []string{"approved cloud platforms and security baseline owner"}, nil),

    confidential(tmpl("incident-response", "A.5.24", "Information Security Incident Response Procedure", "procedure", "Incident Management", "PROCEDURE", "roles.security_owner",
      "Prepare, assess, respond to and learn from information-security events and incidents while preserving evidence and meeting notification duties.",
      []string{
        "Security events are reported through an accessible route and assessed promptly against documented severity and incident criteria.",
        "The incident lead coordinates containment, eradication, recovery, evidence preservation and communications with relevant technical, legal, privacy, business and external stakeholders.",
        "Evidence is handled so integrity and chronology are preserved when disciplinary, contractual, legal or forensic use may be required.",
        "Lessons learned are documented and feed risk assessment, controls, procedures, training and corrective actions.",
      },
      []string{"Incident response procedure", "Incident register", "Completed incident/tabletop evidence"},
      []string{"ISO/IEC 27001:2022 Annex A 5.24-5.28"},
      []string{"severity levels", "security incident contacts and escalation", "external forensic/legal support contacts where applicable"}, nil)),

    tmpl("continuity", "A.5.29", "Information Security Continuity and ICT Readiness Plan", "plan", "Resilience", "PLAN", "roles.security_owner",
      "Maintain security and required ICT capability during disruption and recover important services within approved objectives.",
      []string{
        "Critical information-security controls and ICT dependencies are identified for disruptive scenarios.",
        "Recovery priorities and objectives reflect business impact, contractual and legal requirements.",
        "Recovery arrangements cover people, technology, communications, suppliers, alternate methods and restoration of secure configurations.",
        "Plans are exercised at planned intervals; results, gaps and actions are recorded and tracked.",
      },
      []string{"Continuity/ICT readiness plan", "Recent exercise or recovery test"},
      []string{"ISO/IEC 27001:2022 Annex A 5.29-5.30"},
      []string{"critical services", "recovery objectives", "minimum security controls during disruption"}, nil),

    tmpl("backup", "A.8.13", "Backup, Restore and Recovery Procedure", "procedure", "Resilience", "PROCEDURE", "roles.security_owner",
      "Ensure important information, software and configurations can be recovered from reliable, protected and tested backups.",
      []string{
        "Backup scope, frequency and retention are based on business recovery needs, information sensitivity and legal/contractual requirements.",
        "Backups are protected from unauthorised access and from the same failure/attack that could affect production where practical.",
        "Restore tests are completed at planned intervals and after significant technology changes, with results and remediation recorded.",
        "Backup failures are monitored and escalated according to service importance.",
      },
      []string{"Backup configuration/schedule", "Recent successful restore-test evidence"},
      []string{"ISO/IEC 27001:2022 Annex A 8.13"},
      []string{"backup frequency/retention by critical service", "restore-test frequency"}, nil),

    tmpl("vulnerability", "A.8.8", "Vulnerability and Patch Management Procedure", "procedure", "Technical Security", "PROCEDURE", "roles.security_owner",
      "Identify technical vulnerabilities, assess exposure and apply patches or other treatment within risk-based timeframes.",
      []string{
        "Authoritative vulnerability information is monitored for relevant technology and assets.",
        "Vulnerabilities are prioritised using severity, exploitability, exposure, asset criticality, data sensitivity and compensating controls rather than score alone.",
        "Security patches or mitigations are tested and deployed within approved risk-based service levels; exceptions are documented with temporary controls and an owner/date.",
        "Remediation status is verified and overdue high-risk items are escalated.",
      },
      []string{"Vulnerability scan/advisory evidence", "Patch/remediation records", "Approved exceptions"},
      []string{"ISO/IEC 27001:2022 Annex A 8.8"},
      []string{"remediation service levels by severity/risk", "vulnerability information sources"}, nil),

    confidential(tmpl("logging", "A.8.15", "Logging, Monitoring and Security Alerting Standard", "procedure", "Technical Security", "PROCEDURE", "roles.security_owner",
      "Generate, protect, review and retain relevant logs so suspicious activity, control failures and incidents can be detected and investigated.",
      []string{
        "Logging requirements are defined for critical systems, privileged activity, authentication, security events, sensitive data access and other risk-relevant events.",
        "Log timestamps are sufficiently synchronised for investigation and correlation.",
        "Logs are protected against unauthorised alteration/deletion and access is restricted.",
        "Monitoring and alert rules focus on meaningful threats and control failures; alerts are triaged, investigated and escalated through incident management.",
        "Retention is based on investigation, legal, contractual and operational needs.",
      },
      []string{"Logging/monitoring configuration", "Example security alert/investigation", "Retention setting evidence"},
      []string{"ISO/IEC 27001:2022 Annex A 8.15-8.17"},
      []string{"critical log sources", "log retention", "alert ownership/on-call route"}, nil)),

    tmpl("cryptography", "A.8.24", "Cryptography and Key Management Policy", "policy", "Technical Security", "POLICY", "roles.security_owner",
      "Use appropriate cryptographic controls and manage keys/secrets through a controlled lifecycle.",
      []string{
        "Encryption is applied according to information sensitivity, threat, legal/contractual requirements and technology capability, including data in transit and at rest where risk warrants it.",
        "Approved algorithms/protocols and minimum configurations follow maintained industry/vendor guidance and are reviewed as technology changes.",
        "Keys and secrets are generated, stored, distributed, rotated, revoked, recovered and destroyed using controlled methods with restricted access.",
        "Hard-coded, shared or exposed secrets are avoided; identified exposure triggers prompt rotation and incident assessment.",
      },
      []string{"Encryption/security configuration evidence", "Key/secret management evidence"},
      []string{"ISO/IEC 27001:2022 Annex A 8.24"},
      []string{"approved key/secret management platform", "rotation requirements for high-risk secrets"}, nil),

    tmpl("remote-mobile", "A.6.7", "Remote Working and Mobile Device Security Policy", "policy", "People and Endpoint Security", "POLICY", "roles.security_owner",
      "Protect business information when people work remotely or use portable/mobile devices.",
      []string{
        "Remote access is authorised and uses approved secure connectivity and authentication appropriate to risk.",
        "Managed devices use required security configuration such as screen lock, supported software, malware protection where appropriate, encryption and patching.",
        "Sensitive information is protected from unauthorised viewing, local household/public access, insecure storage and inappropriate printing/disposal.",
        "Loss, theft, suspected compromise or significant device security failure is reported promptly.",
      },
      []string{"Remote-working policy", "Managed-device/security configuration evidence"},
      []string{"ISO/IEC 27001:2022 Annex A 6.7 and relevant A.8 controls"},
      []string{"approved remote access methods", "BYOD rule or prohibition"}, nil),

    tmpl("physical", "A.7.1", "Physical and Environmental Security Standard", "procedure", "Physical Security", "PROCEDURE", "roles.security_owner",
      "Protect facilities, equipment, media and sensitive working areas from unauthorised physical access, damage, interference and environmental threats.",
      []string{
        "Physical security boundaries and entry controls are proportionate to information, equipment and service criticality.",
        "Visitors are controlled in sensitive areas and physical access rights are approved, reviewed and removed when no longer required.",
        "Work areas apply clear-screen/clear-desk and secure storage practices appropriate to sensitivity.",
        "Equipment and media are protected during use, transport, maintenance, reuse and disposal, including secure sanitisation where required.",
        "Environmental threats and utility dependencies are considered for critical facilities/equipment.",
      },
      []string{"Physical access records/review", "Secure disposal/sanitisation evidence", "Environmental/facility controls evidence"},
      []string{"ISO/IEC 27001:2022 Annex A 7.1-7.14"},
      []string{"sensitive areas and physical access approval owner"}, nil),

    tmpl("secure-development", "A.8.25", "Secure Development and Change Management Procedure", "procedure", "Application Security", "PROCEDURE", "roles.security_owner",
      "Build security into system/software changes from requirements through design, development, testing, release and maintenance.",
      []string{
        "Security and privacy requirements are identified during planning and are proportionate to system/data risk.",
        "Development/test/production environments and access are separated as appropriate, and source code/repositories are protected.",
        "Changes are reviewed, tested and approved before release; emergency changes are controlled and retrospectively reviewed.",
        "Security testing includes relevant code/dependency/configuration checks and remediation of material findings before or promptly after release according to risk.",
        "Outsourced development is subject to equivalent security requirements and assurance.",
      },
      []string{"Change records", "Security test/review evidence", "Secure-development requirements"},
      []string{"ISO/IEC 27001:2022 Annex A 8.25-8.33"},
      []string{"development repositories/tools", "required security checks before production release"}, nil),

    confidential(tmpl("data-handling", "A.8.10", "Secure Information Transfer, Retention and Deletion Standard", "procedure", "Data Security", "PROCEDURE", "roles.security_owner",
      "Protect information during transfer, storage, retention, masking and secure deletion according to sensitivity and obligations.",
      []string{
        "Approved transfer methods provide confidentiality, integrity and recipient assurance proportionate to sensitivity.",
        "Retention periods are based on legal, contractual and business requirements and are not extended without reason.",
        "Information and storage media are securely deleted or sanitised when retention ends, assets are reused/disposed, or contracts/services terminate, subject to legitimate preservation needs.",
        "Masking or equivalent minimisation is used where full sensitive values are unnecessary for the purpose or environment.",
        "Data leakage risks are addressed using process and technical controls appropriate to the information and channels involved.",
      },
      []string{"Secure transfer configuration/example", "Retention/deletion evidence", "Data masking/minimisation evidence where applicable"},
      []string{"ISO/IEC 27001:2022 Annex A 5.14, 8.10-8.12"},
      []string{"approved sensitive-data transfer channels", "secure deletion/sanitisation method"}, nil)),

    tmpl("threat-intelligence", "A.5.7", "Threat Intelligence and Security Advisory Procedure", "procedure", "Threat Management", "PROCEDURE", "roles.security_owner",
      "Collect and use relevant threat/advisory information to improve risk decisions, detection and protective controls.",
      []string{
        "Relevant security advisories, supplier notices and credible threat sources are monitored according to the organisation's technology, services and threat profile.",
        "Material intelligence is assessed for applicability and translated into vulnerability action, detection changes, risk updates, awareness or incident investigation as appropriate.",
        "Information sharing respects confidentiality, legal and contractual restrictions.",
      },
      []string{"Threat/advisory review evidence", "Resulting risk/control/remediation action"},
      []string{"ISO/IEC 27001:2022 Annex A 5.7"},
      []string{"primary advisory/threat sources and review owner"}, nil),

    tmpl("network-security", "A.8.20", "Network Security and Segmentation Standard", "procedure", "Technical Security", "PROCEDURE", "roles.security_owner",
      "Protect network services and traffic using secure architecture, controlled connectivity, segmentation and monitored configuration.",
      []string{
        "Network access and services are authorised according to business need and risk, with unnecessary exposure disabled.",
        "Segmentation separates environments or assets with materially different trust, sensitivity or threat levels where it meaningfully reduces risk.",
        "Network security devices/services use controlled secure configuration, administrative access and change management.",
        "Filtering and monitoring are maintained to reduce malicious or unauthorised traffic and support investigation.",
      },
      []string{"Network architecture/segmentation evidence", "Firewall/filter review", "Network change records"},
      []string{"ISO/IEC 27001:2022 Annex A 8.20-8.23"},
      []string{"network zones and internet-facing services", "firewall/rule review frequency"}, nil),

    tmpl("configuration", "A.8.9", "Secure Configuration Management Standard", "procedure", "Technical Security", "PROCEDURE", "roles.security_owner",
      "Establish, approve and maintain secure configurations for important systems, applications, cloud services, network devices and endpoints.",
      []string{
        "Secure baselines are based on vendor or recognised guidance and adjusted for business/risk needs.",
        "Default credentials, unnecessary services/features and insecure settings are removed or changed before production use where applicable.",
        "Configuration changes follow approved change control and material drift is detected/reviewed using suitable manual or automated methods.",
        "Exceptions from secure baselines are documented with risk, owner, compensating controls and review date.",
      },
      []string{"Secure baseline/configuration evidence", "Configuration exception/drift review"},
      []string{"ISO/IEC 27001:2022 Annex A 8.9"},
      []string{"baseline sources and configuration-review frequency"}, nil),

    tmpl("awareness", "A.6.3", "Information Security Awareness and Training Plan", "plan", "People Security", "PLAN", "roles.security_owner",
      "Provide risk-based security awareness and role-specific training and keep evidence that people understand key responsibilities.",
      []string{
        "All relevant workers receive security induction and periodic refresher awareness covering applicable policies, incident reporting, authentication, phishing/social engineering, information handling and remote-working risks.",
        "People in higher-risk or specialist roles receive additional role-specific training appropriate to their responsibilities.",
        "Effectiveness is evaluated using appropriate evidence such as knowledge checks, simulations, observed behaviour, incident trends or manager review.",
      },
      []string{"Training plan", "Completion records", "Effectiveness evaluation"},
      []string{"ISO/IEC 27001:2022 clauses 7.2-7.3 and Annex A 6.3"},
      []string{"training frequency and role-specific topics"}, nil),
  }

  return makePack(
    "iso27001",
    "Comp Plus+ ISO/IEC 27001 Fast Start",
    "2026.08.07",
    "ISO/IEC 27001:2022 + Amendment 1:2024",
    "Ready-made ISMS policies, procedures, plans and registers with a native 93-control Statement of Applicability starter.",
    templates,
    true,
    iso27001AnnexAControls(),
  )
}

func iso27001AnnexAControls() []ControlDefinition {
  groups := []struct {
    prefix string
    count  int
    label  string
  }{
    {"5", 37, "Organisational"},
    {"6", 8, "People"},
    {"7", 14, "Physical"},
    {"8", 34, "Technological"},
  }
  controls := make([]ControlDefinition, 0, 93)
  for _, group := range groups {
    for n := 1; n <= group.count; n++ {
      ref := fmt.Sprintf("A.%s.%d", group.prefix, n)
      controls = append(controls, ControlDefinition{
        ID: ref,
        Name: fmt.Sprintf("%s security control %s", group.label, ref),
        Description: "Annex A reference used by the Comp Plus+ SoA workflow. Review applicability against organisational risk, obligations and context; link implementation measures and evidence before approval.",
      })
    }
  }
  return controls
}
