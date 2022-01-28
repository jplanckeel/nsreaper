package internal

import (
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func DelNs(nameNamespace string) error {

	clientset, ctx, err := clientk8s()
	if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error setting k8s client %s\n", statusError.ErrStatus.Message)
	} else if err != nil {
		return err
	}
	err = clientset.CoreV1().Namespaces().Delete(ctx, nameNamespace, metav1.DeleteOptions{})
	if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error Deleting Namespace %s error: %s\n", nameNamespace, statusError.ErrStatus.Message)
	} else if err != nil {
		return err
	}
	return err
}
