package ref

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/opspec-io/sdk-golang/data"

	"github.com/opspec-io/sdk-golang/model"
	"github.com/opspec-io/sdk-golang/node/core"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Context("Handler", func() {
	Context("NewHandler", func() {
		It("should not return nil", func() {
			/* arrange/act/assert */
			Expect(NewHandler(new(core.Fake))).Should(Not(BeNil()))
		})
	})
	Context("Handle", func() {
		Context("next URL path segment not empty", func() {
			It("should return expected result", func() {
				/* arrange */
				objectUnderTest := _handler{}
				providedHTTPResp := httptest.NewRecorder()

				providedHTTPReq, err := http.NewRequest("dummyMethod", "dummyPath", nil)
				if nil != err {
					panic(err.Error())
				}

				/* act */
				objectUnderTest.Handle(
					"dummyDataRef",
					providedHTTPResp,
					providedHTTPReq,
				)

				/* assert */
				Expect(providedHTTPResp.Code).To(Equal(http.StatusNotFound))
			})
		})
		Context("next URL path segment isn't empty", func() {
			Context("req has BasicAuth", func() {
				It("should call core.ResolveData w/ expected args", func() {
					/* arrange */
					providedDataRef := "dummyDataRef"
					providedUsername := "dummyUsername"
					providedPassword := "dummyPassword"

					fakeCore := new(core.Fake)
					// err to trigger immediate return
					fakeCore.ResolveDataReturns(nil, errors.New("dummyErr"))

					objectUnderTest := _handler{
						core: fakeCore,
					}

					providedHTTPReq, err := http.NewRequest("dummyHttpMethod", "", nil)
					if nil != err {
						panic(err.Error())
					}

					providedHTTPReq.SetBasicAuth(
						providedUsername,
						providedPassword,
					)

					expectedPullCreds := model.PullCreds{
						Username: providedUsername,
						Password: providedPassword,
					}

					/* act */
					objectUnderTest.Handle(
						providedDataRef,
						httptest.NewRecorder(),
						providedHTTPReq,
					)

					/* assert */
					_,
						actualRef,
						actualPullCreds := fakeCore.ResolveDataArgsForCall(0)

					Expect(actualRef).To(Equal(providedDataRef))
					Expect(*actualPullCreds).To(Equal(expectedPullCreds))
				})
			})
			It("should call core.ResolveData w/ expected args", func() {
				/* arrange */
				providedDataRef := "dummyDataRef"

				fakeCore := new(core.Fake)
				// err to trigger immediate return
				fakeCore.ResolveDataReturns(nil, errors.New("dummyErr"))

				objectUnderTest := _handler{
					core: fakeCore,
				}

				providedHTTPReq, err := http.NewRequest("dummyHttpMethod", "", nil)
				if nil != err {
					panic(err.Error())
				}

				/* act */
				objectUnderTest.Handle(
					providedDataRef,
					httptest.NewRecorder(),
					providedHTTPReq,
				)

				/* assert */
				_,
					actualRef,
					actualPullCreds := fakeCore.ResolveDataArgsForCall(0)

				Expect(actualRef).To(Equal(providedDataRef))
				Expect(actualPullCreds).To(BeNil())
			})
			Context("core.ResolveData errs", func() {
				Context("err is ErrDataProviderAuthentication", func() {
					It("should return expected result", func() {
						/* arrange */
						dataRefSegment1 := "dataRefSegment1"
						providedDataRef := strings.Join([]string{dataRefSegment1, "dataRefSegment2"}, "/")

						fakeCore := new(core.Fake)
						fakeCore.ResolveDataReturns(nil, model.ErrDataProviderAuthentication{})

						objectUnderTest := _handler{
							core: fakeCore,
						}
						providedHTTPResp := httptest.NewRecorder()

						providedHTTPReq, err := http.NewRequest("dummyHttpMethod", "", nil)
						if nil != err {
							panic(err.Error())
						}

						/* act */
						objectUnderTest.Handle(
							providedDataRef,
							providedHTTPResp,
							providedHTTPReq,
						)

						/* assert */
						Expect(providedHTTPResp.Code).To(Equal(http.StatusUnauthorized))
						Expect(providedHTTPResp.Header().Get("WWW-Authenticate")).To(Equal(fmt.Sprintf(`Basic realm="%s"`, dataRefSegment1)))
					})
				})
				Context("err is ErrDataProviderAuthorization", func() {
					It("should return expected result", func() {
						/* arrange */
						dataRefSegment1 := "dataRefSegment1"
						providedDataRef := strings.Join([]string{dataRefSegment1, "dataRefSegment2"}, "/")

						fakeCore := new(core.Fake)
						fakeCore.ResolveDataReturns(nil, model.ErrDataProviderAuthorization{})

						objectUnderTest := _handler{
							core: fakeCore,
						}
						providedHTTPResp := httptest.NewRecorder()

						providedHTTPReq, err := http.NewRequest("dummyHttpMethod", "", nil)
						if nil != err {
							panic(err.Error())
						}

						/* act */
						objectUnderTest.Handle(
							providedDataRef,
							providedHTTPResp,
							providedHTTPReq,
						)

						/* assert */
						Expect(providedHTTPResp.Code).To(Equal(http.StatusForbidden))
						Expect(providedHTTPResp.Header().Get("WWW-Authenticate")).To(Equal(fmt.Sprintf(`Basic realm="%s"`, dataRefSegment1)))
					})
				})
				Context("err is ErrDataRefResolution", func() {
					It("should return expected result", func() {
						/* arrange */
						fakeCore := new(core.Fake)
						fakeCore.ResolveDataReturns(nil, model.ErrDataRefResolution{})

						objectUnderTest := _handler{
							core: fakeCore,
						}
						providedHTTPResp := httptest.NewRecorder()

						providedHTTPReq, err := http.NewRequest("dummyHttpMethod", "", nil)
						if nil != err {
							panic(err.Error())
						}

						/* act */
						objectUnderTest.Handle(
							"dummyDataRef",
							providedHTTPResp,
							providedHTTPReq,
						)

						/* assert */
						Expect(providedHTTPResp.Code).To(Equal(http.StatusNotFound))
					})
				})
			})
			Context("core.ResolveData doesn't err", func() {
				It("should call handleGetOrHeader.HandleGetOrHead w/ expected args", func() {
					/* arrange */
					providedDataRef := "dummyDataRef"

					fakeCore := new(core.Fake)
					fakeDataHandle := new(data.FakeHandle)
					fakeCore.ResolveDataReturns(fakeDataHandle, nil)

					fakeHandleGetOrHeader := new(fakeHandleGetOrHeader)

					objectUnderTest := _handler{
						core:              fakeCore,
						handleGetOrHeader: fakeHandleGetOrHeader,
					}

					providedHTTPReq, err := http.NewRequest("dummyHttpMethod", "", nil)
					if nil != err {
						panic(err.Error())
					}

					/* act */
					objectUnderTest.Handle(
						providedDataRef,
						httptest.NewRecorder(),
						providedHTTPReq,
					)

					/* assert */
					actualDataHandle,
						_,
						actualHTTPReq := fakeHandleGetOrHeader.HandleGetOrHeadArgsForCall(0)

					Expect(actualDataHandle).To(Equal(fakeDataHandle))
					Expect(actualHTTPReq.URL.Path).To(BeEmpty())

					// this works because our URL path set mutates the request
					Expect(actualHTTPReq).To(Equal(providedHTTPReq))
				})
			})
		})
	})
})
