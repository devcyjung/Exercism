package erratum

import "errors"

func Use(opener ResourceOpener, input string) (err error) {
    var resource Resource
	resource, err = opener()
    switch t := err.(type) {
    case TransientError:
        return Use(opener, input)
    case any:
        err = errors.Join(t.(error))
        return
    default:
        defer func() {
            switch t := recover().(type) {
            case FrobError:
                resource.Defrob(t.defrobTag)
                err = errors.Join(err, t)
            case any:
                err = errors.Join(err, t.(error))
            }
            err = errors.Join(err, resource.Close())
        }()
        resource.Frob(input)
        return
    }
}