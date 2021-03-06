package programmatic

import (
	. "github.com/l3eegbee/pigs/config/confsources"
	"github.com/l3eegbee/pigs/ioc"
)

func SetEnvForTestsWithPriority(priority int, env map[string]string) {
	ioc.TestPut(
		&SimpleConfigSource{
			Priority: priority,
			Env:      env,
		},
		"ProgrammaticConfigSource", "ConfigSources")
}

func SetEnvForTests(env map[string]string) {
	SetEnvForTestsWithPriority(CONFIG_SOURCE_PRIORITY_TESTS, env)
}
