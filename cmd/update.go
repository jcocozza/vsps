package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

const (
	repoName  = "vsps"
	repoOwner = "jcocozza"
)

// returned by github api
type releaseInfo struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name               string `json:"name"`
		BrowserDownloadUrl string `json:"browser_download_url"`
	} `json:"assets"`
}

func getLatestReleaseInfo() (string, string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", repoOwner, repoName)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return "", "", err
	}
	defer resp.Body.Close()

	var release releaseInfo
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", "", err
	}

	for _, asset := range release.Assets {
		if strings.Contains(asset.Name, runtime.GOOS) && strings.Contains(asset.Name, runtime.GOARCH) {
			return release.TagName, asset.BrowserDownloadUrl, nil
		}
	}
	return release.TagName, "", fmt.Errorf(fmt.Sprintf("failed to find cli binary for %s with architecture %s", runtime.GOOS, runtime.GOARCH))
}

func downloadAndReplaceBinary(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	tmpFile, err := os.CreateTemp("", "vsps")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	if _, err := io.Copy(tmpFile, resp.Body); err != nil {
		return err
	}

	if runtime.GOOS != "windows" {
		if err := tmpFile.Chmod(0755); err != nil {
			return err
		}
	}

	binaryPath, err := os.Executable()
	if err != nil {
		return err
	}

	backupPath := binaryPath + ".bak"
	if err := os.Rename(binaryPath, backupPath); err != nil {
		return fmt.Errorf("failed to backup existing binary: %v", err)
	}

	if err := os.Rename(tmpFile.Name(), binaryPath); err != nil {
		// Attempt to restore backup if rename fails
		os.Rename(backupPath, binaryPath)
		return err
	}

	os.Remove(backupPath)
	return nil
}

var updateVsps = &cobra.Command{
	Use:   "update-vsps",
	Short: "update vsps to the latest version. may require sudo depending on where vsps is stored.",
	Long: `update vsps to the latest version. this will not touch the underlying al file.
even if the update goes wrong, your infomation will be fine.`,
	Run: func(cmd *cobra.Command, args []string) {
		lastestRelease, downloadURL, err := getLatestReleaseInfo()
		if err != nil {
			fmt.Println("update failed: failed to fetch latest version info: " + err.Error())
			return
		}

		if lastestRelease == version {
			fmt.Println("vsps is already up to date!")
			return
		}
		fmt.Printf("updating to latest version: %s -> %s\n", version, lastestRelease)
		if err := downloadAndReplaceBinary(downloadURL); err != nil {
			fmt.Println("update failed:", err)
		} else {
			fmt.Println("update successful")
			fmt.Println("if you have completion enabled, you may need to regenerate the completion script and replace the old one. to do so run:")
			fmt.Println("vsps completion [shell-name]")
		}
	},
}

func init() {
	rootCmd.AddCommand(updateVsps)
}
