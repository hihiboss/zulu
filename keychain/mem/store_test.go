package mem

import (
	"testing"

	"github.com/DE-labtory/zulu/keychain"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name    string
	key     keychain.Key
	isValid bool
}{
	{
		name: "valid key 1",
		key: keychain.Key{
			ID: "28762cb4592cbe44f993032bb1656b201aae9b9a9b730f8f1c37e5c597cbd8cc",
			PrivateKey: []byte{
				0x4c, 0x0c, 0x68, 0x33, 0x88, 0x35, 0x94, 0xab,
				0x03, 0x45, 0xeb, 0x13, 0x53, 0x93, 0xb1, 0x64,
				0x9e, 0x80, 0xfb, 0x8a, 0xa6, 0xfe, 0x96, 0x72,
				0x1e, 0xbf, 0x6f, 0x24, 0xf4, 0x54, 0xe0, 0xcc,
			},
			PublicKey: []byte{
				0x62, 0xf2, 0x32, 0x4e, 0x79, 0xb6, 0x78, 0x59,
				0x7b, 0xdd, 0x7b, 0x5a, 0x48, 0xb6, 0x46, 0xdb,
				0xdc, 0x8c, 0x7f, 0x67, 0x28, 0x33, 0xd0, 0xef,
				0x88, 0xc7, 0x5f, 0xad, 0x69, 0xde, 0x55, 0xf0,
				0xa3, 0xe0, 0xf0, 0xc8, 0x14, 0xc8, 0xa1, 0x30,
				0xcb, 0x59, 0xb6, 0x59, 0xd9, 0x48, 0xd7, 0x95,
				0x4c, 0x33, 0xd3, 0xf2, 0x3a, 0x68, 0x59, 0x9a,
				0xad, 0x02, 0x3b, 0x83, 0xd6, 0x61, 0x45, 0xa6,
			},
		},
		isValid: true,
	},
	{
		name: "valid key 2",
		key: keychain.Key{
			ID: "72c03c069490db94f4618d88267403ab8cfbbe1c8593cac50c954edcf8b28273",
			PrivateKey: []byte{
				0x19, 0xe6, 0xa9, 0x0e, 0xcb, 0x53, 0xd9, 0x55,
				0x6e, 0xb0, 0x6e, 0x5e, 0x8a, 0x34, 0x64, 0x7f,
				0x2b, 0x03, 0x08, 0xcf, 0x52, 0x3d, 0x72, 0x59,
				0xe2, 0x5d, 0x4f, 0xd2, 0xc4, 0x2c, 0xa8, 0x2f,
			},
			PublicKey: []byte{
				0xc8, 0x61, 0xd3, 0x4d, 0x3e, 0xd6, 0x69, 0x74,
				0xe9, 0x5e, 0x0f, 0x39, 0xc7, 0xea, 0xd2, 0x90,
				0xce, 0xc8, 0x4e, 0xed, 0x13, 0x7f, 0xfe, 0xa4,
				0x86, 0xde, 0x55, 0x1d, 0xd1, 0x1e, 0x91, 0x5d,
				0xdc, 0xf9, 0x82, 0xa1, 0xc0, 0x18, 0x74, 0xaa,
				0x0d, 0xa3, 0x14, 0xa7, 0x78, 0xb1, 0xa7, 0x21,
				0xd5, 0x16, 0xcd, 0x03, 0xfd, 0x63, 0x53, 0x25,
				0x43, 0xcc, 0x52, 0xd3, 0x37, 0x9d, 0x35, 0x39,
			},
		},
		isValid: true,
	},
	{
		name: "empty id",
		key: keychain.Key{
			ID: "",
			PrivateKey: []byte{
				0x4c, 0x0c, 0x68, 0x33, 0x88, 0x35, 0x94, 0xab,
				0x03, 0x45, 0xeb, 0x13, 0x53, 0x93, 0xb1, 0x64,
				0x9e, 0x80, 0xfb, 0x8a, 0xa6, 0xfe, 0x96, 0x72,
				0x1e, 0xbf, 0x6f, 0x24, 0xf4, 0x54, 0xe0, 0xcc,
			},
			PublicKey: []byte{
				0x62, 0xf2, 0x32, 0x4e, 0x79, 0xb6, 0x78, 0x59,
				0x7b, 0xdd, 0x7b, 0x5a, 0x48, 0xb6, 0x46, 0xdb,
				0xdc, 0x8c, 0x7f, 0x67, 0x28, 0x33, 0xd0, 0xef,
				0x88, 0xc7, 0x5f, 0xad, 0x69, 0xde, 0x55, 0xf0,
				0xa3, 0xe0, 0xf0, 0xc8, 0x14, 0xc8, 0xa1, 0x30,
				0xcb, 0x59, 0xb6, 0x59, 0xd9, 0x48, 0xd7, 0x95,
				0x4c, 0x33, 0xd3, 0xf2, 0x3a, 0x68, 0x59, 0x9a,
				0xad, 0x02, 0x3b, 0x83, 0xd6, 0x61, 0x45, 0xa6,
			},
		},
		isValid: false,
	},
	{
		name: "invalid id length",
		key: keychain.Key{
			ID: "62cb4592cbe44f993032bb1656b201aae9b9a9b730f8f1c37e5c597cbd8cc",
			PrivateKey: []byte{
				0x4c, 0x0c, 0x68, 0x33, 0x88, 0x35, 0x94, 0xab,
				0x03, 0x45, 0xeb, 0x13, 0x53, 0x93, 0xb1, 0x64,
				0x9e, 0x80, 0xfb, 0x8a, 0xa6, 0xfe, 0x96, 0x72,
				0x1e, 0xbf, 0x6f, 0x24, 0xf4, 0x54, 0xe0, 0xcc,
			},
			PublicKey: []byte{
				0x62, 0xf2, 0x32, 0x4e, 0x79, 0xb6, 0x78, 0x59,
				0x7b, 0xdd, 0x7b, 0x5a, 0x48, 0xb6, 0x46, 0xdb,
				0xdc, 0x8c, 0x7f, 0x67, 0x28, 0x33, 0xd0, 0xef,
				0x88, 0xc7, 0x5f, 0xad, 0x69, 0xde, 0x55, 0xf0,
				0xa3, 0xe0, 0xf0, 0xc8, 0x14, 0xc8, 0xa1, 0x30,
				0xcb, 0x59, 0xb6, 0x59, 0xd9, 0x48, 0xd7, 0x95,
				0x4c, 0x33, 0xd3, 0xf2, 0x3a, 0x68, 0x59, 0x9a,
				0xad, 0x02, 0x3b, 0x83, 0xd6, 0x61, 0x45, 0xa6,
			},
		},
		isValid: false,
	},
	{
		name: "invalid id format",
		key: keychain.Key{
			ID: "^@762cb4592cbe44f993032bb1656b201aae9b9a9b730f8f1c37e5c597cbd8cc",
			PrivateKey: []byte{
				0x4c, 0x0c, 0x68, 0x33, 0x88, 0x35, 0x94, 0xab,
				0x03, 0x45, 0xeb, 0x13, 0x53, 0x93, 0xb1, 0x64,
				0x9e, 0x80, 0xfb, 0x8a, 0xa6, 0xfe, 0x96, 0x72,
				0x1e, 0xbf, 0x6f, 0x24, 0xf4, 0x54, 0xe0, 0xcc,
			},
			PublicKey: []byte{
				0x62, 0xf2, 0x32, 0x4e, 0x79, 0xb6, 0x78, 0x59,
				0x7b, 0xdd, 0x7b, 0x5a, 0x48, 0xb6, 0x46, 0xdb,
				0xdc, 0x8c, 0x7f, 0x67, 0x28, 0x33, 0xd0, 0xef,
				0x88, 0xc7, 0x5f, 0xad, 0x69, 0xde, 0x55, 0xf0,
				0xa3, 0xe0, 0xf0, 0xc8, 0x14, 0xc8, 0xa1, 0x30,
				0xcb, 0x59, 0xb6, 0x59, 0xd9, 0x48, 0xd7, 0x95,
				0x4c, 0x33, 0xd3, 0xf2, 0x3a, 0x68, 0x59, 0x9a,
				0xad, 0x02, 0x3b, 0x83, 0xd6, 0x61, 0x45, 0xa6,
			},
		},
		isValid: false,
	},
	{
		name: "invalid private key length",
		key: keychain.Key{
			ID: "12762cb4592cbe44f993032bb1656b201aae9b9a9b730f8f1c37e5c597cbd8cc",
			PrivateKey: []byte{
				0x4c, 0x0c, 0x68, 0x33, 0x88, 0x35, 0x94, 0xab,
				0x03, 0x45, 0xeb, 0x13, 0x53, 0x93, 0xb1, 0x64,
				0x9e, 0x80, 0xfb, 0x8a, 0xa6, 0xfe, 0x96, 0x72,
				0x1e, 0xbf, 0x6f, 0x24, 0xf4, 0x54, 0xe0,
			},
			PublicKey: []byte{
				0x62, 0xf2, 0x32, 0x4e, 0x79, 0xb6, 0x78, 0x59,
				0x7b, 0xdd, 0x7b, 0x5a, 0x48, 0xb6, 0x46, 0xdb,
				0xdc, 0x8c, 0x7f, 0x67, 0x28, 0x33, 0xd0, 0xef,
				0x88, 0xc7, 0x5f, 0xad, 0x69, 0xde, 0x55, 0xf0,
				0xa3, 0xe0, 0xf0, 0xc8, 0x14, 0xc8, 0xa1, 0x30,
				0xcb, 0x59, 0xb6, 0x59, 0xd9, 0x48, 0xd7, 0x95,
				0x4c, 0x33, 0xd3, 0xf2, 0x3a, 0x68, 0x59, 0x9a,
				0xad, 0x02, 0x3b, 0x83, 0xd6, 0x61, 0x45, 0xa6,
			},
		},
		isValid: false,
	},
	{
		name: "invalid public key length",
		key: keychain.Key{
			ID: "34762cb4592cbe44f993032bb1656b201aae9b9a9b730f8f1c37e5c597cbd8cc",
			PrivateKey: []byte{
				0x4c, 0x0c, 0x68, 0x33, 0x88, 0x35, 0x94, 0xab,
				0x03, 0x45, 0xeb, 0x13, 0x53, 0x93, 0xb1, 0x64,
				0x9e, 0x80, 0xfb, 0x8a, 0xa6, 0xfe, 0x96, 0x72,
				0x1e, 0xbf, 0x6f, 0x24, 0xf4, 0x54, 0xe0, 0xcc,
			},
			PublicKey: []byte{
				0x62, 0xf2, 0x32, 0x4e, 0x79, 0xb6, 0x78, 0x59,
				0x7b, 0xdd, 0x7b, 0x5a, 0x48, 0xb6, 0x46, 0xdb,
				0xdc, 0x8c, 0x7f, 0x67, 0x28, 0x33, 0xd0, 0xef,
				0x88, 0xc7, 0x5f, 0xad, 0x69, 0xde, 0x55, 0xf0,
				0xa3, 0xe0, 0xf0, 0xc8, 0x14, 0xc8, 0xa1, 0x30,
				0xcb, 0x59, 0xb6, 0x59, 0xd9, 0x48, 0xd7, 0x95,
				0x4c, 0x33, 0xd3, 0xf2, 0x3a, 0x68, 0x59, 0x9a,
				0xad, 0x02, 0x3b, 0x83, 0xd6, 0x61, 0x45,
			},
		},
		isValid: false,
	},
}

func TestStore(t *testing.T) {
	ks := NewKeyStore()

	for _, test := range tests {
		k := test.key

		err := ks.Store(k)
		if err != nil {
			if test.isValid {
				t.Errorf("[%s] store a valid key to mem store failed: %v", test.name, err)
			}
			continue
		}
		if !test.isValid {
			t.Errorf("[%s] it should fail but a invalid key has been stored", test.name)
		}
	}
}

func TestStoreKeyFromGenerator(t *testing.T) {
	ks := NewKeyStore()
	g := keychain.NewKeyGenerator()

	k, err := g.Generate()
	assert.NoError(t, err)

	err = ks.Store(k)
	assert.NoError(t, err)
	storedKey, err := ks.Get(k.ID)
	assert.NoError(t, err)

	assert.Equal(t, storedKey, k)
}

func TestGet(t *testing.T) {
	ks := NewKeyStore()

	for _, test := range tests {
		k := test.key

		_ = ks.Store(k)

		storedKey, err := ks.Get(k.ID)
		if err != nil {
			if test.isValid {
				t.Errorf("[%s] valid key couldn't retrieved from mem store: %v", test.name, err)
			}
			continue
		}
		if !test.isValid {
			t.Errorf("[%s] it should fail but a invalid key has been retrieved", test.name)
		}

		assert.Equal(t, storedKey, test.key)
	}
}

func TestGetAll(t *testing.T) {
	ks := NewKeyStore()
	var expected []keychain.Key

	// case: empty key store
	_, err := ks.GetAll()
	assert.EqualError(t, err, "key store is empty")

	// case: key store has 1 key-value pair
	err = ks.Store(tests[0].key)
	expected = append(expected, tests[0].key)
	assert.NoError(t, err)

	keys, err := ks.GetAll()
	assert.NoError(t, err)
	assert.ElementsMatch(t, keys, expected)

	// case: key store has 2 key-value pair
	err = ks.Store(tests[1].key)
	expected = append(expected, tests[1].key)

	keys, err = ks.GetAll()
	assert.NoError(t, err)
	assert.ElementsMatch(t, keys, expected)
}
