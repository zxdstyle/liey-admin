package container

import (
	"reflect"
	"unsafe"
)

func injectFuncArgs(fn interface{}) ([]reflect.Value, error) {
	depFactoryType := reflect.TypeOf(fn)

	if depFactoryType.Kind() != reflect.Func {
		return nil, ErrorDepFactoryNotAFunc
	}

	// Retrieve dependencies for all factory arguments
	args := make([]reflect.Value, depFactoryType.NumIn())
	for i := 0; i < depFactoryType.NumIn(); i++ {
		inType := depFactoryType.In(i)
		resolvedDep, err := InjectT(inType)
		if err != nil {
			// If the field was not resolvable, attempt to fill it as a struct
			possibleStruct := reflect.New(inType)
			fillStruct := possibleStruct

			// If the field is a pointer, the fill struct must be initialized and de-refereced, else we have a double-pointer **Type
			if inType.Kind() == reflect.Pointer {
				possibleStruct.Elem().Set(reflect.New(inType.Elem()))
				fillStruct = possibleStruct.Elem()
			}

			// Attempt to fill in the struct with dependencies
			if err := injectStructFields(fillStruct.Interface()); err != nil {
				return nil, err
			}
			args[i] = reflect.ValueOf(possibleStruct.Elem().Interface())
		} else {
			args[i] = reflect.ValueOf(resolvedDep)
		}
	}
	return args, nil
}

func injectStructFields(strct interface{}) error {
	depFactoryType := reflect.TypeOf(strct)

	if depFactoryType.Kind() != reflect.Ptr {
		return ErrorDepNotAPointer
	}

	if depFactoryType.Elem().Kind() != reflect.Struct {
		return ErrorDepNotAStruct
	}

	// Set struct values to injected dependencies
	depFactoryValue := reflect.ValueOf(strct).Elem()
	for i := 0; i < depFactoryType.Elem().NumField(); i++ {
		field := depFactoryType.Elem().Field(i)
		fieldVal := depFactoryValue.Field(i)
		if lookupType, exists := field.Tag.Lookup("injector"); exists {
			name := DefaultForType
			switch lookupType {
			case "name":
				name = field.Name
			}
			resolvedDep, err := InjectT(field.Type, name)
			if err != nil {
				return err
			}
			ptr := reflect.NewAt(fieldVal.Type(), unsafe.Pointer(fieldVal.UnsafeAddr())).Elem()
			ptr.Set(reflect.ValueOf(resolvedDep))
		}
	}
	return nil
}
