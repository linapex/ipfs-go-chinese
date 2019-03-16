
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:40</date>
//</624460166271602688>

package options

type PubSubPeersSettings struct {
	Topic string
}

type PubSubSubscribeSettings struct {
	Discover bool
}

type PubSubPeersOption func(*PubSubPeersSettings) error
type PubSubSubscribeOption func(*PubSubSubscribeSettings) error

func PubSubPeersOptions(opts ...PubSubPeersOption) (*PubSubPeersSettings, error) {
	options := &PubSubPeersSettings{
		Topic: "",
	}

	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}
	return options, nil
}

func PubSubSubscribeOptions(opts ...PubSubSubscribeOption) (*PubSubSubscribeSettings, error) {
	options := &PubSubSubscribeSettings{
		Discover: false,
	}

	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}
	return options, nil
}

type pubsubOpts struct{}

var PubSub pubsubOpts

func (pubsubOpts) Topic(topic string) PubSubPeersOption {
	return func(settings *PubSubPeersSettings) error {
		settings.Topic = topic
		return nil
	}
}

func (pubsubOpts) Discover(discover bool) PubSubSubscribeOption {
	return func(settings *PubSubSubscribeSettings) error {
		settings.Discover = discover
		return nil
	}
}

