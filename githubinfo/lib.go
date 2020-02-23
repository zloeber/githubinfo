package githubinfo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/zloeber/githubinfo/log"
)

var projectPattern = regexp.MustCompile(`^[a-zA-Z0-9_-]+/[a-zA-Z0-9_-]+$`)
var githubAPI = "https://api.github.com/repos"

// Is the string a valid github vendor/repo combination
func IsValidProject(str string) bool {
	if !projectPattern.MatchString(str) {
		log.Error(fmt.Sprintf("%s does not match the format for a project - vendor/repo!", str))
		return false
	}
	return true
}

// Parse github site for full project information
func ProjectJSON(project string) string {
	response, err := http.Get(fmt.Sprintf("%s/%s", githubAPI, project))
	if err != nil {
		log.Error(fmt.Sprintf("%s", err))
		os.Exit(1)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(fmt.Sprintf("%s", err))
		os.Exit(1)
	}
	return string(contents)
}

// Parse github site for release
func ReleasesJSON(project string) string {
	response, err := http.Get(fmt.Sprintf("%s/%s/releases/latest", githubAPI, project))
	if err != nil {
		log.Error(fmt.Sprintf("%s", err))
		os.Exit(1)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(fmt.Sprintf("%s", err))
		os.Exit(1)
	}
	return string(contents)
}

// Parse for project description
func Description(payload string) string {
	var result map[string]interface{}
	json.Unmarshal([]byte(payload), &result)
	return result["description"].(string)
}

// Parse for project license
func License(payload string) string {
	license := "None Assigned"
	var result map[string]interface{}
	json.Unmarshal([]byte(payload), &result)
	if result["license"] != nil {
		licenseMap := result["license"].(map[string]interface{})
		license = licenseMap["spdx_id"].(string)
	}
	return string(license)
}

// Parse for project license
func ReleaseURLs(payload string) []string {
	var URLs []string
	var result map[string]interface{}
	json.Unmarshal([]byte(payload), &result)
	if result["assets"] != nil {
		assets := result["assets"].([]interface{})
		for _, asset := range assets {
			assetMap := asset.(map[string]interface{})
			URLs = append(URLs, assetMap["download_url"].(string))
		}
	}
	return URLs
}

// Is the string json
func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}
