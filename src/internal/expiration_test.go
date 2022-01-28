package internal

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestExpirationFalse(t *testing.T) {

	var nowTime metav1.Time = metav1.Now()

	testCases := []struct {
		annotationTtl  string
		nsCreationTime metav1.Time
		argCleanTtl    uint
	}{
		{
			annotationTtl:  "",
			nsCreationTime: nowTime,
			argCleanTtl:    10,
		},
	}

	for _, testCase := range testCases {
		e, _ := Expiration(testCase.annotationTtl, testCase.nsCreationTime, testCase.argCleanTtl)
		assert.False(t, e, "The result is not False")
	}
}

func TestExpirationTrue(t *testing.T) {

	var creationTime metav1.Time = metav1.Date(2021, 02, 12, 22, 02, 35, 300, time.UTC)

	testCases := []struct {
		annotationTtl  string
		nsCreationTime metav1.Time
		argCleanTtl    uint
	}{
		// AnnotationTtl not set
		{
			annotationTtl:  "",
			nsCreationTime: creationTime,
			argCleanTtl:    10,
		},
		// AnnotationTtl set
		{
			annotationTtl:  "15",
			nsCreationTime: creationTime,
			argCleanTtl:    10,
		},
		// argCleanTtl not set
		{
			annotationTtl:  "15",
			nsCreationTime: creationTime,
		},
	}

	for _, testCase := range testCases {
		e, _ := Expiration(testCase.annotationTtl, testCase.nsCreationTime, testCase.argCleanTtl)
		assert.True(t, e, "The result is not True")
	}
}

func TestExpirationAnnotationError(t *testing.T) {

	var creationTime metav1.Time= metav1.Date(2021, 02, 12, 22, 02, 35, 300, time.UTC)

	testCases := []struct {
		annotationTtl  string
		nsCreationTime metav1.Time
		argCleanTtl    uint
	}{
		// AnnotationTtl is Integrer
		{
			annotationTtl:  "abc",
			nsCreationTime: creationTime,
		},
	}

	for _, testCase := range testCases {
		_, err := Expiration(testCase.annotationTtl, testCase.nsCreationTime, testCase.argCleanTtl)
		assert.Error(t, err, "The result is not False")
	}
}
