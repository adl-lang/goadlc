// Code generated by goadlc v3 - DO NOT EDIT.
package root

import (
	goadl "github.com/adl-lang/goadl_rt/v3"
	"github.com/adl-lang/goadl_rt/v3/customtypes"
	"github.com/adl-lang/goadl_rt/v3/sys/adlast"
	"github.com/adl-lang/goadl_rt/v3/sys/types"
)

func Texpr_Root() adlast.ATypeExpr[Root] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("cli.root", "Root"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[Root](te)
}

func AST_Root() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"Root",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"Debug",
						"Debug",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Bool",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_just[any](
							false,
						),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("sys.annotations", "Doc"): "Print extra diagnostic information, especially about files being read/written\n"},
					),
					adlast.MakeAll_Field(
						"Cfg",
						"Cfg",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"String",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_just[any](
							"",
						),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("sys.annotations", "Doc"): "Config file in json format (NOTE file entries take precedence over command-line flags & env)\n"},
					),
					adlast.MakeAll_Field(
						"DumpConfig",
						"DumpConfig",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Bool",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_just[any](
							false,
						),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("sys.annotations", "Doc"): "Dump the config to stdout and exits\n"},
					),
				},
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{},
	)
	return adlast.Make_ScopedDecl("cli.root", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("cli.root", "Root"),
		AST_Root(),
	)
}
