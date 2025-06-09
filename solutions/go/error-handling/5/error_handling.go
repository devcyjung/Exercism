package erratum

import "errors"

func Use(opener ResourceOpener, input string) (err error) {
    var resource Resource
	resource, err = opener()
    switch e := err.(type) {
    case TransientError:
        return Use(opener, input)
    case any:
        err = errors.Join(e.(error))
        return
    default:
        defer func() {
            switch e := recover().(type) {
            case FrobError:
                resource.Defrob(e.defrobTag)
                err = errors.Join(err, e)
            case any:
                err = errors.Join(err, e.(error))
            }
            err = errors.Join(err, resource.Close())
        }()
        resource.Frob(input)
        return
    }
}