package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

func SortUsers(users []User, sortField string, orderBy int) []User {
	if orderBy == OrderByAsIs {
		return users
	}

	var ascCompareFunc func(i, j int) bool

	switch sortField {
	case "Id":
		ascCompareFunc = func(i, j int) bool { return users[i].Id < users[j].Id }
	case "Age":
		ascCompareFunc = func(i, j int) bool { return users[i].Age < users[j].Age }
	default:
		ascCompareFunc = func(i, j int) bool { return users[i].Name < users[j].Name }
	}

	sort.Slice(users, func(i, j int) bool {
		isFirstLess := ascCompareFunc(i, j)
		if orderBy == OrderByDesc {
			return !isFirstLess
		}
		return isFirstLess
	})

	return users
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	token := r.Header.Get("AccessToken")
	if token != "token" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	queryParams := r.URL.Query()
	atoiQueryParamParse := func(strParam string) int {
		val, err := strconv.Atoi(strParam)
		if err != nil {
			return 0
		}
		return val
	}

	filters := &SearchRequest{
		Limit:      atoiQueryParamParse(queryParams.Get("limit")),
		Offset:     atoiQueryParamParse(queryParams.Get("offset")),
		Query:      queryParams.Get("query"),
		OrderField: queryParams.Get("order_field"),
		OrderBy:    atoiQueryParamParse(queryParams.Get("order_by")),
	}

	if !(filters.OrderField == "Name" || filters.OrderField == "Age" || filters.OrderField == "Id" || filters.OrderField == "") {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, fmt.Sprintf(`{"error": "%s"}`, INVALID_ORDER_FIELD))
		return
	}

	if _, exists := OrderByMap[filters.OrderBy]; !exists {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, fmt.Sprintf(`{"error": "%s"}`, INVALID_ORDER_BY))
		return
	}

	file, err := os.Open("dataset.xml")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	users := make([]User, 0)
	var xmlUser XmlUser

	for {
		tok, err := decoder.Token()

		if err != nil && err != io.EOF {
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else if err == io.EOF {
			break
		}
		switch se := tok.(type) {
		case xml.StartElement:
			if se.Name.Local == "row" {
				if err := decoder.DecodeElement(&xmlUser, &se); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				user := User{
					Id:     xmlUser.Id,
					Name:   xmlUser.FirstName + " " + xmlUser.LastName,
					Age:    xmlUser.Age,
					About:  xmlUser.About,
					Gender: xmlUser.Gender,
				}

				if filters.Query != "" {
					if strings.Contains(user.Name, filters.Query) || strings.Contains(user.About, filters.Query) {
						users = append(users, user)
					}
				} else {
					users = append(users, user)
				}
			}
		}
	}

	startIndex := min(max(0, filters.Offset), len(users))
	endIndex := min(startIndex+max(0, filters.Limit), len(users))

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SortUsers(users, filters.OrderField, filters.OrderBy)[startIndex:endIndex])
}
