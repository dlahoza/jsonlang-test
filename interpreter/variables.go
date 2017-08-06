package interpreter

import "sync"

// VarScope contains scope of execution variables
type VarScope struct {
	vars map[string]string
	sync.RWMutex
}

// NewVarScope create new clean scope of variables
func NewVarScope() *VarScope {
	return &VarScope{
		vars: make(map[string]string),
	}
}

// Create creates new variable, fails when it does already exist
func (s *VarScope) Create(name, val string) (err error) {
	s.Lock()
	if _, ok := s.vars[name]; !ok {
		s.vars[name] = val
	} else {
		err = ErrorVariableExists
	}
	s.Unlock()
	return
}

// Update updates existing variable, fails when it does not exist
func (s *VarScope) Update(name, val string) (err error) {
	s.Lock()
	if _, ok := s.vars[name]; ok {
		s.vars[name] = val
	} else {
		err = ErrorVariableDoesNotExist
	}
	s.Unlock()
	return
}

// Set creates of updates variable in scope
func (s *VarScope) Set(name, val string) {
	s.Lock()
	s.vars[name] = val
	s.Unlock()
	return
}

// Get returns variable from scope, returns 'undefined' when cannot find it
func (s *VarScope) Get(name string) (val string, err error) {
	s.RLock()
	if _, ok := s.vars[name]; ok {
		val = s.vars[name]
	} else {
		val = "undefined"
		err = ErrorVariableDoesNotExist
	}
	s.RUnlock()
	return
}

// Delete deletes variable from scope, fails when doesn't exist
func (s *VarScope) Delete(name string) (err error) {
	s.Lock()
	if _, ok := s.vars[name]; ok {
		delete(s.vars, name)
	} else {
		err = ErrorVariableDoesNotExist
	}
	s.Unlock()
	return
}
