# Management System Scope

**Document owner:** {{ roles.system_manager }}  
**Approved by:** {{ roles.executive_owner }}  
**Effective date:** {{ document.effective_date }}  
**Next review:** {{ document.next_review_date }}

## 1. Organisation

{{ organization.legal_name }} provides {{ organization.services }} from the following locations and working arrangements:

{{ organization.locations }}

## 2. Purpose

This document defines the boundaries of the integrated management system operated by {{ organization.legal_name }}. It is used to coordinate the compliance packs selected in CompPlus and to make responsibilities, evidence and audit boundaries clear.

## 3. Included activities

The management system covers:

- delivery and support of the products and services described above;
- people, processes, information, technology, facilities and suppliers needed to deliver those services;
- relevant legal, regulatory and contractual obligations;
- activities performed by employees, temporary workers and contractors within the defined scope;
- selected management-system standards and regulatory packs enabled in CompPlus.

## 4. Included locations and systems

The following locations, systems or environments are included:

{{ scope.included_locations_and_systems | confirm }}

## 5. Interfaces and dependencies

The scope depends on the following important internal teams and external providers:

{{ scope.dependencies | confirm }}

## 6. Exclusions

{{ scope.exclusions | confirm }}

For every exclusion, the organisation records why it is outside scope and confirms that the exclusion does not prevent the management system from achieving its intended outcomes.

## 7. Interested parties

The organisation considers the needs of customers, workers, regulators, certification bodies, suppliers, owners and other relevant parties. Detailed needs and obligations are maintained in the Interested Parties and Legal Requirements registers.

## 8. Approval

By approving this document, senior management confirms that the scope accurately reflects the current organisation, services, locations and dependencies.

**Approval decision:** {{ approval.status | confirm }}  
**Approver:** {{ roles.executive_owner }}  
**Approval date:** {{ approval.date | confirm }}
