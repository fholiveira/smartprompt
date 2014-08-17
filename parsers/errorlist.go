package parsers

type ErrorList struct{ errors []error }

func (list ErrorList) Init() ErrorList {
	list.errors = make([]error, 0)
	return list
}

func (list *ErrorList) Append(err error) {
	if nil != err {
		list.errors = append(list.errors, err)
	}
}

func (list *ErrorList) Items() []error {
	if len(list.errors) == 0 {
		return nil
	}

	return list.errors
}
