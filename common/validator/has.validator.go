package validator

import (
	"assignment-movie/common/helper"
	"strings"
)

func hasRequiredTag(validateTag string) bool {
	tags := strings.Split(validateTag, ",")

	for _, tag := range tags {
		if strings.TrimSpace(tag) == helper.Required {
			return true
		}
	}

	return false
}

func hasNumberTag(validateTag string) bool {
	tags := strings.Split(validateTag, ",")

	for _, tag := range tags {
		if strings.TrimSpace(tag) == helper.Number {
			return true
		}
	}

	return false
}

func hasDateTag(validateTag string) bool {
	tags := strings.Split(validateTag, ",")

	for _, tag := range tags {
		if strings.TrimSpace(tag) == helper.Date {
			return true
		}
	}

	return false
}

func hasMaxTag(validateTag string) bool {
	tags := strings.Split(validateTag, ",")

	for _, tag := range tags {
		tagParts := strings.Split(tag, ":")
		tagName := strings.TrimSpace(tagParts[0])

		if tagName == helper.Max {
			return true
		}

	}

	return false
}

func hasMinTag(validateTag string) bool {
	tags := strings.Split(validateTag, ",")

	for _, tag := range tags {
		tagParts := strings.Split(tag, ":")
		tagName := strings.TrimSpace(tagParts[0])

		if tagName == helper.Min {
			return true
		}

	}

	return false
}

func hasEmailTag(validateTag string) bool {
	tags := strings.Split(validateTag, ",")

	for _, tag := range tags {
		if strings.TrimSpace(tag) == helper.Email {
			return true
		}
	}

	return false
}

func hasPasswordTag(validateTag string) bool {
	tags := strings.Split(validateTag, ",")

	for _, tag := range tags {
		if strings.TrimSpace(tag) == helper.Password {
			return true
		}
	}

	return false
}

func hasRequiredIfTag(validateTag string) bool {
	tags := strings.Split(validateTag, ",")

	for _, tag := range tags {
		tagParts := strings.Split(tag, ":")
		tagName := strings.TrimSpace(tagParts[0])

		if tagName == helper.RequiredIf {
			return true
		}

	}

	return false
}

func hasRegexTag(validateTag string) bool {
	tags := strings.Split(validateTag, ",")

	for _, tag := range tags {
		tagParts := strings.Split(tag, ":")
		tagName := strings.TrimSpace(tagParts[0])

		if tagName == helper.Regex {
			return true
		}

	}

	return false
}
