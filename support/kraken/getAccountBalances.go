package kraken

import (
	"reflect"

	"github.com/Beldur/kraken-go-api-client"
	"github.com/lightyeario/kelp/support/exchange/assets"
	"github.com/lightyeario/kelp/support/exchange/number"
)

// GetAccountBalances impl.
func (k krakenExchange) GetAccountBalances(assetList []assets.Asset) (map[assets.Asset]number.Number, error) {
	balanceResponse, e := k.api.Balance()
	if e != nil {
		return nil, e
	}

	m := map[assets.Asset]number.Number{}
	for _, a := range assetList {
		krakenAssetString, e := k.assetConverter.ToString(a)
		if e != nil {
			// discard partially built map for now
			return nil, e
		}
		bal := getFieldValue(*balanceResponse, krakenAssetString)
		m[a] = *number.FromFloat(float64(bal))
	}
	return m, nil
}

func getFieldValue(object krakenapi.BalanceResponse, fieldName string) float32 {
	r := reflect.ValueOf(object)
	f := reflect.Indirect(r).FieldByName(fieldName)
	return f.Interface().(float32)
}