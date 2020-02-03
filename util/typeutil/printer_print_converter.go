package typeutil

import (
	"bytes"
	"fmt"
	"go/types"
	"strings"

	"github.com/hori-ryota/go-strcase"
)

func (p Printer) PrintConverterWitoutErrorCheck(
	sourceName string,
	sourceType types.Type,
	targetType types.Type,
) (string, error) {
	return p.printConverterWitoutErrorCheck(
		sourceName,
		sourceType,
		targetType,
		0,
	)
}

func (p Printer) printConverterWitoutErrorCheck(
	sourceName string,
	sourceType types.Type,
	targetType types.Type,
	nestNum int,
) (string, error) {
	targetTypeStr := p.PrintRelativeType(targetType)

	if namedTarget, ok := targetType.(*types.Named); ok {
		constructor := SearchConstructor(namedTarget)
		if constructor != nil {
			targetValueName := "m"
			constructorStr, err := p.FillConstructor(constructor, sourceName, sourceType, nestNum)
			if err != nil {
				return "", err
			}
			settersStr, err := p.PrintAssignToSetters(sourceName, sourceType, targetValueName, namedTarget, nestNum)
			if err != nil {
				return "", err
			}
			if settersStr != "" {
				return constructorStr, nil
			}
			return fmt.Sprintf(`func() %s {
					%s := %s
					%s
					return %s
				}()`,
				targetTypeStr,
				targetValueName, constructorStr,
				settersStr,
				targetValueName,
			), nil
		}

		targetEnumValues := TypeToEnumValues(targetType)
		sourceEnumValues := TypeToEnumValues(sourceType)
		if len(targetEnumValues) > 0 && len(sourceEnumValues) > 0 {
			b := new(bytes.Buffer)
			for _, sourceEnumValue := range sourceEnumValues {
				var targetEnumValue *types.Const
				for _, v := range targetEnumValues {
					if strings.Contains(
						strcase.ToLowerSnake(sourceEnumValue.Name()),
						strcase.ToLowerSnake(v.Name()),
					) ||
						strings.Contains(
							strcase.ToLowerSnake(v.Name()),
							strcase.ToLowerSnake(sourceEnumValue.Name()),
						) {
						targetEnumValue = v
						break
					}
				}
				if targetEnumValue == nil {
					continue
				}
				fmt.Fprintf(b, "case %s:\n", p.PrintRelativeConst(sourceEnumValue))
				fmt.Fprintf(b, "return %s\n", p.PrintRelativeConst(targetEnumValue))
			}
			return fmt.Sprintf(`func(s %s) %s {
			switch s {
			%s
			default:
				var t%d %s
				return t%d
			}
			}(%s)`,
				p.PrintRelativeType(sourceType), targetTypeStr,
				b.String(),
				nestNum, targetTypeStr,
				nestNum,
				sourceName,
			), nil
		}
	}

	if types.AssignableTo(sourceType, targetType) {
		return sourceName, nil
	}
	if types.ConvertibleTo(sourceType, targetType) {
		return fmt.Sprintf("%s(%s)", targetTypeStr, sourceName), nil
	}

	if sourceSlice, ok := sourceType.Underlying().(*types.Slice); ok {
		targetSlice, ok := targetType.Underlying().(*types.Slice)
		if !ok {
			return "", fmt.Errorf("unknown type pair '%s' '%s'", sourceType, targetType)
		}
		elemConverter, err := p.printConverterWitoutErrorCheck(
			fmt.Sprintf("%s[i%d]", sourceName, nestNum),
			sourceSlice.Elem(),
			targetSlice.Elem(),
			nestNum+1,
		)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(`func() %s {
				t%d := make(%s, len(%s))
				for i%d := range t%d {
					t%d[i%d] = %s
				}
				return t%d
			}()`,
			targetTypeStr,
			nestNum, targetTypeStr, sourceName,
			nestNum, nestNum,
			nestNum, nestNum, elemConverter,
			nestNum,
		), nil
	}

	if structTarget, ok := targetType.Underlying().(*types.Struct); ok {
		b := new(bytes.Buffer)
		fmt.Fprintf(b, "%s{\n", p.PrintRelativeType(targetType))
		for _, field := range StructToFields(structTarget) {
			if !IsContainedInPackage(p.dstPackage, targetType) && !field.Exported() {
				continue
			}
			v, t := SearchGetValueStrAndResultType(sourceType, field.Name(), !IsContainedInPackage(p.dstPackage, sourceType))
			if v == "" || t == nil {
				continue
			}
			getvalueStr, err := p.printConverterWitoutErrorCheck(
				fmt.Sprintf("%s.%s", sourceName, v), t, field.Type(), nestNum,
			)
			if err != nil {
				return "", err
			}
			fmt.Fprintf(b, "%s: %s,\n", field.Name(), getvalueStr)
		}
		fmt.Fprint(b, "}")
		return b.String(), nil
	}

	if targetPointer, ok := targetType.Underlying().(*types.Pointer); ok {
		elemConverter, err := p.printConverterWitoutErrorCheck(
			sourceName,
			sourceType,
			targetPointer.Elem(),
			nestNum,
		)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(`func() %s {
				t%d := %s
				return &t%d
			}()`,
			targetTypeStr,
			nestNum, elemConverter,
			nestNum,
		), nil
	}

	if sourcePointer, ok := sourceType.Underlying().(*types.Pointer); ok {
		sourceValueName := "s"
		elemConverter, err := p.printConverterWitoutErrorCheck(
			"*"+sourceValueName,
			sourcePointer.Elem(),
			targetType,
			nestNum,
		)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(`func(%s %s) %s {
				if %s == nil {
					var t%d %s
					return t%d
				}
				return %s
			}(%s)`,
			sourceValueName, p.PrintRelativeType(sourceType), targetTypeStr,
			sourceValueName,
			nestNum, targetTypeStr,
			nestNum,
			elemConverter,
			sourceName,
		), nil
	}

	if targetType == types.Universe.Lookup("string").Type() {
		return p.ToStringConverter(sourceName, sourceType)
	}

	return "", fmt.Errorf("unknown type pair '%s' '%s'", sourceType, targetType)
}

func (p Printer) FillConstructor(
	constructor *types.Func,
	sourceName string,
	sourceType types.Type,
	nestNum int,
) (string, error) {
	b := new(bytes.Buffer)
	fmt.Fprintf(b, "%s(\n", p.PrintRelativeFuncName(constructor))
	for _, param := range FuncToParams(constructor) {
		v, t := SearchGetValueStrAndResultType(sourceType, param.Name(), !IsContainedInPackage(p.dstPackage, sourceType))
		if v == "" || t == nil {
			return "", fmt.Errorf("mismatched constructor type for '%s' of '%s' from '%s'", param, constructor, sourceType)
		}
		getvalueStr, err := p.printConverterWitoutErrorCheck(
			fmt.Sprintf("%s.%s", sourceName, v), t, param.Type(), nestNum,
		)
		if err != nil {
			return "", err
		}
		fmt.Fprintf(b, "%s,\n", getvalueStr)
	}
	fmt.Fprintln(b, ")")
	return b.String(), nil
}

func (p Printer) PrintAssignToSetters(
	sourceName string,
	sourceType types.Type,
	targetName string,
	targetType *types.Named,
	nestNum int,
) (string, error) {
	b := new(bytes.Buffer)
	for _, setter := range ListSetters(targetType, !IsContainedInPackage(p.dstPackage, targetType)) {
		setterParams := FuncToParams(setter)
		if len(setterParams) != 1 {
			// support only one parameter setter
			continue
		}
		setterParam := setterParams[0]

		paramName := strings.TrimPrefix(strings.TrimPrefix(setter.Name(), "Set"), "set")
		v, t := SearchGetValueStrAndResultType(sourceType, paramName, !IsContainedInPackage(p.dstPackage, sourceType))
		if v == "" || t == nil {
			continue
		}
		getvalueStr, err := p.printConverterWitoutErrorCheck(
			fmt.Sprintf("%s.%s", sourceName, v), t, setterParam.Type(), nestNum,
		)
		if err != nil {
			return "", err
		}
		fmt.Fprintf(b, "%s.%s(%s)\n", targetName, setter.Name(), getvalueStr)
	}
	return b.String(), nil
}

func (p Printer) ToStringConverter(
	sourceName string,
	sourceType types.Type,
) (string, error) {
	printConverter := func(intermediateType string, format string) (string, error) {
		s, err := p.PrintConverterWitoutErrorCheck(sourceName, sourceType, types.Universe.Lookup(intermediateType).Type())
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(format, s), nil
	}

	switch sourceType.Underlying().String() {
	case "string", "[]byte":
		return printConverter("string", "%s")
	case "bool":
		return printConverter("bool", "strconv.FormatBool(%s)")
	case "int", "int8", "int16", "int32", "int64":
		return printConverter("int64", "strconv.FormatInt(%s, 10)")
	case "uint", "uint8", "uint16", "uint32", "uint64":
		return printConverter("uint64", "strconv.FormatUint(%s, 10)")
	case "float32":
		return printConverter("float32", "strconv.FormatFloat(%s, 'g', -1, 32)")
	case "float64":
		return printConverter("float64", "strconv.FormatFloat(%s, 'g', -1, 64)")
	}
	return "", fmt.Errorf("unknown type for stringConverter '%s'", sourceType)
}
