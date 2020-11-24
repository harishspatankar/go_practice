package main

import (
	"encoding/json"
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {
	params := map[string]interface{}{
		"Name":  "Alex",
		"Count": 10,
	}
	fmt.Println(t("dashboard.messages", params, 3))
}

func t(message_id string, params map[string]interface{}, messageCount int) string {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile("en.json")
	localizer := i18n.NewLocalizer(bundle, "en")

	translation := localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    message_id,
		TemplateData: params,
		PluralCount:  messageCount,
	})

	return translation
}
