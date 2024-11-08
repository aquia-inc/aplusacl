package aplusacl

// Statement specifies what principals are allowed or denied to perform the specified actions on the specified resources
type Statement struct {
	principals map[string]interface{}
	resources  map[string]interface{}
	actions    map[string]interface{}
}

// NewStatement returns Statement
func NewStatement() *Statement {
	return &Statement{
		principals: map[string]interface{}{},
		resources:  map[string]interface{}{},
		actions:    map[string]interface{}{},
	}
}

// AddPrincipal adds the provided principal to the statement
func (s *Statement) AddPrincipal(principal string) *Statement {
	s.principals[principal] = nil
	return s
}

// AddPrincipals adds the provided list of principals to the statement
func (s *Statement) AddPrincipals(principals []string) *Statement {
	for _, p := range principals {
		s.principals[p] = nil
	}
	return s
}

// AddResource adds the provided resource to the statement
func (s *Statement) AddResource(resource string) *Statement {
	s.resources[resource] = nil
	return s
}

// AddResources adds the provided list of resources to the statement
func (s *Statement) AddResources(resources []string) *Statement {
	for _, r := range resources {
		s.resources[r] = nil
	}
	return s
}

// AddAction adds the provided action to the statement
func (s *Statement) AddAction(action string) *Statement {
	s.actions[action] = nil
	return s
}

// AddActions adds the provided list of actions to the statement
func (s *Statement) AddActions(actions []string) *Statement {
	for _, a := range actions {
		s.actions[a] = nil
	}
	return s
}
