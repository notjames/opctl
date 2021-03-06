package ref

import (
	"github.com/opspec-io/sdk-golang/data"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Context("HandleGetOrHeader", func() {
	Context("HandleGetOrHead", func() {
		Context("req path non empty", func() {
			It("should return expected result", func() {
				/* arrange */
				objectUnderTest := _handleGetOrHeader{}
				providedHTTPResp := httptest.NewRecorder()

				providedHTTPReq, err := http.NewRequest(http.MethodGet, "dummyPath", nil)
				if nil != err {
					panic(err.Error())
				}

				/* act */
				objectUnderTest.HandleGetOrHead(
					new(data.FakeHandle),
					providedHTTPResp,
					providedHTTPReq,
				)

				/* assert */
				Expect(providedHTTPResp.Code).To(Equal(http.StatusNotFound))
			})
		})
		Context("req.Method not GET or HEAD", func() {
			It("should return expected result", func() {
				/* arrange */
				objectUnderTest := _handleGetOrHeader{}
				providedHTTPResp := httptest.NewRecorder()

				providedHTTPReq, err := http.NewRequest("notGETorHEAD", "", nil)
				if nil != err {
					panic(err.Error())
				}

				/* act */
				objectUnderTest.HandleGetOrHead(
					new(data.FakeHandle),
					providedHTTPResp,
					providedHTTPReq,
				)

				/* assert */
				Expect(providedHTTPResp.Code).To(Equal(http.StatusMethodNotAllowed))
			})
		})
	})
})
