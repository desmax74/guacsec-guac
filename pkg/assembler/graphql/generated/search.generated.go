// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"sync"

	"github.com/99designs/gqlgen/graphql"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

func (ec *executionContext) _CertifyVulnOrCertifyVEXStatement(ctx context.Context, sel ast.SelectionSet, obj model.CertifyVulnOrCertifyVEXStatement) graphql.Marshaler {
	switch obj := (obj).(type) {
	case nil:
		return graphql.Null
	case model.CertifyVuln:
		return ec._CertifyVuln(ctx, sel, &obj)
	case *model.CertifyVuln:
		if obj == nil {
			return graphql.Null
		}
		return ec._CertifyVuln(ctx, sel, obj)
	case model.CertifyVEXStatement:
		return ec._CertifyVEXStatement(ctx, sel, &obj)
	case *model.CertifyVEXStatement:
		if obj == nil {
			return graphql.Null
		}
		return ec._CertifyVEXStatement(ctx, sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNCertifyVulnOrCertifyVEXStatement2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyVulnOrCertifyVEXStatement(ctx context.Context, sel ast.SelectionSet, v model.CertifyVulnOrCertifyVEXStatement) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._CertifyVulnOrCertifyVEXStatement(ctx, sel, v)
}

func (ec *executionContext) marshalNCertifyVulnOrCertifyVEXStatement2ᚕgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyVulnOrCertifyVEXStatementᚄ(ctx context.Context, sel ast.SelectionSet, v []model.CertifyVulnOrCertifyVEXStatement) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNCertifyVulnOrCertifyVEXStatement2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyVulnOrCertifyVEXStatement(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

// endregion ***************************** type.gotpl *****************************