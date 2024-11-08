// Copyright 2024 Aquia, Inc
// https://www.aquia.us

package aplusacl

// Acl provides an Access Control List to provide an authorization mechanism for applications. It is loosely based on AWS IAM policies but simplified for brevity and ease of use.
type ACL struct {
	statements []*Statement
}

// NewAcl returns an Acl
func NewAcl() *ACL {
	return &ACL{}
}

// AddStatement adds a Statement to the Acl
func (a *ACL) AddStatement(s *Statement) *ACL {
	a.statements = append(a.statements, s)
	return a
}

// IsAuthorized determines whether the provided principal is allowed to perform the specified action on the specified resource
func (a *ACL) IsAuthorized(principal, action, resource string) (authorized bool) {
	for _, s := range a.statements {
		_, pOk := s.principals[principal]
		_, aOk := s.actions[action]
		_, rOk := s.resources[resource]

		if pOk && aOk && rOk {
			return true
		}
	}

	return
}
