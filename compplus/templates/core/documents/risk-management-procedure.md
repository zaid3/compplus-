# Integrated Risk Management Procedure

**Owner:** {{ roles.system_manager }}  
**Approved by:** {{ roles.executive_owner }}  
**Review cycle:** Annual and after significant change

## 1. Purpose

{{ organization.legal_name }} uses one consistent process to identify, assess, treat, monitor and communicate risks across information security, quality, privacy, environment and AI governance.

## 2. When a risk assessment is required

A risk assessment is created or reviewed when:

- a new product, service, system, supplier or processing activity is introduced;
- an important change is planned;
- an incident, complaint, audit finding or legal change occurs;
- an existing risk reaches its review date;
- management requests a review.

## 3. Risk description

Each risk record describes:

- the asset, process, objective or person affected;
- the threat, event or failure that may occur;
- the likely consequence;
- existing controls;
- the owner and review date.

## 4. Scoring method

CompPlus uses a default 5 × 5 method:

- **Likelihood:** 1 Rare, 2 Unlikely, 3 Possible, 4 Likely, 5 Almost certain.
- **Impact:** 1 Negligible, 2 Minor, 3 Moderate, 4 Major, 5 Severe.
- **Risk score:** likelihood multiplied by impact.

The organisation confirms its treatment threshold:

{{ risk.appetite | confirm }}

## 5. Treatment options

The risk owner selects one or more actions:

- reduce or mitigate;
- avoid;
- transfer or share;
- accept with documented approval.

Every treatment action has an owner, deadline and required evidence.

## 6. Residual risk

After treatment, the owner reassesses likelihood and impact. Residual risks above the approved threshold require acceptance by {{ roles.executive_owner }}.

## 7. Monitoring and review

High risks are reviewed at least quarterly. Other risks are reviewed at least annually or when material conditions change. Overdue reviews appear on the CompPlus dashboard.

## 8. Communication

Relevant risks and treatment actions are communicated to affected owners, workers, suppliers and senior management. Confidential details are restricted to people with a business need.

## 9. Evidence

Complete or import the initial risk register:

{{ risk.first_register | evidence }}
