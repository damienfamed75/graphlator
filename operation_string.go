// Code generated by "stringer -type=operation -linecomment"; DO NOT EDIT.

package graphlator

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[operationEqual-0]
	_ = x[operationGreaterThanOrEqual-1]
	_ = x[operationLessThanOrEqual-2]
	_ = x[operationGreaterThan-3]
	_ = x[operationLessThan-4]
	_ = x[operationAnyOfTerms-5]
	_ = x[operationAllOfTerms-6]
	_ = x[operationType-7]
	_ = x[operationHas-8]
	_ = x[operationRegexp-9]
	_ = x[operationUIDIn-10]
	_ = x[operationUID-11]
}

const _operation_name = "eqgelegtltanyOfTermsallOfTermstypehasregexpuid_inuid"

var _operation_index = [...]uint8{0, 2, 4, 6, 8, 10, 20, 30, 34, 37, 43, 49, 52}

func (i operation) String() string {
	if i < 0 || i >= operation(len(_operation_index)-1) {
		return "operation(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _operation_name[_operation_index[i]:_operation_index[i+1]]
}
