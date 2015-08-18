package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

var _ = Describe("Angular", func() {
	var page *agouti.Page

	It("should redirect / to #/", func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())
		Expect(page.Navigate("http://localhost:8080")).To(Succeed())
		Expect(page).To(HaveURL("http://localhost:8080/#/"))
	})

	Describe("Phone list view", func() {
		BeforeEach(func() {
			var err error
			page, err = agoutiDriver.NewPage()
			Expect(err).NotTo(HaveOccurred())
			Expect(page.Navigate("http://localhost:8080/#/phones")).To(Succeed())
		})

		AfterEach(func() {
			Expect(page.Destroy()).To(Succeed())
		})

		It("should filter the phone list as a user types into the search box", func() {

			phonesInList := page.Find("ul.phones").All("li")
			query := page.FindByName("query")

			Eventually(phonesInList).Should(HaveCount(20))
			Expect(query.Fill("nexus")).To(Succeed())
			Eventually(phonesInList).Should(HaveCount(1))
			Expect(query.Fill("motorola")).To(Succeed()) // Note: Fill clears automatically
			Eventually(phonesInList).Should(HaveCount(8))
		})

		It("should be possible to control phone order via the drop down select box", func() {

			phoneNameColumn := page.Find("ul.phones").AllByName("phone-name")
			query := page.FindByName("query")
			selected := page.Find("select")

			Expect(query.Fill("tablet")).To(Succeed()) //let's narrow the dataset to make the test assertions shorter
			Eventually(phoneNameColumn.At(0)).Should(HaveText("Motorola XOOM\u2122 with Wi-Fi"))
			Eventually(phoneNameColumn.At(1)).Should(HaveText("MOTOROLA XOOM\u2122"))

			Expect(selected.Select("Alphabetical")).To(Succeed())
			Eventually(phoneNameColumn.At(0)).Should(HaveText("MOTOROLA XOOM\u2122"))
			Eventually(phoneNameColumn.At(1)).Should(HaveText("Motorola XOOM\u2122 with Wi-Fi"))
		})

		It("should render phone specific links", func() {

			query := page.FindByName("query")
			phoneNameColumn := page.Find("ul.phones").AllByName("phone-name")

			Expect(query.Fill("nexus")).To(Succeed())
			Expect(phoneNameColumn.At(0).Click()).To(Succeed())
			Eventually(page).Should(HaveURL("http://localhost:8080/#/phones/nexus-s"))
		})
	})

	Describe("Phone detail view", func() {

		BeforeEach(func() {
			var err error
			page, err = agoutiDriver.NewPage()
			Expect(err).NotTo(HaveOccurred())
			Expect(page.Navigate("http://localhost:8080/#/phones/nexus-s")).To(Succeed())
		})

		AfterEach(func() {
			Expect(page.Destroy()).To(Succeed())
		})

		It("should display nexus-s page", func() {
			Eventually(page.Find("h1")).Should(HaveText("Nexus S"))
		})

		It("should display the first phone image as the main phone image", func() {
			image := page.Find("img.phone")
			Eventually(image).Should(HaveAttribute("ng-src", "static/img/phones/nexus-s.0.jpg"))
		})

		It("should swap main image if a thumbnail image is clicked on", func() {
			images := page.All("ul.phone-thumbs").AllByName("image-see")
			Expect(images.At(2).Click()).To(Succeed())
			Eventually(page.Find("img.phone")).Should(HaveAttribute("ng-src", "static/img/phones/nexus-s.2.jpg"))
			Expect(images.At(0).Click()).To(Succeed())
			Eventually(page.Find("img.phone")).Should(HaveAttribute("ng-src", "static/img/phones/nexus-s.0.jpg"))
		})
	})
})
