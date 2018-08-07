package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Details of knative resources",
}

var describeServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Knative service details",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		output, err := describeService(args)
		if err != nil {
			log.Errorln(err)
		}
		fmt.Println(string(output))
	},
}

var describeConfigurationCmd = &cobra.Command{
	Use:   "configuration",
	Short: "Knative service configuration details",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		output, err := describeConfiguration(args)
		if err != nil {
			log.Errorln(err)
		}
		fmt.Println(string(output))
	},
}

var describeRevisionCmd = &cobra.Command{
	Use:   "revision",
	Short: "Knative revision details",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		output, err := describeRevision(args)
		if err != nil {
			log.Errorln(err)
		}
		fmt.Println(string(output))
	},
}

var describeRouteCmd = &cobra.Command{
	Use:   "route",
	Short: "Knative service route details",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		output, err := describeRoute(args)
		if err != nil {
			log.Errorln(err)
		}
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(describeCmd)
	describeCmd.AddCommand(describeServiceCmd)
	describeCmd.AddCommand(describeConfigurationCmd)
	describeCmd.AddCommand(describeRevisionCmd)
	describeCmd.AddCommand(describeRouteCmd)
}

func describeService(args []string) ([]byte, error) {
	service, err := serving.ServingV1alpha1().Services(namespace).Get(args[0], metav1.GetOptions{})
	if err != nil {
		return []byte{}, err
	}
	if output == "yaml" {
		return yaml.Marshal(service.Status)
	}
	return json.MarshalIndent(service.Status, "", "	")
}

func describeConfiguration(args []string) ([]byte, error) {
	configuration, err := serving.ServingV1alpha1().Configurations(namespace).Get(args[0], metav1.GetOptions{})
	if err != nil {
		return []byte{}, err
	}
	if output == "yaml" {
		return yaml.Marshal(configuration.Status)
	}
	return json.MarshalIndent(configuration.Status, "", "	")
}

func describeRevision(args []string) ([]byte, error) {
	revisions, err := serving.ServingV1alpha1().Revisions(namespace).Get(args[0], metav1.GetOptions{})
	if err != nil {
		return []byte{}, err
	}
	if output == "yaml" {
		return yaml.Marshal(revisions.Status)
	}
	return json.MarshalIndent(revisions.Status, "", "	")
}

func describeRoute(args []string) ([]byte, error) {
	routes, err := serving.ServingV1alpha1().Routes(namespace).Get(args[0], metav1.GetOptions{})
	if err != nil {
		return []byte{}, err
	}
	if output == "yaml" {
		return yaml.Marshal(routes.Status)
	}
	return json.MarshalIndent(routes.Status, "", "	")
}
