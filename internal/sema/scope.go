package sema

// binding is a declared name in some scope. For array declarations Type
// holds the ELEMENT type and IsArray marks the shape (spec section 5).
// IsConst marks a const object, which the C89 3.4 constant gate needs.
type binding struct {
	Type    Type
	IsArray bool
	IsConst bool
}

// scope is a stack of frames; lookup searches innermost-first.
type scope struct {
	frames []map[string]binding
}

func (s *scope) push() {
	s.frames = append(s.frames, map[string]binding{})
}

func (s *scope) pop() {
	if len(s.frames) > 0 {
		s.frames = s.frames[:len(s.frames)-1]
	}
}

func (s *scope) declare(name string, b binding) {
	if len(s.frames) == 0 {
		s.push()
	}
	s.frames[len(s.frames)-1][name] = b
}

func (s *scope) lookup(name string) (binding, bool) {
	for i := len(s.frames) - 1; i >= 0; i-- {
		if b, ok := s.frames[i][name]; ok {
			return b, true
		}
	}
	return binding{}, false
}
