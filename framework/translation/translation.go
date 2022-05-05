package translation

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/zxdstyle/liey-admin/framework/translation/loader"
)

type (
	Translation struct {
		data   *gmap.StrStrMap
		loader loader.Loader
	}
)

var (
	instance        = gmap.NewStrAnyMap(true)
	defaultI18nPath = "resources/lang"
)

func Translator() *Translation {
	return instance.GetOrSetFuncLock("translator", func() interface{} {
		val, _ := g.Cfg("app").Get(context.TODO(), "i18nPath", defaultI18nPath)
		paths := val.Strings()
		t := &Translation{
			data:   gmap.NewStrStrMap(true),
			loader: loader.NewFileLoaderWithPath(paths...),
		}
		return t
	}).(*Translation)
}

// SetLoader  设置多语言资源默认加载器
func (t *Translation) SetLoader(loader loader.Loader) {
	t.loader = loader
}

// LoadDefaultTranslations 加载默认多语言资源
func (t *Translation) LoadDefaultTranslations() {
	t.AddTranslations(t.loader)
}

// AddTranslations 添加新的多语言资源
func (t *Translation) AddTranslations(loader loader.Loader) error {
	data, err := loader.Load()
	if err != nil {
		return err
	}
	t.data.Sets(data)
	return nil
}

// GetLocale 获取系统默认语言配置
func (t *Translation) GetLocale() string {
	ctx := context.Background()
	val, _ := g.Cfg("app").Get(ctx, "locale", "en")
	return val.String()
}

// GetFallbackLocale 获取系统备用语言配置
func (t *Translation) GetFallbackLocale() string {
	ctx := context.Background()
	val, _ := g.Cfg("app").Get(ctx, "fallbackLocale", "en")
	return val.String()
}

// Translate 翻译为系统配置的默认语言或备用语言
func (t *Translation) Translate(key string, params ...interface{}) string {
	locale := t.GetLocale()
	if val := t.doTranslate(locale, key, params...); len(val) != 0 {
		return val
	}

	locale = t.GetFallbackLocale()
	return t.doTranslate(locale, key, params...)
}

// TranslateWithLang 翻译为指定语言
func (t *Translation) TranslateWithLang(lang, key string, params ...interface{}) string {
	return t.doTranslate(lang, key, params...)
}

func (t *Translation) doTranslate(locale string, key string, params ...interface{}) string {
	k := fmt.Sprintf("%s.%s", locale, key)
	val, ok := t.data.Search(k)
	if ok {
		return fmt.Sprintf(val, params...)
	}
	return ""
}
