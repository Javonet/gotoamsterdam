//go:build !excludeBinaries
// +build !excludeBinaries

package binariesextractor

import (
	"archive/zip"
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"unicode"
)

var (
	//go:embed Binaries.zip
	b []byte
)

type BinariesExtractor struct {
	mainWorkingDir  string
	binariesZipPath string
}

func NewBinariesExtractor() *BinariesExtractor {
	return &BinariesExtractor{}
}

func (be *BinariesExtractor) ExtractFromBinariesZip(runtimeName string) error {
	// first letter to upper case
	rr := []rune(runtimeName)
	rr[0] = unicode.ToUpper(rr[0])
	runtimeNameTitle := string(rr)
	be.mainWorkingDir, _ = os.Getwd()
	be.binariesZipPath = filepath.Join(be.mainWorkingDir, "Binaries.zip")

	// remove the binaries folder if it exists
	defer os.Remove(be.binariesZipPath)

	// check if the Binaries.zip file exists, if not create it
	_, err := os.Stat(be.binariesZipPath)
	if os.IsNotExist(err) {
		os.WriteFile(be.binariesZipPath, b, fs.ModeAppend|0777)
	}

	// extract the binaries from the Binaries.zip file
	binariesRuntimePath := be.getBinariesSpecificPath(runtimeNameTitle)
	currentBinariesVersion := be.getCurrentBinariesVersion(runtimeNameTitle)
	zippedBinariesVersion := be.getZippedBinariesVersion()

	r, err := zip.OpenReader(be.binariesZipPath)
	if err != nil {
		return err
	}
	// close the zip file

	defer r.Close()

	// check if file exists and if the version is the same as the zipped version
	if currentBinariesVersion != zippedBinariesVersion {
		os.RemoveAll(fmt.Sprintf("%s/Binaries/%s", be.mainWorkingDir, runtimeNameTitle))
	}

	for _, f := range r.File {
		if strings.Contains(f.Name, binariesRuntimePath) && f.Name != "" {
			thisFilePath := fmt.Sprintf("%s/%s", be.mainWorkingDir, f.Name)
			if !be.fileExists(thisFilePath) || currentBinariesVersion != zippedBinariesVersion {
				if f.FileInfo().IsDir() {
					os.MkdirAll(filepath.Dir(thisFilePath), os.ModePerm)
					continue
				}
				rc, err := f.Open()
				if err != nil {
					return err
				}
				defer rc.Close()

				// if f is not a directory
				outFile, err := os.OpenFile(thisFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
				if err != nil {
					return err
				}
				defer outFile.Close()

				_, err = io.Copy(outFile, rc)
				if err != nil {
					return err
				}

			}
		}
	}

	err = os.WriteFile(fmt.Sprintf("%s/Binaries/%s/version.txt", be.mainWorkingDir, runtimeNameTitle), []byte(zippedBinariesVersion), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (be *BinariesExtractor) getCurrentBinariesVersion(runtimeName string) string {
	currentBinariesVersion := "none"
	currentVersionFile := fmt.Sprintf("%s/Binaries/%s/version.txt", be.mainWorkingDir, runtimeName)

	if be.fileExists(currentVersionFile) {
		file, err := os.Open(currentVersionFile)
		if err != nil {
			return currentBinariesVersion
		}
		defer file.Close()

		reader := bufio.NewReader(file)
		line, _, err := reader.ReadLine()
		if err != nil {
			return currentBinariesVersion
		}

		currentBinariesVersion = string(line)
	}

	return currentBinariesVersion
}

func (be *BinariesExtractor) getZippedBinariesVersion() string {
	//TO DO: implement this function
	zippedBinariesVersion := "2.0.0.0"

	r, err := zip.OpenReader(be.binariesZipPath)
	if err != nil {
		return zippedBinariesVersion
	}
	defer r.Close()

	for _, f := range r.File {
		if f.Name == "Binaries/version.txt" {
			rc, err := f.Open()
			if err != nil {
				return zippedBinariesVersion
			}
			defer rc.Close()

			buf := new(bytes.Buffer)
			buf.ReadFrom(rc)
			zippedBinariesVersion = strings.Trim(buf.String(), "\n")
			break
		}
	}

	return zippedBinariesVersion
}

func (be *BinariesExtractor) getBinariesSpecificPath(runtimeName string) string {
	return fmt.Sprintf("Binaries/%s/%s/%s/", runtimeName, be.getOsName(), be.getEnvBitness())
}

func (be *BinariesExtractor) getEnvBitness() string {
	if runtime.GOARCH == "amd64" {
		return "X64"
	}
	if runtime.GOARCH == "386" {
		return "X86"
	}
	return ""
}

func (be *BinariesExtractor) getOsName() string {
	if runtime.GOOS == "linux" {
		return "Linux"
	}
	if runtime.GOOS == "windows" {
		return "Windows"
	}
	if runtime.GOOS == "darwin" {
		return "MacOs"
	}
	return ""
}

func (be *BinariesExtractor) fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
