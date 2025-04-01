package enginecache

import (
	"fmt"
	"github.com/go-flutter-desktop/hover/internal/log"
	"io"
	"os"
)

// DefaultCachePath tries to resolve the user cache directory. DefaultCachePath
// may return an empty string when none was found, in that case it will print a
// warning to the user.
func DefaultCachePath() string {
	cachePath, err := os.UserCacheDir()
	if err != nil {
		log.Warnf("Failed to resolve cache path: %v", err)
	}
	return cachePath
}

// CopyFile 使用 io.Copy 函数将文件从源路径复制到目标路径
func CopyFile(src, dst string) error {
	// 打开源文件
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	// 确保在函数结束时关闭源文件
	defer sourceFile.Close()

	// 创建目标文件
	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	// 确保在函数结束时关闭目标文件
	defer destFile.Close()

	// 使用 io.Copy 函数复制文件内容
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	// 同步文件内容到磁盘
	err = destFile.Sync()
	if err != nil {
		return err
	}

	// 获取源文件的权限信息
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	// 设置目标文件的权限与源文件一致
	err = os.Chmod(dst, srcInfo.Mode())
	if err != nil {
		return err
	}

	return nil
}

// CreateDirIfNotExists 检查目录是否存在，如果不存在则创建它
func CreateDirIfNotExists(dir string) error {
	// 使用 os.Stat 函数获取目录的信息
	_, err := os.Stat(dir)

	// 如果返回的错误是 os.ErrNotExist，说明目录不存在
	if os.IsNotExist(err) {
		// 使用 os.MkdirAll 函数创建目录，权限设置为 0755
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
		fmt.Printf("目录 %s 已创建\n", dir)
		return nil
	} else if err != nil {
		// 如果返回其他错误，则直接返回该错误
		return err
	}
	fmt.Printf("目录 %s 已存在\n", dir)
	return nil
}
