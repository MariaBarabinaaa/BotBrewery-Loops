var MinDoubleDoubleFunc = &functions.Overload{
	// operator for 2 param
	Operator: "min_double_double",
	Binary: func(lhs, rhs ref.Val) ref.Val {
		lftInt64 := lhs.Value().(float64)
		rgtInt64 := rhs.Value().(float64)
		if lftInt64 > rgtInt64 {
			return types.Double(rgtInt64)
		}
		return types.Double(lftInt64)
	},
}

var MaxIntDoubleFunc = &functions.Overload{
	// operator for 2 param
	Operator: "max_int_double",
	Binary: func(lhs, rhs ref.Val) ref.Val {
		lftInt64 := float64(lhs.Value().(int64))
		rgtInt64 := rhs.Value().(float64)
		if lftInt64 < rgtInt64 {
			return types.Double(rgtInt64)
		}
		return types.Double(lftInt64)
	},
}
