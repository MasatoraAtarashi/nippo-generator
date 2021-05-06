package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

type DefaultConfig struct {
	Template []string
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize nippo config",
	Run: func(cmd *cobra.Command, args []string) {
		err := runInitCmd(cmd, args)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		} else {
			fmt.Println("Initialized nippo config at $HOME/.nippo.yaml")
		}
	},
}

func runInitCmd(cmd *cobra.Command, args []string) (err error) {
	// configファイルがすでに存在していたらその旨を表示して終了
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("nippo config is already initialized")
		os.Exit(0)
	}
	// configファイルを作成する
	fpath, err := makeConfigFile()
	if err != nil {
		// errがあったらconfigファイル消す
		deleteFile(fpath)
		return
	}
	return
}

func makeConfigFile() (fpath string, err error) {
	// defaultのconfigファイルの内容を取得
	defaultContent, err := initDefaultContentOfConfig()
	if err != nil {
		return
	}

	// configファイルを作成
	home := os.Getenv("HOME")
	fpath = filepath.Join(home, ".nippo.yaml")
	if !isFileExist(fpath) {
		err = ioutil.WriteFile(fpath, defaultContent, 0644)
		if err != nil {
			return
		}
	}
	return
}

// init default content of config
func initDefaultContentOfConfig() (defaultContent []byte, err error) {
	template := []string {"今日やったこと", "明日の予定", "所感・連絡事項", "git", "slack"}
	data := DefaultConfig{
		Template: template,
	}
	defaultContent, err = yaml.Marshal(data)
	if err != nil {
		return
	}
	return
}

func init() {
	rootCmd.AddCommand(initCmd)
}
