package tools

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetRandomString2 生成指定长度的随机字符串
func GetRandomString2(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

// GetTempPathFileName 获取一个临时文件名
func GetTempPathFileName() (pathFileName string) {
	return fmt.Sprintf("%s.tmp", GetRandomString2(16))
}

// GetTempPNGPathFileName 获取一个临时文件名，后缀为PNG
func GetTempPNGPathFileName() (pathFileName string) {
	return filepath.Join(os.TempDir(), fmt.Sprintf("%s.png", GetRandomString2(16)))
}

// GetTempPathDirName 获取一个临时目录
func GetTempPathDirName() (pathDirName string) {
	return filepath.Join(os.TempDir(), fmt.Sprintf("%s.dir", GetRandomString2(16)))
}




// DownloadFile 下载文件
func DownloadFile(url, dstPathFile string) (bool, error) {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	// 创建一个文件用于保存
	out, err := os.Create(dstPathFile)
	if err != nil {
		return false, err
	}
	defer out.Close()
	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CheckFileExist 检测文件或目录是否存在
func CheckFileExist(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// MakePath 创建目录，如果目录存在则直接返回
func MakePath(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		if err = os.MkdirAll(filepath, 0777); err == nil {
			return true
		}
	}
	return false
}



type BinShortName string



// GetThirdpartyBinNameByPlatform 根据当前运行平台及架构，生成指定的文件名称
func GetThirdpartyBinNameByPlatform(binShortName BinShortName) (binPlatformName string) {
	binPlatformName = fmt.Sprintf("%s_%s_%s", binShortName, runtime.GOOS, runtime.GOARCH)
	if runtime.GOOS == "windows" {
		binPlatformName += ".exe"
	}
	/*
		https://go.dev/doc/install/source#environment
			$GOOS	$GOARCH
			android   arm
			darwin    386
			darwin    amd64
			darwin    arm
			darwin    arm64
			dragonfly amd64
			freebsd   386
			freebsd   amd64
			freebsd   arm
			linux     386
			linux     amd64
			linux     arm
			linux     arm64
			linux     ppc64
			linux     ppc64le
			linux     mips
			linux     mipsle
			linux     mips64
			linux     mips64le
			netbsd    386
			netbsd    amd64
			netbsd    arm
			openbsd   386
			openbsd   amd64
			openbsd   arm
			plan9     386
			plan9     amd64
			solaris   amd64
			windows   386
			windows   amd64
	*/
	return
}
