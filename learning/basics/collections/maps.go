package collections

import "fmt"

type User struct {
	Id   int64
	Name string
}

func findInSlice(id int64, users []User) *User {
	for _, user := range users {
		if user.Id == id {
			return &user
		}
	}
	return nil
}

func findInMap(id int64, usersMap map[int64]User) *User {
	if user, ok := usersMap[id]; ok {
		return &user
	}
	return nil
}

func Maps() {
	var defaultMap map[int64]string

	fmt.Printf("defaultMap (%T): %#v\n\n", defaultMap, defaultMap)

	mapByMake := make(map[string]string)
	fmt.Printf("mapByMake (%T): %#v\n", mapByMake, mapByMake)

	mapByMakeWithCap := make(map[string]string, 3)
	fmt.Printf("mapByMakeWithCap (%T): %#v\n\n", mapByMakeWithCap, mapByMakeWithCap)

	mapByLiteral := map[int64]string{
		1: "Vasya",
		2: "Petya",
		3: "Ivan",
	}
	fmt.Printf("mapByLiteral (%T): %#v\n", mapByLiteral, mapByLiteral)
	fmt.Printf("len = %v\n\n", len(mapByLiteral))

	mapByNew := *new(map[int64]string)
	fmt.Printf("mapByNew (%T): %#v\n\n", mapByNew, mapByNew)

	exMap := map[string]string{
		"Second": "Petya",
		"Third":  "Ivan",
	}
	// insert value
	exMap["First"] = "Vasya"
	fmt.Printf("exMap insert (%T): %#v\n", exMap, exMap)
	fmt.Printf("len = %v\n\n", len(exMap))

	// update value
	exMap["First"] = "New Vasya"
	fmt.Printf("exMap insert (%T): %#v\n", exMap, exMap)
	fmt.Printf("len = %v\n\n", len(exMap))

	// get value
	fmt.Print("exMap['First'] = ", exMap["First"], "\n")

	// get value by not existed key (default value of type is returned)
	fmt.Print("exMap['NotExistedKey'] = ", exMap["NotExistedKey"], "\n")

	// check if value by key is existed
	value, ok := exMap["NotExistedKey"]
	fmt.Printf("Value of NotExistedKey = %v, ok = %v\n\n", value, ok)

	// delete value
	delete(exMap, "Second")
	fmt.Printf("exMap delete (%T): %#v\n", exMap, exMap)
	fmt.Printf("len = %v\n\n", len(exMap))

	for key, value := range exMap {
		fmt.Printf("%v: %v\n", key, value)
	}
	fmt.Println()

	users := []User{
		{
			Id:   1,
			Name: "Vasya",
		},
		{
			Id:   2,
			Name: "Petya",
		},
		{
			Id:   3,
			Name: "Ivan",
		},
		{
			Id:   2,
			Name: "Petya",
		},
	}

	// map as set
	uniqueUsers := make(map[int64]struct{}, len(users))

	for _, user := range users {
		if _, ok := uniqueUsers[user.Id]; !ok {
			uniqueUsers[user.Id] = struct{}{}
		}
	}
	fmt.Printf("uniqueUsers (%T): %#v\n\n", uniqueUsers, uniqueUsers)

	// use map for avoid O(n) and use O(1)
	usersMap := make(map[int64]User, len(users))

	for _, user := range users {
		if _, ok := usersMap[user.Id]; !ok {
			usersMap[user.Id] = user
		}
	}

	fmt.Println("findInSlice:", findInSlice(3, users))
	fmt.Println("findInMap:", findInMap(3, usersMap))
}
