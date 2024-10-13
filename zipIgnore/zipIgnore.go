package zipIgnore

import (
	"archive/zip"
	"bufio"
	"fmt"
	ignore "github.com/sabhiram/go-gitignore"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// 读取.gitignore文件并返回一个跳过文件夹的数组
func readGitignore(filePath string) (*ignore.GitIgnore, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var patterns []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) != "" && !strings.HasPrefix(line, "#") {
			patterns = append(patterns, strings.TrimSpace(line))
		}
	}
	return ignore.CompileIgnoreLines(patterns...), scanner.Err()
}

// 判断文件夹是否在跳过列表中
func shouldSkip(path string, ign *ignore.GitIgnore) bool {
	return ign.MatchesPath(path)
}

// 压缩文件夹
func zipFolder(source string, target string, ign *ignore.GitIgnore) error {
	zipFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}
		if shouldSkip(path, ign) {
			return nil
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(path, filepath.Dir(source)+"/")
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return nil
}

func ZipFiles() {
	// 读取.gitignore文件
	ign, err := readGitignore(".gitignore")
	if err != nil {
		fmt.Println("Error reading .gitignore:", err)
		return
	}

	// 获取顶层文件夹列表
	root := "/Users/g/Desktop/work/carrier" // 根目录
	dirs, err := os.ReadDir(root)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	var wg sync.WaitGroup
	sem := make(chan struct{}, 45) // 并发限制为5

	for _, dir := range dirs {
		isDirSkip := shouldSkip(dir.Name(), ign)
		fmt.Println(dir.Name(), ign)

		if dir.IsDir() && !isDirSkip {
			wg.Add(1)
			sem <- struct{}{} // 占用一个并发槽

			go func(dirName string) {
				defer wg.Done()
				defer func() { <-sem }() // 释放并发槽

				source := filepath.Join(root, dirName)
				target := filepath.Join(root, dirName+".zip")

				// 压缩文件夹
				err := zipFolder(source, target, ign)
				if err != nil {
					fmt.Println("Error zipping folder:", err)
					return
				}

				// 打印日志
				fmt.Printf("Time: %s, Folder: %s\n", time.Now().Format(time.RFC3339), dirName)

				// 打印被压缩的内容
				fmt.Printf("Compressed contents of %s:\n", dirName)
				filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					fmt.Println(path)
					return nil
				})
			}(dir.Name())
		}
	}

	wg.Wait()
}
