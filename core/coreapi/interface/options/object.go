
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:40</date>
//</624460166141579264>

package options

type ObjectNewSettings struct {
	Type string
}

type ObjectPutSettings struct {
	InputEnc string
	DataType string
	Pin      bool
}

type ObjectAddLinkSettings struct {
	Create bool
}

type ObjectNewOption func(*ObjectNewSettings) error
type ObjectPutOption func(*ObjectPutSettings) error
type ObjectAddLinkOption func(*ObjectAddLinkSettings) error

func ObjectNewOptions(opts ...ObjectNewOption) (*ObjectNewSettings, error) {
	options := &ObjectNewSettings{
		Type: "empty",
	}

	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}
	return options, nil
}

func ObjectPutOptions(opts ...ObjectPutOption) (*ObjectPutSettings, error) {
	options := &ObjectPutSettings{
		InputEnc: "json",
		DataType: "text",
		Pin:      false,
	}

	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}
	return options, nil
}

func ObjectAddLinkOptions(opts ...ObjectAddLinkOption) (*ObjectAddLinkSettings, error) {
	options := &ObjectAddLinkSettings{
		Create: false,
	}

	for _, opt := range opts {
		err := opt(options)
		if err != nil {
			return nil, err
		}
	}
	return options, nil
}

type objectOpts struct{}

var Object objectOpts

//类型是Object.New的选项，它允许更改创建的
//DAG节点。
//
//支持的类型：
//*“空”-空节点
//*'unixfs dir'-空unixfs目录
func (objectOpts) Type(t string) ObjectNewOption {
	return func(settings *ObjectNewSettings) error {
		settings.Type = t
		return nil
	}
}

//inputenc是object.put的一个选项，它指定
//数据。默认为“json”。
//
//支持的编码：
//*“原BUFF”
//*“JSON”
func (objectOpts) InputEnc(e string) ObjectPutOption {
	return func(settings *ObjectPutSettings) error {
		settings.InputEnc = e
		return nil
	}
}

//数据类型是Object.Put的一个选项，它指定数据的编码
//使用JSON或XML输入编码时的字段。
//
//支持的类型：
//*“文本”（默认）
//*“Base64”
func (objectOpts) DataType(t string) ObjectPutOption {
	return func(settings *ObjectPutSettings) error {
		settings.DataType = t
		return nil
	}
}

//pin是对象的一个选项。put指定是否对添加的
//对象，默认值为假
func (objectOpts) Pin(pin bool) ObjectPutOption {
	return func(settings *ObjectPutSettings) error {
		settings.Pin = pin
		return nil
	}
}

//create是object.addlink的一个选项，用于指定是否需要创建
//子目录
func (objectOpts) Create(create bool) ObjectAddLinkOption {
	return func(settings *ObjectAddLinkSettings) error {
		settings.Create = create
		return nil
	}
}

