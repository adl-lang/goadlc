package gen_go

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/adl-lang/goadl_rt/v3/sys/adlast"
	"github.com/adl-lang/goadlc/internal/gen_go/fn/slices"
)

func typeParamsFromDecl(decl adlast.Decl) typeParam {
	return adlast.Handle_DeclType[typeParam](
		decl.Type_,
		func(struct_ adlast.Struct) typeParam {
			return typeParam{
				ps:               struct_.TypeParams,
				type_constraints: []string{},
				added:            false,
			}
		},
		func(union_ adlast.Union) typeParam {
			return typeParam{
				ps:               union_.TypeParams,
				type_constraints: []string{},
				added:            false,
			}
		},
		func(type_ adlast.TypeDef) typeParam {
			return typeParam{
				ps:               type_.TypeParams,
				type_constraints: []string{},
				added:            false,
			}
		},
		func(newtype_ adlast.NewType) typeParam {
			return typeParam{
				ps:               newtype_.TypeParams,
				type_constraints: []string{},
				added:            false,
			}
		},
		nil,
	)
}

type typeParam struct {
	ps               []string
	type_constraints []string
	added            bool
	stdlib           bool
}

func (tp typeParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(tp.ps)
}

func (tp typeParam) AddParam(newp string) typeParam {
	psMap := make(map[string]bool)
	tp0 := make([]string, len(tp.ps)+1)
	for i, p := range tp.ps {
		tp0[i] = p
		psMap[p] = true
	}

	tp0[len(tp.ps)] = newp
	if psMap[tp0[len(tp.ps)]] {
		n := uint64(1)
		for {
			n++
			tp0[len(tp.ps)] = newp + strconv.FormatUint(n, 10)
			if !psMap[tp0[len(tp.ps)]] {
				break
			}
		}
	}
	return typeParam{
		ps: tp0,
		// tp.isTypeParam,
		added: true,
	}
}
func (tp typeParam) Has() bool {
	return (!tp.added && len(tp.ps) != 0) || len(tp.ps) != 1
}
func (tp typeParam) Last() string {
	if len(tp.ps) == 0 {
		return ""
	}
	return tp.ps[len(tp.ps)-1]
}
func (tp typeParam) LSide() string {
	// if tp.isTypeParam {
	// 	return ""
	// }
	if len(tp.ps) == 0 {
		return ""
	}
	return "[" + strings.Join(slices.MapI(tp.ps, func(e string, i int) string {
		if i+1 <= len(tp.type_constraints) {
			return e + " " + tp.type_constraints[i]
		}
		return e + " any"
	}), ", ") + "]"
}
func (tp typeParam) RSide() string {
	if len(tp.ps) == 0 {
		return ""
	}
	return "[" + strings.Join(tp.ps, ",") + "]"
}
func (tp typeParam) TexprArgs() string {
	if len(tp.ps) == 0 {
		return ""
	}
	// if tp.stdlib {
	// 	return strings.Join(slices.Map(tp.ps, func(e string) string { return fmt.Sprintf("%s ATypeExpr[%s]", strings.ToLower(e), e) }), ", ")
	// }
	return strings.Join(slices.Map(tp.ps, func(e string) string { return fmt.Sprintf("%s adlast.ATypeExpr[%s]", strings.ToLower(e), e) }), ", ")
}
func (tp typeParam) TexprValues() string {
	if len(tp.ps) == 0 {
		return ""
	}
	return strings.Join(slices.Map(tp.ps, func(e string) string { return fmt.Sprintf("%s.Value", strings.ToLower(e)) }), ", ")
}

func (tp typeParam) TpArgs() string {
	if len(tp.ps) == 0 {
		return ""
	}
	return "[any" + strings.Repeat(",any", len(tp.ps)-1) + "]"
}
