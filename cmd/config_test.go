package cmd

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

type config struct {
	username          string
	preferredLanguage string
}

func TestConfigActions(t *testing.T) {
	currentConfig := config{
		username:          "username",
		preferredLanguage: "preferred_language",
	}

	initConfig()
	viper.Set(usernameConfigKey, currentConfig.username)
	viper.Set(preferredLanguageConfigKey, currentConfig.preferredLanguage)
	viper.WriteConfig()

	testCases := []struct {
		name   string
		config config
	}{
		{
			name: "NoConfigDefined",
			config: config{
				username:          "",
				preferredLanguage: "",
			},
		},
		{
			name: "UsernameNotDefined",
			config: config{
				username:          "",
				preferredLanguage: "cpp",
			},
		},
		{
			name: "LanguageNotDefined",
			config: config{
				username:          "nguyenducloc",
				preferredLanguage: "",
			},
		},
		{
			name: "BothDefined",
			config: config{
				username:          "nguyenducloc",
				preferredLanguage: "cpp",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var out bytes.Buffer

			err := configActions(&out, tc.config.username, tc.config.preferredLanguage)
			require.NoError(t, err)

			err = viper.ReadInConfig()
			require.NoError(t, err)

			if tc.config.username == "" {
				require.Equal(t, currentConfig.username, viper.GetString(usernameConfigKey))
			} else {
				require.Equal(t, tc.config.username, viper.GetString(usernameConfigKey))
			}

			if tc.config.preferredLanguage == "" {
				require.Equal(t, currentConfig.preferredLanguage, viper.GetString(preferredLanguageConfigKey))
			} else {
				require.Equal(t, tc.config.preferredLanguage, viper.GetString(preferredLanguageConfigKey))
			}

			if tc.config.username != "" && tc.config.preferredLanguage != "" {
				require.Equal(t, tc.config.username, viper.GetString(usernameConfigKey))
				require.Equal(t, tc.config.preferredLanguage, viper.GetString(preferredLanguageConfigKey))
			}

			currentConfig = config{
				username:          viper.GetString(usernameConfigKey),
				preferredLanguage: viper.GetString(preferredLanguageConfigKey),
			}

			require.Equal(t, out.String(), fmt.Sprintf("Config written in: %s\n", viper.ConfigFileUsed()))
		})
	}
}
