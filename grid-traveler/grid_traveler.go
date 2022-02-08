package grid_traveler

import "fmt"

func GridTraveler(rows int, columns int) uint {
	memo := map[string]uint{}

	return gridTraveler(rows, columns, memo)
}

func gridTraveler(rows int, columns int, memo map[string]uint) uint {
	if rows <= 0 || columns <= 0 {
		return 0
	}

	if rows == 1 && columns == 1 {
		return 1
	}

	key := fmt.Sprintf("%d,%d", rows, columns)

	if value, exists := memo[key]; exists {
		return value
	}

	value := gridTraveler(rows-1, columns, memo) + gridTraveler(rows, columns-1, memo)

	memo[key] = value

	return value
}
