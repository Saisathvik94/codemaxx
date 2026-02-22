package files

import "os"

func WriteFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}
