package parser

// IsPureSpecifierAllowed implements the grammar's semantic predicate
// { IsPureSpecifierAllowed() }?, used to disambiguate a pure specifier
// ("= 0") in a member declarator from a default member initializer.
//
// A pure specifier is only valid inside a class member declaration. For
// the current syntax-only phase the predicate is always true.
func IsPureSpecifierAllowed() bool {
	return true
}
