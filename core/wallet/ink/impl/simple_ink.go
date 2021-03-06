/*
Copyright Ziggurat Corp. 2017 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package impl

import "math/big"

type SimpleInkAlg struct {
	inkRatio float32
}

func NewSimpleInkAlg() *SimpleInkAlg {
	return &SimpleInkAlg{inkRatio: 0.01}
}
func (s *SimpleInkAlg) CalcInk(textLength int) (*big.Int, error) {
	ink := big.NewInt(int64(float32(textLength) * s.inkRatio))
	return ink, nil
}
