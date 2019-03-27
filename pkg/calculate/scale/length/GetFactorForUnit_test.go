package length_test

import (
	"github.com/bmaximilian/gutils/pkg/calculate/scale/length"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetFactorForUnit", func() {
	Describe("Returns the factor to normalize", func () {
		Context("when passing a valid factor", func () {
			It("Should return -1000 for mm", func() {
				unit, err := length.GetFactorForUnit("mm")
				Expect(unit).To(Equal(-1000))
				Expect(err).To(BeNil())
			})

			It("Should return -100 for cm", func() {
				unit, err := length.GetFactorForUnit("cm")
				Expect(unit).To(Equal(-100))
				Expect(err).To(BeNil())
			})

			It("Should return 1 for m", func() {
				unit, err := length.GetFactorForUnit("m")
				Expect(unit).To(Equal(1))
				Expect(err).To(BeNil())
			})

			It("Should return 10 for dm", func() {
				unit, err := length.GetFactorForUnit("dm")
				Expect(unit).To(Equal(10))
				Expect(err).To(BeNil())
			})

			It("Should return 1000 for km", func() {
				unit, err := length.GetFactorForUnit("km")
				Expect(unit).To(Equal(1000))
				Expect(err).To(BeNil())
			})
		})
	})
	Describe("Returns an Error", func() {
		Context("When passing a not supported unit", func() {
			It("Should throw when passing unit 'xm'", func() {
				unit, err := length.GetFactorForUnit("xm")
				Expect(unit).To(Equal(1))
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
