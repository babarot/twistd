package twistd

import ()

type Twistd struct {
	Option *Option
}

type Option struct {
	Child      bool
	Config     string
	Foreground bool
}

func NewTwistd(opt *Option) (*Twistd, error) {
	return &Twistd{opt}, nil
}
