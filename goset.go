package goset

type set struct {
	data map[any]bool
}

type Set interface {
	Add(item any)
	Remove(item any)
	Get() []any
}

func (s *set) Add(item any) {
	if _, ok := s.data[item]; !ok {
		s.data[item] = true
	}
}

func (s *set) Remove(item any) {
	delete(s.data, item)
}

func (s *set) Get() []any {
	items := make([]any, 0, len(s.data))
	for item := range s.data {
		items = append(items, item)
	}
	return items
}

func NewSet() Set {
	return &set{
		data: make(map[any]bool),
	}
}
