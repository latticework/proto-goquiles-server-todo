package jalicore

func IsNilOrEmpty(s *string) {
	return s == nil || len(s) == 0
}
