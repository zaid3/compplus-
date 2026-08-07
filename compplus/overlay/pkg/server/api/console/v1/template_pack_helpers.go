package console_v1

import compplustemplates "go.probo.inc/probo/compplus/templates"

func compiledTemplateWorkCounts(compiled *compplustemplates.CompiledPack) (int, int) {
  tasks := 0
  evidences := 0
  for _, measure := range compiled.Measures {
    tasks += len(measure.Tasks)
    for _, task := range measure.Tasks {
      evidences += len(task.RequestedEvidences)
    }
  }
  return tasks, evidences
}
