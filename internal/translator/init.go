package translator

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
	"os"
	"path/filepath"
)

var bundle *i18n.Bundle
var availableKeys []Key

const LocalePath = "locales"

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	err := filepath.Walk(LocalePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			res, err := bundle.LoadMessageFile(path)
			if err != nil {
				return err
			}
			if res.Tag == language.English {
				for _, message := range res.Messages {
					availableKeys = append(availableKeys, Key(message.ID))
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalf("failed loading translations: %v", err)
	}
}
