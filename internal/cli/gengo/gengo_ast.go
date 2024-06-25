// Code generated by goadlc v3 - DO NOT EDIT.
package gengo

import (
	goadl "github.com/adl-lang/goadl_rt/v3"
	"github.com/adl-lang/goadl_rt/v3/customtypes"
	"github.com/adl-lang/goadl_rt/v3/sys/adlast"
	"github.com/adl-lang/goadl_rt/v3/sys/types"
)

func Texpr_GenGo() adlast.ATypeExpr[GenGo] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("cli.gengo", "GenGo"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[GenGo](te)
}

func AST_GenGo() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"GenGo",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"root",
						"-",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Nullable",
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"cli.root",
											"Root",
										),
									),
									[]adlast.TypeExpr{},
								),
							},
						),
						types.Make_Maybe_just[any](
							nil,
						),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"Loader",
						"Loader",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Nullable",
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"cli.loader",
											"Loader",
										),
									),
									[]adlast.TypeExpr{},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"Mod",
						"Mod",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Nullable",
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"cli.gomod",
											"GoModule",
										),
									),
									[]adlast.TypeExpr{},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("sys.annotations", "Doc"): "Used to find the go.mod file and calculate the codes module-path and sub-dir of generated code\n"},
					),
					adlast.MakeAll_Field(
						"GoTypes",
						"GoTypes",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Nullable",
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"cli.gotypes",
											"GoTypes",
										),
									),
									[]adlast.TypeExpr{},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"GoApis",
						"GoApis",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Nullable",
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"Vector",
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_reference(
												adlast.MakeAll_ScopedName(
													"cli.goapi",
													"GoApi",
												),
											),
											[]adlast.TypeExpr{},
										),
									},
								),
							},
						),
						types.Make_Maybe_just[any](
							nil,
						),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
				},
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{},
	)
	return adlast.Make_ScopedDecl("cli.gengo", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("cli.gengo", "GenGo"),
		AST_GenGo(),
	)
}
