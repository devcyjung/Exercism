package erratum

import "errors"

func Use(opener ResourceOpener, input string) (err error) {
	var tranErr TransientError
	var frobErr FrobError
	var resource Resource

	defer func() {
		if resource != nil {
			cerr := resource.Close()
            if cerr != nil {
                err = cerr
            }
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			if errors.As(err, &frobErr) {
				if resource != nil {
					resource.Defrob(err.(FrobError).defrobTag)
				}
			}
		}
	}()

	resource, err = opener()
	for err != nil {
		if errors.As(err, &tranErr) {
			resource, err = opener()
		} else {
			return
		}
	}

	resource.Frob(input)
	return
}

