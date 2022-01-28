package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/jplanckeel/nsreaper/internal"
	"k8s.io/apimachinery/pkg/api/errors"
)

var cleanNamespace string
var cleanDryrun bool
var cleanTtl uint

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean Preview Namespace on ",
	Run: func(cmd *cobra.Command, args []string) {

		clean()
	},
}

func clean() {

	ns, err := internal.GetNs(cleanNamespace)
	if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error in function get namespaces %s\n", statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	}

	for i := range ns {
		if ns[i].ObjectMeta.Annotations["nsreaper/type"] == "preview" {
			nsName := ns[i].ObjectMeta.Name
			nsCreationTime := ns[i].ObjectMeta.CreationTimestamp
			nsTtl := ns[i].ObjectMeta.Annotations["nsreaper/ttl"]
			nsRepo := ns[i].ObjectMeta.Annotations["nsreaper/repository"]
			nsPr := ns[i].ObjectMeta.Annotations["nsreaper/pull_request_id"]

			nsexpiration, err := internal.Expiration(nsTtl, nsCreationTime, cleanTtl)
			if statusError, isStatus := err.(*errors.StatusError); isStatus {
				fmt.Printf("Error in function Expiration %s\n", statusError.ErrStatus.Message)
			} else if err != nil {
				panic(err.Error())
			}

			if nsexpiration {
				fmt.Printf("Namespace: %s, Creationdate: %s, Repo %s, Pr:%s will be delete \n", nsName, nsCreationTime, nsRepo, nsPr)

				if !cleanDryrun {
					err := internal.DelNs(nsName)
					if errors.IsNotFound(err) {
						fmt.Printf("Namespace not found\n")
					} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
						fmt.Printf("Error deleting Namespace %s\n", statusError.ErrStatus.Message)
					} else if err != nil {
						panic(err.Error())
					}
				}

			}
		}
	}
}

func init() {
	rootCmd.AddCommand(cleanCmd)
	cleanCmd.PersistentFlags().StringVarP(&cleanNamespace, "namespace", "n", "", "Select Namespace for clean")
	cleanCmd.PersistentFlags().BoolVarP(&cleanDryrun, "dryrun", "d", false, "run clean namespace on dry-run mode")
	cleanCmd.PersistentFlags().UintVarP(&cleanTtl, "ttl", "t", 10, "run clean namespace on dry-run mode")
}
