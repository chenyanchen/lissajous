// Revision history:
//     Init: 2019/12/1    Jon Snow

package worker

import (
	"os"
)

func DirOpener(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	filesName, err := file.Readdirnames(-1)
	if err != nil {
		return nil, err
	}

	for i := range filesName {
		filesName[i] = path + "/" + filesName[i]
	}

	return filesName, nil
}
