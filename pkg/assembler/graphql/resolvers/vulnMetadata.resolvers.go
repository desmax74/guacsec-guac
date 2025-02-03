package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"
	"strings"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// IngestVulnerabilityMetadata is the resolver for the ingestVulnerabilityMetadata field.
func (r *mutationResolver) IngestVulnerabilityMetadata(ctx context.Context, vulnerability model.IDorVulnerabilityInput, vulnerabilityMetadata model.VulnerabilityMetadataInputSpec) (string, error) {
	funcName := "IngestVulnerabilityMetadata"

	if vulnerability.VulnerabilityInput != nil {
		err := validateNoVul(*vulnerability.VulnerabilityInput)
		if err != nil {
			return "", gqlerror.Errorf("%v ::  %s", funcName, err)
		}

		err = validateVulnerabilityIDInputSpec(*vulnerability.VulnerabilityInput)
		if err != nil {
			return "", gqlerror.Errorf("%v ::  %s", funcName, err)
		}
		return r.Backend.IngestVulnerabilityMetadata(ctx, model.IDorVulnerabilityInput{
			VulnerabilityTypeID: vulnerability.VulnerabilityTypeID,
			VulnerabilityNodeID: vulnerability.VulnerabilityNodeID,
			VulnerabilityInput:  &model.VulnerabilityInputSpec{Type: strings.ToLower(vulnerability.VulnerabilityInput.Type), VulnerabilityID: strings.ToLower(vulnerability.VulnerabilityInput.VulnerabilityID)},
		}, vulnerabilityMetadata)
	} else {
		return r.Backend.IngestVulnerabilityMetadata(ctx, vulnerability, vulnerabilityMetadata)
	}
}

// IngestBulkVulnerabilityMetadata is the resolver for the ingestBulkVulnerabilityMetadata field.
func (r *mutationResolver) IngestBulkVulnerabilityMetadata(ctx context.Context, vulnerabilities []*model.IDorVulnerabilityInput, vulnerabilityMetadataList []*model.VulnerabilityMetadataInputSpec) ([]string, error) {
	funcName := "IngestVulnerabilityMetadatas"
	if len(vulnerabilities) != len(vulnerabilityMetadataList) {
		return []string{}, gqlerror.Errorf("%v :: uneven vulnerabilities and vulnerabilityMetadata for ingestion", funcName)
	}

	// vulnerability input (type and vulnerability ID) will be enforced to be lowercase
	var lowercaseVulnInputList []*model.IDorVulnerabilityInput
	for _, v := range vulnerabilities {
		if v.VulnerabilityInput == nil {
			lowercaseVulnInputList = append(lowercaseVulnInputList, v)
			continue
		}
		err := validateNoVul(*v.VulnerabilityInput)
		if err != nil {
			return []string{}, gqlerror.Errorf("%v ::  %s", funcName, err)
		}

		err = validateVulnerabilityIDInputSpec(*v.VulnerabilityInput)
		if err != nil {
			return []string{}, gqlerror.Errorf("%v ::  %s", funcName, err)
		}

		lowercaseVulnInput := model.VulnerabilityInputSpec{
			Type:            strings.ToLower(v.VulnerabilityInput.Type),
			VulnerabilityID: strings.ToLower(v.VulnerabilityInput.VulnerabilityID),
		}
		lowercaseVulnInputList = append(lowercaseVulnInputList, &model.IDorVulnerabilityInput{
			VulnerabilityTypeID: v.VulnerabilityTypeID,
			VulnerabilityNodeID: v.VulnerabilityNodeID,
			VulnerabilityInput:  &lowercaseVulnInput,
		})
	}
	return r.Backend.IngestBulkVulnerabilityMetadata(ctx, lowercaseVulnInputList, vulnerabilityMetadataList)
}

// VulnerabilityMetadata is the resolver for the vulnerabilityMetadata field.
func (r *queryResolver) VulnerabilityMetadata(ctx context.Context, vulnerabilityMetadataSpec model.VulnerabilityMetadataSpec) ([]*model.VulnerabilityMetadata, error) {
	funcName := "VulnerabilityMetadata"
	// vulnerability input (type and vulnerability ID) will be enforced to be lowercase

	if vulnerabilityMetadataSpec.Comparator != nil && vulnerabilityMetadataSpec.ScoreValue == nil {
		return []*model.VulnerabilityMetadata{}, gqlerror.Errorf("%v :: comparator cannot be set without a score value specified", funcName)
	}

	if vulnerabilityMetadataSpec.Vulnerability != nil {

		var typeLowerCase *string = nil
		var vulnIDLowerCase *string = nil
		if vulnerabilityMetadataSpec.Vulnerability.Type != nil {
			lower := strings.ToLower(*vulnerabilityMetadataSpec.Vulnerability.Type)
			typeLowerCase = &lower
		}
		if vulnerabilityMetadataSpec.Vulnerability.VulnerabilityID != nil {
			lower := strings.ToLower(*vulnerabilityMetadataSpec.Vulnerability.VulnerabilityID)
			vulnIDLowerCase = &lower
		}

		err := validateVulnerabilitySpec(*vulnerabilityMetadataSpec.Vulnerability)
		if err != nil {
			return []*model.VulnerabilityMetadata{}, gqlerror.Errorf("%v ::  %s", funcName, err)
		}

		lowercaseVulnFilter := model.VulnerabilitySpec{
			ID:              vulnerabilityMetadataSpec.Vulnerability.ID,
			Type:            typeLowerCase,
			VulnerabilityID: vulnIDLowerCase,
			NoVuln:          vulnerabilityMetadataSpec.Vulnerability.NoVuln,
		}

		lowercaseVulnerabilityMetadataSpec := model.VulnerabilityMetadataSpec{
			ID:            vulnerabilityMetadataSpec.ID,
			Vulnerability: &lowercaseVulnFilter,
			ScoreType:     vulnerabilityMetadataSpec.ScoreType,
			ScoreValue:    vulnerabilityMetadataSpec.ScoreValue,
			Comparator:    vulnerabilityMetadataSpec.Comparator,
			Timestamp:     vulnerabilityMetadataSpec.Timestamp,
			Origin:        vulnerabilityMetadataSpec.Origin,
			Collector:     vulnerabilityMetadataSpec.Collector,
		}
		return r.Backend.VulnerabilityMetadata(ctx, &lowercaseVulnerabilityMetadataSpec)
	} else {
		return r.Backend.VulnerabilityMetadata(ctx, &vulnerabilityMetadataSpec)
	}
}

// VulnerabilityMetadataList is the resolver for the vulnerabilityMetadataList field.
func (r *queryResolver) VulnerabilityMetadataList(ctx context.Context, vulnerabilityMetadataSpec model.VulnerabilityMetadataSpec, after *string, first *int) (*model.VulnerabilityMetadataConnection, error) {
	funcName := "VulnerabilityMetadataList"
	// vulnerability input (type and vulnerability ID) will be enforced to be lowercase

	if vulnerabilityMetadataSpec.Comparator != nil && vulnerabilityMetadataSpec.ScoreValue == nil {
		return nil, gqlerror.Errorf("%v :: comparator cannot be set without a score value specified", funcName)
	}

	if vulnerabilityMetadataSpec.Vulnerability != nil {

		var typeLowerCase *string = nil
		var vulnIDLowerCase *string = nil
		if vulnerabilityMetadataSpec.Vulnerability.Type != nil {
			lower := strings.ToLower(*vulnerabilityMetadataSpec.Vulnerability.Type)
			typeLowerCase = &lower
		}
		if vulnerabilityMetadataSpec.Vulnerability.VulnerabilityID != nil {
			lower := strings.ToLower(*vulnerabilityMetadataSpec.Vulnerability.VulnerabilityID)
			vulnIDLowerCase = &lower
		}

		err := validateVulnerabilitySpec(*vulnerabilityMetadataSpec.Vulnerability)
		if err != nil {
			return nil, gqlerror.Errorf("%v ::  %s", funcName, err)
		}

		lowercaseVulnFilter := model.VulnerabilitySpec{
			ID:              vulnerabilityMetadataSpec.Vulnerability.ID,
			Type:            typeLowerCase,
			VulnerabilityID: vulnIDLowerCase,
			NoVuln:          vulnerabilityMetadataSpec.Vulnerability.NoVuln,
		}

		lowercaseVulnerabilityMetadataSpec := model.VulnerabilityMetadataSpec{
			ID:            vulnerabilityMetadataSpec.ID,
			Vulnerability: &lowercaseVulnFilter,
			ScoreType:     vulnerabilityMetadataSpec.ScoreType,
			ScoreValue:    vulnerabilityMetadataSpec.ScoreValue,
			Comparator:    vulnerabilityMetadataSpec.Comparator,
			Timestamp:     vulnerabilityMetadataSpec.Timestamp,
			Origin:        vulnerabilityMetadataSpec.Origin,
			Collector:     vulnerabilityMetadataSpec.Collector,
		}
		return r.Backend.VulnerabilityMetadataList(ctx, lowercaseVulnerabilityMetadataSpec, after, first)
	} else {
		return r.Backend.VulnerabilityMetadataList(ctx, vulnerabilityMetadataSpec, after, first)
	}
}
