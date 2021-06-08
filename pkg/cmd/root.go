package cmd

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/StevenLeRoux/dirt/pkg/dirt"
	mod "github.com/StevenLeRoux/dirt/pkg/mod"
	"github.com/gofrs/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	projectName = "dirt"
	seed        = "636afbf8-df68-41b2-adc2-7a8ecf06172c"
	conf        *mod.Config
	registry    *prometheus.Registry
)

func init() {

	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().Int32P("log-level", "l", 4, "set logging level between 0 and 5")
	RootCmd.PersistentFlags().StringP("config", "c", "config.yml", "set config file")
	RootCmd.PersistentFlags().BoolP("debug", "d", false, "set debug")
	RootCmd.PersistentFlags().BoolP("bootstrap", "b", false, "define as a bootstrap node which won't try to join any other node")

	if err := viper.BindPFlags(RootCmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}

	if err := viper.BindPFlags(RootCmd.Flags()); err != nil {
		log.Fatal(err)
	}

	registry = prometheus.NewRegistry()

}

func initConfig() {

	// Set viper defaults
	viper.SetDefault("log-level", 1)
	viper.SetDefault("bootstrap", false)
	viper.SetDefault("metrics.bind", "127.0.0.1")
	viper.SetDefault("metrics.port", 9199)
	viper.SetDefault("name", "local")
	viper.SetDefault("discovery.bind", "0.0.0.0")
	viper.SetDefault("discovery.port", 7946)
	viper.SetDefault("discovery.adv-address", "0.0.0.0")
	viper.SetDefault("discovery.adv-port", 7946)
	viper.SetDefault("server.bind", "0.0.0.0")
	viper.SetDefault("server.port", 2112)
	viper.SetDefault("group", "default")
	viper.SetDefault("peers", []string{"default"})
	viper.SetDefault("rack", "default")

	// Environment variables management
	viper.SetEnvPrefix(projectName)
	viper.SetEnvKeyReplacer(strings.NewReplacer("_", "."))
	viper.AutomaticEnv()

	// Set config search path
	viper.AddConfigPath("/etc/" + projectName + "/")
	viper.AddConfigPath("$HOME/." + projectName)
	viper.AddConfigPath(".")

	// Set a default ID for the current instance.
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	u, err := uuid.FromString(seed)
	if err != nil {
		log.Error(err)
	}

	id := uuid.NewV5(u, hostname)
	viper.SetDefault("id", id.String)

	// Load config
	viper.SetConfigName("config")
	if err := viper.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Panicf("Fatal error in config file: %v \n", err)
		}
	}

	// Load user defined config
	cfgFile := viper.GetString("config")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		err := viper.ReadInConfig()
		if err != nil {
			log.Panicf("Fatal error in config file: %v \n", err)
		}
	}

	if viper.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
		log.Info("Dirt: log level set to  " + string(log.GetLevel()))
	}

	if ok := viper.IsSet("join"); !ok {
		log.Fatal(errors.New("cannot find the bootstrap address to connect to the cluster (check 'join' key in your configuration"))
	}

	conf = &mod.Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

}

// RootCmd is the cli root command
var RootCmd = &cobra.Command{
	Use: projectName,
	Run: func(cmd *cobra.Command, arguments []string) {

		go func() {
			listen := fmt.Sprintf("%s:%d", conf.Metrics.Bind, conf.Metrics.Port)
			log.Infof("Start metrics endpint on on %s", listen)
			http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
			if err := http.ListenAndServe(listen, nil); err != nil && err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}()

		d, err := dirt.Create(registry, conf)
		if err != nil {
			log.Panicf("Dirt: fatal error while trying to start : %v \n", err)
		}
		d.Run()

		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 5 seconds.
		quit := make(chan os.Signal, 2)

		signal.Notify(quit, syscall.SIGTERM)
		signal.Notify(quit, syscall.SIGINT)

		<-quit

		d.Close()

	},
}
