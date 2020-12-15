package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/suryakencana007/mimir"
	"{{cookiecutter.repo_path}}/{{cookiecutter.repo_name}}/config"
	"{{cookiecutter.repo_path}}/{{cookiecutter.repo_name}}/rest"
)

var (
	flags   = flag.NewFlagSet("{{cookiecutter.repo_name}}", flag.ExitOnError)
	help    = flags.Bool("h", false, "print help")
	version = flags.Bool("version", false, "print version")
)

func main() {
	cfg := &config.Config{}
	if err := mimir.Config(mimir.ConfigOpts{
		Config:   cfg,
		Filename: "app.config",
		Paths:    []string{"./config"},
	}, func(v *viper.Viper) error {
		_ = v.BindEnv("db.dsn_secondary")
		return v.BindEnv("db.dsn_main")
	}); err != nil {
		fmt.Printf("%v ", err)
		return
	}

	logger := mimir.With(mimir.Field(cfg.App.Name, cfg.App.Version))

	flags.Usage = usage
	if err := flags.Parse(os.Args[1:]); err != nil {
		logger.Errorf("%v", err)
		return
	}

	if *version {
		logger.Infof("%s:%s", cfg.App.Name, cfg.App.Version)
		return
	}

	args := flags.Args()
	if len(args) == 0 || *help {
		flags.Usage()
		return
	}

	switch args[0] {
	case "rest":
		if err := rest.Application(cfg, logger); err != nil {
			logger.Errorf("%v", err)
		}
		return
	}
}

func usage() {
	fmt.Println(usagePrefix)
	flags.PrintDefaults()
}

var (
	usagePrefix = `Usage: {{cookiecutter.repo_name}} [OPTIONS] COMMAND

Examples:
	{{cookiecutter.repo_name}} rest
	{{cookiecutter.repo_name}} event-store
	{{cookiecutter.repo_name}} dispatcher

Options:`
)
