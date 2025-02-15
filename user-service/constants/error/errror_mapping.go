package error

func ErrMapping(err error) bool {
	allErrors := append(GeneralErrors, UserErrors...)

	for _, e := range allErrors {
		if err == e {
			return true
		}
	}

	return false
}
