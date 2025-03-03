package rule

import (
	"context"
	"regexp"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gvalid"
)

func RegisterRule() {
	g.Dump("注册自定义校验规则")

	decimal := "custom-decimal"
	gvalid.RegisterRule(decimal, customDecimal)
	pinyin := "custom-pinyin"
	gvalid.RegisterRule(pinyin, customPinyin)
	password := "custom-password"
	gvalid.RegisterRule(password, customPassword)
}

func customDecimal(ctx context.Context, in gvalid.RuleFuncInput) error {
	regex := regexp.MustCompile(`^(?:-)?\d+(?:\.\d{1,2})?$`)
	value := in.Value.String()
	if ok := regex.MatchString(value); ok {
		return nil
	}
	return gerror.New(in.Message)
}

func customPinyin(ctx context.Context, in gvalid.RuleFuncInput) error {
	regex := regexp.MustCompile(`^[0-9A-Z !@#$%^&*()-_=+\[\]{}|;:'\",.<>?/]{1,30}$`)
	value := in.Value.String()
	if ok := regex.MatchString(value); ok {
		return nil
	}
	return gerror.New(in.Message)
}
func customPassword(ctx context.Context, in gvalid.RuleFuncInput) error {
	regex := regexp.MustCompile(`^[a-zA-Z][0-9A-Z!@#$%^&*()-_=+\[\]{}|;:'\",.<>?/]{5,19}$`)
	value := in.Value.String()
	if ok := regex.MatchString(value); ok {
		return nil
	}
	return gerror.New(in.Message)
}
