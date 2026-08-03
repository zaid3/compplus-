// Copyright (c) 2026 CompPlus.
// Use of this source code is governed by the MIT license in the repository root.

package console_v1

import compplustemplates "go.probo.inc/probo/compplus/templates"

func compiledWorkCounts(compiled *compplustemplates.CompiledPack) (int, int) {
	tasksCount := 0
	evidenceRequestsCount := 0

	for _, measure := range compiled.Measures {
		tasksCount += len(measure.Tasks)
		for _, task := range measure.Tasks {
			evidenceRequestsCount += len(task.RequestedEvidences)
		}
	}

	return tasksCount, evidenceRequestsCount
}
