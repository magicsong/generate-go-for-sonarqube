package generate_test

import (
	"os"

	. "github.com/magicsong/generate-go-for-sonarqube/pkg/generate"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetCurrentRepo", func() {
	gopath := os.Getenv("GOPATH")
	BeforeEach(func() {
		if gopath == "" {
			_, err := GetCurrentRepo("sonargo", "")
			Expect(err).To(HaveOccurred())
		}
	})
	Describe("Test Error Case", func() {
		It("Error should be not nil", func() {
			_, err := GetCurrentRepo("sonargo", "/whatever/src/github.com/magicsong/generate-go-for-sonarqube")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("current location is not in $GOPATH/src"))
		})
	})
	Describe("Test Right Case", func() {
		It("Error should be nil", func() {
			str, err := GetCurrentRepo("sonargo", gopath+"/src/github.com/magicsong/generate-go-for-sonarqube")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(str).To(Equal("github.com/magicsong/generate-go-for-sonarqube/sonargo"))
		})
	})
})
