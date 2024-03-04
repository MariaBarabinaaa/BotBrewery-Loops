var RandFuncDecls = decls.NewFunction("rand",
	decls.NewOverload("rand_int",
		[]*exprpb.Type{decls.Int},
		decls.Int),
	decls.NewOverload("rand",
		[]*exprpb.Type{},
		decls.Double),
)


var Log2FuncDecls = decls.NewFunction("log2",
	decls.NewOverload("log2_double",
		[]*exprpb.Type{decls.Double},
		decls.Double),
	decls.NewOverload("log2_int",
		[]*exprpb.Type{decls.Int},
		decls.Double),
)
