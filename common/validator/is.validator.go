package validator

import (
	"assignment-movie/common/helper"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func isNumber(s interface{}) bool {
	str := fmt.Sprintf("%s", s)
	_, err := strconv.Atoi(str)
	return err == nil
}

func isDateFormat(s string) bool {
	format := "2006-01-02"
	_, err := time.Parse(format, s)
	return err == nil
}

func isMaxValue(fieldValue, validateTag string) string {
	tags := strings.Split(validateTag, ",")

	var message string

	for _, tag := range tags {
		tagParts := strings.Split(tag, ":")

		if len(tagParts) == 2 {
			maxValue, err := strconv.Atoi(tagParts[1])
			if err != nil {
				return err.Error()
			}

			amountValue := len(fieldValue)
			if !validateMax(amountValue, maxValue) {
				message = fmt.Sprintf(helper.MaxMessage, tagParts[1])
				return message
			}
		}

	}

	return ""
}

func isMinValue(fieldValue, validateTag string) string {
	tags := strings.Split(validateTag, ",")

	var message string

	for _, tag := range tags {
		tagParts := strings.Split(tag, ":")

		if len(tagParts) == 2 {
			minValue, err := strconv.Atoi(tagParts[1])
			if err != nil {
				return err.Error()
			}

			amountValue := len(fieldValue)
			if !validateMin(amountValue, minValue) {
				message = fmt.Sprintf(helper.MinMessage, tagParts[1])
				return message
			}
		}

	}

	return ""
}

func isPassword(s string) bool {
	uppercaseRegex := regexp.MustCompile(`[A-Z]`)
	hasUppercase := uppercaseRegex.MatchString(s)

	lowercaseRegex := regexp.MustCompile(`[a-z]`)
	hasLowercase := lowercaseRegex.MatchString(s)

	digitRegex := regexp.MustCompile(`\d`)
	hasDigit := digitRegex.MatchString(s)

	return hasUppercase && hasLowercase && hasDigit
}

func isEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return emailRegex.MatchString(email)
}

func isRegexValid(fieldValue, validateTag string) bool {
	tags := strings.Split(validateTag, ",")

	for _, tag := range tags {
		tagParts := strings.Split(tag, ":")

		if len(tagParts) == 2 {
			regexValue := fmt.Sprintf(`%s`, tagParts[1])
			dataRegex := regexp.MustCompile(regexValue)

			return dataRegex.MatchString(fieldValue)
		}

	}

	return false
}

func isRequiredIfValid(dataValue, validateTag string, itemType reflect.Type, itemValue reflect.Value) string {
	tags := strings.Split(validateTag, ",")

	for _, tag := range tags {
		tagParts := strings.Split(tag, ":")

		if len(tagParts) == 2 {
			tagName := tagParts[1]

			if itemType.Kind() == reflect.Slice {
				if errTag := checkSliceElements(itemType, itemValue, tagName, dataValue); errTag != "" {
					return errTag
				}
			} else {
				if errTag := checkSingleElement(itemType, itemValue, tagName, dataValue); errTag != "" {
					return errTag
				}
			}
		}
	}

	return ""
}

// Helper
func tagContains(tag, firstTagPart string) bool {
	return tag == firstTagPart
}

func validateMax(value interface{}, max int) bool {
	switch val := value.(type) {
	case int:
		return val <= max
	default:
		return false
	}
}

func validateMin(value interface{}, min int) bool {
	switch val := value.(type) {
	case int:
		return val >= min
	default:
		return false
	}
}

func checkSliceElements(itemType reflect.Type, itemValue reflect.Value, tagName, dataValue string) string {
	for i := 0; i < itemValue.Len(); i++ {
		field, found := itemType.Elem().FieldByName(tagName)
		if !found {
			return tagName
		}

		jsonTag := field.Tag.Get("json")
		element := itemValue.Index(i)

		if element.Kind() == reflect.Struct {
			fieldValue := element.FieldByName(tagName)

			if fieldValue.String() != "" && dataValue == "" {
				return jsonTag
			}
		}
	}
	return ""
}

func checkSingleElement(itemType reflect.Type, itemValue reflect.Value, tagName, dataValue string) string {
	field, found := itemType.FieldByName(tagName)
	if !found {
		return tagName
	}

	jsonTag := field.Tag.Get("json")
	fieldValue := itemValue.FieldByName(tagName)

	if fieldValue.String() != "" && dataValue == "" {
		return jsonTag
	}

	return ""
}
