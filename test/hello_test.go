package test

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

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

func Test_001(t *testing.T) {

	if err := CopyFile("/Users/apple/opt/Jetbrain/hover/FlutterEmbedder.framework.zip", "/Users/apple/Desktop/0108/long/FlutterEmbedder.framework.zip"); err != nil {
		fmt.Println("复制文件失败")
	}

}

func Test_002(t *testing.T) {

	// 创建框架目录结构
	frameworkDestPath := "/Users/apple/Library/Caches/hover/engine/darwin-debug_unopt/FlutterEmbedder.framework"
	err := os.MkdirAll(frameworkDestPath+"/Versions/A", 0755)
	if err != nil {

	}

	//// 然后创建符号链接
	//createSymLink("A", frameworkDestPath+"/Versions/Current")
	//createSymLink("Versions/Current/FlutterEmbedder", frameworkDestPath+"/FlutterEmbedder")
	//// ... 其他符号链接

}

func Test_003(t *testing.T) {
	frameworkPath := "/Users/apple/Library/Caches/hover/engine/darwin-debug_unopt/FlutterEmbedder.framework"

	// 检查 FlutterEmbedder.framework 文件夹是否存在
	if _, err := os.Stat(frameworkPath); os.IsNotExist(err) {
		fmt.Println("FlutterEmbedder.framework 文件夹不存在")
		return
	}

	// 检查 FlutterEmbedder.framework 内的文件和符号链接
	filesToCheck := []string{
		"Versions/Current",
		"FlutterEmbedder",
		"Headers",
		"Modules",
		"Resources",
	}

	for _, file := range filesToCheck {
		fullPath := filepath.Join(frameworkPath, file)
		if _, err := os.Lstat(fullPath); os.IsNotExist(err) {
			fmt.Printf("缺少文件或符号链接: %s\n", fullPath)
		} else {
			fmt.Printf("文件或符号链接存在: %s\n", fullPath)
		}
	}
}

func Test_006(t *testing.T) {
	// 验证 FlutterEmbedder.framework 及其所有必要组件
	frameworkPath := "/Users/apple/Library/Caches/hover/engine/darwin-debug_unopt/FlutterEmbedder.framework"

	// 检查框架是��存在
	info, err := os.Stat(frameworkPath)
	if os.IsNotExist(err) {
		t.Fatalf("框架不存在: %s", frameworkPath)
	}
	fmt.Printf("框架存在，权限: %v\n", info.Mode())

	// 验证二进制文件
	binaryPath := filepath.Join(frameworkPath, "Versions/A/FlutterEmbedder")
	_, err = os.Stat(binaryPath)
	if os.IsNotExist(err) {
		t.Fatalf("引擎二进制文件不存在: %s", binaryPath)
	}

	// 检查是否包含 OpenGL 依赖
	cmd := exec.Command("otool", "-L", binaryPath)
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("无法检查二进制依赖: %v", err)
	}

	fmt.Println("引擎二进制依赖:")
	fmt.Println(string(output))

	// 检查是否有 OpenGL 相关库
	if !strings.Contains(string(output), "OpenGL") {
		fmt.Println("警告: 引擎二进制文件可能不包含 OpenGL 依赖")
	}
}
