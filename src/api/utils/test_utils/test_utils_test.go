package test_utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

var (
	a = AtomicInt{}
	counter = 0
	lock sync.Mutex
)

type AtomicInt struct {
	value int
	lock sync.Mutex
}

func(i *AtomicInt) Increase() {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.value ++
}

func(i *AtomicInt) Decrease() {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.value--
}

func(i *AtomicInt) Value() int {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.value
}

func TestGetMockedContext(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/something", nil)
	request.Header = http.Header{"X-Mock": {"true"}}
	c := GetMockedContext(request, response)
	assert.NotNil(t, c)
	assert.EqualValues(t, http.MethodGet, c.Request.Method)
	assert.EqualValues(t, "8080", c.Request.URL.Port())
	assert.EqualValues(t, "/something", c.Request.URL.Path)
	assert.EqualValues(t, "http", c.Request.URL.Scheme)
	assert.EqualValues(t, 1, len(c.Request.Header))
	assert.EqualValues(t, "true", c.GetHeader("X-Mock"))
}

func TestMutex(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go updateAtomic(&wg)
	}
	wg.Wait()
	fmt.Println(a.Value())
	fmt.Println(counter)

}

func updateAtomic(w *sync.WaitGroup) {
	defer lock.Unlock()
	lock.Lock()
	a.Increase()
	counter++
	w.Done()
}