
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:37</date>
//</624460153894211584>

package util

import (
	"context"
	"io"
)

type ctxCloser context.CancelFunc

func (c ctxCloser) Close() error {
	c()
	return nil
}

func SetupInterruptHandler(ctx context.Context) (io.Closer, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	return ctxCloser(cancel), ctx
}

