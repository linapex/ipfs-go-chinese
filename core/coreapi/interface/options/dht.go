
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:40</date>
//</624460165789257728>

package options

type DhtProvideSettings struct {
	Recursive bool
}

type DhtFindProvidersSettings struct {
	NumProviders int
}

type DhtProvideOption func(*DhtProvideSettings) error
type DhtFindProvidersOption func(*DhtFindProvidersSettings) error

func DhtProvideOptions(opts ...DhtProvideOption) (*DhtProvideSettings, error) {
	options := &DhtProvideSettings{
		Recursive: false,
	}

	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}
	return options, nil
}

func DhtFindProvidersOptions(opts ...DhtFindProvidersOption) (*DhtFindProvidersSettings, error) {
	options := &DhtFindProvidersSettings{
		NumProviders: 20,
	}

	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}
	return options, nil
}

type dhtOpts struct{}

var Dht dhtOpts

//递归是dht.provide的一个选项，指定是否提供
//给定路径递归
func (dhtOpts) Recursive(recursive bool) DhtProvideOption {
	return func(settings *DhtProvideSettings) error {
		settings.Recursive = recursive
		return nil
	}
}

//numProviders是dht.findproviders的一个选项，它指定
//要查找的对等数。默认值为20
func (dhtOpts) NumProviders(numProviders int) DhtFindProvidersOption {
	return func(settings *DhtFindProvidersSettings) error {
		settings.NumProviders = numProviders
		return nil
	}
}

