
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:43</date>
//</624460177822715904>

package keystore

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	ci "gx/ipfs/QmNiJiXwWE3kRhZrC5ej3kSjWHm337pYfhjLGSCDNKJP2s/go-libp2p-crypto"
	logging "gx/ipfs/QmcuXC5cxs79ro2cUuHs4HQ2bkDLJUYokwL8aivcX6HW3C/go-log"
)

var log = logging.Logger("keystore")

//密钥库提供密钥管理接口
type Keystore interface {
//has返回密钥库中是否存在密钥
	Has(string) (bool, error)
//Put在密钥库中存储密钥，如果已存在同名密钥，则返回errkeyexists
	Put(string, ci.PrivKey) error
//get从密钥存储库中检索密钥（如果存在），并返回errnosuchkey
//否则。
	Get(string) (ci.PrivKey, error)
//删除从密钥库中删除密钥
	Delete(string) error
//list返回键标识符列表
	List() ([]string, error)
}

var ErrNoSuchKey = fmt.Errorf("no key by the given name was found")
var ErrKeyExists = fmt.Errorf("key by that name already exists, refusing to overwrite")

//fskeystore是一个由存储在磁盘上的给定目录中的文件支持的密钥库。
type FSKeystore struct {
	dir string
}

func validateName(name string) error {
	if name == "" {
		return fmt.Errorf("key names must be at least one character")
	}

	if strings.Contains(name, "/") {
		return fmt.Errorf("key names may not contain slashes")
	}

	if strings.HasPrefix(name, ".") {
		return fmt.Errorf("key names may not begin with a period")
	}

	return nil
}

func NewFSKeystore(dir string) (*FSKeystore, error) {
	_, err := os.Stat(dir)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
		if err := os.Mkdir(dir, 0700); err != nil {
			return nil, err
		}
	}

	return &FSKeystore{dir}, nil
}

//has返回密钥库中是否存在密钥
func (ks *FSKeystore) Has(name string) (bool, error) {
	kp := filepath.Join(ks.dir, name)

	_, err := os.Stat(kp)

	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	if err := validateName(name); err != nil {
		return false, err
	}

	return true, nil
}

//Put在密钥库中存储密钥，如果已存在同名密钥，则返回errkeyexists
func (ks *FSKeystore) Put(name string, k ci.PrivKey) error {
	if err := validateName(name); err != nil {
		return err
	}

	b, err := k.Bytes()
	if err != nil {
		return err
	}

	kp := filepath.Join(ks.dir, name)

	_, err = os.Stat(kp)
	if err == nil {
		return ErrKeyExists
	} else if !os.IsNotExist(err) {
		return err
	}

	fi, err := os.Create(kp)
	if err != nil {
		return err
	}
	defer fi.Close()

	_, err = fi.Write(b)

	return err
}

//get从密钥存储库中检索密钥（如果存在），并返回errnosuchkey
//否则。
func (ks *FSKeystore) Get(name string) (ci.PrivKey, error) {
	if err := validateName(name); err != nil {
		return nil, err
	}

	kp := filepath.Join(ks.dir, name)

	data, err := ioutil.ReadFile(kp)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNoSuchKey
		}
		return nil, err
	}

	return ci.UnmarshalPrivateKey(data)
}

//删除从密钥库中删除密钥
func (ks *FSKeystore) Delete(name string) error {
	if err := validateName(name); err != nil {
		return err
	}

	kp := filepath.Join(ks.dir, name)

	return os.Remove(kp)
}

//list返回键标识符列表
func (ks *FSKeystore) List() ([]string, error) {
	dir, err := os.Open(ks.dir)
	if err != nil {
		return nil, err
	}

	dirs, err := dir.Readdirnames(0)
	if err != nil {
		return nil, err
	}

	list := make([]string, 0, len(dirs))

	for _, name := range dirs {
		err := validateName(name)
		if err == nil {
			list = append(list, name)
		} else {
			log.Warningf("Ignoring the invalid keyfile: %s", name)
		}
	}

	return list, nil
}

