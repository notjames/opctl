package boolean

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/data"
	"github.com/opspec-io/sdk-golang/data/coerce"
	"github.com/opspec-io/sdk-golang/model"
	"github.com/opspec-io/sdk-golang/op/interpreter/interpolater"
	"github.com/pkg/errors"
)

var _ = Context("Interpret", func() {
	Context("expression is boolean", func() {
		It("should return expected result", func() {
			/* arrange */
			providedExpression := true

			objectUnderTest := _interpreter{}

			/* act */
			actualBoolean, actualErr := objectUnderTest.Interpret(
				map[string]*model.Value{},
				providedExpression,
				new(data.FakeHandle),
			)

			/* assert */
			Expect(*actualBoolean).To(Equal(model.Value{Boolean: &providedExpression}))
			Expect(actualErr).To(BeNil())
		})
	})
	Context("expression is string", func() {
		It("should call interpolater.Interpolate w/ expected args", func() {
			/* arrange */
			providedScope := map[string]*model.Value{"dummyName": {}}
			providedExpression := "dummyExpression"
			providedOpRef := new(data.FakeHandle)

			fakeInterpolater := new(interpolater.Fake)
			// err to trigger immediate return
			fakeInterpolater.InterpolateReturns("", errors.New("dummyError"))

			objectUnderTest := _interpreter{
				interpolater: fakeInterpolater,
			}

			/* act */
			objectUnderTest.Interpret(
				providedScope,
				providedExpression,
				providedOpRef,
			)

			/* assert */
			actualExpression,
				actualScope,
				actualOpRef := fakeInterpolater.InterpolateArgsForCall(0)

			Expect(actualExpression).To(Equal(providedExpression))
			Expect(actualScope).To(Equal(providedScope))
			Expect(actualOpRef).To(Equal(providedOpRef))

		})
		Context("interpolater.Interpolate errs", func() {
			It("should return expected err", func() {
				/* arrange */
				fakeInterpolater := new(interpolater.Fake)
				interpolateErr := errors.New("dummyError")
				fakeInterpolater.InterpolateReturns("", interpolateErr)

				objectUnderTest := _interpreter{
					interpolater: fakeInterpolater,
				}

				/* act */
				_, actualErr := objectUnderTest.Interpret(
					map[string]*model.Value{},
					"dummyExpression",
					new(data.FakeHandle),
				)

				/* assert */
				Expect(actualErr).To(Equal(interpolateErr))

			})
		})
		Context("interpolater.Interpolate doesn't err", func() {
			It("should call coerce.ToBoolean w/ expected args & return result", func() {
				/* arrange */
				fakeInterpolater := new(interpolater.Fake)

				interpolatedValue := "dummyString"
				fakeInterpolater.InterpolateReturns(interpolatedValue, nil)

				fakeCoerce := new(coerce.Fake)

				coercedValue := model.Value{Boolean: new(bool)}
				fakeCoerce.ToBooleanReturns(&coercedValue, nil)

				objectUnderTest := _interpreter{
					coerce:       fakeCoerce,
					interpolater: fakeInterpolater,
				}

				/* act */
				actualBoolean, actualErr := objectUnderTest.Interpret(
					map[string]*model.Value{},
					"dummyExpression",
					new(data.FakeHandle),
				)

				/* assert */
				Expect(*fakeCoerce.ToBooleanArgsForCall(0)).To(Equal(model.Value{String: &interpolatedValue}))

				Expect(*actualBoolean).To(Equal(coercedValue))
				Expect(actualErr).To(BeNil())
			})
		})
	})
})
