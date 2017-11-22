package actions

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
	namespace string
	logLevel  string
	kubeConf  string
	keyName   string
	keyValue  string
)

func init() {
	kc := fmt.Sprintf("%s/%s", os.Getenv("HOME"), DEFAULT_KUBE_CONFIG)

	RootCmd.PersistentFlags().StringVarP(&Cluster, "cluster", "c", "", "cluster for secret manipulation")
	RootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "default", "namespace for secret manipulation")
	RootCmd.Flags().StringVar(&logLevel, "loglevel", "info", "log level")
	RootCmd.Flags().StringVar(&kubeConf, "kube-conf", kc, "kubectl config file to use")
}

var RootCmd = &cobra.Command{
	Use:   "kubetool",
	Short: "kubetool -  CLI tool for kubernetes data object manipulation",
	Long: `A basic CLI tool for CRUD like operations on Secrets and ConfigMap
objects in a Kubernetes cluster`,
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
