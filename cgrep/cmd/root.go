/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cgrep/errors"
	"cgrep/result"
	"cgrep/search"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/spf13/cobra"
)

var dir string
var withContent bool

var rootCmd = &cobra.Command{
	Use:   "cgrep [flags] [args]",
	Short: "Search for file names containing a argument",
	Long: `Search file names contains argument.
Arguments are treated as regular expressions.

Args:
  A search string that can be compiled as a regular expression`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fullPath, err := filepath.Abs(dir)
		if err != nil {
			return err
		}

		if err := ExecSearch(fullPath, args[0]); err != nil {
			return err
		}

		if err := errors.Error(); err != nil {
			return err
		}

		Render(os.Stdout)
		return nil
	},
}

func ExecSearch(fullPath, regexpWord string) error {
	var wg sync.WaitGroup

	re, err := regexp.Compile(regexpWord)
	if err != nil {
		return err
	}

	dir, err := search.New(&wg, fullPath, re)
	if err != nil {
		return err
	}

	wg.Add(1)
	go dir.Search()

	wg.Wait()
	return nil
}

func Render(w io.Writer) {
	if withContent {
		result.RenderWithContent(w)
	} else {
		result.RenderFiles(w)
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&dir, "dir", "d", "./", "searching directory")
	rootCmd.Flags().BoolVarP(&withContent, "with-content", "c", false, "render with matched content lines")
}
