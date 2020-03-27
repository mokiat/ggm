package sprec_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mokiat/gomath/sprec"
	. "github.com/mokiat/gomath/testing/sprectest"
)

var _ = Describe("Util", func() {

	Specify("Abs", func() {
		Expect(Abs(-0.1)).To(EqualFloat32(0.1))
		Expect(Abs(-13.57)).To(EqualFloat32(13.57))
		Expect(Abs(11.01)).To(EqualFloat32(11.01))
	})

	Specify("Max", func() {
		Expect(Max(1.0, 2.0)).To(EqualFloat32(2.0))
		Expect(Max(1.0, -1.0)).To(EqualFloat32(1.0))
		Expect(Max(5.0, 5.0)).To(EqualFloat32(5.0))
	})

	Specify("Min", func() {
		Expect(Min(1.0, 2.0)).To(EqualFloat32(1.0))
		Expect(Min(1.0, -1.0)).To(EqualFloat32(-1.0))
		Expect(Min(5.0, 5.0)).To(EqualFloat32(5.0))
	})

	Specify("Eq", func() {
		Expect(Eq(0.000001, 0.000001)).To(BeTrue())
		Expect(Eq(0.000001, 0.000002)).To(BeFalse())
		Expect(Eq(0.0000003, 0.0000002)).To(BeTrue()) // outside precision
	})

	Specify("EqEps", func() {
		Expect(EqEps(0.01, 0.01, 0.01)).To(BeTrue())
		Expect(EqEps(0.01, 0.02, 0.01)).To(BeFalse())
		Expect(EqEps(0.003, 0.002, 0.01)).To(BeTrue()) // outside precision
	})

	Specify("Sqrt", func() {
		Expect(Sqrt(17.64)).To(EqualFloat32(4.2))
	})

	Specify("Cos", func() {
		Expect(Cos(0.0)).To(EqualFloat32(1.0))
		Expect(Cos(Pi / 3)).To(EqualFloat32(0.5))
		Expect(Cos(Pi / 2)).To(EqualFloat32(0.0))
	})

	Specify("Sin", func() {
		Expect(Sin(0.0)).To(EqualFloat32(0.0))
		Expect(Sin(Pi / 6)).To(EqualFloat32(0.5))
		Expect(Sin(Pi / 2)).To(EqualFloat32(1.0))
	})

	Specify("Sign", func() {
		Expect(Sign(0.1)).To(EqualFloat32(1.0))
		Expect(Sign(-0.1)).To(EqualFloat32(-1.0))
	})

})
