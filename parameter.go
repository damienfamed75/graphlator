package graphlator

type Parameter struct {
	Operation operation
	Predicate string
	Value     interface{}
}

func buildParamater(o operation, k string, v interface{}) Parameter {
	return Parameter{
		Operation: o,
		Predicate: k,
		Value:     v,
	}
}

func (p Parameter) AsFilter() *Filters {
	return &Filters{
		constraint: invalidConstraint,
		params:     []Parameter{p},
	}
}

func Equal(predicate string, value interface{}) Parameter {
	return buildParamater(operationEqual, predicate, value)
}

func GreaterThanOrEqual(predicate string, value interface{}) Parameter {
	return buildParamater(operationGreaterThanOrEqual, predicate, value)
}

func LessThanOrEqual(predicate string, value interface{}) Parameter {
	return buildParamater(operationLessThanOrEqual, predicate, value)
}

func GreaterThan(predicate string, value interface{}) Parameter {
	return buildParamater(operationGreaterThan, predicate, value)
}

func LessThan(predicate string, value interface{}) Parameter {
	return buildParamater(operationLessThan, predicate, value)
}

func AnyOfTerms(predicate string, value interface{}) Parameter {
	return buildParamater(operationAnyOfTerms, predicate, value)
}

func AllOfTerms(predicate string, value interface{}) Parameter {
	return buildParamater(operationAllOfTerms, predicate, value)
}

func Type(typeAsString string) Parameter {
	return buildParamater(operationType, typeAsString, nil)
}

func Has(predicate string) Parameter {
	return buildParamater(operationHas, predicate, nil)
}

func Regexp(predicate string, regex string) Parameter {
	return buildParamater(operationRegexp, predicate, regex)
}

func UIDIn(predicate string, uid interface{}) Parameter {
	return buildParamater(operationUIDIn, predicate, uid)
}

func UID(uid string) Parameter {
	return buildParamater(operationUID, uid, nil)
}
