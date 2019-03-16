
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:40</date>
//</624460165902503936>

package options

type ApiSettings struct {
	Offline     bool
	FetchBlocks bool
}

type ApiOption func(*ApiSettings) error

func ApiOptions(opts ...ApiOption) (*ApiSettings, error) {
	options := &ApiSettings{
		Offline:     false,
		FetchBlocks: true,
	}

	return ApiOptionsTo(options, opts...)
}

func ApiOptionsTo(options *ApiSettings, opts ...ApiOption) (*ApiSettings, error) {
	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}
	return options, nil
}

type apiOpts struct{}

var Api apiOpts

func (apiOpts) Offline(offline bool) ApiOption {
	return func(settings *ApiSettings) error {
		settings.Offline = offline
		return nil
	}
}

//如果设置为false，则fetchBlocks会阻止API从
//网络，同时允许其他服务（如IPN）仍然在线
func (apiOpts) FetchBlocks(fetch bool) ApiOption {
	return func(settings *ApiSettings) error {
		settings.FetchBlocks = fetch
		return nil
	}
}

