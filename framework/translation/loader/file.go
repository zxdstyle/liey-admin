package loader

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gyaml"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

type FileLoader struct {
	paths []string
}

func NewFileLoaderWithPath(paths ...string) *FileLoader {
	return &FileLoader{paths: paths}
}

func (f FileLoader) Load() (data map[string]string, err error) {
	for _, path := range f.paths {
		if err = f.doLoad(path, &data); err != nil {
			return nil, err
		}
	}
	return
}

func (f FileLoader) doLoad(p string, data *map[string]string) error {
	realPath, _ := gfile.Search(p)
	if len(realPath) == 0 {
		return fmt.Errorf(`%s does not exist`, p)
	}

	files, _ := gfile.ScanDirFile(realPath, "*.yaml", true)
	if len(files) == 0 {
		return nil
	}
	var (
		path string
		ctx  = context.Background()
	)
	for _, file := range files {
		path = file[len(realPath)+1:]
		key := gstr.ReplaceByMap(path, map[string]string{
			"/":     ".",
			".yaml": "",
		})

		var content map[string]string
		if err := gyaml.DecodeTo(gfile.GetBytes(file), &content); err != nil {
			g.Log().Fatal(ctx, err)
		}

		for attr, format := range content {
			key = fmt.Sprintf("%s.%s", key, attr)
			(*data)[key] = format
		}
	}
	return nil
}
