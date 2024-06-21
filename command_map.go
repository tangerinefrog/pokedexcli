package main

import (
	"fmt"
	"math"
	"net/url"
	"strconv"

	"github.com/tangerinefrog/pokedexcli/internal/pokeapi"
)

func mapCallback(paging *pagingParam, _ string) error {
	if paging.urlNext == "" && paging.urlPrev != "" {
		return fmt.Errorf("there's no more location areas left")
	}

	err := printLocationAreas(paging, true)
	if err != nil {
		return err
	}

	return nil
}

func mapbCallback(paging *pagingParam, _ string) error {
	if paging.urlPrev == "" {
		return fmt.Errorf("cannot go back")
	}

	err := printLocationAreas(paging, false)
	if err != nil {
		return err
	}

	return nil
}

func printLocationAreas(paging *pagingParam, isForward bool) error {

	var url string = paging.urlNext
	if !isForward {
		url = paging.urlPrev
	}

	offset, limit, err := paramsFromUrl(url)
	if err != nil {
		return err
	}

	areasResp, err := pokeapi.ListLocationAreas(offset, limit)
	if err != nil {
		return err
	}

	if areasResp.UrlNext != nil {
		paging.urlNext = *areasResp.UrlNext
	} else {
		paging.urlNext = ""
	}
	if areasResp.UrlPrev != nil {
		paging.urlPrev = *areasResp.UrlPrev
	} else {
		paging.urlPrev = ""
	}

	totalPages := int64(math.Ceil(float64(areasResp.Count) / float64(PageSize)))
	fmt.Printf("Location areas, page %d of %d:\n\n", currentPage(offset), totalPages)

	for _, v := range areasResp.Results {
		fmt.Printf("\t%s\n", v.Name)
	}

	return nil
}

func currentPage(offset int) int {
	return offset/PageSize + 1
}

func paramsFromUrl(urlstr string) (offset int, limit int, err error) {
	if urlstr == "" {
		return 0, PageSize, nil
	}

	u, err := url.Parse(urlstr)
	if err != nil {
		return
	}
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return
	}

	if offsetVal := q.Get("offset"); offsetVal != "" {
		offset, err = strconv.Atoi(offsetVal)
		if err != nil {
			return
		}
	}
	if limitVal := q.Get("limit"); limitVal != "" {
		limit, err = strconv.Atoi(limitVal)
		if err != nil {
			return
		}
	}
	if limit == 0 {
		limit = PageSize
	}

	return
}
