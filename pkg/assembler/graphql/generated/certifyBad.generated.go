// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

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

func (ec *executionContext) _CertifyBad_id(ctx context.Context, field graphql.CollectedField, obj *model.CertifyBad) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyBad_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyBad_id(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyBad",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyBad_subject(ctx context.Context, field graphql.CollectedField, obj *model.CertifyBad) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyBad_subject(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Subject, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(model.PackageSourceOrArtifact)
	fc.Result = res
	return ec.marshalNPackageSourceOrArtifact2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPackageSourceOrArtifact(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyBad_subject(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyBad",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type PackageSourceOrArtifact does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyBad_justification(ctx context.Context, field graphql.CollectedField, obj *model.CertifyBad) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyBad_justification(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Justification, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyBad_justification(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyBad",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyBad_knownSince(ctx context.Context, field graphql.CollectedField, obj *model.CertifyBad) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyBad_knownSince(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.KnownSince, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(time.Time)
	fc.Result = res
	return ec.marshalNTime2timeᚐTime(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyBad_knownSince(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyBad",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Time does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyBad_origin(ctx context.Context, field graphql.CollectedField, obj *model.CertifyBad) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyBad_origin(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Origin, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyBad_origin(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyBad",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyBad_collector(ctx context.Context, field graphql.CollectedField, obj *model.CertifyBad) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyBad_collector(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Collector, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyBad_collector(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyBad",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyBad_documentRef(ctx context.Context, field graphql.CollectedField, obj *model.CertifyBad) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyBad_documentRef(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DocumentRef, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyBad_documentRef(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyBad",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyBadConnection_totalCount(ctx context.Context, field graphql.CollectedField, obj *model.CertifyBadConnection) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyBadConnection_totalCount(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.TotalCount, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyBadConnection_totalCount(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyBadConnection",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyBadConnection_pageInfo(ctx context.Context, field graphql.CollectedField, obj *model.CertifyBadConnection) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyBadConnection_pageInfo(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PageInfo, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.PageInfo)
	fc.Result = res
	return ec.marshalNPageInfo2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPageInfo(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyBadConnection_pageInfo(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyBadConnection",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "hasNextPage":
				return ec.fieldContext_PageInfo_hasNextPage(ctx, field)
			case "startCursor":
				return ec.fieldContext_PageInfo_startCursor(ctx, field)
			case "endCursor":
				return ec.fieldContext_PageInfo_endCursor(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type PageInfo", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyBadConnection_edges(ctx context.Context, field graphql.CollectedField, obj *model.CertifyBadConnection) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyBadConnection_edges(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Edges, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]*model.CertifyBadEdge)
	fc.Result = res
	return ec.marshalNCertifyBadEdge2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyBadEdgeᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyBadConnection_edges(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyBadConnection",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "cursor":
				return ec.fieldContext_CertifyBadEdge_cursor(ctx, field)
			case "node":
				return ec.fieldContext_CertifyBadEdge_node(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type CertifyBadEdge", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyBadEdge_cursor(ctx context.Context, field graphql.CollectedField, obj *model.CertifyBadEdge) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyBadEdge_cursor(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Cursor, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyBadEdge_cursor(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyBadEdge",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CertifyBadEdge_node(ctx context.Context, field graphql.CollectedField, obj *model.CertifyBadEdge) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CertifyBadEdge_node(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Node, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.CertifyBad)
	fc.Result = res
	return ec.marshalNCertifyBad2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyBad(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CertifyBadEdge_node(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CertifyBadEdge",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_CertifyBad_id(ctx, field)
			case "subject":
				return ec.fieldContext_CertifyBad_subject(ctx, field)
			case "justification":
				return ec.fieldContext_CertifyBad_justification(ctx, field)
			case "knownSince":
				return ec.fieldContext_CertifyBad_knownSince(ctx, field)
			case "origin":
				return ec.fieldContext_CertifyBad_origin(ctx, field)
			case "collector":
				return ec.fieldContext_CertifyBad_collector(ctx, field)
			case "documentRef":
				return ec.fieldContext_CertifyBad_documentRef(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type CertifyBad", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputCertifyBadInputSpec(ctx context.Context, obj any) (model.CertifyBadInputSpec, error) {
	var it model.CertifyBadInputSpec
	asMap := map[string]any{}
	for k, v := range obj.(map[string]any) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"justification", "knownSince", "origin", "collector", "documentRef"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "justification":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("justification"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Justification = data
		case "knownSince":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("knownSince"))
			data, err := ec.unmarshalNTime2timeᚐTime(ctx, v)
			if err != nil {
				return it, err
			}
			it.KnownSince = data
		case "origin":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("origin"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Origin = data
		case "collector":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("collector"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Collector = data
		case "documentRef":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("documentRef"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.DocumentRef = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputCertifyBadSpec(ctx context.Context, obj any) (model.CertifyBadSpec, error) {
	var it model.CertifyBadSpec
	asMap := map[string]any{}
	for k, v := range obj.(map[string]any) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"id", "subject", "justification", "knownSince", "origin", "collector", "documentRef"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "id":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
			data, err := ec.unmarshalOID2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.ID = data
		case "subject":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("subject"))
			data, err := ec.unmarshalOPackageSourceOrArtifactSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPackageSourceOrArtifactSpec(ctx, v)
			if err != nil {
				return it, err
			}
			it.Subject = data
		case "justification":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("justification"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.Justification = data
		case "knownSince":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("knownSince"))
			data, err := ec.unmarshalOTime2ᚖtimeᚐTime(ctx, v)
			if err != nil {
				return it, err
			}
			it.KnownSince = data
		case "origin":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("origin"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.Origin = data
		case "collector":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("collector"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.Collector = data
		case "documentRef":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("documentRef"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.DocumentRef = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputMatchFlags(ctx context.Context, obj any) (model.MatchFlags, error) {
	var it model.MatchFlags
	asMap := map[string]any{}
	for k, v := range obj.(map[string]any) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"pkg"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "pkg":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("pkg"))
			data, err := ec.unmarshalNPkgMatchType2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPkgMatchType(ctx, v)
			if err != nil {
				return it, err
			}
			it.Pkg = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputPackageSourceOrArtifactInput(ctx context.Context, obj any) (model.PackageSourceOrArtifactInput, error) {
	var it model.PackageSourceOrArtifactInput
	asMap := map[string]any{}
	for k, v := range obj.(map[string]any) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"package", "source", "artifact"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "package":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("package"))
			data, err := ec.unmarshalOIDorPkgInput2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐIDorPkgInput(ctx, v)
			if err != nil {
				return it, err
			}
			it.Package = data
		case "source":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("source"))
			data, err := ec.unmarshalOIDorSourceInput2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐIDorSourceInput(ctx, v)
			if err != nil {
				return it, err
			}
			it.Source = data
		case "artifact":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("artifact"))
			data, err := ec.unmarshalOIDorArtifactInput2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐIDorArtifactInput(ctx, v)
			if err != nil {
				return it, err
			}
			it.Artifact = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputPackageSourceOrArtifactInputs(ctx context.Context, obj any) (model.PackageSourceOrArtifactInputs, error) {
	var it model.PackageSourceOrArtifactInputs
	asMap := map[string]any{}
	for k, v := range obj.(map[string]any) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"packages", "sources", "artifacts"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "packages":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("packages"))
			data, err := ec.unmarshalOIDorPkgInput2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐIDorPkgInputᚄ(ctx, v)
			if err != nil {
				return it, err
			}
			it.Packages = data
		case "sources":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("sources"))
			data, err := ec.unmarshalOIDorSourceInput2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐIDorSourceInputᚄ(ctx, v)
			if err != nil {
				return it, err
			}
			it.Sources = data
		case "artifacts":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("artifacts"))
			data, err := ec.unmarshalOIDorArtifactInput2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐIDorArtifactInputᚄ(ctx, v)
			if err != nil {
				return it, err
			}
			it.Artifacts = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputPackageSourceOrArtifactSpec(ctx context.Context, obj any) (model.PackageSourceOrArtifactSpec, error) {
	var it model.PackageSourceOrArtifactSpec
	asMap := map[string]any{}
	for k, v := range obj.(map[string]any) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"package", "source", "artifact"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "package":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("package"))
			data, err := ec.unmarshalOPkgSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPkgSpec(ctx, v)
			if err != nil {
				return it, err
			}
			it.Package = data
		case "source":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("source"))
			data, err := ec.unmarshalOSourceSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSourceSpec(ctx, v)
			if err != nil {
				return it, err
			}
			it.Source = data
		case "artifact":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("artifact"))
			data, err := ec.unmarshalOArtifactSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐArtifactSpec(ctx, v)
			if err != nil {
				return it, err
			}
			it.Artifact = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

func (ec *executionContext) _PackageSourceOrArtifact(ctx context.Context, sel ast.SelectionSet, obj model.PackageSourceOrArtifact) graphql.Marshaler {
	switch obj := (obj).(type) {
	case nil:
		return graphql.Null
	case model.Package:
		return ec._Package(ctx, sel, &obj)
	case *model.Package:
		if obj == nil {
			return graphql.Null
		}
		return ec._Package(ctx, sel, obj)
	case model.Source:
		return ec._Source(ctx, sel, &obj)
	case *model.Source:
		if obj == nil {
			return graphql.Null
		}
		return ec._Source(ctx, sel, obj)
	case model.Artifact:
		return ec._Artifact(ctx, sel, &obj)
	case *model.Artifact:
		if obj == nil {
			return graphql.Null
		}
		return ec._Artifact(ctx, sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var certifyBadImplementors = []string{"CertifyBad", "Node"}

func (ec *executionContext) _CertifyBad(ctx context.Context, sel ast.SelectionSet, obj *model.CertifyBad) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, certifyBadImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("CertifyBad")
		case "id":
			out.Values[i] = ec._CertifyBad_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "subject":
			out.Values[i] = ec._CertifyBad_subject(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "justification":
			out.Values[i] = ec._CertifyBad_justification(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "knownSince":
			out.Values[i] = ec._CertifyBad_knownSince(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "origin":
			out.Values[i] = ec._CertifyBad_origin(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "collector":
			out.Values[i] = ec._CertifyBad_collector(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "documentRef":
			out.Values[i] = ec._CertifyBad_documentRef(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var certifyBadConnectionImplementors = []string{"CertifyBadConnection"}

func (ec *executionContext) _CertifyBadConnection(ctx context.Context, sel ast.SelectionSet, obj *model.CertifyBadConnection) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, certifyBadConnectionImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("CertifyBadConnection")
		case "totalCount":
			out.Values[i] = ec._CertifyBadConnection_totalCount(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "pageInfo":
			out.Values[i] = ec._CertifyBadConnection_pageInfo(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "edges":
			out.Values[i] = ec._CertifyBadConnection_edges(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var certifyBadEdgeImplementors = []string{"CertifyBadEdge"}

func (ec *executionContext) _CertifyBadEdge(ctx context.Context, sel ast.SelectionSet, obj *model.CertifyBadEdge) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, certifyBadEdgeImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("CertifyBadEdge")
		case "cursor":
			out.Values[i] = ec._CertifyBadEdge_cursor(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "node":
			out.Values[i] = ec._CertifyBadEdge_node(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNCertifyBad2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyBadᚄ(ctx context.Context, sel ast.SelectionSet, v []*model.CertifyBad) graphql.Marshaler {
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
			ret[i] = ec.marshalNCertifyBad2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyBad(ctx, sel, v[i])
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

func (ec *executionContext) marshalNCertifyBad2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyBad(ctx context.Context, sel ast.SelectionSet, v *model.CertifyBad) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._CertifyBad(ctx, sel, v)
}

func (ec *executionContext) marshalNCertifyBadEdge2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyBadEdgeᚄ(ctx context.Context, sel ast.SelectionSet, v []*model.CertifyBadEdge) graphql.Marshaler {
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
			ret[i] = ec.marshalNCertifyBadEdge2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyBadEdge(ctx, sel, v[i])
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

func (ec *executionContext) marshalNCertifyBadEdge2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyBadEdge(ctx context.Context, sel ast.SelectionSet, v *model.CertifyBadEdge) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._CertifyBadEdge(ctx, sel, v)
}

func (ec *executionContext) unmarshalNCertifyBadInputSpec2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyBadInputSpec(ctx context.Context, v any) (model.CertifyBadInputSpec, error) {
	res, err := ec.unmarshalInputCertifyBadInputSpec(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNCertifyBadInputSpec2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyBadInputSpecᚄ(ctx context.Context, v any) ([]*model.CertifyBadInputSpec, error) {
	var vSlice []any
	if v != nil {
		vSlice = graphql.CoerceList(v)
	}
	var err error
	res := make([]*model.CertifyBadInputSpec, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalNCertifyBadInputSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyBadInputSpec(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) unmarshalNCertifyBadInputSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyBadInputSpec(ctx context.Context, v any) (*model.CertifyBadInputSpec, error) {
	res, err := ec.unmarshalInputCertifyBadInputSpec(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNCertifyBadSpec2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyBadSpec(ctx context.Context, v any) (model.CertifyBadSpec, error) {
	res, err := ec.unmarshalInputCertifyBadSpec(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNMatchFlags2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐMatchFlags(ctx context.Context, v any) (model.MatchFlags, error) {
	res, err := ec.unmarshalInputMatchFlags(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNPackageSourceOrArtifact2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPackageSourceOrArtifact(ctx context.Context, sel ast.SelectionSet, v model.PackageSourceOrArtifact) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._PackageSourceOrArtifact(ctx, sel, v)
}

func (ec *executionContext) marshalNPackageSourceOrArtifact2ᚕgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPackageSourceOrArtifactᚄ(ctx context.Context, sel ast.SelectionSet, v []model.PackageSourceOrArtifact) graphql.Marshaler {
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
			ret[i] = ec.marshalNPackageSourceOrArtifact2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPackageSourceOrArtifact(ctx, sel, v[i])
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

func (ec *executionContext) unmarshalNPackageSourceOrArtifactInput2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPackageSourceOrArtifactInput(ctx context.Context, v any) (model.PackageSourceOrArtifactInput, error) {
	res, err := ec.unmarshalInputPackageSourceOrArtifactInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNPackageSourceOrArtifactInputs2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPackageSourceOrArtifactInputs(ctx context.Context, v any) (model.PackageSourceOrArtifactInputs, error) {
	res, err := ec.unmarshalInputPackageSourceOrArtifactInputs(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalNPkgMatchType2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPkgMatchType(ctx context.Context, v any) (model.PkgMatchType, error) {
	var res model.PkgMatchType
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNPkgMatchType2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPkgMatchType(ctx context.Context, sel ast.SelectionSet, v model.PkgMatchType) graphql.Marshaler {
	return v
}

func (ec *executionContext) marshalOCertifyBadConnection2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCertifyBadConnection(ctx context.Context, sel ast.SelectionSet, v *model.CertifyBadConnection) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._CertifyBadConnection(ctx, sel, v)
}

func (ec *executionContext) unmarshalOPackageSourceOrArtifactSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPackageSourceOrArtifactSpec(ctx context.Context, v any) (*model.PackageSourceOrArtifactSpec, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputPackageSourceOrArtifactSpec(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

// endregion ***************************** type.gotpl *****************************
