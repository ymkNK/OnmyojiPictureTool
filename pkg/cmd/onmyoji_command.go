package cmd

import (
	"fmt"
	"os"

	"otool/pkg/onmyoji"

	"github.com/spf13/cobra"
)

var onmyojiCmd = &cobra.Command{
	Use:     "sync",
	Aliases: []string{"s"},
	Short:   "A simple tool for onmyoji",
	Long:    `A simple tool for onmyoji, you can get the source code in https://github.com/ymkNK/otool`,
	Run:     syncPic,
}

var downloadPath = ""

func init() {
	onmyojiCmd.PersistentFlags().StringVarP(&downloadPath, "downloadPath", "d", "", "download path for the pictures, default: ./OnmyojiPictures")

	if downloadPath == "" {
		downloadPath = "./OnmyojiPictures"
	}
}

func syncPic(c *cobra.Command, args []string) {
	for _, arg := range args {
		fmt.Println("arg: ", arg)
	}

	fmt.Println("download path: ", downloadPath)

	err := os.MkdirAll(downloadPath, 0777)
	if err != nil {
		fmt.Println("os.MkdirAll err: ", err.Error())

		os.Exit(1)
	}

	entries, err := os.ReadDir(downloadPath)
	if err != nil {
		fmt.Println("os.ReadDir err: ", err.Error())

		os.Exit(1)
	}

	existedFileMap := make(map[string]bool)

	for _, entry := range entries {
		existedFileMap[entry.Name()] = true
	}

	err = onmyoji.RefreshOnmyojiMedia(existedFileMap, downloadPath)
	if err != nil {
		fmt.Println("call onmyoji.RefreshOnmyojiMedia err: ", err.Error())

		os.Exit(1)
	}
}
