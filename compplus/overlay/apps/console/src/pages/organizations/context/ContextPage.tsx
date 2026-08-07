import {
  Button,
  Card,
  IconCheckmark1,
  IconCrossLargeX,
  IconPencil,
  Markdown,
  Textarea,
} from "@probo/ui";
import { useRef, useState } from "react";
import { useTranslation } from "react-i18next";
import { useFragment } from "react-relay";
import { Link } from "react-router";
import { graphql } from "relay-runtime";

import type { ContextPage_UpdateMutation } from "#/__generated__/core/ContextPage_UpdateMutation.graphql";
import type { ContextPageFragment$key } from "#/__generated__/core/ContextPageFragment.graphql";
import { useMutationWithToasts } from "#/hooks/useMutationWithToasts";

const fragment = graphql`
  fragment ContextPageFragment on Organization {
    id
    canUpdateContext: permission(action: "core:organization-context:update")
    context {
      product
      architecture
      team
      processes
      customers
    }
  }
`;

const updateMutation = graphql`
  mutation ContextPage_UpdateMutation(
    $input: UpdateOrganizationContextInput!
  ) {
    updateOrganizationContext(input: $input) {
      context {
        organizationId
        product
        architecture
        team
        processes
        customers
      }
    }
  }
`;

type SectionKey = "product" | "architecture" | "team" | "processes" | "customers";

type SectionConfig = {
  key: SectionKey;
  title: string;
  description: string;
  placeholder: string;
};

type Props = {
  organization: ContextPageFragment$key;
};

const startSteps = [
  {
    number: "1",
    title: "Tell us what your organisation does",
    text: "Complete the Context sections below. Plain English is enough: what you provide, the systems you use, your team, main processes and customers.",
    to: null,
    action: "You are here",
  },
  {
    number: "2",
    title: "Choose the standards you need",
    text: "Open Documents → Template Library. Select ISO 27001, ISO 9001, UK GDPR, ISO 14001 or ISO 42001. The shared Core pack is included automatically.",
    to: "../documents",
    action: "Open Documents",
  },
  {
    number: "3",
    title: "Review the ready-made paperwork",
    text: "ISOpilot prepares the policies, procedures, registers and forms. Read them and change only the [CONFIRM] points so the wording matches what your organisation really does.",
    to: "../documents",
    action: "Review documents",
  },
  {
    number: "4",
    title: "Review risks and controls",
    text: "Open the framework and risks. Confirm which requirements apply, record real risks, assign owners and work through the measures and tasks ISOpilot created.",
    to: "../risks",
    action: "Review risks",
  },
  {
    number: "5",
    title: "Attach proof",
    text: "Policies describe what should happen. Evidence proves it actually happens. Upload approvals, training records, supplier reviews, screenshots, logs, completed forms and other real records.",
    to: "../tasks",
    action: "Open tasks",
  },
  {
    number: "6",
    title: "Audit and close gaps",
    text: "Run an internal audit, record findings and corrective actions, then verify that each action fixed the underlying problem.",
    to: "../audits",
    action: "Open audits",
  },
  {
    number: "7",
    title: "Management review and maintain",
    text: "Approve the final documents, complete management review and keep evidence current. ISOpilot helps you stay audit-ready; certification is still awarded by an independent certification body.",
    to: "../documents",
    action: "Continue",
  },
] as const;

export default function ContextPage(props: Props) {
  const { t } = useTranslation();
  const organization = useFragment(fragment, props.organization);
  const context = organization.context;

  const sections: SectionConfig[] = [
    {
      key: "product",
      title: t("context.sections.product.title"),
      description: t("context.sections.product.description"),
      placeholder: t("context.sections.product.placeholder"),
    },
    {
      key: "architecture",
      title: t("context.sections.architecture.title"),
      description: t("context.sections.architecture.description"),
      placeholder: t("context.sections.architecture.placeholder"),
    },
    {
      key: "team",
      title: t("context.sections.team.title"),
      description: t("context.sections.team.description"),
      placeholder: t("context.sections.team.placeholder"),
    },
    {
      key: "processes",
      title: t("context.sections.processes.title"),
      description: t("context.sections.processes.description"),
      placeholder: t("context.sections.processes.placeholder"),
    },
    {
      key: "customers",
      title: t("context.sections.customers.title"),
      description: t("context.sections.customers.description"),
      placeholder: t("context.sections.customers.placeholder"),
    },
  ];

  const values: Record<SectionKey, string | null> = {
    product: context?.product ?? null,
    architecture: context?.architecture ?? null,
    team: context?.team ?? null,
    processes: context?.processes ?? null,
    customers: context?.customers ?? null,
  };

  const contextCompleted = Object.values(values).filter(Boolean).length;

  return (
    <div className="space-y-6">
      <Card padded>
        <div className="space-y-5">
          <div className="flex flex-wrap items-start justify-between gap-3">
            <div>
              <div className="text-xs font-semibold uppercase tracking-wide text-txt-accent">ISOpilot guided setup</div>
              <h2 className="mt-1 text-xl font-semibold text-txt-primary">Start here — no compliance experience needed</h2>
              <p className="mt-1 max-w-3xl text-sm text-txt-secondary">
                Your compliance co-pilot shows the next move. Follow these steps in order; you confirm what is true, ISOpilot organises the paperwork, tasks and evidence around it.
              </p>
            </div>
            <div className="rounded-full border border-border-low bg-subtle px-3 py-1.5 text-xs font-medium text-txt-secondary">
              Context {contextCompleted}/5 completed
            </div>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-3">
            {startSteps.map(step => (
              <div key={step.number} className="rounded-xl border border-border-low bg-level-1 p-4">
                <div className="flex items-start gap-3">
                  <div className="flex size-7 shrink-0 items-center justify-center rounded-full bg-active text-xs font-semibold text-txt-accent">
                    {step.number}
                  </div>
                  <div className="min-w-0">
                    <h3 className="text-sm font-semibold text-txt-primary">{step.title}</h3>
                    <p className="mt-1 text-xs leading-5 text-txt-secondary">{step.text}</p>
                    {step.to
                      ? (
                          <Link to={step.to} className="mt-2 inline-block text-xs font-medium text-txt-accent underline underline-offset-2">
                            {step.action} →
                          </Link>
                        )
                      : (
                          <span className="mt-2 inline-block text-xs font-medium text-txt-secondary">{step.action}</span>
                        )}
                  </div>
                </div>
              </div>
            ))}
          </div>

          <div className="grid grid-cols-1 lg:grid-cols-2 gap-3">
            <div className="rounded-xl border border-border-low bg-subtle p-3 text-xs text-txt-secondary">
              <strong className="text-txt-primary">Policy vs evidence</strong><span className="block mt-1">A policy says what the organisation intends to do. Evidence shows it actually happened—for example a completed access review, training record, supplier assessment or audit record.</span>
            </div>
            <div className="rounded-xl border border-border-low bg-subtle p-3 text-xs text-txt-secondary">
              <strong className="text-txt-primary">What “audit-ready” means here</strong><span className="block mt-1">Requirements are mapped, documents reflect real practice, tasks have owners, evidence is current, gaps are tracked and management has reviewed the system.</span>
            </div>
          </div>
        </div>
      </Card>

      {sections.map(section => (
        <ContextSection
          key={section.key}
          section={section}
          organizationId={organization.id}
          value={values[section.key]}
          canEdit={organization.canUpdateContext}
        />
      ))}
    </div>
  );
}

function ContextSection({
  section,
  organizationId,
  value,
  canEdit,
}: {
  section: SectionConfig;
  organizationId: string;
  value: string | null;
  canEdit: boolean;
}) {
  const { t } = useTranslation();
  const [isEditing, setIsEditing] = useState(false);
  const [text, setText] = useState(value ?? "");
  const [displayedValue, setDisplayedValue] = useState(value ?? "");
  const justSavedRef = useRef(false);

  const [updateContext, isUpdating]
    = useMutationWithToasts<ContextPage_UpdateMutation>(
      updateMutation,
      {
        successMessage: t("context.messages.updated"),
        errorMessage: t("context.errors.update"),
      },
    );

  const handleSave = async () => {
    const valueToSave = text.trim();
    const previousValue = value ?? "";
    setDisplayedValue(valueToSave);
    justSavedRef.current = true;

    const valueToSend = valueToSave.length > 0 ? valueToSave : null;

    await updateContext({
      variables: {
        input: {
          organizationId,
          [section.key]: valueToSend,
        },
      },
      onError: () => {
        setDisplayedValue(previousValue);
        justSavedRef.current = false;
      },
      onCompleted: (_, errors) => {
        if (errors?.length) {
          setDisplayedValue(previousValue);
          justSavedRef.current = false;
        }
        setIsEditing(false);
      },
    });
  };

  const handleCancel = () => {
    setText(value ?? "");
    setIsEditing(false);
  };

  return (
    <Card padded>
      {isEditing
        ? (
            <div className="space-y-4">
              <div>
                <h3 className="text-sm font-semibold">{section.title}</h3>
                <p className="text-xs text-txt-tertiary mt-1">
                  {section.description}
                </p>
              </div>
              <Textarea
                value={text}
                onChange={e => setText(e.target.value)}
                autogrow
                className="min-h-32 font-mono text-sm"
                placeholder={section.placeholder}
              />
              <div className="flex gap-2 justify-end">
                <Button
                  variant="secondary"
                  icon={IconCrossLargeX}
                  onClick={handleCancel}
                  disabled={isUpdating}
                >
                  {t("context.actions.cancel")}
                </Button>
                <Button
                  icon={IconCheckmark1}
                  onClick={() => void handleSave()}
                  disabled={isUpdating}
                >
                  {t("context.actions.save")}
                </Button>
              </div>
            </div>
          )
        : (
            <div className="space-y-2">
              <div className="flex items-center justify-between">
                <div>
                  <h3 className="text-sm font-semibold">{section.title}</h3>
                  <p className="text-xs text-txt-tertiary mt-1">
                    {section.description}
                  </p>
                </div>
                {canEdit && (
                  <Button
                    variant="quaternary"
                    icon={IconPencil}
                    onClick={() => {
                      setText(value ?? "");
                      setIsEditing(true);
                    }}
                  >
                    {t("context.actions.edit")}
                  </Button>
                )}
              </div>
              <div className="w-full">
                {displayedValue
                  ? (
                      <div className="prose prose-sm max-w-none w-full [&_.prose]:max-w-none">
                        <Markdown content={displayedValue} />
                      </div>
                    )
                  : (
                      <div className="text-txt-tertiary text-sm italic">
                        {t("context.empty")}
                      </div>
                    )}
              </div>
            </div>
          )}
    </Card>
  );
}
