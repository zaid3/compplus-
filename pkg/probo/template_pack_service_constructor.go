// Copyright (c) 2026 CompPlus.
// Use of this source code is governed by the MIT license in the repository root.

package probo

// NewTemplatePackService creates the CompPlus template pack compiler and
// installer around an existing Probo service instance.
func NewTemplatePackService(svc *Service) *TemplatePackService {
	return &TemplatePackService{svc: svc}
}