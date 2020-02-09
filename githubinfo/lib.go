package githubinfo

import "fmt"
import "json"
import "strings"
import "regexp"
import "github.com/zloeber/githubinfo/log"

var projectPattern = regexp.MustCompile(`^[a-zA-Z0-9_-]+/[a-zA-Z0-9_-]+$`)

// Is the string a valid github vendor/repo combination
func IsValidProject(str string) bool {
	if !projectPattern.MatchString(str) {
		log.Error("%s does not contain '/'", str)
		return false
	}
	return true
}

// Is the string json
func IsJSON(str string) bool {
    var js json.RawMessage
    return json.Unmarshal([]byte(str), &js) == nil
}