package argument

import "os"

const metadataArgument = "--metadata"

type Argument struct {
	urls           []string
	isShowMetadata bool
}

func (Argument) Arguments() []string {
	return os.Args
}

func (a *Argument) prepare() {
	for i, arg := range os.Args {
		if i == 0 {
			continue
		}

		if arg == metadataArgument {
			a.isShowMetadata = true
			continue
		}
		a.urls = append(a.urls, arg)
	}
}

func (a *Argument) IsShowMetadata() bool {
	return a.isShowMetadata
}

func (a *Argument) Urls() []string {
	return a.urls
}

func NewArgument() Argument {
	argument := Argument{[]string{}, false}
	argument.prepare()
	return argument
}
