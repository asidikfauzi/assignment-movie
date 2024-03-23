package helper

// Tag Validator
const (
	Required   = "required"
	Number     = "number"
	Date       = "date"
	Max        = "max"
	Min        = "min"
	Email      = "email"
	Password   = "password"
	RequiredIf = "requiredif"
	Regex      = "regex"
)

// Message Validator
const (
	RequiredMessage    = " must be required"
	NumberMessage      = " must be a number"
	MaxMessage         = " must be less than or equal to %s"
	MinMessage         = " must be greater than or equal to %s"
	DateFormatMessage  = " must be in the date format (ex: 2006-01-02)"
	EmailMessage       = " must be a valid email address"
	PasswordMessage    = " must have at least one uppercase letter, one lowercase letter, and one digit"
	RequiredIfMessage  = " is required when '%s' is present"
	RegexMessage       = " not valid"
	MaxFileSizeMessage = "File size must be less than 2MB"
)

// Message Response
const (
	SuccessGetData     = "Successfully Get Data!"
	SuccessCreatedData = "Successfully Created Data!"
	SuccessUpdatedData = "Successfully Updated Data!"
	SuccessDeletedData = "Successfully Deleted Data!"
)

// Storage
const (
	MaxFileSize  = 2 << 20 // 2 MB in bytes
	StorageImage = "common/public/assets/img/"
)
