package kit

import (
	"math/rand"
	"strconv"
	"time"
)

const Letter = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func RandString(count int) string {
	if count <= 0 {
		return ""
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	sets := []byte(Letter)
	var ret = make([]byte, count)
	setLen := len(Letter)
	for i := 0; i < count; i++ {
		ret[i] = sets[r.Intn(setLen-1)]
	}
	return string(ret)
}

func UUID(prefix string) string {
	return prefix + "_" + strconv.Itoa(time.Now().Second()) + strconv.Itoa(time.Now().Nanosecond()) + RandString(1)
}
