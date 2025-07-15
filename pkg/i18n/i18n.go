package i18n

import (
	"app/lang"
	"io/fs"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

// 初始化 i18n 包
func InitBundle() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("lang/active.zh_tw.toml")

	fs.WalkDir(lang.LocaleFS, ".", func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".toml" {
			data, _ := lang.LocaleFS.ReadFile(path)
			bundle.ParseMessageFileBytes(data, path)
		}
		return nil
	})
}

// 建立 Localizer 實例
func Localize(lang string, message string) (string, error) {
	localizer := i18n.NewLocalizer(bundle, lang)

	return localizer.Localize(&i18n.LocalizeConfig{
		MessageID: message,
	})
}
