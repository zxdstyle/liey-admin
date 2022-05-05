package loader

// Loader 翻译资源加载器
type Loader interface {
	Load() (map[string]string, error)
}
