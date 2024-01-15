package errorCollections

type ErrorCollection struct {
	err   error
	count int
}

func (ec *ErrorCollection) Error() error {
	return ec.err
}

func (ec *ErrorCollection) Count() int {
	return ec.count
}

type ErrorCollections struct {
	ErrorCollections []*ErrorCollection
}

func NewErrorCollections() *ErrorCollections {
	return &ErrorCollections{}
}

func (ec *ErrorCollections) AddError(err error) {
	if err == nil {
		return
	}
	for _, e := range ec.ErrorCollections {
		if e.err.Error() == err.Error() {
			e.count++
			return
		}
	}
	ec.ErrorCollections = append(ec.ErrorCollections, &ErrorCollection{
		err:   err,
		count: 1,
	})
}

func (ec *ErrorCollections) Clear() {
	ec.ErrorCollections = []*ErrorCollection{}
}
