// Code generated by goadlc v3 - DO NOT EDIT.
package gomod

import (
	goadl "github.com/adl-lang/goadl_rt/v3"
	"github.com/adl-lang/goadl_rt/v3/customtypes"
	"github.com/adl-lang/goadl_rt/v3/sys/adlast"
	"github.com/adl-lang/goadl_rt/v3/sys/types"
)

func Texpr_GoModResult() adlast.ATypeExpr[GoModResult] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("cli.gomod", "GoModResult"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[GoModResult](te)
}

func AST_GoModResult() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"GoModResult",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"ModulePath",
						"ModulePath",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"String",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"MidPath",
						"MidPath",
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
	return adlast.Make_ScopedDecl("cli.gomod", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("cli.gomod", "GoModResult"),
		AST_GoModResult(),
	)
}

func Texpr_GoModule() adlast.ATypeExpr[GoModule] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("cli.gomod", "GoModule"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[GoModule](te)
}

func AST_GoModule() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"GoModule",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_union_(
			adlast.MakeAll_Union(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"ModulePath",
						"ModulePath",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"String",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("sys.annotations", "Doc"): "The path of the Go module for the generated code.\n"},
					),
					adlast.MakeAll_Field(
						"GoModFile",
						"GoModFile",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"String",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("sys.annotations", "Doc"): "Path of a go.mod file\n"},
					),
					adlast.MakeAll_Field(
						"Outputdir",
						"Outputdir",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"String",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("sys.annotations", "Doc"): "Used to find the go.mod file and calculate the codes module-path and sub-dir of generated code\n"},
					),
				},
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{},
	)
	return adlast.Make_ScopedDecl("cli.gomod", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("cli.gomod", "GoModule"),
		AST_GoModule(),
	)
}
