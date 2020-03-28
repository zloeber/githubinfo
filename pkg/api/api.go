package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/zloeber/githubinfo/pkg/log"
)

var (
	projectPattern = regexp.MustCompile(`^[a-zA-Z0-9_-]+/[a-zA-Z0-9_-]+$`)
	githubAPI      = "https://api.github.com/repos"
)

// IsValidProject determines if the string is a valid github vendor/repo combination
func IsValidProject(str string) bool {
	if !projectPattern.MatchString(str) {
		log.Error(fmt.Sprintf("%s does not match the format for a project - vendor/repo!", str))
		return false
	}
	return true
}

// ProjectJSON will parse github site for full project information
func ProjectJSON(project string) string {
	response, err := http.Get(fmt.Sprintf("%s/%s", githubAPI, project))
	if err != nil {
		log.Error(fmt.Sprintf("%s", err))
		os.Exit(1)
	}
	defer CloseQuietly(response.Body)
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(fmt.Sprintf("%s", err))
		os.Exit(1)
	}
	return string(contents)
}

// ReleasesJSON will parse github site for release
func ReleasesJSON(project string) string {
	response, err := http.Get(fmt.Sprintf("%s/%s/releases/latest", githubAPI, project))
	if err != nil {
		log.Error(fmt.Sprintf("%s", err))
		os.Exit(1)
	}
	//defer response.Body.Close()
	defer CloseQuietly(response.Body)
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(fmt.Sprintf("%s", err))
		os.Exit(1)
	}
	return string(contents)
}

// Description will parse for project description
func Description(payload string) string {
	var result map[string]interface{}
	bytes := []byte(payload)
	err := json.Unmarshal(bytes, &result)
	if err != nil {
		log.Error(fmt.Sprintf("%s", err))
		os.Exit(1)
	}
	return result["description"].(string)
}

// License will parse for project license
func License(payload string) string {
	license := "None Assigned"
	var result map[string]interface{}
	bytes := []byte(payload)
	err := json.Unmarshal(bytes, &result)
	if err != nil {
		log.Error(fmt.Sprintf("%s", err))
		os.Exit(1)
	}
	if result["license"] != nil {
		licenseMap := result["license"].(map[string]interface{})
		license = licenseMap["spdx_id"].(string)
	}
	return string(license)
}

// ReleaseURLs will parse for project license
func ReleaseURLs(payload string) []string {
	var URLs []string
	var result map[string]interface{}
	bytes := []byte(payload)
	err := json.Unmarshal(bytes, &result)
	if err != nil {
		log.Error(fmt.Sprintf("%s", err))
		os.Exit(1)
	}
	if result["assets"] != nil {
		assets := result["assets"].([]interface{})
		for _, asset := range assets {
			assetMap := asset.(map[string]interface{})
			URLs = append(URLs, assetMap["download_url"].(string))
		}
	}
	return URLs
}

// IsJSON will determine if the string is valid json
func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// CloseQuietly closes `io.Closer` quietly. Very handy and helpful for code
// quality coverage testing (but not readability).
func CloseQuietly(v interface{}) {
	if d, ok := v.(io.Closer); ok {
		_ = d.Close()
	}
}
