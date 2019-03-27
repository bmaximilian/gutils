package length_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLength(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Length Suite")
}
