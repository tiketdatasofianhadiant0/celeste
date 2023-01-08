package celeste

const (
	ConjunctionAnd = "$and"
	ConjunctionOr  = "$or"

	OperatorEqual        = "$eq"
	OperatorNotEqual     = "$neq"
	OperatorGreater      = "$gt"
	OperatorGreaterEqual = "$gte"
	OperatorLess         = "$lt"
	OperatorLessEqual    = "$lte"
	OperatorIn           = "$in"
	OperatorNotIn        = "$not_in"
	OperatorLike         = "$like"
	OperatorNotLike      = "$not_like"

	FunctionLower = "$lower"
	FunctionUpper = "$upper"
	FunctionTrim  = "$trim"

	KeywordAnd = " AND "
	KeywordOr  = " OR "
)

var (
	operatorMap = map[string]string{
		OperatorEqual:        "=",
		OperatorNotEqual:     "<>",
		OperatorGreater:      ">",
		OperatorGreaterEqual: ">=",
		OperatorLess:         "<",
		OperatorLessEqual:    "<=",
		OperatorIn:           "IN",
		OperatorNotIn:        "NOT IN",
		OperatorLike:         "LIKE",
		OperatorNotLike:      "NOT LIKE",
	}

	functionMap = map[string]string{
		FunctionLower: "LOWER",
		FunctionUpper: "UPPER",
		FunctionTrim:  "TRIM",
	}
)
