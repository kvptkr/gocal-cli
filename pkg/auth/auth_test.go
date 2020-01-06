package auth_test

import (
	"github.com/kvptkr/gocal-cli/pkg/auth"
	. "github.com/onsi/gomega"
	"testing"
)

func testAuth(t *testing.T) {
	g := NewGomegaWithT(t)
	_, err := auth.GetClient()
	g.Expect(err).To(BeNil())
}
