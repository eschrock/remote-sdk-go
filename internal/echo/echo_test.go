/*
 * Copyright The Titan Project Contributors.
 */
package echo

import (
	"github.com/stretchr/testify/assert"
	"github.com/titan-data/remote-sdk-go/remote"
	"testing"
)

func TestType(t *testing.T) {
	e := EchoRemote{}
	typ, err := e.Type()
	if assert.NoError(t, err) {
		assert.Equal(t, "echo", typ)
	}
}

func TestToURL(t *testing.T) {
	e := EchoRemote{}
	u, props, err := e.ToURL(map[string]interface{}{"a": "b"})
	if assert.NoError(t, err) {
		assert.Equal(t, "echo://echo", u)
		assert.Len(t, props, 1)
		assert.Equal(t, "b", props["a"])
	}
}

func TestFromURL(t *testing.T) {
	e := EchoRemote{}
	res, err := e.FromURL("echo://echo", map[string]string{"a": "b"})
	if assert.NoError(t, err) {
		assert.Len(t, res, 2)
		assert.Equal(t, "b", res["a"])
		assert.Equal(t, "echo://echo", res["url"])
	}
}

func TestGetParameters(t *testing.T) {
	e := EchoRemote{}
	res, err := e.GetParameters(map[string]interface{}{"a": "b"})
	if assert.NoError(t, err) {
		assert.Len(t, res, 1)
		assert.Equal(t, "b", res["a"])
	}
}

func TestValidateRemote(t *testing.T) {
	e := EchoRemote{}
	err := e.ValidateRemote(map[string]interface{}{})
	assert.NoError(t, err)
}

func TestValidateParameters(t *testing.T) {
	e := EchoRemote{}
	err := e.ValidateParameters(map[string]interface{}{})
	assert.NoError(t, err)
}

func TestListCommits(t *testing.T) {
	e := EchoRemote{}
	commits, err := e.ListCommits(map[string]interface{}{}, map[string]interface{}{}, []remote.Tag{})
	if assert.NoError(t, err) {
		assert.Len(t, commits, 2)
		assert.Equal(t, "two", commits[0].Id)
		assert.Equal(t, "two", commits[0].Properties["tags"].(map[string]string)["name"])
		assert.Equal(t, "one", commits[1].Id)
		assert.Equal(t, "one", commits[1].Properties["tags"].(map[string]string)["name"])
	}
}

func TestListCommitsFilter(t *testing.T) {
	e := EchoRemote{}
	search := "one"
	commits, err := e.ListCommits(map[string]interface{}{}, map[string]interface{}{}, []remote.Tag{{Key: "name", Value: &search}})
	if assert.NoError(t, err) {
		assert.Len(t, commits, 1)
		assert.Equal(t, "one", commits[0].Id)
	}
}

func TestGetCommit(t *testing.T) {
	e := EchoRemote{}
	commit, err := e.GetCommit(map[string]interface{}{}, map[string]interface{}{}, "echo")
	if assert.NoError(t, err) {
		assert.Equal(t, "echo", commit.Id)
	}
}

func TestGetMissingCommit(t *testing.T) {
	e := EchoRemote{}
	commit, err := e.GetCommit(map[string]interface{}{}, map[string]interface{}{}, "foo")
	if assert.NoError(t, err) {
		assert.Nil(t, commit)
	}
}
