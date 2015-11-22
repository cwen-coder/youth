package check

import (
	"regexp"
	"strconv"
	"strings"
	"youth/model"
)

func CheckUserName(name string) (string, bool) {
	name = strings.TrimSpace(name)
	if len(name) == 0 || len(name) > 12 {
		return name, false
	}
	if m, _ := regexp.MatchString("^\\p{Han}+$", name); !m {
		return name, false
	}
	return name, true
}

func CheckNumber(number string) (string, bool) {
	number = strings.TrimSpace(number)
	_, err := strconv.Atoi(number)
	if err != nil {
		return number, false
	}
	if len(number) != 9 {
		return number, false
	}
	result := model.CheckNumber(number)
	if result != true {
		return number, false
	}
	return number, true
}

func CheckSex(sex string) (string, bool) {
	sexs := []string{"man", "woman"}
	for _, v := range sexs {
		if v == sex {
			return sex, true
		}
	}
	return sex, false
}

func CheckBrand(brand string) (string, bool) {
	brands := []string{"add", "nick"}
	for _, v := range brands {
		if v == brand {
			return brand, true
		}
	}
	return brand, false
}

func CheckColor(color string) (string, bool) {
	colors := []string{"red", "black"}
	for _, v := range colors {
		if v == color {
			return color, true
		}
	}
	return color, false
}

func CheckSize(size string) (string, bool) {
	sizes := []string{"l", "xl", "xxl"}
	for _, v := range sizes {
		if v == size {
			return size, true
		}
	}
	return size, false
}
