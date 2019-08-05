package typeutil

import "go/types"

func FuncToParams(s *types.Func) []*types.Var {
	if s == nil {
		return nil
	}
	tuple := s.Type().(*types.Signature).Params()
	params := make([]*types.Var, tuple.Len())
	for i := range params {
		params[i] = tuple.At(i)
	}
	return params
}

func FuncToResults(s *types.Func) []*types.Var {
	if s == nil {
		return nil
	}
	tuple := s.Type().(*types.Signature).Results()
	results := make([]*types.Var, tuple.Len())
	for i := range results {
		results[i] = tuple.At(i)
	}
	return results
}
