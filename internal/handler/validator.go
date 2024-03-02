package handler

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"time"

	"github.com/go-playground/locales/ja_JP"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ja_translations "github.com/go-playground/validator/v10/translations/ja"
)

var validate *validator.Validate

// Bool is a custom type for JSON unmarshalling
type Bool bool

// UnmarshalJSON is a custom unmarshaller for Bool
func (b *Bool) UnmarshalJSON(data []byte) error {
	var s interface{}
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch v := s.(type) {
	case string:
		switch v {
		case "true", "1":
			*b = true
		case "false", "0":
			*b = false
		default:
			return errors.New("invalid value for boolean")
		}
	case float64: // UnmarshalJSONは数値をデフォルトでfloat64に変換するため、キャストする
		switch v {
		case 1:
			*b = true
		case 0:
			*b = false
		default:
			return errors.New("invalid value for boolean")
		}
	case bool:
		*b = Bool(v)
	default:
		return errors.New("invalid value for boolean")
	}
	return nil
}

type fieldError struct {
	Field string `json:"field" example:"id"`
	Error string `json:"error" example:"idは必須です"`
}

func fieldErrors(err error) []fieldError {
	var fls []fieldError

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return fls
	}

	for _, err := range validationErrors {
		fe := fieldError{
			Field: err.Field(),
			Error: err.Tag(),
		}
		fls = append(fls, fe)
	}

	return fls
}

type CustomValidator struct {
	Validator  *validator.Validate
	Translator ut.Translator
}

func init() {
	v := validator.New()

	ja := ja_JP.New()
	uni := ut.New(ja, ja)
	trans, _ := uni.GetTranslator("ja")

	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		fieldName := field.Tag.Get("ja")
		if fieldName == "-" {
			return ""
		}
		return fieldName
	})

	_ = ja_translations.RegisterDefaultTranslations(v, trans)
	_ = v.RegisterValidation("isAfterNow", isAfterNow)
	_ = v.RegisterValidation("isAfterField", isAfterOrEqualField)
	_ = v.RegisterValidation("boolean", isBoolean)
	_ = v.RegisterValidation("isTime", isTimeFormat)
	_ = v.RegisterValidation("isDate", isDate)
	_ = v.RegisterValidation("isInt", isInt)

	validate = v
}

func validateBranchID(branchID string) error {
	type path struct {
		BranchID string `validate:"required,len=3"`
	}
	return validate.Struct(path{BranchID: branchID})
}

func validateAppointmentID(appointmentID string) error {
	type path struct {
		AppointmentID string `validate:"required,len=36"`
	}
	return validate.Struct(path{AppointmentID: appointmentID})
}

func validateBannerImageID(bannerImageID string) error {
	type path struct {
		BannerImageID string `validate:"required,uuid4"`
	}
	return validate.Struct(path{BannerImageID: bannerImageID})
}

func validateID(id string) error {
	type path struct {
		ID string `validate:"required"`
	}
	return validate.Struct(path{ID: id})
}

// Custom validation function to check if the date is greater than or equal to the current date.
func isAfterNow(fl validator.FieldLevel) bool {
	if dateField, ok := fl.Field().Interface().(time.Time); ok {
		now := time.Now()
		return now.Before(dateField) || now.Equal(dateField)
	}

	return false
}

// Custom validation function to check if one time field is after or equal to another.
func isAfterOrEqualField(fl validator.FieldLevel) bool {
	// Get the field and parameter names from the validator
	paramName := fl.Param()

	// Get the value of the fields
	fieldValue := fl.Field().Interface().(time.Time)
	paramValue := fl.Parent().FieldByName(paramName).Interface().(time.Time)

	// Check if end date is after or equal to start date
	return fieldValue.After(paramValue) || fieldValue.Equal(paramValue)
}

// Custom validation function to check if the field is a boolean.
func isBoolean(fl validator.FieldLevel) bool {
	_, ok := fl.Field().Interface().(bool)
	return ok
}

// isTimeFormat validates if a string is a valid time format of "15:04"
func isTimeFormat(fl validator.FieldLevel) bool {
	_, err := time.Parse("15:04", fl.Field().String())
	return err == nil
}

func isDate(fl validator.FieldLevel) bool {
	_, err := time.Parse("2006-01-02", fl.Field().String())
	return err == nil
}

func isInt(fl validator.FieldLevel) bool {
	value := fl.Field().Interface()

	switch value := value.(type) {
	case int:
		// 値がint型の場合はバリデーションを通過
		return true
	case string:
		// 値が文字列の場合は文字列からintに変換してバリデーションを行う
		strValue := value
		intValue, err := strconv.Atoi(strValue)
		if err != nil {
			// 文字列からintへの変換に失敗した場合はバリデーションエラー
			return false
		}
		return intValue > 0
	default:
		// 上記以外の型の場合はバリデーションエラー
		return false
	}
}
