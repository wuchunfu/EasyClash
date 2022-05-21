package update

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io/ioutil"
	"net/http"
	"time"

	version "github.com/hashicorp/go-version"
)

var appVersion = "v1.1"

const latestReleaseURL = "https://api.github.com/repos/daodao97/easyclash/releases/latest"

var noUpdate = errors.New("no update available")

type releaseResponse struct {
	TagName string `json:"tag_name"`
	HTMLURL string `json:"html_url"`
}

type releaseInfo struct {
	OldVersion string `json:"old_version"`
	NewVersion string `json:"new_version"`
	URL        string `json:"url"`
}

var ghClient *http.Client

func init() {
	ghClient = &http.Client{
		Timeout: 5 * time.Second,
	}
}

type Options struct {
	CurrentVersion string
	ReleaseUrl     string
}

func New(opt *Options) *update {
	return &update{
		currentVersion: opt.CurrentVersion,
		releaseUrl:     opt.ReleaseUrl,
	}
}

type update struct {
	currentVersion string
	releaseUrl     string
}

func (u *update) Check(ctx context.Context) {
	nev, err := u.checkForUpdate()
	if err != noUpdate {
		result, err := runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
			Type:          runtime.WarningDialog,
			Title:         "提醒",
			Message:       "发现新版本" + nev.NewVersion,
			Buttons:       []string{"忽略", "更新"},
			DefaultButton: "更新",
			CancelButton:  "忽略",
			Icon:          nil,
		})
		if err != nil {
			return
		}
		if result == "更新" {
			runtime.BrowserOpenURL(ctx, nev.URL)
		}
		return
	}
}

func (u *update) checkForUpdate() (*releaseInfo, error) {
	req, err := http.NewRequest("GET", u.releaseUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	resp, err := ghClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %v", resp.StatusCode)
	}

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	r := &releaseResponse{}
	if err := json.Unmarshal(raw, r); err != nil {
		return nil, err
	}

	if versionGreaterThanOrEqual(u.currentVersion, r.TagName) {
		return nil, noUpdate
	}

	return &releaseInfo{
		OldVersion: u.currentVersion,
		NewVersion: r.TagName,
		URL:        r.HTMLURL,
	}, nil
}

func versionGreaterThanOrEqual(v, w string) bool {
	vv, ve := version.NewVersion(v)
	vw, we := version.NewVersion(w)

	return ve == nil && we == nil && vv.GreaterThanOrEqual(vw)
}
