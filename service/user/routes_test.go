package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nadeem-baig/go-auth/config"
	"github.com/nadeem-baig/go-auth/types"
)

// MockStore is a mock implementation of the UserStore interface
type MockStore struct {
	getUserByEmailFunc func(email string) (*types.User, error)
	createUserFunc     func(user types.User) error
}

func (m *MockStore) GetUserByEmail(email string) (*types.User, error) {
	if m.getUserByEmailFunc != nil {
		return m.getUserByEmailFunc(email)
	}
	return nil, errors.New("GetUserByEmail not implemented")
}

func (m *MockStore) CreateUser(user types.User) error {
	if m.createUserFunc != nil {
		return m.createUserFunc(user)
	}
	return errors.New("CreateUser not implemented")
}

// Ensure MockStore implements UserStore
var _ UserStore = (*MockStore)(nil)

func TestRegisterHandler(t *testing.T) {
	mockHandler := &config.Handler{}

	tests := []struct {
		name           string
		payload        types.RegisterUserPayload
		expectedStatus int
		expectedBody   string
		setupMock      func(*MockStore)
	}{
		{
			name: "Successful Registration",
			payload: types.RegisterUserPayload{
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john@example.com",
				Password:  "password123",
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   `{"message":"User registered successfully"}`,
			setupMock: func(m *MockStore) {
				m.getUserByEmailFunc = func(email string) (*types.User, error) {
					return nil, ErrUserNotFound
				}
				m.createUserFunc = func(user types.User) error {
					return nil
				}
			},
		},
		{
			name: "User Already Exists",
			payload: types.RegisterUserPayload{
				FirstName: "Jane",
				LastName:  "Doe",
				Email:     "jane@example.com",
				Password:  "password123",
			},
			expectedStatus: http.StatusConflict,
			expectedBody:   `{"message":"User already exists"}`,
			setupMock: func(m *MockStore) {
				m.getUserByEmailFunc = func(email string) (*types.User, error) {
					return &types.User{}, nil
				}
			},
		},
		{
			name: "Invalid JSON Payload",
			payload: types.RegisterUserPayload{
				// Intentionally left empty to cause JSON parsing error
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"message":"EOF"}`,
			setupMock: func(m *MockStore) {
				// No need to set up mock functions for this case
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock store and setup expectations
			mockStore := &MockStore{}
			tt.setupMock(mockStore)

			// Create a new request
			payloadBytes, _ := json.Marshal(tt.payload)
			req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(payloadBytes))
			if err != nil {
				t.Fatal(err)
			}

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			// Call the handler
			handler := RegisterHandler(mockHandler, mockStore)
			handler.ServeHTTP(rr, req)

			// Check the status code
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			// Check the response body
			if rr.Body.String() != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tt.expectedBody)
			}
		})
	}
}

// ErrUserNotFound is a custom error for when a user is not found
var ErrUserNotFound = errors.New("user not found")
