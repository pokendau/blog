package api

import "fmt"

func WithComData(pairs ...any) (map[any]any, error) {
	if len(pairs)%2 == 1 {
		return nil, fmt.Errorf("WithComData needs a key and value pair")
	}

	m := make(map[any]any, len(pairs)/2)
	for i := 0; i < len(pairs)/2; i += 2 {
		m[pairs[i]] = pairs[i+1]
	}

	return m, nil
}
