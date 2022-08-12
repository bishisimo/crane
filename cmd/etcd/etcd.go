/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package etcd

import (
	"crane/app/etcd"
	"github.com/spf13/cobra"
)

var EtcdCmd = etcdCmd

// etcdCmd represents the etcd command
var etcdCmd = &cobra.Command{
	Use:   "etcd",
	Short: "管理etcd",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}
var (
	etcdEntity = new(etcd.Etcd)
)

func init() {
	etcdCmd.Flags().StringVarP(&etcdEntity.Endpoints, "endpoints", "e", "https://127.0.0.1:2379", "etcd接入点url")
	etcdCmd.Flags().StringVar(&etcdEntity.CaPath, "ca", "/etc/kubernetes/pki/etcd/ca.crt", "etcd ca证书路径")
	etcdCmd.Flags().StringVar(&etcdEntity.CertPath, "cert", "/etc/kubernetes/pki/etcd/healthcheck-client.crt", "etcd 健康检查证书路径")
	etcdCmd.Flags().StringVar(&etcdEntity.KeyPath, "key", "/etc/kubernetes/pki/etcd/healthcheck-client.key", "etcd 健康检查密钥路径")
}
