package api

import (
	"reflect"
	"testing"
)

func TestCalculateBoxes(t *testing.T) {
	items := [][]deliveryRequest{
		{
			{1, 3, 2},
			{2, 2, 3},
			{3, 8, 6},
			{4, 6, 1},
			{5, 3, 9},
			{6, 1, 4},
		},
		{
			{1, 1, 4},
			{2, 8, 1},
			{3, 7, 2},
			{4, 4, 10},
			{5, 3, 3},
			{6, 2, 5},
		},
	}

	expectedResponse := []deliveryResponse{
		{
			IDs:          []int{4, 6, 2},
			DeliveryDays: 8,
		},
		{
			IDs:          []int{2, 1},
			DeliveryDays: 5,
		},
	}

	for key, value := range items {
		actualResponse := calculateBoxes(value)

		if !reflect.DeepEqual(expectedResponse[key], actualResponse) {
			t.Errorf("Response was incorrect, got: %+v, want: %+v.", actualResponse, expectedResponse[key])
		}
	}
}
