package pageUtil

import "gorm.io/gorm"

// Response: Base response of query
type Response struct {
	PageNow    int         //PageNow: current page of query
	PageCount  int         //PageCount: total page of the query
	RawCount   int         //RawCount: total raw of query
	RawPerPage int         //RawPerPage: rpp
	ResultSet  interface{} //ResultSet: result data
	FirstPage  bool        //FirstPage: if the result is the first page
	LastPage   bool        //LastPage: if the result is the last page
	Empty      bool        //Empty: if the result is empty
	StartRow   int         //The number of first record the the resultSet
	EndRow     int         //The number of last record the the resultSet
	handler    *gorm.DB    //the handler of gorm Query
}
type recoveryHandler func()
var recovery recoveryHandler
var use0Page bool = false
var defaultRpp int

// getLimitOffset (private) get LIMIT and OFFSET keyword in SQL
func getLimitOffset(page, rpp int) (limit, offset int) {
	if page < 0 {
		page = 0
	}
	if rpp < 1 {
		rpp = defaultRpp
	}
	return rpp, page * rpp
}

func PageQuery(page int, rawPerPage int, queryHandler *gorm.DB, resultPtr interface{}) (*Response, error) {
	//recovery
	//	defer recovery()
	count := 0
	// get limit and offSet
	var limit, offset int
	if !use0Page {
		limit, offset = getLimitOffset(page-1, rawPerPage)
	} else {
		limit, offset = getLimitOffset(page, rawPerPage)
	}
	// get total count of the table
	a:=int64(count)
	queryHandler.Count(&a)
	// get result set by param
	queryHandler.Limit(limit).Offset(offset).Find(resultPtr)
	// handle DB error
	if err := queryHandler.Error; err != nil {
		return nil, err
	}
	// get page count
	count=int(a)
	PageCount := count/ rawPerPage
	if count%rawPerPage != 0 {
		PageCount++
	}
	startRow, endRow, empty, lastPage := 0, 0, (offset > count) || count == 0, page == PageCount
	if !empty {
		startRow = (page-1) * rawPerPage +1
		if !lastPage {
			endRow = page*rawPerPage
		} else {
			endRow = count
		}
	}
	// prepare base response
	return &Response{
		PageNow:    page,
		PageCount:  PageCount,
		RawPerPage: rawPerPage,
		RawCount:   count,
		ResultSet:  resultPtr,
		FirstPage:  page == 1,
		LastPage:   lastPage,
		Empty:      empty,
		StartRow:   startRow,
		EndRow:     endRow,
		handler:    queryHandler,
	}, nil
}