package length_test

import (
	"github.com/bmaximilian/gutils/pkg/calculate/scale/length"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"
)

var _ = Describe("ConvertForScale", func() {
	Describe("Converts the units", func () {
		Context("When passing a bigger source than destination unit", func() {
			It("Should convert km to dm", func() {
				converted, err := length.ConvertForScale(1.00, "km", 1.00, "dm")
				Expect(converted).To(Equal(100.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(2.00, "km", 1.00, "dm")
				Expect(converted).To(Equal(200.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-10.00, "km", 1.00, "dm")
				Expect(converted).To(Equal(-1000.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(12.3456, "km", 1.00, "dm")
				Expect(math.Round(converted)).To(Equal(math.Round(1234.56)))
				Expect(err).To(BeNil())
			})
			It("Should convert km to m", func() {
				converted, err := length.ConvertForScale(1.00, "km", 1.00, "m")
				Expect(converted).To(Equal(1000.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(2.00, "km", 1.00, "m")
				Expect(converted).To(Equal(2000.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-10.00, "km", 1.00, "m")
				Expect(converted).To(Equal(-10000.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(12.3456, "km", 1.00, "m")
				Expect(math.Round(converted)).To(Equal(math.Round(12345.6)))
				Expect(err).To(BeNil())
			})
			It("Should convert km to cm", func() {
				converted, err := length.ConvertForScale(1.00, "km", 1.00, "cm")
				Expect(converted).To(Equal(100000.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(2.00, "km", 1.00, "cm")
				Expect(converted).To(Equal(200000.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-10.00, "km", 1.00, "cm")
				Expect(converted).To(Equal(-1000000.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(12.345678, "km", 1.00, "cm")
				Expect(math.Round(converted)).To(Equal(math.Round(1234567.8)))
				Expect(err).To(BeNil())
			})
			It("Should convert km to mm", func() {
				converted, err := length.ConvertForScale(1.00, "km", 1.00, "mm")
				Expect(converted).To(Equal(1000000.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(2.00, "km", 1.00, "mm")
				Expect(converted).To(Equal(2000000.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-10.00, "km", 1.00, "mm")
				Expect(converted).To(Equal(-10000000.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(12.3456789, "km", 1.00, "mm")
				Expect(math.Round(converted)).To(Equal(math.Round(12345678.9)))
				Expect(err).To(BeNil())
			})

			It("Should convert dm to m", func() {
				converted, err := length.ConvertForScale(1.00, "dm", 1.00, "m")
				Expect(converted).To(Equal(10.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(2.00, "dm", 1.00, "m")
				Expect(converted).To(Equal(20.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-10.00, "dm", 1.00, "m")
				Expect(converted).To(Equal(-100.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(12.3456789, "dm", 1.00, "m")
				Expect(math.Round(converted)).To(Equal(math.Round(123.456789)))
				Expect(err).To(BeNil())
			})
			It("Should convert dm to cm", func() {
				converted, err := length.ConvertForScale(1.00, "dm", 1.00, "cm")
				Expect(converted).To(Equal(1000.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(2.00, "dm", 1.00, "cm")
				Expect(converted).To(Equal(2000.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-10.00, "dm", 1.00, "cm")
				Expect(converted).To(Equal(-10000.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(12.3456789, "dm", 1.00, "cm")
				Expect(math.Round(converted)).To(Equal(math.Round(12345.6789)))
				Expect(err).To(BeNil())
			})
			It("Should convert dm to mm", func() {
				converted, err := length.ConvertForScale(1.00, "dm", 1.00, "mm")
				Expect(converted).To(Equal(10000.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(2.00, "dm", 1.00, "mm")
				Expect(converted).To(Equal(20000.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-10.00, "dm", 1.00, "mm")
				Expect(converted).To(Equal(-100000.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(12.3456789, "dm", 1.00, "mm")
				Expect(math.Round(converted)).To(Equal(math.Round(123456.789)))
				Expect(err).To(BeNil())
			})

			It("Should convert m to cm", func() {
				converted, err := length.ConvertForScale(1.00, "m", 1.00, "cm")
				Expect(converted).To(Equal(100.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(2.00, "m", 1.00, "cm")
				Expect(converted).To(Equal(200.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-10.00, "m", 1.00, "cm")
				Expect(converted).To(Equal(-1000.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(12.3456789, "m", 1.00, "cm")
				Expect(math.Round(converted)).To(Equal(math.Round(1234.56789)))
				Expect(err).To(BeNil())
			})
			It("Should convert m to mm", func() {
				converted, err := length.ConvertForScale(1.00, "m", 1.00, "mm")
				Expect(converted).To(Equal(1000.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(2.00, "m", 1.00, "mm")
				Expect(converted).To(Equal(2000.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-10.00, "m", 1.00, "mm")
				Expect(converted).To(Equal(-10000.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(12.3456789, "m", 1.00, "mm")
				Expect(math.Round(converted)).To(Equal(math.Round(12345.6789)))
				Expect(err).To(BeNil())
			})

			It("Should convert cm to mm", func() {
				converted, err := length.ConvertForScale(1.00, "cm", 1.00, "mm")
				Expect(converted).To(Equal(10.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(2.00, "cm", 1.00, "mm")
				Expect(converted).To(Equal(20.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-10.00, "cm", 1.00, "mm")
				Expect(converted).To(Equal(-100.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(12.3456789, "cm", 1.00, "mm")
				Expect(math.Round(converted)).To(Equal(math.Round(123.456789)))
				Expect(err).To(BeNil())
			})
		})

		Context("When passing a smaller source than destination unit", func() {
			It("Should convert mm to cm", func() {
				converted, err := length.ConvertForScale(10.00, "mm", 1.00, "cm")
				Expect(converted).To(Equal(1.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(20.00, "mm", 1.00, "cm")
				Expect(converted).To(Equal(2.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-100.00, "mm", 1.00, "cm")
				Expect(converted).To(Equal(-10.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(123.456789, "mm", 1.00, "cm")
				Expect(math.Round(converted)).To(Equal(math.Round(12.3456789)))
				Expect(err).To(BeNil())
			})
			It("Should convert mm to m", func() {
				converted, err := length.ConvertForScale(1000.00, "mm", 1.00, "m")
				Expect(converted).To(Equal(1.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(2000.00, "mm", 1.00, "m")
				Expect(converted).To(Equal(2.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-10000.00, "mm", 1.00, "m")
				Expect(converted).To(Equal(-10.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(12345.6789, "mm", 1.00, "m")
				Expect(math.Round(converted)).To(Equal(math.Round(12.3456789)))
				Expect(err).To(BeNil())
			})
			It("Should convert mm to dm", func() {
				converted, err := length.ConvertForScale(10000.00, "mm", 1.00, "dm")
				Expect(converted).To(Equal(1.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(20000.00, "mm", 1.00, "dm")
				Expect(converted).To(Equal(2.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-100000.00, "mm", 1.00, "dm")
				Expect(converted).To(Equal(-10.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(123456.789, "mm", 1.00, "dm")
				Expect(math.Round(converted)).To(Equal(math.Round(12.3456789)))
				Expect(err).To(BeNil())
			})
			It("Should convert mm to km", func() {
				converted, err := length.ConvertForScale(1000000.00, "mm", 1.00, "km")
				Expect(converted).To(Equal(1.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(2000000.00, "mm", 1.00, "km")
				Expect(converted).To(Equal(2.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-10000000.00, "mm", 1.00, "km")
				Expect(converted).To(Equal(-10.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(12345678.9, "mm", 1.00, "km")
				Expect(math.Round(converted)).To(Equal(math.Round(12.3456789)))
				Expect(err).To(BeNil())
			})

			It("Should convert cm to m", func() {
				converted, err := length.ConvertForScale(100.00, "cm", 1.00, "m")
				Expect(converted).To(Equal(1.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(200.00, "cm", 1.00, "m")
				Expect(converted).To(Equal(2.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-1000.00, "cm", 1.00, "m")
				Expect(converted).To(Equal(-10.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(1234.56789, "cm", 1.00, "m")
				Expect(math.Round(converted)).To(Equal(math.Round(12.3456789)))
				Expect(err).To(BeNil())
			})
			It("Should convert cm to dm", func() {
				converted, err := length.ConvertForScale(1000.00, "cm", 1.00, "dm")
				Expect(converted).To(Equal(1.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(2000.00, "cm", 1.00, "dm")
				Expect(converted).To(Equal(2.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-10000.00, "cm", 1.00, "dm")
				Expect(converted).To(Equal(-10.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(12345.6789, "cm", 1.00, "dm")
				Expect(math.Round(converted)).To(Equal(math.Round(12.3456789)))
				Expect(err).To(BeNil())
			})
			It("Should convert cm to km", func() {
				converted, err := length.ConvertForScale(100000.00, "cm", 1.00, "km")
				Expect(converted).To(Equal(1.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(200000.00, "cm", 1.00, "km")
				Expect(converted).To(Equal(2.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-1000000.00, "cm", 1.00, "km")
				Expect(converted).To(Equal(-10.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(1234567.89, "cm", 1.00, "km")
				Expect(math.Round(converted)).To(Equal(math.Round(12.3456789)))
				Expect(err).To(BeNil())
			})

			It("Should convert m to dm", func() {
				converted, err := length.ConvertForScale(10.00, "m", 1.00, "dm")
				Expect(converted).To(Equal(1.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(20.00, "m", 1.00, "dm")
				Expect(converted).To(Equal(2.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-100.00, "m", 1.00, "dm")
				Expect(converted).To(Equal(-10.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(123.456789, "m", 1.00, "dm")
				Expect(math.Round(converted)).To(Equal(math.Round(12.3456789)))
				Expect(err).To(BeNil())
			})
			It("Should convert m to km", func() {
				converted, err := length.ConvertForScale(1000.00, "m", 1.00, "km")
				Expect(converted).To(Equal(1.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(2000.00, "m", 1.00, "km")
				Expect(converted).To(Equal(2.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-10000.00, "m", 1.00, "km")
				Expect(converted).To(Equal(-10.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(12345.6789, "m", 1.00, "km")
				Expect(math.Round(converted)).To(Equal(math.Round(12.3456789)))
				Expect(err).To(BeNil())
			})

			It("Should convert dm to km", func() {
				converted, err := length.ConvertForScale(100.00, "dm", 1.00, "km")
				Expect(converted).To(Equal(1.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(200.00, "dm", 1.00, "km")
				Expect(converted).To(Equal(2.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-1000.00, "dm", 1.00, "km")
				Expect(converted).To(Equal(-10.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(1234.56789, "dm", 1.00, "km")
				Expect(math.Round(converted)).To(Equal(math.Round(12.3456789)))
				Expect(err).To(BeNil())
			})
		})

		Context("When passing a equal source and destination unit", func() {
			It("Should convert km to km", func() {
				converted, err := length.ConvertForScale(1.00, "km", 1.00, "km")
				Expect(converted).To(Equal(1.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(2.00, "km", 1.00, "km")
				Expect(converted).To(Equal(2.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-10.00, "km", 1.00, "km")
				Expect(converted).To(Equal(-10.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(12.3456, "km", 1.00, "km")
				Expect(math.Round(converted)).To(Equal(math.Round(12.3456)))
				Expect(err).To(BeNil())
			})
			It("Should convert dm to dm", func() {
				converted, err := length.ConvertForScale(1.00, "dm", 1.00, "dm")
				Expect(converted).To(Equal(1.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(2.00, "dm", 1.00, "dm")
				Expect(converted).To(Equal(2.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-10.00, "dm", 1.00, "dm")
				Expect(converted).To(Equal(-10.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(12.3456789, "dm", 1.00, "dm")
				Expect(math.Round(converted)).To(Equal(math.Round(12.3456789)))
				Expect(err).To(BeNil())
			})
			It("Should convert m to m", func() {
				converted, err := length.ConvertForScale(1.00, "m", 1.00, "m")
				Expect(converted).To(Equal(1.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(2.00, "m", 1.00, "m")
				Expect(converted).To(Equal(2.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-10.00, "m", 1.00, "m")
				Expect(converted).To(Equal(-10.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(12.3456789, "m", 1.00, "m")
				Expect(math.Round(converted)).To(Equal(math.Round(12.3456789)))
				Expect(err).To(BeNil())
			})
			It("Should convert cm to cm", func() {
				converted, err := length.ConvertForScale(1.00, "cm", 1.00, "cm")
				Expect(converted).To(Equal(1.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(2.00, "cm", 1.00, "cm")
				Expect(converted).To(Equal(2.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-10.00, "cm", 1.00, "cm")
				Expect(converted).To(Equal(-10.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(12.3456789, "cm", 1.00, "cm")
				Expect(math.Round(converted)).To(Equal(math.Round(12.3456789)))
				Expect(err).To(BeNil())
			})
			It("Should convert mm to mm", func() {
				converted, err := length.ConvertForScale(1.00, "mm", 1.00, "mm")
				Expect(converted).To(Equal(1.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(2.00, "mm", 1.00, "mm")
				Expect(converted).To(Equal(2.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(-10.00, "mm", 1.00, "mm")
				Expect(converted).To(Equal(-10.00))
				Expect(err).To(BeNil())

				converted, err = length.ConvertForScale(12.3456789, "mm", 1.00, "mm")
				Expect(math.Round(converted)).To(Equal(math.Round(12.3456789)))
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Scales the conversion", func () {
		Context("When passing a smaller source than destination unit", func() {
			It("Should convert mm to cm", func() {
				converted, err := length.ConvertForScale(10.00, "mm", 3.00, "cm")
				Expect(converted).To(Equal(3.00))
				Expect(err).To(BeNil())
			})
		})
		Context("When passing a bigger source than destination unit", func() {
			It("Should convert m to cm", func() {
				converted, err := length.ConvertForScale(3.50, "m", 4.00, "cm")
				Expect(converted).To(Equal(1400.00))
				Expect(err).To(BeNil())
			})
		})
		Context("When passing a equal source and destination unit", func() {
			It("Should convert m to m", func() {
				converted, err := length.ConvertForScale(3.50, "m", 4.00, "m")
				Expect(converted).To(Equal(14.00))
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Returns an Error", func() {
		Context("When passing a not supported unit", func() {
			It("Should throw when passing source unit 'xm'", func() {
				converted, err := length.ConvertForScale(200.00,"xm", 0.5, "m")
				Expect(converted).To(Equal(200.00))
				Expect(err).NotTo(BeNil())
			})
			It("Should throw when passing destination unit 'xm'", func() {
				converted, err := length.ConvertForScale(200.00,"m", 0.5, "xm")
				Expect(converted).To(Equal(200.00))
				Expect(err).NotTo(BeNil())
			})
			It("Should throw when passing source unit 'ym' and destination unit 'xm'", func() {
				converted, err := length.ConvertForScale(200.00,"ym", 0.5, "xm")
				Expect(converted).To(Equal(200.00))
				Expect(err).NotTo(BeNil())
			})
			It("Should not throw when passing valid source and destination units", func() {
				converted, err := length.ConvertForScale(200.00,"km", 0.5, "cm")
				Expect(converted).To(BeAssignableToTypeOf(0.00))
				Expect(err).To(BeNil())
			})
		})
	})
})
