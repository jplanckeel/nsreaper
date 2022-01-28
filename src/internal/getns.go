package internal

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNs(CleanNamespace string) ([]v1.Namespace, error) {

	clientset, ctx, err := clientk8s()

	if CleanNamespace == "" {
		ns, err := clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Namespace not found\n")
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error listing Namespace %s\n", statusError.ErrStatus.Message)
		} else if err != nil {
			return nil, err
		} else {
			return ns.Items, nil
		}
	} else if CleanNamespace != "" {
		ns, err := clientset.CoreV1().Namespaces().Get(ctx, CleanNamespace, metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Namespace not found\n")
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error listing Namespace %s\n", statusError.ErrStatus.Message)
		} else if err != nil {
			return nil, err
		} else {
			return []v1.Namespace{*ns}, nil
		}
	} else {
		return nil, err
	}
	return nil, err
}
