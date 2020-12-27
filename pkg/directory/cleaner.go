// Copyright 2020 Singularity, Inc. All rights reserved.

package directory

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func RemoveEmpty(path string) error {
	return filepath.Walk(path, cleaner)
}

func cleaner(path string, info os.FileInfo, err error) error {
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		log.Printf("clean directory failed: %v", err)
		return nil
	}

	if info.IsDir() {
		if info.Size() == 0 {
			return os.Remove(path)
		}
		infos, err := ioutil.ReadDir(path)
		if err != nil {
			log.Printf("failed to open dir: %v", err)
			return nil
		}
		if len(infos) > 0 {
			return nil
		}
		if err = os.Remove(path); err != nil {
			log.Printf("failed to remove dir: %v", err)
			return nil
		}
		return filepath.Walk(strings.TrimSuffix(path, info.Name()), cleaner)
	}

	return nil
}
