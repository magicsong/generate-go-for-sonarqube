package generate_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGenerate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Generate Sonar API")
}

func TestGetCurrentRepo(t *testing.T) {

	// g := gomega.NewGomegaWithT(t)
	// str, err := GetCurrentRepo("sonargo", "/home/magicsong/go/src/github.com/magicsong/generate-go-for-sonarqube")
	// g.Expect(err).NotTo(gomega.HaveOccurred())
	// g.Expect(str).To(gomega.Equal("github.com/magicsong/generate-go-for-sonarqube/sonargo"))
	// str, err = GetCurrentRepo("sonargo", "/home/go/src/github.com/magicsong/generate-go-for-sonarqube")
	// g.Expect(err).To(gomega.HaveOccurred())
	// g.Expect(str).To(gomega.Equal(""))
}
