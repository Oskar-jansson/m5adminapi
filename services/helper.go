package services

// coalesce returns the first non-nil pointer value, or nil if both are nil
func coalesce[T any](a, b *T) *T {
	if a != nil {
		return a
	}
	return b
}

func coalesceString(a, b *string) *string {
	if a != nil {
		return a
	}
	return b
}

func coalesceUint32(a, b *uint32) *uint32 {
	if a != nil {
		return a
	}
	return b
}

func coalesceBool(a, b *bool) *bool {
	if a != nil {
		return a
	}
	return b
}

func coalesceInt(a, b *int) *int {
	if a != nil {
		return a
	}
	return b
}

func coalesceFloat64(a, b *float64) *float64 {
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
