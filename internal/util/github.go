package util

import (
	"io"
	"net/http"
	"os"

	"github.com/Optum/dce-cli/configs"
	"github.com/Optum/dce-cli/internal/constants"
	observ "github.com/Optum/dce-cli/internal/observation"
)

type GithubUtil struct {
	Config      *configs.Root
	Observation *observ.ObservationContainer
}

func (u *GithubUtil) DownloadGithubReleaseAsset(assetName string) {
	assetDownloadURL := constants.GithubAssetDownloadURL + assetName
	req, err := http.NewRequest("GET", assetDownloadURL, nil)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer resp.Body.Close()

	out, err := os.Create(assetName)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
