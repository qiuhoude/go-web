package models

import "strings"

func HandleTagsListData(tags []string) map[string]int {
	var tagsMap = make(map[string]int)
	for _, tag := range tags {
		tagList := strings.Split(tag, "&")
		for _, v := range tagList {
			if _, ok := tagsMap[v]; !ok {
				tagsMap[v] = 0
			}
			tagsMap[v]++
		}
	}
	return tagsMap
}
