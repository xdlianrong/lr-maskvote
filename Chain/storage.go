//用于和链进行交互（模拟），主要负责存储
package Chain

import "errors"

type EthDemo interface {
	Set(int, string) error
	Get(int) (string, error)
}

//存储公钥的map
type PubKey struct {
	pubKeyMap map[int]string
}

//公开系数承诺的map
type FactorCommit struct {
	factorCommitMap map[int]string
}

//各节点间发送密文的map
type EncryptSecret struct {
	secretMap map[int]map[int]string
}

//公开对

func (pubKey *PubKey) New() {
	pubKey.pubKeyMap = make(map[int]string)
}

func (pubKey *PubKey) Set(index int, publicKey string) error {
	pubKey.pubKeyMap[index] = publicKey
	return nil
}

func (pubKey *PubKey) Get(index int) (string, error) {
	publicKey, ok := pubKey.pubKeyMap[index]
	if ok {
		return publicKey, nil
	}
	return "", errors.New("the public key is not found")
}

func (factorCommit *FactorCommit) New() {
	factorCommit.factorCommitMap = make(map[int]string)
}

func (factorCommit *FactorCommit) Set(index int, commit string) error {
	factorCommit.factorCommitMap[index] = commit
	return nil
}

func (factorCommit *FactorCommit) Get(index int) (string, error) {
	commit, ok := factorCommit.factorCommitMap[index]
	if ok {
		return commit, nil
	}
	return "",errors.New("the commit is not found")
}

func (encryptSecret *EncryptSecret)New() {
	encryptSecret.secretMap = make(map[int]map[int]string)
}

func (encryptSecret *EncryptSecret)Set(from int, to int, secret string) error {
	encryptSecret.secretMap[from][to] = secret
	return nil
}

func (encryptSecret *EncryptSecret)Get(from int, to int) (string,error) {
	secret,ok := encryptSecret.secretMap[from][to]
	if ok {
		return secret, nil
	}
	return "",errors.New("the secret is not found")
}