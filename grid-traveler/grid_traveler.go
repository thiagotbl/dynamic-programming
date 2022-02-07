package grid_traveler

func GridTraveler(rows int, columns int) uint {
	memo := map[int]map[int]uint{}

	return gridTraveler(rows, columns, memo)
}

func gridTraveler(rows int, columns int, memo map[int]map[int]uint) uint {
	if rows <= 0 || columns <= 0 {
		return 0
	}

	if rows == 1 && columns == 1 {
		return 1
	}

	value, exists := memo[rows][columns]

	if exists {
		return value
	}

	newValue := gridTraveler(rows-1, columns, memo) + gridTraveler(rows, columns-1, memo)

	addToMemo(rows, columns, newValue, memo)

	return newValue
}

func addToMemo(rows int, columns int, value uint, memo map[int]map[int]uint) {
	innerMap, ok := memo[rows]

	if !ok {
		memo[rows] = map[int]uint{columns: value}
	} else {
		innerMap[columns] = value
	}
}
