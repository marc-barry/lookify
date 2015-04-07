package util

import (
	"bufio"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/Shopify/exabgp-util/types"
)

func ReadMessage(r io.Reader) (*types.ExaBGPReaderMessage, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	m, err := types.UnmarshalExaBGPReaderMessage(b)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func ScanMessage(r io.Reader) <-chan *types.ExaBGPReaderMessage {
	scanner := bufio.NewScanner(r)
	mChan := make(chan *types.ExaBGPReaderMessage)

	go func() {
		for scanner.Scan() {
			var message types.ExaBGPReaderMessage
			err := json.Unmarshal(scanner.Bytes(), &message)
			if err == nil {
				mChan <- &message
			}
		}
		close(mChan)
	}()

	return mChan
}
