package intersection

import (
	"cmp"
	"slices"
)

func SolutionWithSet[T cmp.Ordered](
	first []T,
	second []T,
) []T {
	set := make(map[T]struct{}, len(first))

	for _, value := range first {
		set[value] = struct{}{}
	}

	added := make(map[T]struct{})
	result := make([]T, 0)

	for _, value := range second {
		_, exists := set[value]
		_, alreadyAdded := added[value]

		if exists && !alreadyAdded {
			result = append(result, value)
			added[value] = struct{}{}
		}
	}

	return result
}

func SolutionWithSort[T cmp.Ordered](
	first []T,
	second []T,
) []T {
	slices.Sort(first)
	slices.Sort(second)

	firstPtr, secondPtr := 0, 0
	result := make([]T, 0)

	for firstPtr < len(first) && secondPtr < len(second) {
		if first[firstPtr] == second[secondPtr] {
			result = append(result, first[firstPtr])

			for firstPtr+1 < len(first) {
				if first[firstPtr] != first[firstPtr+1] {
					break
				}
				firstPtr++
			}

			for secondPtr+1 < len(second) {
				if second[secondPtr] != second[secondPtr+1] {
					break
				}
				secondPtr++
			}

			firstPtr++
			secondPtr++

		} else {
			if first[firstPtr] < second[secondPtr] {
				for firstPtr < len(first) && first[firstPtr] < second[secondPtr] {
					firstPtr++
				}
			} else {
				for secondPtr < len(second) && second[secondPtr] < first[firstPtr] {
					secondPtr++
				}
			}
		}
	}

	return result
}
