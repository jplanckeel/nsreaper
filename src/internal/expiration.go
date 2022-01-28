package internal

import (
	"strconv"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Expiration(annotationTtl string, nsCreationTime metav1.Time, argCleanTtl uint) (bool, error) {
	var ttl uint = 10

	if annotationTtl != "" {
		ttlUint, err := strconv.ParseUint(annotationTtl, 10, 0)
		ttl = uint(ttlUint)
		if err != nil {
			return false, err
		}
	} else if argCleanTtl != 0 {
		ttl = argCleanTtl
	}

	today := time.Now()
	pastWeek := nsCreationTime.AddDate(0, 0, int(ttl))

	return today.After(pastWeek), nil

}
