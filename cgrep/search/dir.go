package search

import (
	"bufio"
	"cgrep/errors"
	"cgrep/result"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

var (
	currentDir string
	gitRegExp  = regexp.MustCompile(`\.git$`)
)

type Dir interface {
	Search()
}

type dir struct {
	wg            *sync.WaitGroup
	path          string
	regexp        *regexp.Regexp
	subDirs       []Dir
	fileFullPaths []string
}

// ディレクトリごとに検索用オブジェクトを生成するファクトリ関数
func New(wg *sync.WaitGroup, fullPath string, re *regexp.Regexp) (Dir, error) {
	d := &dir{wg: wg, path: fullPath, regexp: re}
	if d.isGitDri() {
		return d, nil
	}

	d.Scan()
	return d, nil
}

// func New() を実行した際、自身のサブディレクトリとファイル郡をスキャンする処理
func (d *dir) Scan() error {
	fs, err := ioutil.ReadDir(d.path)
	if err != nil {
		return err
	}

	for _, f := range fs {
		path := filepath.Join(d.path, f.Name())
		if f.IsDir() {
			subDir, err := New(d.wg, path, d.regexp)
			if err != nil {
				return err
			}
			d.subDirs = append(d.subDirs, subDir)
			continue
		}

		d.fileFullPaths = append(d.fileFullPaths, path)
	}

	return nil
}

func (d *dir) Search() {
	defer d.wg.Done()

	for _, subDir := range d.subDirs {
		d.wg.Add(1)
		go subDir.Search()
	}

	if err := d.GrepFiles(); err != nil {
		errors.Set(err)
	}
}

func (d *dir) GrepFiles() error {
	for _, fileFullPath := range d.fileFullPaths {
		file, err := os.Open(fileFullPath)
		if err != nil {
			return err
		}
		defer file.Close()

		lineNumber := 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lineNumber++
			line := scanner.Text()
			if !d.regexp.MatchString(line) {
				continue
			}

			relPath, err := relativePath(file)
			if err != nil {
				return err
			}

			result.Set(relPath, line, lineNumber)
		}
		if err := scanner.Err(); err != nil {
			return err
		}
	}

	return nil
}

// 自身が .git ディレクトリであるかを検証するメソッド
func (d *dir) isGitDri() bool {
	return gitRegExp.MatchString(d.path)
}

// *os.File を渡すと、ファイル名にカレントディレクトリからそのファイルまでのフルパスを添えて返す関数
func relativePath(file *os.File) (string, error) {
	return filepath.Rel(currentDir, file.Name())
}

// 処理の開始時に実行される関数
func init() {
	currentDir, _ = os.Getwd()
}
