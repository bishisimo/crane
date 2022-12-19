package init

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"path"
)

func initValues() {
	err := viper.BindEnv("HOME")
	if err != nil {
		log.Fatal().Msg("fail to get HOME")
	}
	viper.SetDefault("crane.mainDir", path.Join(viper.GetString("HOME"), ".crane"))
	viper.SetDefault("crane.kubeSpace", path.Join(viper.GetString("crane.mainDir"), "kube"))
	viper.SetDefault("sys.sshKeyDir", path.Join(viper.GetString("HOME"), ".ssh"))
}
