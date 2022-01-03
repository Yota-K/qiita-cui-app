package setting

import (
	"fmt"
	"os"
	"path/filepath"
)

const Dir = ".qiita"

func InitSetting(t string) {
	if t == "" {
		fmt.Println("Arguments has not been set.")
		return
	}

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// /Users/userName/.qiita
	homeDir := getAbsPath(fmt.Sprintf("%s/%s", home, Dir))

	makeDir(homeDir)
	writeApiToken(home, t)
}

// 相対パスを絶対パスに変換
func getAbsPath(relPath string) string {
	path, err := filepath.Abs(relPath)
	if err != nil {
		fmt.Print(err.Error())
	}
	return path
}

// .qiitaディレクトリを作成
func makeDir(homeDir string) {
	// ディレクトリがない場合は作成
	if _, err := os.Stat(homeDir); os.IsNotExist(err) {
		os.Mkdir(homeDir, 0777)
	}
}

// API Tokenをconfigに書き込む
func writeApiToken(home, t string) {
	absPath := getAbsPath(home)

	// os.O_RDWR・・・読み書きどっちもするとき
	FileName, err := os.OpenFile(fmt.Sprintf("%s/.qiita/config", absPath), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err)
	}

	defer FileName.Close()

	fmt.Fprintln(FileName, t)

	fmt.Println("Success!")
}
