package api

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// .qiita/configに書き込まれたtokenを取得
func GetApiToken() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fileName, err := filepath.Abs(fmt.Sprintf("%s/.qiita/config", home))
	if err != nil {
		fmt.Print(err.Error())
	}

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print(err.Error())
	}

	t := strings.TrimRight(string(bytes), "\n")
	return t
}
