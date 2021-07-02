package galaxylib

type GalaxyError struct {
	Code    int
	Message string
}

var DefaultGalaxyError = &GalaxyError{}

func (self GalaxyError) Error() string {
	return self.Message
}

func (self GalaxyError) FromError(code int, err error) *GalaxyError {
	if err == nil {
		return nil
	}
	return &GalaxyError{
		Code:    code,
		Message: err.Error(),
	}
}

func (self GalaxyError) FromText(code int, err string) *GalaxyError {
	return &GalaxyError{code, err}
}

func (self GalaxyError) GalaxyOutput() *ApiOutput {
	return &ApiOutput{self.Code, self.Message}
}
