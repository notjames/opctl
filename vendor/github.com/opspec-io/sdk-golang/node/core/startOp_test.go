package core

import (
	"context"
	"errors"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/data"
	"github.com/opspec-io/sdk-golang/model"
	"github.com/opspec-io/sdk-golang/node/core/containerruntime"
	"github.com/opspec-io/sdk-golang/op/dotyml"
	"github.com/opspec-io/sdk-golang/util/pubsub"
	"github.com/opspec-io/sdk-golang/util/uniquestring"
)

var _ = Context("core", func() {
	Context("StartOp", func() {
		It("should call data.NewGitProvider w/ expected args", func() {
			/* arrange */
			providedStartOpReq := model.StartOpReq{
				Op: model.StartOpReqOp{
					PullCreds: &model.PullCreds{
						Username: "dummyUsername",
						Password: "dummyPassword",
					},
				},
			}
			providedDataCachePath := "providedDataCachePath"

			fakeData := new(data.Fake)
			// err to trigger immediate return
			fakeData.ResolveReturns(nil, errors.New("dummyErr"))

			objectUnderTest := _core{
				data:          fakeData,
				dataCachePath: providedDataCachePath,
			}

			/* act */
			objectUnderTest.StartOp(
				context.Background(),
				providedStartOpReq,
			)

			/* assert */
			actualCachePath,
				actualPullCreds := fakeData.NewGitProviderArgsForCall(0)

			Expect(actualCachePath).To(Equal(providedDataCachePath))
			Expect(actualPullCreds).To(Equal(providedStartOpReq.Op.PullCreds))
		})
		It("should call data.Resolve w/ expected args", func() {
			/* arrange */
			providedCtx := context.Background()
			providedStartOpReq := model.StartOpReq{
				Op: model.StartOpReqOp{
					Ref: "dummyOpRef",
				},
			}

			fakeData := new(data.Fake)
			// err to trigger immediate return
			fakeData.ResolveReturns(nil, errors.New("dummyErr"))

			fsProvider := new(data.FakeProvider)
			fakeData.NewFSProviderReturns(fsProvider)

			gitProvider := new(data.FakeProvider)
			fakeData.NewGitProviderReturns(gitProvider)

			expectedProviders := []data.Provider{
				fsProvider,
				gitProvider,
			}

			objectUnderTest := _core{
				data: fakeData,
			}

			/* act */
			objectUnderTest.StartOp(
				providedCtx,
				providedStartOpReq,
			)

			/* assert */
			actualCtx,
				actualOpRef,
				actualProviders := fakeData.ResolveArgsForCall(0)

			Expect(actualCtx).To(Equal(providedCtx))
			Expect(actualOpRef).To(Equal(providedStartOpReq.Op.Ref))
			Expect(actualProviders).To(Equal(expectedProviders))
		})
		Context("data.Resolve errs", func() {
			It("should return expected result", func() {

				/* arrange */
				providedCtx := context.Background()
				providedStartOpReq := model.StartOpReq{
					Op: model.StartOpReqOp{
						Ref: "dummyOpRef",
					},
				}

				fakeData := new(data.Fake)
				expectedErr := errors.New("dummyErr")
				fakeData.ResolveReturns(nil, expectedErr)

				objectUnderTest := _core{
					data: fakeData,
				}

				/* act */
				_, actualErr := objectUnderTest.StartOp(
					providedCtx,
					providedStartOpReq,
				)

				/* assert */
				Expect(actualErr).To(Equal(expectedErr))
			})
		})
		Context("data.Resolve doesn't err", func() {
			It("should call data.Get w/ expected args", func() {
				/* arrange */
				providedCtx := context.Background()
				fakeData := new(data.Fake)
				fakeDataHandle := new(data.FakeHandle)
				fakeData.ResolveReturns(fakeDataHandle, nil)

				fakeDotYmlGetter := new(dotyml.FakeGetter)
				// err to trigger immediate return
				fakeDotYmlGetter.GetReturns(nil, errors.New("dummyError"))

				objectUnderTest := _core{
					data:                fakeData,
					dotYmlGetter:        fakeDotYmlGetter,
					uniqueStringFactory: new(uniquestring.Fake),
				}

				/* act */
				objectUnderTest.StartOp(
					providedCtx,
					model.StartOpReq{},
				)

				/* assert */
				actualCtx,
					actualDataHandle := fakeDotYmlGetter.GetArgsForCall(0)

				Expect(actualCtx).To(Equal(providedCtx))
				Expect(actualDataHandle).To(Equal(fakeDataHandle))
			})
			Context("data.Get errs", func() {
				It("should return expected error", func() {
					/* arrange */
					fakeData := new(data.Fake)
					fakeDataHandle := new(data.FakeHandle)
					fakeData.ResolveReturns(fakeDataHandle, nil)

					fakeDotYmlGetter := new(dotyml.FakeGetter)
					expectedErr := errors.New("dummyError")
					fakeDotYmlGetter.GetReturns(&model.OpDotYml{}, expectedErr)

					objectUnderTest := _core{
						data:                fakeData,
						dotYmlGetter:        fakeDotYmlGetter,
						uniqueStringFactory: new(uniquestring.Fake),
					}

					/* act */
					_, actualErr := objectUnderTest.StartOp(
						context.Background(),
						model.StartOpReq{},
					)

					/* assert */
					Expect(actualErr).To(Equal(expectedErr))
				})
			})
			Context("data.Get doesn't err", func() {
				It("should call opCaller.Call w/ expected args", func() {
					/* arrange */
					providedArg1String := "dummyArg1Value"
					providedArg2Dir := "dummyArg2Value"
					providedArg3Dir := "dummyArg3Value"
					providedArg4Dir := "dummyArg4Value"
					providedReq := model.StartOpReq{
						Args: map[string]*model.Value{
							"dummyArg1Name": {String: &providedArg1String},
							"dummyArg2Name": {Dir: &providedArg2Dir},
							"dummyArg3Name": {Dir: &providedArg3Dir},
							"dummyArg4Name": {Dir: &providedArg4Dir},
						},
						Op: model.StartOpReqOp{
							Ref: "dummyOpRef",
						},
					}

					fakeData := new(data.Fake)
					fakeDataHandle := new(data.FakeHandle)
					fakeData.ResolveReturns(fakeDataHandle, nil)

					opDotYml := &model.OpDotYml{
						Outputs: map[string]*model.Param{
							"dummyOutput1": nil,
							"dummyOutput2": nil,
						},
					}

					fakeDotYmlGetter := new(dotyml.FakeGetter)
					fakeDotYmlGetter.GetReturns(opDotYml, nil)

					expectedSCGOpCall := &model.SCGOpCall{
						Ref:     fakeDataHandle.Ref(),
						Inputs:  map[string]interface{}{},
						Outputs: map[string]string{},
					}
					for name := range providedReq.Args {
						expectedSCGOpCall.Inputs[name] = ""
					}
					for name := range opDotYml.Outputs {
						expectedSCGOpCall.Outputs[name] = ""
					}

					expectedOpID := "dummyOpID"

					fakeOpCaller := new(fakeOpCaller)

					fakeUniqueStringFactory := new(uniquestring.Fake)
					fakeUniqueStringFactory.ConstructReturns(expectedOpID, nil)

					objectUnderTest := _core{
						containerRuntime:    new(containerruntime.Fake),
						pubSub:              new(pubsub.Fake),
						data:                fakeData,
						dotYmlGetter:        fakeDotYmlGetter,
						opCaller:            fakeOpCaller,
						dcgNodeRepo:         new(fakeDCGNodeRepo),
						uniqueStringFactory: fakeUniqueStringFactory,
					}

					/* act */
					objectUnderTest.StartOp(
						context.Background(),
						providedReq,
					)

					/* assert */
					// Call happens in go routine; wait 500ms to allow it to occur
					time.Sleep(time.Millisecond * 500)
					actualInboundScope,
						actualOpID,
						actualOpHandle,
						actualRootOpID,
						actualSCGOpCall := fakeOpCaller.CallArgsForCall(0)

					Expect(actualInboundScope).To(Equal(providedReq.Args))
					Expect(actualOpID).To(Equal(expectedOpID))
					Expect(actualOpHandle).To(Equal(fakeDataHandle))
					Expect(actualRootOpID).To(Equal(actualOpID))
					Expect(actualSCGOpCall).To(Equal(expectedSCGOpCall))
				})
			})
		})
	})
})
