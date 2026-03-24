package services

// coalesce returns the first non-nil pointer value, or nil if both are nil
func coalesce[T any](a, b *T) *T {
	if a != nil {
		return a
	}
	return b
}

// coalesceSlice returns the first non-empty slice, or the second if the first is empty
func coalesceSlice[T any](a, b []T) []T {
	if len(a) > 0 {
		return a
	}
	return b
}
