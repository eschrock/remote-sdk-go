/*
 * Copyright The Titan Project Contributors.
 */
package echo

import "github.com/titan-data/remote-sdk-go/remote"

type EchoRemote struct {
}

func (m EchoRemote) Type() (string, error) {
	return "echo", nil
}

func (m EchoRemote) FromURL(url string, additionalProperties map[string]string) (map[string]interface{}, error) {
	ret := map[string]interface{}{
		"url": url,
	}
	for k, v := range additionalProperties {
		ret[k] = v
	}
	return ret, nil
}

func (m EchoRemote) ToURL(properties map[string]interface{}) (string, map[string]string, error) {
	ret := map[string]string{}
	for k, v := range properties {
		ret[k] = v.(string)
	}
	return "echo://echo", ret, nil
}

func (m EchoRemote) GetParameters(remoteProperties map[string]interface{}) (map[string]interface{}, error) {
	return remoteProperties, nil
}

func (m EchoRemote) ValidateRemote(properties map[string]interface{}) error {
	return nil
}

func (m EchoRemote) ValidateParameters(parameters map[string]interface{}) error {
	return nil
}

func (m EchoRemote) ListCommits(properties map[string]interface{}, parameters map[string]interface{}, tags []remote.Tag) ([]remote.Commit, error) {
	panic("implement me")
}

func (m EchoRemote) GetCommit(properties map[string]interface{}, parameters map[string]interface{}, commitId string) (*remote.Commit, error) {
	panic("implement me")
}
