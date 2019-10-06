package api

import (
	"sort"
)

type deliveryResponse struct {
	IDs          []int `json:"ids"`
	DeliveryDays int   `json:"deliveryDays"`
}

type deliveryBox struct {
	Items        []deliveryRequest
	DeliveryDays int
}

const maxWeight = 10

func mapBoxToIDs(vs []deliveryRequest) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = v.ID
	}
	return vsm
}

func createBox(items []deliveryRequest) deliveryBox {
	deliveryDays := 0

	for _, v := range items {
		deliveryDays = deliveryDays + v.DeliveryDays
	}

	box := deliveryBox{
		Items:        items,
		DeliveryDays: deliveryDays,
	}

	return box
}

func calculateBox(items []deliveryRequest) []deliveryBox {
	// Create an array of boxes
	var boxes []deliveryBox
	highestKey := 0

	// Loop backwards through the items, so we start at the largest
	for key := range items {
		key = len(items) - 1 - key
		value := items[key]

		// If we've already reached this key from the bottom, break, otherwise we'll
		// end up with boxes that have just the lowest weights in, which we don't want.
		if key <= highestKey {
			break
		}

		var itemsThatFit []deliveryRequest

		remainingWeight := maxWeight - value.Weight

		if remainingWeight > 0 {
			itemsThatFit = append(itemsThatFit, value)

			// We need to loop over the items from the smallest first, and work up.
			for lowestKey, lowestValue := range items {
				if key == lowestKey {
					break
				}

				if lowestValue.Weight <= remainingWeight {
					itemsThatFit = append(itemsThatFit, lowestValue)
					remainingWeight = remainingWeight - lowestValue.Weight
					highestKey = lowestKey
				}
			}
		}

		if len(itemsThatFit) > 0 {
			boxes = append(boxes, createBox(itemsThatFit))
		}
	}

	return boxes
}

func calculateBoxes(items []deliveryRequest) deliveryResponse {
	// Copy the items, so we're not mutating
	sortedItems := []deliveryRequest{}
	sortedItems = append([]deliveryRequest(nil), items...)

	// First sort the items by their weight, smallest to heaviest
	sort.Slice(sortedItems, func(i, j int) bool {
		return sortedItems[i].Weight < sortedItems[j].Weight
	})

	return convertBoxesToResponse(calculateBox(sortedItems))
}

func convertBoxesToResponse(boxes []deliveryBox) deliveryResponse {
	// Copy the items, so we're not mutating
	sortedItems := []deliveryBox{}
	sortedItems = append([]deliveryBox(nil), boxes...)

	// Sort the items by delivery date
	sort.Slice(sortedItems, func(i, j int) bool {
		return sortedItems[i].DeliveryDays < sortedItems[j].DeliveryDays
	})

	if len(sortedItems) == 0 {
		return deliveryResponse{}
	}

	smallestBox := sortedItems[0]

	return deliveryResponse{
		IDs:          mapBoxToIDs(smallestBox.Items),
		DeliveryDays: smallestBox.DeliveryDays,
	}
}
