package future

type SuccessFunc func(string)
type FailFunc func(error)
type ExecuteStringFunc func() (string, error)

type MaybeString struct {
	successFunc SuccessFunc
	failFunc    FailFunc
}

func (ms *MaybeString) Success(sf SuccessFunc) *MaybeString {
	ms.successFunc = sf
	return ms
}

func (ms *MaybeString) Fail(ff FailFunc) *MaybeString {
	ms.failFunc = ff
	return ms
}

func (ms *MaybeString) Execute(esf ExecuteStringFunc) {
	go func(ms *MaybeString) {
		str, err := esf()

		if err != nil {
			ms.failFunc(err)
		} else {
			ms.successFunc(str)
		}
	}(ms)
}
