package file_init

import (
	. "github.com/l3eegbee/pigs/config/confsources"
	. "github.com/l3eegbee/pigs/config/confsources/file"
	. "github.com/l3eegbee/pigs/config/confsources/file/properties"
)

func init() {

	RegisterFileConfig(
		CONFIG_SOURCE_PRIORITY_FILE_PROPERTIES,
		".properties",
		ParsePropertiesToEnv,
		"PropertiesFileConfigSource")

}
