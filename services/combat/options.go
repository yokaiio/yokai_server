package combat

import (
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func NewFlags() []cli.Flag {
	return []cli.Flag{
		// combat settings
		altsrc.NewBoolFlag(&cli.BoolFlag{Name: "debug", Usage: "debug mode"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "log_level", Usage: "log level"}),
		altsrc.NewIntFlag(&cli.IntFlag{Name: "combat_id", Usage: "combat server unique id(0 - 1024)"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "https_listen_addr", Usage: "https listen address"}),

		// rate limit
		altsrc.NewDurationFlag(&cli.DurationFlag{Name: "rate_limit_interval", Usage: "rpc server rate limit interval"}),
		altsrc.NewIntFlag(&cli.IntFlag{Name: "rate_limit_capacity", Usage: "rpc server rate limit capacity"}),

		// cert
		altsrc.NewStringFlag(&cli.StringFlag{Name: "cert_path_debug", Usage: "debug tls cert_pem path"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "key_path_debug", Usage: "debug tls server_key path"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "cert_path_release", Usage: "release tls cert_pem path"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "key_path_release", Usage: "release tls server_key path"}),

		// db
		altsrc.NewStringFlag(&cli.StringFlag{Name: "db_dsn", Usage: "db data source name"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "database", Usage: "database name"}),

		// micro service
		altsrc.NewStringFlag(&cli.StringFlag{Name: "registry_debug", Usage: "micro service registry in debug mode"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "registry_release", Usage: "micro service registry in release mode"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "broker_debug", Usage: "micro service broker in debug mode"}),
		altsrc.NewStringFlag(&cli.StringFlag{Name: "broker_release", Usage: "micro service broker in release mode"}),

		altsrc.NewStringFlag(&cli.StringFlag{Name: "config_file", Usage: "combat config path"}),
	}
}
