package main

import (
	"encoding/json"
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {
	loc := Init("en.json")

	params := map[string]interface{}{
		"Name":  "Alex",
		"Count": 10,
	}

	fmt.Println(Translate(loc, "dashboard.messages", params, 3))
	fmt.Println(Translate(loc, "hello_world", params, 3))
}

func Init(fileName string) *i18n.Localizer {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile(fileName)
	loc := i18n.NewLocalizer(bundle, "en")
	return loc
}

func Translate(loc *i18n.Localizer, message_id string, params map[string]interface{}, messageCount int) string {
	translation := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    message_id,
		TemplateData: params,
		PluralCount:  messageCount,
	})

	return translation
}
