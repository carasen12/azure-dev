package installer

import (
	"os"
	"path/filepath"
	"strings"
)

const cInstalledByFileName = ".installed-by.txt"

type InstallType string

const InstallTypeUnknown InstallType = ""
const InstallTypePs InstallType = "install-azd.ps1"
const InstallTypeSh InstallType = "install-azd.sh"
const InstallTypeBrew InstallType = "brew"
const InstallTypeChoco InstallType = "choco"
const InstallTypeWinget InstallType = "winget"
const InstallTypeDeb InstallType = "deb"
const InstallTypeRpm InstallType = "rpm"

func InstalledBy() InstallType {
	raw := RawInstalledBy()

	switch raw {
	case string(InstallTypePs):
		return InstallTypePs
	case string(InstallTypeSh):
		return InstallTypeSh
	case string(InstallTypeBrew):
		return InstallTypeBrew
	case string(InstallTypeChoco):
		return InstallTypeChoco
	case string(InstallTypeWinget):
		return InstallTypeWinget
	case string(InstallTypeDeb):
		return InstallTypeDeb
	case string(InstallTypeRpm):
		return InstallTypeRpm
	default:
		return InstallTypeUnknown
	}
}

func RawInstalledBy() string {
	exePath, err := os.Executable()

	if err != nil {
		return ""
	}

	resolvedPath, err := filepath.EvalSymlinks(exePath)
	if err != nil {
		return ""
	}

	exeDir := filepath.Dir(resolvedPath)
	installedByFile := filepath.Join(exeDir, cInstalledByFileName)

	bytes, err := os.ReadFile(installedByFile)
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(bytes))
}
