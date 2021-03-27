package main

type stream struct {
	list []interface{}
}

func Stream(arrs ...interface{}) *stream {

	st := new(stream)

	if len(arrs) > 0 {
		if x, ok := arrs[0].([]interface{}); ok {
			st.list = make([]interface{}, len(x))
			copy(st.list, x)
		} else {
			st.list = make([]interface{}, len(arrs))
			copy(st.list, arrs)
		}
	}

	return st
}

func (s *stream) Filter(fn func(each interface{}) bool) *stream {
	list := make([]interface{}, 0, len(s.list))
	for _, x := range s.list {
		if fn(x) {
			list = append(list, x)
		}
	}
	s.list = list
	return s
}

func (s *stream) ForEach(fn func(each interface{})) {
	list := s.list
	for _, x := range list {
		fn(x)
	}
}

func (s *stream) Reduce(initialValue interface{}, fn func(pre interface{}, cur interface{}) interface{}) interface{} {
	for i := 0; i < len(s.list); i++ {
		initialValue = fn(initialValue, s.list[i])
	}
	return initialValue
}
