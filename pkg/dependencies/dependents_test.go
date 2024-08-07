//
// Copyright 2024 The GUAC Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build integration

package dependencies

import (
	"context"
	"fmt"
	"testing"
	"time"

	clients "github.com/guacsec/guac/internal/testing/graphqlClients"

	model "github.com/guacsec/guac/pkg/assembler/clients/generated"

	"github.com/Khan/genqlient/graphql"
)

type testPkgInputSpec struct {
	id       string
	typeName string
	name     string
	version  string
}

type createIsDependency struct {
	pkg    specInfo
	depPkg specInfo
}

type specInfo struct {
	name     string
	version  string
	typeName string
}

func Test_findAllDependencies(t *testing.T) {

	t.Run("default", func(t *testing.T) {
		gqlClient := clients.SetupTest(t)
		ctx := context.Background()

		createNodes(t, ctx, gqlClient,
			[]specInfo{{"A1", "1", "test"}, {"B1", "1", "test"}, {"C1", "1", "test"},
				{"D1", "1", "test"}, {"E1", "1", "test"}, {"F1", "1", "test"}, {"G1", "1", "test"}},
			[]createIsDependency{
				{pkg: specInfo{"B1", "1", "test"}, depPkg: specInfo{"A1", "1", "test"}},
				{pkg: specInfo{"C1", "1", "test"}, depPkg: specInfo{"A1", "1", "test"}},
				{pkg: specInfo{"D1", "1", "test"}, depPkg: specInfo{"B1", "1", "test"}},
				{pkg: specInfo{"E1", "1", "test"}, depPkg: specInfo{"B1", "1", "test"}},
				{pkg: specInfo{"F1", "1", "test"}, depPkg: specInfo{"C1", "1", "test"}},
				{pkg: specInfo{"G1", "1", "test"}, depPkg: specInfo{"C1", "1", "test"}},
			}, specInfo{name: "A1", version: "1", typeName: "test"})

		/*
			The graph should look like:

			D1  E1 F1  G1
			 \ /    \ /
			  B1   C1
			   \  /
				A1

			D1 and E1 depend on B1
			F1 and G1 depend on C1
			D1, E1, F1, G1, B1 and C1 all depend on A1
		*/

		check(t, false, map[string]int{"pkg:test/A1": 6, "pkg:test/B1": 2, "pkg:test/C1": 2}, gqlClient)
	})

	t.Run("two separate graphs", func(t *testing.T) {
		gqlClient := clients.SetupTest(t)
		ctx := context.Background()

		createNodes(t, ctx, gqlClient,
			[]specInfo{{"A2", "1", "test"}, {"B2", "1", "test"}, {"C2", "1", "test"}, {"D2", "1", "test"},
				{"E2", "1", "test"}, {"F2", "1", "test"}, {"G2", "1", "test"}, {"B2", "2", "test"}, {"I2", "1", "test"}},
			[]createIsDependency{
				{pkg: specInfo{"B2", "1", "test"}, depPkg: specInfo{"A2", "1", "test"}},
				{pkg: specInfo{"C2", "1", "test"}, depPkg: specInfo{"A2", "1", "test"}},
				{pkg: specInfo{"D2", "1", "test"}, depPkg: specInfo{"B2", "1", "test"}},
				{pkg: specInfo{"E2", "1", "test"}, depPkg: specInfo{"B2", "1", "test"}},
				{pkg: specInfo{"F2", "1", "test"}, depPkg: specInfo{"C2", "1", "test"}},
				{pkg: specInfo{"G2", "1", "test"}, depPkg: specInfo{"C2", "1", "test"}},
				{pkg: specInfo{"I2", "1", "test"}, depPkg: specInfo{"B2", "2", "test"}},
			}, specInfo{name: "A2", version: "1", typeName: "test"})

		/*
			The graph should look like:

			D2  E2 F2  G2    I2
			 \ /    \ /       \
			  B2   C2         B2
			   \  /
				A2

			The graph is now made up of two separate graphs that aren't connected.

			There are two B2s.
			The first one has 2 dependencies D2 and E2 and is dependent on A2.
			The second one has 1 dependency I2 and isn't dependent on anything.

			The first B2 has a version of 1 and the second B2 has a version of 2.

			Even through they have two different versions and are two different nodes
			they will both be grouped together while calculating dependencies because
			they have the same Name.
		*/

		check(t, false, map[string]int{"pkg:test/A2": 6, "pkg:test/B2": 3, "pkg:test/C2": 2}, gqlClient)
	})

	t.Run("two different versions in same graph", func(t *testing.T) {
		gqlClient := clients.SetupTest(t)
		ctx := context.Background()

		createNodes(t, ctx, gqlClient,
			[]specInfo{
				{"A3", "1", "test"},
				{"B3", "1", "test"}, // B3 version 1
				{"C3", "1", "test"},
				{"D3", "1", "test"},
				{"E3", "1", "test"},
				{"F3", "1", "test"},
				{"G3", "1", "test"},
				{"B3", "2", "test"}, // B3 version 2
				{"I3", "1", "test"},
			},
			[]createIsDependency{
				{pkg: specInfo{"B3", "1", "test"}, depPkg: specInfo{"A3", "1", "test"}},
				{pkg: specInfo{"C3", "1", "test"}, depPkg: specInfo{"A3", "1", "test"}},
				{pkg: specInfo{"D3", "1", "test"}, depPkg: specInfo{"B3", "1", "test"}},
				{pkg: specInfo{"E3", "1", "test"}, depPkg: specInfo{"B3", "1", "test"}},
				{pkg: specInfo{"F3", "1", "test"}, depPkg: specInfo{"C3", "1", "test"}},
				{pkg: specInfo{"G3", "1", "test"}, depPkg: specInfo{"C3", "1", "test"}},
				{pkg: specInfo{"I3", "1", "test"}, depPkg: specInfo{"B3", "2", "test"}},
				{pkg: specInfo{"B3", "2", "test"}, depPkg: specInfo{"A3", "1", "test"}},
			}, specInfo{name: "A3", version: "1", typeName: "test"})

		/*
				The graph should look like:

				D3  E3 F3  G3   I3
				 \ /    \ /    /
				  B3    C3    B3
				   \   /      |
			        A3 -------+
		*/

		check(t, false, map[string]int{"pkg:test/A3": 8, "pkg:test/B3": 3, "pkg:test/C3": 2}, gqlClient)
	})

	t.Run("different versions same graph 2", func(t *testing.T) {
		gqlClient := clients.SetupTest(t)
		ctx := context.Background()

		createNodes(t, ctx, gqlClient,
			[]specInfo{
				{"A4", "1", "test"},
				{"B4", "1", "test"}, // B4 version 1
				{"C4", "1", "test"},
				{"D4", "1", "test"},
				{"E4", "1", "test"},
				{"F4", "1", "test"},
				{"G4", "1", "test"},
				{"B4", "2", "test"}, // B4 version 2
				{"I4", "1", "test"},
			},
			[]createIsDependency{
				{pkg: specInfo{"B4", "1", "test"}, depPkg: specInfo{"A4", "1", "test"}},
				{pkg: specInfo{"C4", "1", "test"}, depPkg: specInfo{"A4", "1", "test"}},
				{pkg: specInfo{"D4", "1", "test"}, depPkg: specInfo{"B4", "1", "test"}},
				{pkg: specInfo{"E4", "1", "test"}, depPkg: specInfo{"B4", "1", "test"}},
				{pkg: specInfo{"F4", "1", "test"}, depPkg: specInfo{"C4", "1", "test"}},
				{pkg: specInfo{"G4", "1", "test"}, depPkg: specInfo{"C4", "1", "test"}},
				{pkg: specInfo{"I4", "1", "test"}, depPkg: specInfo{"B4", "2", "test"}},
				{pkg: specInfo{"B4", "2", "test"}, depPkg: specInfo{"F4", "1", "test"}},
			}, specInfo{name: "A4", version: "1", typeName: "test"})

		/*
				The graph should look like:

					I4
					 \
					  B4
					   \
				D4  E4 F4  G4
				 \ /    \ /
				  B4    C4
				   \   /
			        A4
		*/

		check(t, false, map[string]int{"pkg:test/A4": 8, "pkg:test/B4": 3, "pkg:test/C4": 4, "pkg:test/F4": 2}, gqlClient)

	})
}

// createNodes creates all the nodes for each test
func createNodes(t *testing.T, ctx context.Context, gqlClient graphql.Client, pkgSpecInfo []specInfo, dependencyConnections []createIsDependency, info specInfo) {
	packages := map[string]testPkgInputSpec{}

	packageIds, err := createPackageNodes(ctx, gqlClient, packages, pkgSpecInfo)

	if err != nil {
		t.Fatalf("failed to create package nodes: %v", err)
	}

	dependencyIds, err := createDependencyNodes(ctx, gqlClient, packages, dependencyConnections)

	if err != nil {
		t.Fatalf("failed to create dependency nodes: %v", err)
	}

	err, id := createOccurrenceAndArtifact(t, ctx, gqlClient, info)

	if err != nil {
		t.Fatalf("failed to create occurrence and artifact: %v", err)
	}

	err = ingestHasSBOM(ctx, gqlClient, dependencyIds, packageIds, []string{id.IngestOccurrence})

	if err != nil {
		t.Fatalf("failed to ingest hasSBOM: %v", err)
	}
}

func ingestHasSBOM(ctx context.Context, client graphql.Client, dependencyIds, softwareIds, occIds []string) error {
	// logger := logging.FromContext(ctx)
	tm, _ := time.Parse(time.RFC3339, "2022-11-21T17:45:50.52Z")
	opensslNs := "openssl.org"
	opensslVersion := "3.0.3"
	ingestHasSBOMStruct := []struct {
		name     string
		pkg      *model.PkgInputSpec
		artifact *model.ArtifactInputSpec
		hasSBOM  model.HasSBOMInputSpec
		includes model.HasSBOMIncludesInputSpec
	}{
		{
			name: "test",
			pkg: &model.PkgInputSpec{
				Type:       "conan",
				Namespace:  &opensslNs,
				Name:       "test",
				Version:    &opensslVersion,
				Qualifiers: []model.PackageQualifierInputSpec{{Key: "user", Value: "bincrafters"}, {Key: "channel", Value: "stable"}},
			},
			hasSBOM: model.HasSBOMInputSpec{
				Uri:              "uri:location of package SBOM",
				Algorithm:        "sha256",
				Digest:           "6bbb0da1891646e58eb3e6a63af3a6fc3c8eb5a0d44824cba581d2e14a0450cf",
				DownloadLocation: "uri: download location of the SBOM",
				Origin:           "Demo ingestion",
				Collector:        "Demo ingestion",
				KnownSince:       tm,
			},
			includes: model.HasSBOMIncludesInputSpec{
				Dependencies: dependencyIds,
				Occurrences:  occIds,
				Packages:     softwareIds,
				Artifacts:    []string{},
			},
		},
	}
	for _, ingest := range ingestHasSBOMStruct {
		if ingest.pkg != nil {
			if _, err := model.IngestPackage(ctx, client, model.IDorPkgInput{PackageInput: ingest.pkg}); err != nil {
				return fmt.Errorf("Error in ingesting package: %v\n", err)
			}

			if _, err := model.IngestHasSBOMPkg(ctx, client, model.IDorPkgInput{PackageInput: ingest.pkg}, ingest.hasSBOM, ingest.includes); err != nil {
				return fmt.Errorf("Error in ingesting: %v\n", err)
			}
		} else {
			fmt.Printf("input missing for package or source")
		}
	}

	return nil
}

// createOccurrenceAndArtifact creates a test artifact and a test occurrence.
// The values in artifact and occurrence are fake values and won't work for anything other than tests
func createOccurrenceAndArtifact(t *testing.T, ctx context.Context, gqlClient graphql.Client, info specInfo) (error, *model.IngestIsOccurrencePkgResponse) {
	_, err := model.IngestArtifact(ctx, gqlClient, model.IDorArtifactInput{
		ArtifactInput: &model.ArtifactInputSpec{
			Algorithm: "sha265",
			Digest:    "123",
		},
	})

	if err != nil {
		t.Fatalf("%v", err)
	}

	id, err := model.IngestIsOccurrencePkg(ctx, gqlClient, model.IDorPkgInput{
		PackageInput: &model.PkgInputSpec{
			Type:    info.typeName,
			Name:    info.name,
			Version: &info.version,
		},
	}, model.IDorArtifactInput{
		ArtifactInput: &model.ArtifactInputSpec{
			Algorithm: "sha265",
			Digest:    "123",
		},
	}, model.IsOccurrenceInputSpec{
		Justification: "test-justification",
		Origin:        "test-origin",
		Collector:     "test-collector",
	})

	if err != nil {
		t.Fatalf("%v", err)
	}

	return err, id
}

func check(t *testing.T, wantErr bool, want map[string]int, gqlClient graphql.Client) {
	ctx := context.Background()

	got, err := findDependentsOfDependencies(ctx, gqlClient)
	if (err != nil) != wantErr {
		t.Errorf("findDependentsOfDependencies() error = %v, wantErr %v", err, wantErr)
		return
	}

	for k, v := range want {
		if v != len(got[k].dependents) {
			t.Errorf("findDependentsOfDependencies() for node %v, got %v dependencies, want %v dependencies", k, len(got[k].dependents), v)
		}
	}
}

func createPackageNodes(ctx context.Context, gqlClient graphql.Client, packages map[string]testPkgInputSpec, pkgSpecInfo []specInfo) ([]string, error) {
	var res []string
	for _, info := range pkgSpecInfo {
		spec := model.IDorPkgInput{
			PackageInput: &model.PkgInputSpec{
				Type:    info.typeName,
				Name:    info.name,
				Version: &info.version,
			},
		}
		id, err := model.IngestPackage(ctx, gqlClient, spec)
		if err != nil {
			return nil, fmt.Errorf("failed to ingest package: %v", err)
		}
		packages[spec.PackageInput.Type+spec.PackageInput.Name+"Version"+*spec.PackageInput.Version] = testPkgInputSpec{
			id:       id.IngestPackage.PackageVersionID,
			name:     spec.PackageInput.Name,
			typeName: spec.PackageInput.Type,
			version:  info.version,
		}
		res = append(res, id.IngestPackage.PackageVersionID)
	}

	return res, nil
}

func createDependencyNodes(ctx context.Context, gqlClient graphql.Client, packages map[string]testPkgInputSpec, dependencyConnections []createIsDependency) ([]string, error) {
	var ids []string
	for _, connection := range dependencyConnections {
		pkgName := connection.pkg.typeName + connection.pkg.name + "Version" + connection.pkg.version
		depPkgName := connection.depPkg.typeName + connection.depPkg.name + "Version" + connection.depPkg.version

		pkgVersion := packages[pkgName].version
		pkSpec := model.IDorPkgInput{
			PackageInput: &model.PkgInputSpec{
				Type:    packages[pkgName].typeName,
				Name:    packages[pkgName].name,
				Version: &pkgVersion,
			},
		}

		depPkgVersion := packages[depPkgName].version
		depPkgSpec := model.IDorPkgInput{
			PackageInput: &model.PkgInputSpec{
				Type:    packages[depPkgName].typeName,
				Name:    packages[depPkgName].name,
				Version: &depPkgVersion,
			},
		}

		id, err := model.IngestIsDependency(
			ctx,
			gqlClient,
			pkSpec,
			depPkgSpec,
			model.IsDependencyInputSpec{
				Justification:  "test",
				Origin:         "test",
				Collector:      "test",
				DependencyType: model.DependencyTypeUnknown,
			})

		if err != nil {
			return nil, fmt.Errorf("failed to ingest isDependency: %v", err)
		}
		ids = append(ids, id.IngestDependency)
	}

	return ids, nil
}
