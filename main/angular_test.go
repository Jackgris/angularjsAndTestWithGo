package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

var _ = Describe("Angular", func() {
	var page *agouti.Page

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	It("should filter the phone list as a user types into the search box", func() {

		Expect(page.Navigate("http://localhost:8080")).To(Succeed())
		phonesInList := page.Find(".phones").All(".phone")
		query := page.Find(".query")

		Eventually(phonesInList).Should(HaveCount(3))
		Expect(query.Fill("nexus")).To(Succeed())
		Eventually(phonesInList).Should(HaveCount(1))
		Expect(query.Fill("motorola")).To(Succeed()) // Note: Fill clears automatically
		Eventually(phonesInList).Should(HaveCount(2))
	})
})
