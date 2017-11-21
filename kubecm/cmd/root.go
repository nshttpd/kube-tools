package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/spf13/cobra"
)

const (
	DEFAULT_KUBE_CONFIG = ".kube/config"
)

var (
	Cluster   string
	Namespace string
	logLevel  string
	kubeConf  string
)

var RootCmd = &cobra.Command{
	Use:   "kubecm",
	Short: "kubecm CLI tool for ConfigMAp manipulation",
	Long:  `A basic CLI tool for CRUD like operations on ConfigMaps`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		l, err := log.ParseLevel(logLevel)
		if err != nil {
			fmt.Printf("error setting log level : %v", err)
			os.Exit(1)
		}
		log.SetLevel(l)

		if err != nil {
			log.Error("error creating kube client")
			log.Error(err)
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	kc := fmt.Sprintf("%s/%s", os.Getenv("HOME"), DEFAULT_KUBE_CONFIG)

	RootCmd.PersistentFlags().StringVarP(&Cluster, "cluster", "c", "", "cluster for secret manipulation")
	RootCmd.PersistentFlags().StringVarP(&Namespace, "namespace", "n", "default", "namespace for secret manipulation")
	RootCmd.Flags().StringVar(&logLevel, "loglevel", "info", "log level")
	RootCmd.Flags().StringVar(&kubeConf, "kube-conf", kc, "kubectl config file to use")
}

func loadClient() *kubernetes.Clientset {
	var cs *kubernetes.Clientset

	lr := &clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConf}
	co := &clientcmd.ConfigOverrides{}
	if Cluster != "" {
		co.CurrentContext = Cluster
	}

	kc := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(lr, co)

	// grab the raw config to make sure the selected cluster/context is valid.
	rc, err := kc.RawConfig()
	if err == nil {
		if Cluster != "" {
			if _, ok := rc.Contexts[Cluster]; !ok {
				var ctx []string
				for k := range rc.Contexts {
					ctx = append(ctx, k)
				}
				err = fmt.Errorf("invalid cluster of '%s' specified, valid ones are %s", Cluster, ctx)
			}
		}
		if err == nil {
			cc, err := kc.ClientConfig()
			if err == nil {
				cs, err = kubernetes.NewForConfig(cc)
			}
		}
	}

	log.WithField("context", rc.CurrentContext).Debug("context set")

	if err != nil {
		log.Error("error creating kubernetes client")
		log.Error(err)
		os.Exit(1)
	}

	return cs
}
