package config

import (
	"log"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
)

type Target struct {
	Name string `json:"name"`
	BundleId string `json:"bundle_id"`
}
type BuildOption struct {
	Id int `json:"id"`
	Project string `json:"project"`
	Targets []Target `json:"targets"`
	Path string `json:"source_dir"`
}

var (
	BuildOptions []BuildOption
	HttpEndPoint string
	WsEndPoint string
	HttpPort string
	ScriptPath string
)

const (
	ManifestFormat = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>items</key>
	<array>
		<dict>
			<key>assets</key>
			<array>
				<dict>
					<key>kind</key>
					<string>software-package</string>
					<key>url</key>
					<string>%v</string>
				</dict>
			</array>
			<key>metadata</key>
			<dict>
				<key>bundle-identifier</key>
				<string>%v</string>
				<key>bundle-version</key>
				<string>1.0</string>
				<key>kind</key>
				<string>software</string>
				<key>subtitle</key>
				<string>%v</string>
				<key>title</key>
				<string>%v</string>
			</dict>
		</dict>
	</array>
</dict>
</plist>
`
)

func prepareViper(prefix string) {
	viper.AutomaticEnv()
	viper.SetEnvPrefix(prefix)

	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/" + prefix)
	viper.AddConfigPath("$HOME/." + prefix)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml") // because there is no file extension in a stream of bytes

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("fatal error config file: %s \n", err)
	}
}

func init() {
	prepareViper("")

	HttpEndPoint = viper.GetString("web_endpoint")
	WsEndPoint = viper.GetString("ws_endpoint")
	HttpPort = viper.GetString("web_port")
	ScriptPath = viper.GetString("script_path")
}

func ParseFlags(cmd *cobra.Command) {
	portVal := cmd.Flag("port").Value
	if portVal != nil {
		HttpPort = portVal.String()
	}
}
