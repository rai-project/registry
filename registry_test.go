package registry

import (
	"os"
	"testing"
	"time"

	"github.com/rai-project/config"
	"github.com/rai-project/libkv/store"
	"github.com/stretchr/testify/assert"
)

func TestOperations(t *testing.T) {
	st, err := New()
	assert.NoError(t, err)
	assert.NotEmpty(t, st)

	defer st.Close()

	key := "testing"
	value := []byte("TestOperations")

	err = st.Put(key, value, &store.WriteOptions{IsDir: false, TTL: 2 * time.Minute})
	assert.NoError(t, err)

	kv, err := st.Get(key)
	assert.NoError(t, err)
	assert.NotEmpty(t, kv)
	assert.Equal(t, string(value), string(kv.Value))

	lst, err := st.List("/")
	assert.NoError(t, err)
	assert.NotEmpty(t, lst)

	err = st.Delete(key)
	assert.NoError(t, err)
}

func TestMain(m *testing.M) {
	config.Init(
		config.AppName("carml"),
		config.DebugMode(true),
		config.VerboseMode(true),
	)
	os.Exit(m.Run())
}
