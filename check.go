package checkParams

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type CheckFunc func(param interface{}) error

var (
	ErrNotAStructPtr = errors.New("Expected a pointer to a Struct")
	ErrNoCheckName   = errors.New("Register check name is empty.")

	// check_name to check funcations
	name2CheckFuncs = make(map[string][]CheckFunc)
)

func CheckParams(v interface{}) error {
	ptrRef := reflect.ValueOf(v)
	if ptrRef.Kind() != reflect.Ptr {
		return ErrNotAStructPtr
	}
	ref := ptrRef.Elem()
	if ref.Kind() != reflect.Struct {
		return ErrNotAStructPtr
	}

	return doCheck(ref)
}

func doCheck(ref reflect.Value) error {
	refType := ref.Type()
	for i := 0; i < refType.NumField(); i++ {
		key := refType.Field(i).Tag.Get("check")
		if key == "" {
			continue
		}
		value := ref.Field(i).Interface()

		checkers, ok := name2CheckFuncs[key]
		if ok {
			for i := 0; i < len(checkers); i++ {
				err := checkers[i](value)
				if err != nil {
					return fmt.Errorf("param is error. [name: %+v]", key)
				}
			}
		}

		// fmt.Printf("test: %+v, %+v \n", key, value)
	}

	return nil
}

func RegisterCheckFunc(check_name string, funcation CheckFunc) error {
	name := strings.TrimSpace(check_name)
	if check_name == "" {
		return ErrNoCheckName
	}

	if funcation == nil {
		return ErrNoCheckName
	}

	funcations, ok := name2CheckFuncs[name]
	if !ok {
		funcations = make([]CheckFunc, 0)
	}
	funcations = append(funcations, funcation)
	name2CheckFuncs[name] = funcations

	return nil
}
