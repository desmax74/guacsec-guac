// Code generated by ent, DO NOT EDIT.

package license

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.License {
	return predicate.License(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.License {
	return predicate.License(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.License {
	return predicate.License(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.License {
	return predicate.License(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.License {
	return predicate.License(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.License {
	return predicate.License(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.License {
	return predicate.License(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.License {
	return predicate.License(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.License {
	return predicate.License(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.License {
	return predicate.License(sql.FieldEQ(FieldName, v))
}

// Inline applies equality check predicate on the "inline" field. It's identical to InlineEQ.
func Inline(v string) predicate.License {
	return predicate.License(sql.FieldEQ(FieldInline, v))
}

// ListVersion applies equality check predicate on the "list_version" field. It's identical to ListVersionEQ.
func ListVersion(v string) predicate.License {
	return predicate.License(sql.FieldEQ(FieldListVersion, v))
}

// InlineHash applies equality check predicate on the "inline_hash" field. It's identical to InlineHashEQ.
func InlineHash(v string) predicate.License {
	return predicate.License(sql.FieldEQ(FieldInlineHash, v))
}

// ListVersionHash applies equality check predicate on the "list_version_hash" field. It's identical to ListVersionHashEQ.
func ListVersionHash(v string) predicate.License {
	return predicate.License(sql.FieldEQ(FieldListVersionHash, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.License {
	return predicate.License(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.License {
	return predicate.License(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.License {
	return predicate.License(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.License {
	return predicate.License(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.License {
	return predicate.License(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.License {
	return predicate.License(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.License {
	return predicate.License(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.License {
	return predicate.License(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.License {
	return predicate.License(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.License {
	return predicate.License(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.License {
	return predicate.License(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.License {
	return predicate.License(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.License {
	return predicate.License(sql.FieldContainsFold(FieldName, v))
}

// InlineEQ applies the EQ predicate on the "inline" field.
func InlineEQ(v string) predicate.License {
	return predicate.License(sql.FieldEQ(FieldInline, v))
}

// InlineNEQ applies the NEQ predicate on the "inline" field.
func InlineNEQ(v string) predicate.License {
	return predicate.License(sql.FieldNEQ(FieldInline, v))
}

// InlineIn applies the In predicate on the "inline" field.
func InlineIn(vs ...string) predicate.License {
	return predicate.License(sql.FieldIn(FieldInline, vs...))
}

// InlineNotIn applies the NotIn predicate on the "inline" field.
func InlineNotIn(vs ...string) predicate.License {
	return predicate.License(sql.FieldNotIn(FieldInline, vs...))
}

// InlineGT applies the GT predicate on the "inline" field.
func InlineGT(v string) predicate.License {
	return predicate.License(sql.FieldGT(FieldInline, v))
}

// InlineGTE applies the GTE predicate on the "inline" field.
func InlineGTE(v string) predicate.License {
	return predicate.License(sql.FieldGTE(FieldInline, v))
}

// InlineLT applies the LT predicate on the "inline" field.
func InlineLT(v string) predicate.License {
	return predicate.License(sql.FieldLT(FieldInline, v))
}

// InlineLTE applies the LTE predicate on the "inline" field.
func InlineLTE(v string) predicate.License {
	return predicate.License(sql.FieldLTE(FieldInline, v))
}

// InlineContains applies the Contains predicate on the "inline" field.
func InlineContains(v string) predicate.License {
	return predicate.License(sql.FieldContains(FieldInline, v))
}

// InlineHasPrefix applies the HasPrefix predicate on the "inline" field.
func InlineHasPrefix(v string) predicate.License {
	return predicate.License(sql.FieldHasPrefix(FieldInline, v))
}

// InlineHasSuffix applies the HasSuffix predicate on the "inline" field.
func InlineHasSuffix(v string) predicate.License {
	return predicate.License(sql.FieldHasSuffix(FieldInline, v))
}

// InlineIsNil applies the IsNil predicate on the "inline" field.
func InlineIsNil() predicate.License {
	return predicate.License(sql.FieldIsNull(FieldInline))
}

// InlineNotNil applies the NotNil predicate on the "inline" field.
func InlineNotNil() predicate.License {
	return predicate.License(sql.FieldNotNull(FieldInline))
}

// InlineEqualFold applies the EqualFold predicate on the "inline" field.
func InlineEqualFold(v string) predicate.License {
	return predicate.License(sql.FieldEqualFold(FieldInline, v))
}

// InlineContainsFold applies the ContainsFold predicate on the "inline" field.
func InlineContainsFold(v string) predicate.License {
	return predicate.License(sql.FieldContainsFold(FieldInline, v))
}

// ListVersionEQ applies the EQ predicate on the "list_version" field.
func ListVersionEQ(v string) predicate.License {
	return predicate.License(sql.FieldEQ(FieldListVersion, v))
}

// ListVersionNEQ applies the NEQ predicate on the "list_version" field.
func ListVersionNEQ(v string) predicate.License {
	return predicate.License(sql.FieldNEQ(FieldListVersion, v))
}

// ListVersionIn applies the In predicate on the "list_version" field.
func ListVersionIn(vs ...string) predicate.License {
	return predicate.License(sql.FieldIn(FieldListVersion, vs...))
}

// ListVersionNotIn applies the NotIn predicate on the "list_version" field.
func ListVersionNotIn(vs ...string) predicate.License {
	return predicate.License(sql.FieldNotIn(FieldListVersion, vs...))
}

// ListVersionGT applies the GT predicate on the "list_version" field.
func ListVersionGT(v string) predicate.License {
	return predicate.License(sql.FieldGT(FieldListVersion, v))
}

// ListVersionGTE applies the GTE predicate on the "list_version" field.
func ListVersionGTE(v string) predicate.License {
	return predicate.License(sql.FieldGTE(FieldListVersion, v))
}

// ListVersionLT applies the LT predicate on the "list_version" field.
func ListVersionLT(v string) predicate.License {
	return predicate.License(sql.FieldLT(FieldListVersion, v))
}

// ListVersionLTE applies the LTE predicate on the "list_version" field.
func ListVersionLTE(v string) predicate.License {
	return predicate.License(sql.FieldLTE(FieldListVersion, v))
}

// ListVersionContains applies the Contains predicate on the "list_version" field.
func ListVersionContains(v string) predicate.License {
	return predicate.License(sql.FieldContains(FieldListVersion, v))
}

// ListVersionHasPrefix applies the HasPrefix predicate on the "list_version" field.
func ListVersionHasPrefix(v string) predicate.License {
	return predicate.License(sql.FieldHasPrefix(FieldListVersion, v))
}

// ListVersionHasSuffix applies the HasSuffix predicate on the "list_version" field.
func ListVersionHasSuffix(v string) predicate.License {
	return predicate.License(sql.FieldHasSuffix(FieldListVersion, v))
}

// ListVersionIsNil applies the IsNil predicate on the "list_version" field.
func ListVersionIsNil() predicate.License {
	return predicate.License(sql.FieldIsNull(FieldListVersion))
}

// ListVersionNotNil applies the NotNil predicate on the "list_version" field.
func ListVersionNotNil() predicate.License {
	return predicate.License(sql.FieldNotNull(FieldListVersion))
}

// ListVersionEqualFold applies the EqualFold predicate on the "list_version" field.
func ListVersionEqualFold(v string) predicate.License {
	return predicate.License(sql.FieldEqualFold(FieldListVersion, v))
}

// ListVersionContainsFold applies the ContainsFold predicate on the "list_version" field.
func ListVersionContainsFold(v string) predicate.License {
	return predicate.License(sql.FieldContainsFold(FieldListVersion, v))
}

// InlineHashEQ applies the EQ predicate on the "inline_hash" field.
func InlineHashEQ(v string) predicate.License {
	return predicate.License(sql.FieldEQ(FieldInlineHash, v))
}

// InlineHashNEQ applies the NEQ predicate on the "inline_hash" field.
func InlineHashNEQ(v string) predicate.License {
	return predicate.License(sql.FieldNEQ(FieldInlineHash, v))
}

// InlineHashIn applies the In predicate on the "inline_hash" field.
func InlineHashIn(vs ...string) predicate.License {
	return predicate.License(sql.FieldIn(FieldInlineHash, vs...))
}

// InlineHashNotIn applies the NotIn predicate on the "inline_hash" field.
func InlineHashNotIn(vs ...string) predicate.License {
	return predicate.License(sql.FieldNotIn(FieldInlineHash, vs...))
}

// InlineHashGT applies the GT predicate on the "inline_hash" field.
func InlineHashGT(v string) predicate.License {
	return predicate.License(sql.FieldGT(FieldInlineHash, v))
}

// InlineHashGTE applies the GTE predicate on the "inline_hash" field.
func InlineHashGTE(v string) predicate.License {
	return predicate.License(sql.FieldGTE(FieldInlineHash, v))
}

// InlineHashLT applies the LT predicate on the "inline_hash" field.
func InlineHashLT(v string) predicate.License {
	return predicate.License(sql.FieldLT(FieldInlineHash, v))
}

// InlineHashLTE applies the LTE predicate on the "inline_hash" field.
func InlineHashLTE(v string) predicate.License {
	return predicate.License(sql.FieldLTE(FieldInlineHash, v))
}

// InlineHashContains applies the Contains predicate on the "inline_hash" field.
func InlineHashContains(v string) predicate.License {
	return predicate.License(sql.FieldContains(FieldInlineHash, v))
}

// InlineHashHasPrefix applies the HasPrefix predicate on the "inline_hash" field.
func InlineHashHasPrefix(v string) predicate.License {
	return predicate.License(sql.FieldHasPrefix(FieldInlineHash, v))
}

// InlineHashHasSuffix applies the HasSuffix predicate on the "inline_hash" field.
func InlineHashHasSuffix(v string) predicate.License {
	return predicate.License(sql.FieldHasSuffix(FieldInlineHash, v))
}

// InlineHashIsNil applies the IsNil predicate on the "inline_hash" field.
func InlineHashIsNil() predicate.License {
	return predicate.License(sql.FieldIsNull(FieldInlineHash))
}

// InlineHashNotNil applies the NotNil predicate on the "inline_hash" field.
func InlineHashNotNil() predicate.License {
	return predicate.License(sql.FieldNotNull(FieldInlineHash))
}

// InlineHashEqualFold applies the EqualFold predicate on the "inline_hash" field.
func InlineHashEqualFold(v string) predicate.License {
	return predicate.License(sql.FieldEqualFold(FieldInlineHash, v))
}

// InlineHashContainsFold applies the ContainsFold predicate on the "inline_hash" field.
func InlineHashContainsFold(v string) predicate.License {
	return predicate.License(sql.FieldContainsFold(FieldInlineHash, v))
}

// ListVersionHashEQ applies the EQ predicate on the "list_version_hash" field.
func ListVersionHashEQ(v string) predicate.License {
	return predicate.License(sql.FieldEQ(FieldListVersionHash, v))
}

// ListVersionHashNEQ applies the NEQ predicate on the "list_version_hash" field.
func ListVersionHashNEQ(v string) predicate.License {
	return predicate.License(sql.FieldNEQ(FieldListVersionHash, v))
}

// ListVersionHashIn applies the In predicate on the "list_version_hash" field.
func ListVersionHashIn(vs ...string) predicate.License {
	return predicate.License(sql.FieldIn(FieldListVersionHash, vs...))
}

// ListVersionHashNotIn applies the NotIn predicate on the "list_version_hash" field.
func ListVersionHashNotIn(vs ...string) predicate.License {
	return predicate.License(sql.FieldNotIn(FieldListVersionHash, vs...))
}

// ListVersionHashGT applies the GT predicate on the "list_version_hash" field.
func ListVersionHashGT(v string) predicate.License {
	return predicate.License(sql.FieldGT(FieldListVersionHash, v))
}

// ListVersionHashGTE applies the GTE predicate on the "list_version_hash" field.
func ListVersionHashGTE(v string) predicate.License {
	return predicate.License(sql.FieldGTE(FieldListVersionHash, v))
}

// ListVersionHashLT applies the LT predicate on the "list_version_hash" field.
func ListVersionHashLT(v string) predicate.License {
	return predicate.License(sql.FieldLT(FieldListVersionHash, v))
}

// ListVersionHashLTE applies the LTE predicate on the "list_version_hash" field.
func ListVersionHashLTE(v string) predicate.License {
	return predicate.License(sql.FieldLTE(FieldListVersionHash, v))
}

// ListVersionHashContains applies the Contains predicate on the "list_version_hash" field.
func ListVersionHashContains(v string) predicate.License {
	return predicate.License(sql.FieldContains(FieldListVersionHash, v))
}

// ListVersionHashHasPrefix applies the HasPrefix predicate on the "list_version_hash" field.
func ListVersionHashHasPrefix(v string) predicate.License {
	return predicate.License(sql.FieldHasPrefix(FieldListVersionHash, v))
}

// ListVersionHashHasSuffix applies the HasSuffix predicate on the "list_version_hash" field.
func ListVersionHashHasSuffix(v string) predicate.License {
	return predicate.License(sql.FieldHasSuffix(FieldListVersionHash, v))
}

// ListVersionHashIsNil applies the IsNil predicate on the "list_version_hash" field.
func ListVersionHashIsNil() predicate.License {
	return predicate.License(sql.FieldIsNull(FieldListVersionHash))
}

// ListVersionHashNotNil applies the NotNil predicate on the "list_version_hash" field.
func ListVersionHashNotNil() predicate.License {
	return predicate.License(sql.FieldNotNull(FieldListVersionHash))
}

// ListVersionHashEqualFold applies the EqualFold predicate on the "list_version_hash" field.
func ListVersionHashEqualFold(v string) predicate.License {
	return predicate.License(sql.FieldEqualFold(FieldListVersionHash, v))
}

// ListVersionHashContainsFold applies the ContainsFold predicate on the "list_version_hash" field.
func ListVersionHashContainsFold(v string) predicate.License {
	return predicate.License(sql.FieldContainsFold(FieldListVersionHash, v))
}

// HasDeclaredInCertifyLegals applies the HasEdge predicate on the "declared_in_certify_legals" edge.
func HasDeclaredInCertifyLegals() predicate.License {
	return predicate.License(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, DeclaredInCertifyLegalsTable, DeclaredInCertifyLegalsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeclaredInCertifyLegalsWith applies the HasEdge predicate on the "declared_in_certify_legals" edge with a given conditions (other predicates).
func HasDeclaredInCertifyLegalsWith(preds ...predicate.CertifyLegal) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		step := newDeclaredInCertifyLegalsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDiscoveredInCertifyLegals applies the HasEdge predicate on the "discovered_in_certify_legals" edge.
func HasDiscoveredInCertifyLegals() predicate.License {
	return predicate.License(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, DiscoveredInCertifyLegalsTable, DiscoveredInCertifyLegalsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDiscoveredInCertifyLegalsWith applies the HasEdge predicate on the "discovered_in_certify_legals" edge with a given conditions (other predicates).
func HasDiscoveredInCertifyLegalsWith(preds ...predicate.CertifyLegal) predicate.License {
	return predicate.License(func(s *sql.Selector) {
		step := newDiscoveredInCertifyLegalsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.License) predicate.License {
	return predicate.License(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.License) predicate.License {
	return predicate.License(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.License) predicate.License {
	return predicate.License(sql.NotPredicates(p))
}
