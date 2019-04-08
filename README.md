# go-map-utils

Utility functions for golang maps.

**WARNING**: Heavy use of reflection, so none of these functions should be used critical code paths.

## Functions

All of the provided functions expect a map as first argument and a pointer to a slice as second
argument. The data type of the slice has to match the key/value type of the map.

`SortedKeys` can only sort basic golang data types like numbers or strings.

```go
Keys(interface{}, interface{})

SortedKeys(interface{}, interface{})

Values(interface{}, interface{})
```
