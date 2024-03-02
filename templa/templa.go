package templa

import (
	"os"
	"path/filepath"
	"text/template"
)

func getenvs() map[string]string {
	envMap := make(map[string]string)
	for _, env := range os.Environ() {
		for i := range len(env) {
			if env[i] == '=' {
				envMap[env[:i]] = env[i+1:]
				break
			}
		}
	}
	return envMap
}

func Run(srcRoot, dstRoot string) error {
	envMap := getenvs()
	return filepath.Walk(srcRoot, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcRoot, srcPath)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dstRoot, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		dstFile, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		tmpl, err := template.ParseFiles(srcPath)
		if err != nil {
			return err
		}
		return tmpl.Execute(dstFile, envMap)
	})
}
