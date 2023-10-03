package syscmd

import (
	"fmt"
	"os"
	"path"

	log "github.com/golang/glog"
)

// SaveToFile saves the input string into a file in the file system.
// It creates the file and all parent folder if not exist.
func SaveToFile(folderPath, fileName, content string) error {
	fileFullPath := path.Join(folderPath, fileName)
	log.Infof("Saving content to file %v...", fileFullPath)

	fileInfo, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		// Config folder does not exist, create one.
		if err = os.MkdirAll(folderPath, 0755); err != nil {
			log.Errorf("Creating folder %v failed. Error: %v.", folderPath, err)
			return err
		}
		log.Infof("Created folder %v.", folderPath)
	} else if !fileInfo.Mode().IsDir() {
		log.Errorf("Unable to save content to %v, since %v points an existing file.", fileFullPath, folderPath)
		return fmt.Errorf("%s points an existing file", folderPath)
	}

	if err := os.WriteFile(fileFullPath, []byte(content), 0600); err != nil {
		log.Errorf("Saving content to file %v failed. Error: %v.", fileFullPath, err)
		return err
	}
	log.Infof("Saved content to file %v.", fileFullPath)
	return nil
}
