// Copyright 2021 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package trie

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func generateRandBytes(size int) []byte {
	buf := make([]byte, rand.Intn(size)+1)
	rand.Read(buf)
	return buf
}

func generateRand(size int) [][]byte {
	rt := make([][]byte, size)
	for i := range rt {
		buf := make([]byte, rand.Intn(379)+1)
		rand.Read(buf)
		rt[i] = buf
	}
	return rt
}

func TestHashLeaf(t *testing.T) {
	n := &leaf{key: generateRandBytes(380), value: generateRandBytes(64)}

	buffer := bytes.NewBuffer(nil)
	const parallel = false

	err := encodeNode(n, buffer, parallel)

	require.NoError(t, err)
	assert.NotZero(t, buffer.Len())
}

func TestHashBranch(t *testing.T) {
	n := &branch{key: generateRandBytes(380), value: generateRandBytes(380)}
	n.children[3] = &leaf{key: generateRandBytes(380), value: generateRandBytes(380)}

	buffer := bytes.NewBuffer(nil)
	const parallel = false

	err := encodeNode(n, buffer, parallel)

	require.NoError(t, err)
	assert.NotZero(t, buffer.Len())
}

func TestHashShort(t *testing.T) {
	n := &leaf{
		key:   generateRandBytes(2),
		value: generateRandBytes(3),
	}

	encodingBuffer := bytes.NewBuffer(nil)
	const parallel = false
	err := encodeNode(n, encodingBuffer, parallel)
	require.NoError(t, err)

	digestBuffer := bytes.NewBuffer(nil)
	err = hashNode(n, digestBuffer)
	require.NoError(t, err)
	assert.Equal(t, encodingBuffer.Bytes(), digestBuffer.Bytes())
}

func Test_encodeChildsSequentially(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		children [16]node
		expected []byte
		err      error
	}{
		"nil children": {},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			buffer := bytes.NewBuffer(nil)
			err := encodeChildsSequentially(testCase.children, buffer)

			if testCase.err != nil {
				assert.EqualError(t, err, testCase.err.Error())
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, testCase.expected, buffer.Bytes())
		})
	}
}

func Test_encodeChild(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		child    node
		expected []byte
		err      error
	}{
		"nil node": {},
		"nil leaf": {
			child: (*leaf)(nil),
		},
		"nil branch": {
			child: (*branch)(nil),
		},
		"empty leaf child": {
			child:    &leaf{},
			expected: []byte{0x8, 0x40, 0x0},
		},
		"empty branch child": {
			child:    &branch{},
			expected: []byte{0xc, 0x80, 0x0, 0x0},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			buffer := bytes.NewBuffer(nil)
			err := encodeChild(testCase.child, buffer)

			if testCase.err != nil {
				assert.EqualError(t, err, testCase.err.Error())
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, testCase.expected, buffer.Bytes())
		})
	}
}
