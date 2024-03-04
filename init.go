func migrateParamsStore(ctx sdk.Context, paramstore paramtypes.Subspace) {
	if paramstore.HasKeyTable() {
		paramstore.Set(ctx, types.ParamStoreKeyMaxTxsInBlock, types.DefaultMaxTxsInBlock)
	} else {
		paramstore.WithKeyTable(types.ParamKeyTable())
		paramstore.Set(ctx, types.ParamStoreKeyMaxTxsInBlock, types.DefaultMaxTxsInBlock)
	}
}

func MigrateStore(ctx sdk.Context, paramstore paramtypes.Subspace) error {
	migrateParamsStore(ctx, paramstore)
	return nil
}
