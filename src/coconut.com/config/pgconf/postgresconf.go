package pgconf

import (
	"github.com/jackc/pgx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

const (
	host        = "deploygate_host"
	defaultHost = "localhost"

	database        = "deploygate_database"
	defaultDatabase = "deploygate"

	user        = "cc_user"
	defaultUser = "cc"

	password        = "cc_password"
	defaultPassword = "cc"
)

func init() {
	viper.SetDefault(host, defaultHost)
	viper.SetDefault(database, defaultDatabase)
	viper.SetDefault(user, defaultUser)
	viper.SetDefault(password, defaultPassword)
}


// SetFlags adds a local postgres_uri flag to cmd.
func SetFlags(cmd *cobra.Command) {
	setDbFlag(cmd, host, defaultHost, "pg host")
	setDbFlag(cmd, database, defaultDatabase, "pg database")
	setDbFlag(cmd, user, defaultUser, "pg user")
	setDbFlag(cmd, password, defaultPassword, "pg password")
}

func setDbFlag(cmd *cobra.Command, key, value, usage string) {
	f := cmd.Flags()
	f.String(key, value, usage)
	localName := cmd.Name() + "_" + key
	if err := viper.BindPFlag(localName, f.Lookup(key)); err != nil {
		log.Fatalf("unable to bind key: %v", key)
	}
}

// Config returns the postgres config object
func Config(cmd *cobra.Command) (config pgx.ConnPoolConfig) {
	var err error
	var vHost, vDatabase, vUser, vPassword = host, database, user, password
	if cmd != nil {
		vHost = cmd.Name() + "_" + host
		vDatabase = cmd.Name() + "_" + database
		vUser = cmd.Name() + "_" + user
		vPassword = cmd.Name() + "_" + password
	}
	config.ConnConfig, err = pgx.ParseEnvLibpq()
	if err != nil {
		return config
	}
	config.Host = viper.GetString(vHost)
	config.Database = viper.GetString(vDatabase)
	config.User = viper.GetString(vUser)
	config.Password = viper.GetString(vPassword)

	return
}
