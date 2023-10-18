package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.37

import (
	"context"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// FindSoftware is the resolver for the findSoftware field.
func (r *queryResolver) FindSoftware(ctx context.Context, searchText string) ([]model.PackageSourceOrArtifact, error) {
	return r.Backend.FindSoftware(ctx, searchText)
}

// FindTopLevelPackagesRelatedToVulnerability is the resolver for the findTopLevelPackagesRelatedToVulnerability field.
func (r *queryResolver) FindTopLevelPackagesRelatedToVulnerability(ctx context.Context, vulnerabilityID string) ([][]model.Node, error) {
	return r.Backend.FindTopLevelPackagesRelatedToVulnerability(ctx, vulnerabilityID)
}

// FindVulnerability is the resolver for the findVulnerability field.
func (r *queryResolver) FindVulnerability(ctx context.Context, purl string) ([]model.CertifyVulnOrCertifyVEXStatement, error) {
	return r.Backend.FindVulnerability(ctx, purl)
	//return nil, nil
}
