package erratum

import "errors"

func Use(opener ResourceOpener, input string) (err error) {
	var tranErr TransientError
	var frobErr FrobError
	var resource Resource

	defer func() {
		if resource != nil {
			e := resource.Close()
			if e != nil {
				err = errors.Join(err, e)
			}
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			e := r.(error)
            err = errors.Join(err, e)
			if errors.As(e, &frobErr) {
				if resource != nil {
					resource.Defrob(e.(FrobError).defrobTag)
				}
			}
		}
	}()

	resource, e := opener()
	for e != nil {
		if errors.As(e, &tranErr) {
			resource, e = opener()
		} else {
            err = errors.Join(err, e)
			return
		}
	}

	resource.Frob(input)
	return
}
