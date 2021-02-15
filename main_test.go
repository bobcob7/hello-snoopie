package main

import (
	"crypto/rand"
	"fmt"
	"testing"
)

func Test_getNonce(t *testing.T) {
	for i:=100;i>=0;i--{
		t.Run(fmt.Sprintf("Len %d", i), func(t *testing.T) {
			nonce := getNonce(rand.Reader, i)
			if len(nonce) != i {
				t.Errorf("getNonce() wrong size = %v", nonce)
			}
		})
	}
}
