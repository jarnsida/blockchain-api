package handlerblock

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/evt/blockchain-api/internal/pkg/models"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
)

// TestGet tests Get
func TestGet(t *testing.T) {
	const (
		testBlockNumber int64 = 1
	)

	log.SetOutput(ioutil.Discard)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	blockService := NewMockBlockService(ctrl)
	blockHandler := New(blockService)

	app := fiber.New()
	app.Get("/blocks/:id", blockHandler.Get)

	tests := []struct {
		name   string
		path   string
		expect func()
		assert func([]byte)
	}{
		{
			name: "by block number (success)",
			path: "/blocks/1",
			expect: func() {
				blockService.EXPECT().GetBlockByNumber(gomock.Any(), big.NewInt(testBlockNumber)).Return(&models.Block{
					Header: &types.Header{
						ParentHash:  common.HexToHash("0x3dd4dc843801af12c0a6dd687642467a3ce835dca09159734dec03109a1c1f1f"),
						UncleHash:   common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
						Coinbase:    common.HexToAddress("0xc2fa6dcef5a1fbf70028c5636e7f64cd46e7cfd4"),
						Root:        common.HexToHash("0xf5f18c33ddff06efa928d22a2432fb34a11e6f62cce825cdad1c78e1068e6b7b"),
						TxHash:      common.HexToHash("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"),
						ReceiptHash: common.HexToHash("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"),
						Difficulty:  big.NewInt(827755),
						Number:      big.NewInt(100),
						GasLimit:    15217318,
						GasUsed:     0,
						Time:        1479653850,
						Extra:       []byte("0xd783010502846765746887676f312e362e33856c696e7578"),
						MixDigest:   common.HexToHash("0x3172866e675b057a294d3f474e9141b588d5a0c622b4d8049e272c6a001e9c4e"),
						Nonce:       types.EncodeNonce(7892755374462247712),
					},
				}, nil)
			},
			assert: func(content []byte) {
				assert.Equal(t, []byte(`{"Header":{"parentHash":"0x3dd4dc843801af12c0a6dd687642467a3ce835dca09159734dec03109a1c1f1f","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","miner":"0xc2fa6dcef5a1fbf70028c5636e7f64cd46e7cfd4","stateRoot":"0xf5f18c33ddff06efa928d22a2432fb34a11e6f62cce825cdad1c78e1068e6b7b","transactionsRoot":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","receiptsRoot":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","difficulty":"0xca16b","number":"0x64","gasLimit":"0xe832a6","gasUsed":"0x0","timestamp":"0x5831b9da","extraData":"0x3078643738333031303530323834363736353734363838373637366633313265333632653333383536633639366537353738","mixHash":"0x3172866e675b057a294d3f474e9141b588d5a0c622b4d8049e272c6a001e9c4e","nonce":"0x6d88b33209e0a320","baseFeePerGas":null,"hash":"0x61004ffd9447fac2f9849cd75f362d68a73c97766b121e9fe8cecc7e92cec5a5"}}`), content)
			},
		},
		{
			name: "by block number (error)",
			path: "/blocks/1",
			expect: func() {
				blockService.EXPECT().GetBlockByNumber(gomock.Any(), big.NewInt(testBlockNumber)).Return(nil, errors.New("test error"))
			},
			assert: func(content []byte) {
				assert.Equal(t, []byte(`{"error":"test error"}`), content)
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if tc.expect != nil {
				tc.expect()
			}

			req := httptest.NewRequest("GET", tc.path, nil)

			res, err := app.Test(req, 1000)
			if err != nil {
				t.Error(err)
			}
			defer res.Body.Close()

			content, err := io.ReadAll(res.Body)
			if err != nil {
				t.Error(err)
			}

			if tc.assert != nil {
				tc.assert(content)
			}
		})
	}
}
