package collection

// KeyValue is a tuple of a Key and Value pair.
type KeyValue[K comparable, V any] struct {
	Key   K
	Value V
}

// JustValues return a slice of just the values from the given slice of KeyValue pairs.
func JustValues[K comparable, V any](pairs []KeyValue[K, V]) []V {
	result := make([]V, 0, len(pairs))
	for _, kv := range pairs {
		result = append(result, kv.Value)
	}
	return result
}
