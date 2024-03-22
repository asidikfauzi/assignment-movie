package validator

import (
	"assignment-movie/common/helper"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func ValidatorMessage(item interface{}) []string {
	var messages []string

	itemType := reflect.TypeOf(item)
	itemValue := reflect.ValueOf(item)

	for i := 0; i < itemType.NumField(); i++ {
		index := strconv.Itoa(i)

		fieldTag := itemType.Field(i).Tag.Get("json")
		fieldTagValidate := itemType.Field(i).Tag.Get("validate")
		fieldValue := itemValue.Field(i)

		if fieldValue.Kind() == reflect.Struct || fieldValue.Kind() == reflect.Slice {
			messages = append(messages, validateImplementation(index, fieldTag, fieldTagValidate, itemType, itemValue, fieldValue.Interface())...)
		} else {
			messages = append(messages, validateImplementation("-", fieldTag, fieldTagValidate, itemType, itemValue, fieldValue.Interface())...)
		}

	}
	return messages
}

func validateImplementation(idx string, tag, validate string, itemType reflect.Type, itemValue reflect.Value, data interface{}) []string {
	var messages []string

	sliceTypes := reflect.TypeOf(data)
	sliceValues := reflect.ValueOf(data)

	tagParts := strings.Split(tag, ".")
	firstTagPart := tagParts[0]

	if sliceValues.Kind() == reflect.Slice {
		if hasRequiredTag(validate) && sliceValues.Len() == 0 {
			messages = append(messages, validateImplementation(idx, tag, validate, sliceTypes, sliceValues, sliceValues)...)
			return messages
		}

		for i := 0; i < sliceValues.Len(); i++ {

			sliceValue := sliceValues.Index(i)

			if sliceValue.Kind() == reflect.String && sliceValue.Interface() == "" {
				if tagContains(tag, firstTagPart) {
					messages = append(messages, validateImplementation("-", tag, validate, sliceTypes, sliceValues, sliceValue.Interface())...)
				}
			}

			if sliceValue.Kind() == reflect.Int && sliceValue.Interface() == 0 {
				if tagContains(tag, firstTagPart) {
					messages = append(messages, validateImplementation("-", tag, validate, sliceTypes, sliceValues, sliceValue.Interface())...)
				}
			}

			if sliceValue.Kind() == reflect.Float64 && sliceValue.Interface().(float64) == 0 {
				if tagContains(tag, firstTagPart) {
					messages = append(messages, validateImplementation("-", tag, validate, sliceTypes, sliceValues, sliceValue.Interface())...)
				}
			}

			if sliceValue.Kind() == reflect.Struct {
				for j := 0; j < sliceValue.NumField(); j++ {
					index := strconv.Itoa(i)

					fieldTag := sliceValue.Type().Field(j).Tag.Get("json")
					fieldTagValidate := sliceValue.Type().Field(j).Tag.Get("validate")
					fieldValue := sliceValue.Field(j)

					fieldTag = fmt.Sprintf("%s.%d.%s", tag, i, fieldTag)

					if fieldValue.Kind() == reflect.Slice && hasRequiredTag(fieldTagValidate) {
						if fieldValue.Len() == 0 || (fieldValue.Len() == 1 && fieldValue.Index(0).Interface() == "") {
							message := fmt.Sprintf("Field '%s.%s' must be required", fieldTag, index)
							messages = append(messages, message)
						}
					}

					messages = append(messages, validateImplementation(index, fieldTag, fieldTagValidate, sliceTypes, sliceValues, fieldValue.Interface())...)
				}
			}
		}

	} else if sliceValues.Kind() == reflect.Struct {

		if hasRequiredTag(validate) && sliceValues.NumField() == 0 {
			messages = append(messages, validateImplementation(idx, tag, validate, sliceTypes, sliceValues, sliceValues.Interface())...)
		}

		for i := 0; i < sliceValues.NumField(); i++ {
			fieldTag := sliceValues.Type().Field(i).Tag.Get("json")
			fieldTagValidate := sliceValues.Type().Field(i).Tag.Get("validate")
			fieldValue := sliceValues.Field(i)

			messages = append(messages, validateImplementation("-", fieldTag, fieldTagValidate, sliceTypes, sliceValues, fieldValue.Interface())...)
		}

	} else {

		var message string
		if idx == "-" {
			message = fmt.Sprintf("Field '%s'", tag)
		} else {
			message = fmt.Sprintf("Field '%s.%s'", tag, idx)
		}

		if sliceTypes.Kind() == reflect.Float64 && hasRequiredTag(validate) && (data == nil || data.(float64) == 0) {
			messages = append(messages, message+helper.RequiredMessage)
		}

		if sliceTypes.Kind() == reflect.Int && hasRequiredTag(validate) && data == nil || data == 0 {
			messages = append(messages, message+helper.RequiredMessage)
		}

		if sliceTypes.Kind() == reflect.String {
			if hasRequiredTag(validate) && (data == nil || data.(string) == "") {
				messages = append(messages, message+helper.RequiredMessage)
			}

			if hasNumberTag(validate) && !isNumber(data) {
				messages = append(messages, message+helper.NumberMessage)
			}

			if hasDateTag(validate) && !isDateFormat(data.(string)) {
				messages = append(messages, message+helper.DateFormatMessage)
			}

			if hasMaxTag(validate) {
				messageMax := isMaxValue(data.(string), validate)
				if messageMax != "" {
					messages = append(messages, message+messageMax)
				}
			}

			if hasMinTag(validate) {
				messageMin := isMinValue(data.(string), validate)
				if messageMin != "" {
					messages = append(messages, message+messageMin)
				}
			}

			if hasEmailTag(validate) {
				email := isEmailValid(data.(string))
				if !email {
					messages = append(messages, message+helper.EmailMessage)
				}
			}

			if hasPasswordTag(validate) {
				password := isPassword(data.(string))
				if !password {
					messages = append(messages, message+helper.PasswordMessage)
				}
			}

			if hasRegexTag(validate) {
				regex := isRegexValid(data.(string), validate)
				if !regex {
					messages = append(messages, message+helper.RegexMessage)
				}
			}

			if hasRequiredIfTag(validate) {
				requiredIf := isRequiredIfValid(data.(string), validate, itemType, itemValue)
				if requiredIf != "" {
					messageReqIf := fmt.Sprintf(helper.RequiredIfMessage, requiredIf)
					messages = append(messages, message+messageReqIf)
				}
			}
		}

	}

	return messages
}
