# Supplier Due Diligence and Monitoring Procedure

**Owner:** {{ roles.system_manager }}  
**Review cycle:** Annual

## 1. Purpose

{{ organization.legal_name }} assesses suppliers before use and monitors suppliers that can affect service quality, information security, privacy, environmental performance or AI-related outcomes.

## 2. Supplier classification

Suppliers are classified as low, medium or critical based on access to information, personal data, operational dependency, spend, substitutability, regulatory impact and potential harm.

Critical-supplier rules:

{{ supplier.criticality_rules | confirm }}

## 3. Due diligence before approval

The owner records:

- service and business owner;
- data, systems, sites or processes accessed;
- security, privacy, quality, environmental and AI risks that apply;
- certifications or independent assurance where relevant;
- contract, confidentiality, data-processing, service-level, incident-notification and exit requirements;
- approval decision and conditions.

## 4. Contract requirements

Contracts must include requirements proportionate to risk, including responsibilities, access restrictions, confidentiality, incident notification, audit or assurance rights, business continuity, subcontracting, data return or deletion and termination support where relevant.

## 5. Ongoing monitoring

Critical suppliers are reviewed at least annually and after a major incident or change. Monitoring may include performance results, incidents, complaints, assurance reports, certifications, financial or operational concerns and unresolved actions.

## 6. Offboarding

When a supplier relationship ends, the owner confirms account removal, data return or deletion, asset return, service transition, final invoices and closure of open risks.

## 7. Evidence

Attach one current critical-supplier agreement, due-diligence record or review:

{{ supplier.contract_sample | evidence }}
