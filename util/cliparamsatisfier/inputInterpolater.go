package cliparamsatisfier

//go:generate counterfeiter -o ./fakeInputInterpolater.go --fake-name FakeInputInterpolater ./ InputInterpolater

import (
	"context"
	"github.com/golang-interfaces/ios"
	"github.com/opspec-io/sdk-golang/data"
	"github.com/opspec-io/sdk-golang/model"
	"github.com/opspec-io/sdk-golang/op/interpreter/interpolater"
	"os"
	"strings"
)

type InputInterpolater interface {
	// Interpolates the input expression
	Interpolate(
		inputExpression string,
	) (string, error)
}

func newInputInterpolater(
	sources ...InputSrc,
) InputInterpolater {
	return inputInterpolater{
		data:         data.New(),
		interpolater: interpolater.New(),
		os:           ios.New(),
	}
}

type inputInterpolater struct {
	data         data.Data
	interpolater interpolater.Interpolater
	os           ios.IOS
}

func (is inputInterpolater) getScopeFromEnvironment() map[string]*model.Value {
	environmentScope := map[string]*model.Value{}
	for _, envVar := range os.Environ() {
		envVarKeyValue := strings.Split(envVar, "=")

		environmentScope[envVarKeyValue[0]] = &model.Value{
			String: &envVarKeyValue[1],
		}
	}

	return environmentScope
}

func (is inputInterpolater) Interpolate(
	inputExpression string,
) (string, error) {
	scope := is.getScopeFromEnvironment()

	opDirHandle, err := is.data.Resolve(
		context.TODO(),
		"/",
		is.data.NewFSProvider(),
	)
	if nil != err {
		return "", err
	}

	// @TODO: this currently will not interpolate relative path references
	// relative paths don't make sense inside of ops but they do in the context
	// of a CLI where you want to source inputs from your local environment
	return is.interpolater.Interpolate(
		inputExpression,
		scope,
		opDirHandle,
	)
}
