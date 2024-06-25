// Code generated by goadlc v3 - DO NOT EDIT.
package goapi

import (
	goadl "github.com/adl-lang/goadl_rt/v3"
	"github.com/adl-lang/goadl_rt/v3/customtypes"
	"github.com/adl-lang/goadl_rt/v3/sys/adlast"
	"github.com/adl-lang/goadl_rt/v3/sys/types"
)

func Texpr_GoApi() adlast.ATypeExpr[GoApi] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("cli.goapi", "GoApi"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[GoApi](te)
}

func AST_GoApi() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"GoApi",
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
						"loader",
						"-",
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
						types.Make_Maybe_just[any](
							nil,
						),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"goMod",
						"-",
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
						types.Make_Maybe_just[any](
							nil,
						),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"ApiStruct",
						"ApiStruct",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"String",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
				},
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{},
	)
	return adlast.Make_ScopedDecl("cli.goapi", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("cli.goapi", "GoApi"),
		AST_GoApi(),
	)
}
