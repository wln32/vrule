package vrule

import (
	"fmt"
)

var (
	i18nSupportLanguage = []string{
		"en", "zh-CN", "cn",
	}
)

func i18nSupport(arr []string, s string) bool {
	for _, e := range arr {
		if e == s {
			return true
		}
	}
	return false
}

func (v *Validator) I18nSetLanguage(language string) *Validator {
	if i18nSupport(i18nSupportLanguage, language) == false {
		panic(fmt.Errorf("language not support %s", language))
	}
	if language == "cn" {
		language = "zh-CN"
	}

	v.i18n.SetLanguage(language)
	return v
}

func (v *Validator) I18nSetPath(path string) *Validator {
	err := v.i18n.SetPath(path)
	if err != nil {
		return nil
	}
	return v
}
