package codegen_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"

	"github.com/alexeyco/yqlx/internal/codegen"
)

// nolint:funlen
func TestCodegen_Add(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	fs := afero.NewMemMapFs()

	const (
		output    = "internal/yqlx/schema"
		tablePath = "table"
	)

	// nolint:goerr113
	expectedErr := errors.New("error")

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()

		validatorMock := NewMockValidator(ctrl)
		addGeneratorMock := NewMockAddGenerator(ctrl)

		validatorMock.EXPECT().ValidateAdd(fs, output, tablePath).Return(nil)
		addGeneratorMock.EXPECT().Generate(fs, output, tablePath).Return(nil)

		err := codegen.New(
			codegen.WithFs(fs),
			codegen.WithValidator(validatorMock),
			codegen.WithAddGenerator(addGeneratorMock),
		).Add(output, tablePath)

		assert.NoError(t, err)
	})

	t.Run("GeneratorError", func(t *testing.T) {
		t.Parallel()

		validatorMock := NewMockValidator(ctrl)
		addGeneratorMock := NewMockAddGenerator(ctrl)

		validatorMock.EXPECT().ValidateAdd(fs, output, tablePath).Return(nil)
		addGeneratorMock.EXPECT().Generate(fs, output, tablePath).Return(expectedErr)

		err := codegen.New(
			codegen.WithFs(fs),
			codegen.WithValidator(validatorMock),
			codegen.WithAddGenerator(addGeneratorMock),
		).Add(output, tablePath)

		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("GeneratorError", func(t *testing.T) {
		t.Parallel()

		validatorMock := NewMockValidator(ctrl)
		addGeneratorMock := NewMockAddGenerator(ctrl)

		validatorMock.EXPECT().ValidateAdd(fs, output, tablePath).Return(expectedErr)

		err := codegen.New(
			codegen.WithFs(fs),
			codegen.WithValidator(validatorMock),
			codegen.WithAddGenerator(addGeneratorMock),
		).Add(output, tablePath)

		assert.ErrorIs(t, err, expectedErr)
	})
}
