/*

  Copyright 2017 Loopring Project Ltd (Loopring Foundation).

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

package matchengine

import (
	"github.com/Loopring/ringminer/types"
	"math/big"
)

type LegalCurrency int

const (
	_ LegalCurrency = iota
	CNY
	USD
)

const (
	LRC_ADDRESS = "0x5132a8ce9a61b13b9cAEcd2261abF95323056423"
)

//todo:获取法币汇率
func GetLegalRate(currency LegalCurrency, tokenAddress types.Address) *types.EnlargedInt {
	return &types.EnlargedInt{Value:big.NewInt(100), Decimals:big.NewInt(100)}
}
