package galaxylib

type ApiOutput struct {
	Code int
	Data interface{}
}

var DefaultApiOutput = &ApiOutput{}

func (self ApiOutput) Out(code int, data interface{}) *ApiOutput {
	return &ApiOutput{code, data}
}

type IGalaxyOutput interface {
	GalaxyOutput() *ApiOutput
}
