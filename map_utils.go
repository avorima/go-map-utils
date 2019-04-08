package map_utils

import (
	"fmt"
	"reflect"
	"sort"
)

// Keys copies the keys from a map into a slice pointer
func Keys(dict interface{}, keys interface{}) error {
	dictVal, empty, err := dictValue(dict)
	if err != nil {
		return err
	} else if empty {
		return nil
	}
	keysVal, err := slicePtrValue(keys)
	if err != nil {
		return err
	}

	dictType := dictVal.Type()
	keysType := keysVal.Type()
	if dictType.Key() != keysType.Elem() {
		return fmt.Errorf("mismatching types, map: %v, slice: %v", dictType, keysType)
	}

	keysVal.Set(reflect.Append(keysVal, dictVal.MapKeys()...))
	return nil
}

// Values copies the values from a map into a slice pointer
func Values(dict interface{}, values interface{}) error {
	dictVal, empty, err := dictValue(dict)
	if err != nil {
		return err
	} else if empty {
		return nil
	}
	valuesVal, err := slicePtrValue(values)
	if err != nil {
		return err
	}

	dictType := dictVal.Type()
	valuesType := valuesVal.Type()
	if dictType.Elem() != valuesType.Elem() {
		return fmt.Errorf("mismatching types, map: %v, slice: %v", dictType, valuesType)
	}

	for _, key := range dictVal.MapKeys() {
		valuesVal.Set(reflect.Append(valuesVal, dictVal.MapIndex(key)))
	}
	return nil
}

func dictValue(dict interface{}) (dictVal reflect.Value, empty bool, err error) {
	dictVal = reflect.ValueOf(dict)
	if dictVal.Kind() == reflect.Ptr {
		dictVal = dictVal.Elem()
	}
	if dictVal.Kind() != reflect.Map {
		return dictVal, false, fmt.Errorf("can't get keys of type %v", dictVal)
	}
	empty = len(dictVal.MapKeys()) == 0
	return
}

func slicePtrValue(slice interface{}) (reflect.Value, error) {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() != reflect.Ptr {
		return sliceVal, fmt.Errorf("can't append slice to a slice value")
	}
	sliceVal = sliceVal.Elem()
	if sliceVal.Kind() != reflect.Slice {
		return sliceVal, fmt.Errorf("can't append slice to type %v", sliceVal)
	}
	return sliceVal, nil
}

type sortableSlice []reflect.Value

func (s sortableSlice) Len() int      { return len(s) }
func (s sortableSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s sortableSlice) Less(i, j int) (less bool) {
	k := s[i].Kind()
	if k != s[j].Kind() {
		panic("incompatible types for comparison")
	}
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		less = s[i].Int() < s[j].Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		less = s[i].Uint() < s[j].Uint()
	case reflect.Float32, reflect.Float64:
		less = s[i].Float() < s[j].Float()
	case reflect.String:
		less = s[i].String() < s[j].String()
	default:
		panic("invalid type for comparison")
	}
	return
}

// SortedKeys copies the keys from a map into a slice pointer and sorts it.
func SortedKeys(dict interface{}, keys interface{}) error {
	dictVal, empty, err := dictValue(dict)
	if err != nil {
		return err
	} else if empty {
		return nil
	}
	keysVal, err := slicePtrValue(keys)
	if err != nil {
		return err
	}

	dictType := dictVal.Type()
	keysType := keysVal.Type()
	if dictType.Key() != keysType.Elem() {
		return fmt.Errorf("mismatching types, map: %v, slice: %v", dictType, keysType)
	}

	var sortable sortableSlice
	for _, key := range dictVal.MapKeys() {
		sortable = append(sortable, key)
	}
	sort.Sort(sortable)
	keysVal.Set(reflect.Append(keysVal, []reflect.Value(sortable)...))

	return nil
}
