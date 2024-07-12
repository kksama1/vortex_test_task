package postgres

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"vortex/internal/model"
)

var pool = CreateConnection("localhost", 5432, "vortexTaskDB", "kksama", "kksama1", "disable")
var pd = &PostgresDriver{Pool: pool}

func TestAddClient(t *testing.T) {
	tests := []struct {
		name     string
		input    model.Client
		expected error
	}{
		{
			name: "no error",
			input: model.Client{
				ClientName:  "ww",
				Version:     12,
				Image:       "ww",
				CPU:         "ww",
				Memory:      "ww",
				Priority:    0.1,
				NeedRestart: false,
			},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pd.AddClient(&tt.input)
			assert.NoError(t, result, tt.expected)
		})
	}
}

func TestUpdateClient(t *testing.T) {
	tests := []struct {
		name  string
		input model.Client
	}{
		{
			name: "no error",
			input: model.Client{
				ID:          777,
				ClientName:  "ww",
				Version:     12,
				Image:       "ww",
				CPU:         "ww",
				Memory:      "ww",
				Priority:    0.1,
				NeedRestart: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pd.UpdateClient(&tt.input)
			assert.NoError(t, result)
		})
	}
}

func TestUpdateAlgorithmStatus(t *testing.T) {
	tests := []struct {
		name  string
		input model.Algorithm
	}{
		{
			name: "0 0 0",
			input: model.Algorithm{
				VWAP: false,
				TWAP: false,
				HFT:  false,
			},
		},
		{
			name: "0 0 1",
			input: model.Algorithm{
				VWAP: false,
				TWAP: false,
				HFT:  true,
			},
		},
		{
			name: "0 1 0",
			input: model.Algorithm{
				VWAP: false,
				TWAP: true,
				HFT:  false,
			},
		},
		{
			name: "0 1 1",
			input: model.Algorithm{
				VWAP: false,
				TWAP: true,
				HFT:  true,
			},
		},
		{
			name: "1 0 0",
			input: model.Algorithm{
				VWAP: true,
				TWAP: false,
				HFT:  false,
			},
		},
		{
			name: "1 0 1",
			input: model.Algorithm{
				VWAP: true,
				TWAP: false,
				HFT:  true,
			},
		},
		{
			name: "1 1 0",
			input: model.Algorithm{
				VWAP: true,
				TWAP: true,
				HFT:  false,
			},
		},
		{
			name: "1 1 1",
			input: model.Algorithm{
				VWAP: true,
				TWAP: true,
				HFT:  true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pd.UpdateAlgorithmStatus(&tt.input)
			assert.NoError(t, result)
		})
	}
}

func TestDeleteClient(t *testing.T) {
	tests := []struct {
		name  string
		input model.Client
	}{
		{
			name: "no error",
			input: model.Client{
				ID:          2,
				ClientName:  "ww",
				Version:     12,
				Image:       "ww",
				CPU:         "ww",
				Memory:      "ww",
				Priority:    0.1,
				NeedRestart: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pd.DeleteClient(&tt.input)
			assert.NoError(t, result)
		})
	}
}

//func TestGetInActiveAlgorithms(t *testing.T) {
//	log.Println(pd.GetInActiveAlgorithms())
//
//	tests := []struct {
//		name        string
//		expected    []model.Algorithm
//		expectedErr error
//	}{
//		{
//			name: "no error",
//			expected: []model.Algorithm{
//				{
//					ID:       1,
//					ClientID: 1,
//					VWAP:     false,
//					TWAP:     false,
//					HFT:      false,
//				},
//				{
//					ID:       3,
//					ClientID: 3,
//					VWAP:     false,
//					TWAP:     false,
//					HFT:      false,
//				},
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			result, err := pd.GetInActiveAlgorithms()
//			reflect.DeepEqual(result, tt.expected)
//			assert.NoError(t, err)
//		})
//	}
//}

func TestGetAllAlgorithms(t *testing.T) {
	result := pd.GetAllAlgorithms()
	assert.NoError(t, result)
}
