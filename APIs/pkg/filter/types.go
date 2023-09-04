package filter

const (
	Equal        = "equal"
	NotEqual     = "not-equal"
	Greater      = "greater"
	Less         = "less"
	GreaterEqual = "greater-equal"
	LessEqual    = "less-equal"
	Like         = "like"
	NotLike      = "not-like"
	In           = "in"
	NotIn        = "not-in"
	Between      = "between"
)

type Operator struct {
	Type string
}

type ExpectedParam struct {
	Param          string
	ValidOperators []Operator
}

type QueryComposition struct {
	Key      string
	Value    string
	Operator Operator
}
