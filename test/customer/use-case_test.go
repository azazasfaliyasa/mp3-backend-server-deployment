package customer

import (
	"errors"
	"mini-project-3-azaz/internal/customer"
	"mini-project-3-azaz/internal/customer/mocks"
	"reflect"
	"testing"
)

func TestUseCase_ShowAll(t *testing.T) {
	type fields struct {
		R customer.RepositoryInterface
	}

	successCase := []customer.Customer{}
	errCase := errors.New("some error")

	for i := 0; i < 5; i++ {
		successCase = append(successCase, customer.Customer{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "johndoe@example.com",
			Avatar:    "",
		})
	}

	mockRepository := mocks.NewRepositoryInterface(t)
	mockRepository.EXPECT().ShowAll().Return(successCase, nil).Once()
	mockRepository.EXPECT().ShowAll().Return(nil, errCase).Once()

	tests := []struct {
		name    string
		fields  fields
		want    []customer.Customer
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success case",
			fields: fields{
				R: mockRepository,
			},
			want:    successCase,
			wantErr: false,
		},
		{
			name: "error case",
			fields: fields{
				R: mockRepository,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &customer.UseCase{
				R: tt.fields.R,
			}
			got, err := u.ShowAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.ShowAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.ShowAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_Create(t *testing.T) {
	type fields struct {
		R customer.RepositoryInterface
	}
	type args struct {
		customer *customer.Customer
	}

	successCase := customer.Customer{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "ychag@example.com",
		Avatar:    "",
	}
	nilRequest := customer.Customer{}
	errorCase := errors.New("some error")

	mockRepository := mocks.NewRepositoryInterface(t)
	mockRepository.EXPECT().Create(&successCase).Return(nil).Once()
	mockRepository.EXPECT().Create(&nilRequest).Return(errorCase).Once()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				R: mockRepository,
			},
			args: args{
				customer: &successCase,
			},
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				R: mockRepository,
			},
			args: args{
				customer: &nilRequest,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &customer.UseCase{
				R: tt.fields.R,
			}
			if err := u.Create(tt.args.customer); (err != nil) != tt.wantErr {
				t.Errorf("UseCase.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_Show(t *testing.T) {
	type fields struct {
		R customer.RepositoryInterface
	}
	type args struct {
		ID string
	}
	successCase := customer.Customer{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "ychag@example.com",
		Avatar:    "",
	}
	nilData := customer.Customer{}
	errorCase := errors.New("some error")

	mockRepository := mocks.NewRepositoryInterface(t)
	mockRepository.EXPECT().Show("1").Return(&successCase, nil).Once()
	mockRepository.EXPECT().Show("").Return(&nilData, errorCase).Once()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *customer.Customer
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				R: mockRepository,
			},
			args: args{
				ID: "1",
			},
			want:    &successCase,
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				R: mockRepository,
			},
			args: args{
				ID: "",
			},
			want:    &nilData,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &customer.UseCase{
				R: tt.fields.R,
			}
			got, err := u.Show(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.Show() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.Show() = %v, want %v", got, tt.want)
			}
		})
	}
}
