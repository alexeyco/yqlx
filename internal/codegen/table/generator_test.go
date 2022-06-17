package table_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"

	"github.com/alexeyco/yqlx/internal/codegen/table"
)

func TestGenerator_To(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)

	const (
		tableName = "table_name"
		directory = "/path/to/directory"
	)

	expectedError := errors.New("error")

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()

		fsMock := afero.NewMemMapFs()
		validatorMock := NewMockValidator(ctrl)
		writerMock := NewMockWriter(ctrl)
		loggerMock := NewMockLogger(ctrl)

		validatorMock.EXPECT().Validate(fsMock, tableName, directory).Return(nil)
		writerMock.EXPECT().
			Write(fsMock, tableName, directory, gomock.Any()).
			DoAndReturn(func(_ afero.Fs, _ string, _ string, fn func(s string)) error {
				fn("fileName")

				return nil
			})
		loggerMock.EXPECT().Infof("Created file %q", "fileName")

		err := table.Generate(
			tableName,
			table.WithFs(fsMock),
			table.WithValidator(validatorMock),
			table.WithWriter(writerMock),
			table.WithLogger(loggerMock),
		).To(directory)

		assert.NoError(t, err)
	})

	t.Run("WriterError", func(t *testing.T) {
		t.Parallel()

		fsMock := afero.NewMemMapFs()
		validatorMock := NewMockValidator(ctrl)
		writerMock := NewMockWriter(ctrl)

		validatorMock.EXPECT().Validate(fsMock, tableName, directory).Return(nil)
		writerMock.EXPECT().Write(fsMock, tableName, directory, gomock.Any()).Return(expectedError)

		err := table.Generate(
			tableName,
			table.WithFs(fsMock),
			table.WithValidator(validatorMock),
			table.WithWriter(writerMock),
		).To(directory)

		assert.ErrorIs(t, err, expectedError)
	})

	t.Run("ValidatorError", func(t *testing.T) {
		t.Parallel()

		fsMock := afero.NewMemMapFs()
		validatorMock := NewMockValidator(ctrl)
		writerMock := NewMockWriter(ctrl)

		validatorMock.EXPECT().Validate(fsMock, tableName, directory).Return(expectedError)

		err := table.Generate(
			tableName,
			table.WithFs(fsMock),
			table.WithValidator(validatorMock),
			table.WithWriter(writerMock),
		).To(directory)

		assert.ErrorIs(t, err, expectedError)
	})
}
