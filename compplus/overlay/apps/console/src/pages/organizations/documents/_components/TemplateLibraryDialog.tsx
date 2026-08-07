import { formatError } from "@probo/helpers";
import {
  Badge,
  Breadcrumb,
  Button,
  Dialog,
  DialogContent,
  DialogFooter,
  Input,
  Label,
  Textarea,
  useDialogRef,
  useToast,
} from "@probo/ui";
import { type ReactNode, useMemo, useState } from "react";
import { graphql, useMutation } from "react-relay";
import { z } from "zod";

import type { TemplateLibraryDialogInstallMutation } from "#/__generated__/core/TemplateLibraryDialogInstallMutation.graphql";
import { useFormWithSchema } from "#/hooks/useFormWithSchema";
import { useOrganizationId } from "#/hooks/useOrganizationId";

const installTemplatePackMutation = graphql`
  mutation TemplateLibraryDialogInstallMutation($input: InstallTemplatePackInput!) {
    installTemplatePack(input: $input) {
      packId
      version
      documentsCreated
      measuresCreated
      tasksCreated
      evidenceRequestsCreated
      statementOfApplicabilityCreated
      alreadyInstalled
      framework {
        id
        name
      }
    }
  }
`;

const packOptions = [
  {
    id: "core",
    name: "Core management pack",
    standard: "Shared foundation",
    description: "Context, scope, roles, risk, document control, training, suppliers, incidents, audit, management review and corrective action.",
  },
  {
    id: "iso27001",
    name: "ISO/IEC 27001 Fast Start",
    standard: "2022 + Amendment 1:2024",
    description: "ISMS policies, risk methodology/register/treatment, security procedures and a native 93-control Statement of Applicability starter.",
  },
  {
    id: "iso9001",
    name: "ISO 9001 Fast Start",
    standard: "Current: 2015 + Amendment 1:2024",
    description: "Quality policy, process map, objectives, customer/supplier controls, service delivery, complaints, audit, corrective action and improvement.",
  },
  {
    id: "uk-gdpr",
    name: "UK GDPR Fast Start",
    standard: "Current UK law + DUAA changes in force in 2026",
    description: "ROPA, lawful basis, DPIA, privacy notices, retention, rights/SARs, complaints, breaches, processors, data sharing and transfers.",
  },
  {
    id: "iso14001",
    name: "ISO 14001 Fast Start",
    standard: "2026 edition",
    description: "EMS policy, aspects/impacts, compliance obligations, objectives, waste/resources, emergencies, monitoring, audit and management review.",
  },
  {
    id: "iso42001",
    name: "ISO/IEC 42001 Fast Start",
    standard: "2023 edition",
    description: "AI inventory, impact/risk, data, human oversight, suppliers, testing, fairness, security, monitoring, change and incidents.",
  },
] as const;

type PackId = (typeof packOptions)[number]["id"];

const formSchema = z.object({
  legalName: z.string().trim().min(1, "Company or organisation name is required"),
  services: z.string().trim().min(1, "Add a short description of your products or services"),
  locations: z.string().trim().min(1, "Add the locations or remote operation covered"),
  executiveOwner: z.string().trim().min(1, "Add the senior owner"),
  systemManager: z.string().trim().min(1, "Add the compliance/management-system owner"),
});

type Props = {
  trigger?: ReactNode;
  organizationName: string;
};

export function TemplateLibraryDialog({ trigger, organizationName }: Props) {
  const organizationId = useOrganizationId();
  const dialogRef = useDialogRef();
  const { toast } = useToast();
  const [selected, setSelected] = useState<Set<PackId>>(new Set<PackId>(["core"]));
  const [installPack, isInstalling] = useMutation<TemplateLibraryDialogInstallMutation>(installTemplatePackMutation);

  const { register, handleSubmit, formState } = useFormWithSchema(formSchema, {
    defaultValues: {
      legalName: organizationName,
      services: "",
      locations: "",
      executiveOwner: "",
      systemManager: "",
    },
  });

  const selectedPacks = useMemo(
    () => packOptions.filter(pack => selected.has(pack.id)),
    [selected],
  );

  const togglePack = (id: PackId) => {
    if (id === "core") return;
    setSelected(current => {
      const next = new Set(current);
      if (next.has(id)) next.delete(id);
      else next.add(id);
      return next;
    });
  };

  const installOne = (packId: PackId, data: z.infer<typeof formSchema>) => new Promise<{ alreadyInstalled: boolean; documentsCreated: number; soa: boolean }>((resolve, reject) => {
    installPack({
      variables: {
        input: {
          organizationId,
          packId,
          answers: {
            legalName: data.legalName,
            services: data.services,
            locations: data.locations,
            executiveOwner: data.executiveOwner,
            systemManager: data.systemManager,
            securityOwner: null,
            privacyOwner: null,
            qualityOwner: null,
            environmentalOwner: null,
            aiOwner: null,
          },
        },
      },
      onCompleted(response, errors) {
        if (errors?.length) {
          reject(new Error(formatError("Template pack could not be installed", errors)));
          return;
        }
        resolve({
          alreadyInstalled: response.installTemplatePack.alreadyInstalled,
          documentsCreated: response.installTemplatePack.documentsCreated,
          soa: response.installTemplatePack.statementOfApplicabilityCreated,
        });
      },
      onError(error) {
        reject(error);
      },
    });
  });

  const onSubmit = async (data: z.infer<typeof formSchema>) => {
    try {
      let documentsCreated = 0;
      let installed = 0;
      let skipped = 0;
      let soaCreated = false;

      for (const pack of selectedPacks) {
        const result = await installOne(pack.id, data);
        if (result.alreadyInstalled) skipped += 1;
        else installed += 1;
        documentsCreated += result.documentsCreated;
        soaCreated ||= result.soa;
      }

      toast({
        title: "ISOpilot templates installed",
        description: `${installed} pack${installed === 1 ? "" : "s"} installed, ${documentsCreated} ready-made document${documentsCreated === 1 ? "" : "s"} created${soaCreated ? ", including the ISO 27001 Statement of Applicability starter" : ""}${skipped ? `. ${skipped} already-installed pack${skipped === 1 ? " was" : "s were"} left unchanged.` : "."}`,
        variant: "success",
      });
      dialogRef.current?.close();
      window.location.reload();
    } catch (error) {
      toast({
        title: "Template installation failed",
        description: error instanceof Error ? error.message : "The selected template packs could not be installed.",
        variant: "error",
      });
    }
  };

  return (
    <Dialog
      ref={dialogRef}
      trigger={trigger}
      title={<Breadcrumb items={["Documents", "ISOpilot Template Library"]} />}
    >
      <form onSubmit={event => void handleSubmit(onSubmit)(event)}>
        <DialogContent padded className="max-h-[76vh] overflow-y-auto space-y-7">
          <section className="space-y-3">
            <div>
              <div className="text-xs font-semibold uppercase tracking-wide text-txt-accent">Your compliance co-pilot</div>
              <h2 className="mt-1 text-xl font-semibold text-txt-primary">Start with ready-made compliance paperwork</h2>
              <p className="mt-1 max-w-3xl text-sm text-txt-secondary">
                You do not need to know which policy or register comes first. Choose the standards you need, answer five organisation questions, and ISOpilot prepares the working documents, controls, tasks and evidence requests for you.
              </p>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-3 gap-2">
              <div className="rounded-lg border border-border-low bg-subtle p-3 text-sm"><strong className="text-txt-primary">1. Review</strong><span className="block mt-1 text-xs text-txt-secondary">Read the prepared wording and confirm it matches your organisation.</span></div>
              <div className="rounded-lg border border-border-low bg-subtle p-3 text-sm"><strong className="text-txt-primary">2. Personalise</strong><span className="block mt-1 text-xs text-txt-secondary">Only change items marked [CONFIRM] or organisation-specific details.</span></div>
              <div className="rounded-lg border border-border-low bg-subtle p-3 text-sm"><strong className="text-txt-primary">3. Prove</strong><span className="block mt-1 text-xs text-txt-secondary">Complete tasks and attach real evidence before approving the document.</span></div>
            </div>
          </section>

          <section className="space-y-3">
            <div className="flex items-center justify-between gap-3">
              <div>
                <h3 className="font-semibold text-txt-primary">1. Choose your standards</h3>
                <p className="text-sm text-txt-secondary">The shared Core management pack is always included so common work is completed once and reused.</p>
              </div>
              <Badge variant="neutral" size="md">{selectedPacks.length} selected</Badge>
            </div>

            <div className="grid grid-cols-1 lg:grid-cols-2 gap-3">
              {packOptions.map(pack => {
                const checked = selected.has(pack.id);
                const required = pack.id === "core";
                return (
                  <label
                    key={pack.id}
                    className={`block rounded-xl border p-4 cursor-pointer transition-colors ${checked ? "border-border-active bg-active" : "border-border-low bg-level-1 hover:bg-subtle"}`}
                  >
                    <div className="flex items-start gap-3">
                      <input
                        type="checkbox"
                        checked={checked}
                        disabled={required}
                        onChange={() => togglePack(pack.id)}
                        className="mt-1 size-4 accent-[#2F6BFF]"
                      />
                      <div className="min-w-0">
                        <div className="flex flex-wrap items-center gap-2">
                          <span className="font-medium text-txt-primary">{pack.name}</span>
                          {required && <Badge variant="neutral" size="sm">Included</Badge>}
                        </div>
                        <div className="mt-0.5 text-xs font-medium text-txt-accent">{pack.standard}</div>
                        <p className="mt-1.5 text-sm text-txt-secondary">{pack.description}</p>
                      </div>
                    </div>
                  </label>
                );
              })}
            </div>
          </section>

          <section className="space-y-4">
            <div>
              <h3 className="font-semibold text-txt-primary">2. Tell ISOpilot once</h3>
              <p className="text-sm text-txt-secondary">These five answers are reused across every selected document. Specialist owners default to the compliance owner and can be changed later.</p>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <Field label="Organisation name" error={formState.errors.legalName?.message}>
                <Input {...register("legalName")} required />
              </Field>
              <Field label="Senior / executive owner" error={formState.errors.executiveOwner?.message}>
                <Input {...register("executiveOwner")} required placeholder="e.g. Managing Director" />
              </Field>
              <Field label="Products and services" error={formState.errors.services?.message}>
                <Textarea {...register("services")} autogrow required placeholder="e.g. education and student support services" />
              </Field>
              <Field label="Locations covered" error={formState.errors.locations?.message}>
                <Textarea {...register("locations")} autogrow required placeholder="e.g. London office and approved remote-working locations" />
              </Field>
              <Field label="Compliance / management-system owner" error={formState.errors.systemManager?.message}>
                <Input {...register("systemManager")} required placeholder="Person or job title" />
              </Field>
            </div>
          </section>

          <div className="rounded-xl border border-border-low bg-subtle p-4 text-sm text-txt-secondary">
            <strong className="text-txt-primary">What happens after Install?</strong> ISOpilot creates editable documents, framework controls, measures, tasks and evidence requests in the normal workspace. ISO 27001 also creates a native Statement of Applicability starter. Company details and owners are prefilled; remaining decisions are clearly marked <strong>[CONFIRM]</strong>. The templates use original practical wording rather than reproducing copyrighted ISO standard text.
          </div>
        </DialogContent>
        <DialogFooter>
          <Button type="submit" disabled={isInstalling || selectedPacks.length === 0}>
            {isInstalling ? "Preparing your workspace…" : `Install ${selectedPacks.length} selected pack${selectedPacks.length === 1 ? "" : "s"}`}
          </Button>
        </DialogFooter>
      </form>
    </Dialog>
  );
}

function Field({ label, error, children }: { label: string; error?: string; children: ReactNode }) {
  return (
    <label className="block space-y-1.5">
      <Label>{label}</Label>
      {children}
      {error && <span className="block text-xs text-danger">{error}</span>}
    </label>
  );
}
