package interpreter

import "sync"

type VarScope struct {
	vars map[string]string
	sync.RWMutex
}

func NewVarScope() *VarScope {
	return &VarScope{
		vars: make(map[string]string),
	}
}

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

func (s *VarScope) Set(name, val string) {
	s.Lock()
	s.vars[name] = val
	s.Unlock()
	return
}

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
