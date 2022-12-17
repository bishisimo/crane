// Describe:
package validate

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"reflect"
	"strings"
	"sync"
)

var validateOnce sync.Once
var Validator *Validate

func init() {
	validateOnce.Do(func() {
		var err error
		Validator, err = NewValidate()
		if err != nil {
			log.Fatal().Send()
		}
	})
}

type Validate struct {
	trans ut.Translator
	*validator.Validate
	BookingTrans map[string]string
}

func NewValidate() (*Validate, error) {
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()

	//通过label标签返回自定义错误内容
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("json")
		if label == "" {
			return field.Name
		}
		return label
	})

	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return nil, err
	}
	v := &Validate{
		trans:        trans,
		Validate:     validate,
		BookingTrans: BookingTrans,
	}
	return v, nil
}

func (v Validate) ValidateErr(err error) error {
	if err == nil {
		return nil
	}
	e := errors.New("")
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, err := range errs {
			errTrans := v.CustomerErr(err)
			log.Error().Msg(errTrans)
			e = errors.WithMessage(e, errTrans)
		}
	}
	return e
}

func (v Validate) CustomerErr(err validator.FieldError) string {
	errTrans := err.Translate(v.trans)
	if strings.Contains(errTrans, "Field validation for") {
		if format, ok := v.BookingTrans[err.Tag()]; ok {
			errTrans = fmt.Sprintf(format, err.Field(), err.Value())
		}
	}
	return errTrans
}

// Struct 验证Struct并翻译
func (v Validate) Struct(s interface{}) error {
	err := v.Validate.Struct(s)
	if err != nil {
		return v.ValidateErr(err)
	}
	return nil
}
