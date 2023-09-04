package filter

// var operators = NewOperator()

// func ValidateOperator(operator string) (string, error) {
// 	finalOperator, err := operators.GetOperator(operator)
// 	return finalOperator, err
// }

// // the param comes from the url like this: /users/?id=equal.1
// // the patern is: key=operator.value
// // the operator `In` and `NotIn` are special cases, because the value is a list of values
// // the pattern is: key=operator.value1,value2,value3

// func ParseParams(r *http.Request, validator Validator) ([]string, error) {
// 	params := r.URL.Query().Get("filter")
// 	if params == "" {
// 		return nil, nil
// 	}

// 	splitParams := strings.Split(params, ",")

// 	return splitParams, nil
// }

// func ParseParam(param string) (key string, operator string, value string, err error) {
// 	splitKey := strings.Split(param, "=")
// 	splitValue := strings.Split(splitKey[1], ".")

// 	key = splitKey[0]
// 	value = splitValue[1]
// 	operator = splitValue[0]

// 	finalOperator, err := ValidateOperator(operator)
// 	if err != nil {
// 		return "", "", "", err
// 	}

// 	return key, finalOperator, value, nil
// }

// func ApplyFiters(db *gorm.DB, params []string) *gorm.DB {
// 	for _, param := range params {
// 		key, operator, value, err := ParseParam(param)
// 		if err != nil {
// 			continue
// 		}

// 		db = db.Where(key+" "+operator+" ?", value)
// 	}

// 	return db
// }
