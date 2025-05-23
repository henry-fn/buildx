package userfunc

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

// A ContextFunc is a callback used to produce the base EvalContext for
// running a particular set of functions.
//
// This is a function rather than an EvalContext directly to allow functions
// to be decoded before their context is complete. This will be true, for
// example, for applications that wish to allow functions to refer to themselves.
//
// The simplest use of a ContextFunc is to give user functions access to the
// same global variables and functions available elsewhere in an application's
// configuration language, but more complex applications may use different
// contexts to support lexical scoping depending on where in a configuration
// structure a function declaration is found, etc.
type ContextFunc func() *hcl.EvalContext

// NewFunction creates a new function instance from preparsed HCL expressions.
func NewFunction(paramsExpr, varParamExpr, resultExpr hcl.Expression, getBaseCtx func() *hcl.EvalContext) (function.Function, hcl.Diagnostics) {
	var params []string
	var varParam string

	paramExprs, paramsDiags := hcl.ExprList(paramsExpr)
	if paramsDiags.HasErrors() {
		return function.Function{}, paramsDiags
	}
	for _, paramExpr := range paramExprs {
		param := hcl.ExprAsKeyword(paramExpr)
		if param == "" {
			return function.Function{}, hcl.Diagnostics{{
				Severity: hcl.DiagError,
				Summary:  "Invalid param element",
				Detail:   "Each parameter name must be an identifier.",
				Subject:  paramExpr.Range().Ptr(),
			}}
		}
		params = append(params, param)
	}

	if varParamExpr != nil {
		varParam = hcl.ExprAsKeyword(varParamExpr)
		if varParam == "" {
			return function.Function{}, hcl.Diagnostics{{
				Severity: hcl.DiagError,
				Summary:  "Invalid variadic_params",
				Detail:   "The variadic parameter name must be an identifier.",
				Subject:  varParamExpr.Range().Ptr(),
			}}
		}
	}

	spec := &function.Spec{}
	for _, paramName := range params {
		spec.Params = append(spec.Params, function.Parameter{
			Name: paramName,
			Type: cty.DynamicPseudoType,
		})
	}
	if varParamExpr != nil {
		spec.VarParam = &function.Parameter{
			Name: varParam,
			Type: cty.DynamicPseudoType,
		}
	}
	impl := func(args []cty.Value) (cty.Value, error) {
		ctx := getBaseCtx()
		ctx = ctx.NewChild()
		ctx.Variables = make(map[string]cty.Value)

		// The cty function machinery guarantees that we have at least
		// enough args to fill all of our params.
		for i, paramName := range params {
			ctx.Variables[paramName] = args[i]
		}
		if spec.VarParam != nil {
			varArgs := args[len(params):]
			ctx.Variables[varParam] = cty.TupleVal(varArgs)
		}

		result, diags := resultExpr.Value(ctx)
		if diags.HasErrors() {
			// Smuggle the diagnostics out via the error channel, since
			// a diagnostics sequence implements error. Caller can
			// type-assert this to recover the individual diagnostics
			// if desired.
			return cty.DynamicVal, diags
		}
		return result, nil
	}
	spec.Type = func(args []cty.Value) (cty.Type, error) {
		val, err := impl(args)
		return val.Type(), err
	}
	spec.Impl = func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		return impl(args)
	}
	return function.New(spec), nil
}
