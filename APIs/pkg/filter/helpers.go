package filter

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

// import "errors"

// func (o *Operators) GetOperator(key string) (string, error) {

// 	switch key {
// 	case o.Equal.Param:
// 		return o.Equal.Operator, nil
// 	case o.NotEqual.Param:
// 		return o.NotEqual.Operator, nil
// 	case o.GreaterThan.Param:
// 		return o.GreaterThan.Operator, nil
// 	case o.LessThan.Param:
// 		return o.LessThan.Operator, nil
// 	case o.GreaterThanOrEqual.Param:
// 		return o.GreaterThanOrEqual.Operator, nil
// 	case o.LessThanOrEqual.Param:
// 		return o.LessThanOrEqual.Operator, nil
// 	case o.Like.Param:
// 		return o.Like.Operator, nil
// 	case o.In.Param:
// 		return o.In.Operator, nil
// 	case o.NotIn.Param:
// 		return o.NotIn.Operator, nil
// 	case o.Between.Param:
// 		return o.Between.Operator, nil
// 	case o.NotBetween.Param:
// 		return o.NotBetween.Operator, nil
// 	case o.And.Param:
// 		return o.And.Operator, nil
// 	case o.Or.Param:
// 		return o.Or.Operator, nil
// 	case o.Not.Param:
// 		return o.Not.Operator, nil
// 	default:
// 		return "", errors.New("Operator not found")
// 	}

// }

func NewOperator() map[string]Operator {
	return map[string]Operator{
		Equal:        {Type: "="},
		NotEqual:     {Type: "!="},
		Greater:      {Type: ">"},
		Less:         {Type: "<"},
		GreaterEqual: {Type: ">="},
		LessEqual:    {Type: "<="},
		Like:         {Type: "like"},
		NotLike:      {Type: "not like"},
		In:           {Type: "in"},
		NotIn:        {Type: "not in"},
		Between:      {Type: "between"},
	}
}

func (o *Operator) ToSQL() (string, error) {
	switch o.Type {
	case Equal:
		return "=", nil
	case NotEqual:
		return "!=", nil
	case Greater:
		return ">", nil
	case Less:
		return "<", nil
	case GreaterEqual:
		return ">=", nil
	case LessEqual:
		return "<=", nil
	case Like:
		return "like", nil
	case NotLike:
		return "not like", nil
	case In:
		return "in", nil
	case NotIn:
		return "not in", nil
	case Between:
		return "between", nil
	default:
		return "", errors.New("Operator not found")
	}
}

func (e *ExpectedParam) ValidateExpectedParam(recived string) (sql string, err error) {
	for _, v := range e.ValidOperators {
		if v.Type == recived {
			sql, err := v.ToSQL()
			if err != nil {
				return "", err
			}

			return sql, nil
		}
	}
	return "", errors.New("Invalid Operator")
}

// the pather is key=operator.value

func ParseParam(param string) (operator string, value string, err error) {
	splited := strings.Split(param, ".")
	if len(splited) != 2 {
		return "", "", errors.New("Invalid Param Format")
	}

	operator = splited[0]
	value = splited[1]

	return operator, value, nil
}

func (e *ExpectedParam) GetFromRequest(request *http.Request) (query QueryComposition, err error) {
	queryValue := request.URL.Query().Get(e.Param)

	if queryValue == "" {
		return query, errors.New("Param not found")
	}

	operator, value, err := ParseParam(queryValue)
	if err != nil {
		return query, err
	}

	operator, err = e.ValidateExpectedParam(operator)

	query = QueryComposition{
		Key:      e.Param,
		Value:    value,
		Operator: Operator{Type: operator},
	}

	return query, nil

}

// create Gorm Query

func GenerateQueryGORM(entity interface{}, db *gorm.DB, comp QueryComposition) {

	val, _ := db.Model(entity).Where(comp.Key+" "+comp.Operator.Type+" ?", comp.Value).Rows()
	log.Println(&val)

}
