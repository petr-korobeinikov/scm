package history

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/petr-korobeinikov/scm/internal"
)

func LastWrite(histEntry HistEntry) error {
	// todo override default path
	path, err := internal.ExpandHomeDir(DefaultLastFile)
	if err != nil {
		return err
	}

	// fixme hardcoded file mode
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)

	marshalled, err := json.Marshal(histEntry)
	if err != nil {
		return err
	}

	if _, err := file.Write(marshalled); err != nil {
		return err
	}

	return nil
}

func LastRead() error {
	path, err := internal.ExpandHomeDir(DefaultLastFile)
	if err != nil {
		return err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var histEntry HistEntry
	err = json.Unmarshal(data, &histEntry)
	if err != nil {
		return err
	}

	// todo remote or local?
	// todo return instead of print
	fmt.Println(histEntry.Local)

	return nil
}

type (
	HistEntry struct {
		Remote string `json:"remote"`
		Local  string `json:"local"`
	}
)

const (
	DefaultLastFile = "~/.config/scm/last"
)
