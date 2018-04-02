package config

import (
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

const fileConfig = "server.yaml"

// Init add command configuration to inject cobra command
func Init(command *cobra.Command) {
	// read bool flags from command
	for _, f := range boolFlags {
		flag.BoolP(f.name, f.short, f.defaultValue, f.description)
	}
	flag.Parse()
	viper.BindPFlags(flag.CommandLine)

	flagsSetting := getAdditionalSettings()

	// fill command flags from cmd flags if are set, fill from file/env otherwise
	for _, f := range boolFlags {
		var value bool

		// if flag was added by cmd use it; if not use form file/env
		if flag.Lookup(f.name).Changed {
			value = viper.GetBool(f.name)
		} else {
			value = flagsSetting[f.name].Get()
		}
		command.PersistentFlags().BoolP(f.name, f.short, value, f.description)
	}
}

// create box for settings from file/env
// fill the box by data if is set
func getAdditionalSettings() map[string]*BoolConfig {
	v := readConfigFile()
	flagsSetting := map[string]*BoolConfig{}

	for _, f := range boolFlags {
		// create config (single flag) option
		c := InitBoolConfig(f.name)
		flagsSetting[f.name] = c

		// fill config from file
		if v.Get(shortName(f.name)) != nil {
			c.SetFileValue(v.Get(shortName(f.name)).(bool))
		}

		// TODO: implement env values
	}

	return flagsSetting
}

// Read config file (file name comes from const)
func readConfigFile() *viper.Viper {
	viper.SetConfigFile(fileConfig)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Cannot read config file: (%s). Default or Flags config will be used.", err)
	}

	return viper.GetViper()
}
