package erratum

import (
    "errors"
    "fmt"
)

func Use(opener ResourceOpener, input string) (err error) {
    var resource Resource
	resource, err = opener()
    switch openErr := err.(type) {
    case TransientError:
        return Use(opener, input)
    case error:
        err = errors.Join(openErr)
        return
    default:
        defer func() {
            err = handlePanic(recover(), resource)
        }()
        resource.Frob(input)
        return
    }
}

func handlePanic(recoverResult any, resource Resource) (err error) {
    switch panicMsg := recoverResult.(type) {
        case FrobError:
        resource.Defrob(panicMsg.defrobTag)
        err = errors.Join(err, panicMsg)
        case error:
        err = errors.Join(err, panicMsg)
        case string:
        err = errors.Join(err, errors.New(panicMsg))
        case any:
        err = errors.Join(err, errors.New(fmt.Sprint(panicMsg)))
    }
    err = errors.Join(err, resource.Close())
    return
}