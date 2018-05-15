package args

import (
	"os"
	"regexp"

	. "github.com/l3eegbee/pigs/config/confsources"
	"github.com/l3eegbee/pigs/ioc"
)

var valueRegexpSimple *regexp.Regexp = regexp.MustCompile("^--([^=]+)='(.*)'$")
var valueRegexpDouble *regexp.Regexp = regexp.MustCompile("^--([^=]+)=\"(.*)\"$")

var boolRegexp *regexp.Regexp = regexp.MustCompile("^--([^=]+)$")
var noboolRegexp *regexp.Regexp = regexp.MustCompile("^--no-([^=]+)$")

func NewArgsConfigSource(args []string) *SimpleConfigSource {

	env := make(map[string]string)

	for _, arg := range args {

		var match []string

		match = valueRegexpSimple.FindStringSubmatch(arg)
		if match == nil {
			match = valueRegexpDouble.FindStringSubmatch(arg)
		}
		if match != nil {
			env[match[1]] = match[2]
			continue
		}

		match = noboolRegexp.FindStringSubmatch(arg)
		if match != nil {
			env[match[1]] = "false"
		}

		match = boolRegexp.FindStringSubmatch(arg)
		if match != nil {
			env[match[1]] = "true"
		}

	}

	return &SimpleConfigSource{
		Priority: CONFIG_SOURCE_PRIORITY_ARGS,
		Env:      env,
	}

}

func init() {
	ioc.Put(NewArgsConfigSource(os.Args[1:]), "ArgsConfigSource", "ConfigSources")
}