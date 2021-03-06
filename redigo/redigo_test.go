package redigo

import (
	"os"
	"testing"
)

var c *DB

func raise(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func getConn(t *testing.T) {
	if c == nil {
		rawurl := os.Getenv("redis_url")
		if rawurl == "" {
			t.SkipNow()
		} else {
			db, err := New(rawurl)
			raise(t, err)
			c = db
		}
	}
}

/*
func TestDBString(t *testing.T) {
	getConn(t)

	k := "test"

	ping, err := c.Ping()
	raise(t, err)
	if !ping {
		t.Fatal("ping error")
	}

	ok, err := c.Set(k, "value")
	raise(t, err)
	if !ok {
		t.Fatal("set error")
	}

	v, err := c.Get(k)
	raise(t, err)
	if v != "value" {
		t.Fatal("get error")
	}

	ks, err := c.Keys("*")
	raise(t, err)
	if !inSlice(k, ks) {
		t.Fatal("keys error")
	}

	typ, err := c.Type(k)
	raise(t, err)
	if typ != "string" {
		t.Fatal("string type error")
	}

	has, err := c.Exsits(k)
	raise(t, err)
	if !has {
		t.Fatal("exists error")
	}

	ok, err = c.Expire(k, 10)
	raise(t, err)
	if !ok {
		t.Fatal("expire error")
	}

	ttl, err := c.TTL(k)
	raise(t, err)
	if ttl <= 0 || !(ttl > 5) {
		t.Fatal("ttl error")
	}

	_, err = c.Del(k)
	raise(t, err)
	has, _ = c.Exsits(k)
	if has {
		t.Fatal("del error")
	}
}

func TestList(t *testing.T) {
	getConn(t)

	k := "list"
	v := "v2"
	vt := "top"
	ve := "end"

	_, err := c.Del(k)
	raise(t, err)

	oknum, err := c.RPush(k, vt, v, ve)
	raise(t, err)
	if oknum != 3 {
		t.Fatal("rpush error")
	}

	lv, err := c.LRange(k, 0, -1)
	raise(t, err)
	length, err := c.LLen(k)
	raise(t, err)
	if length != uint64(len(lv)) {
		t.Fatal("list key number should be equal")
	}

	_ve, err := c.RPop(k)
	raise(t, err)
	if _ve != ve {
		t.Fatal("list key value should be equal(rpop)")
	}
	_vt, err := c.LPop(k)
	raise(t, err)
	if _vt != vt {
		t.Fatal("list key value should be equal(lpop)")
	}

	delLen, _ := c.LLen(k)
	if delLen != (length - 2) {
		t.Fatal("RPop/LPop error")
	}

	typ, err := c.Type(k)
	raise(t, err)
	if typ != "list" {
		t.Fatal("list type error")
	}
}

func TestSet(t *testing.T) {
	getConn(t)

	k := "set"
	v := "v3"

	_, err := c.Del(k)
	raise(t, err)

	oknum, err := c.SAdd(k, v, "1", "2")
	raise(t, err)
	if oknum != 3 {
		t.Fatal("sadd error")
	}

	delnum, err := c.SRem(k, "1", "2")
	raise(t, err)
	if delnum != 2 {
		t.Fatal("srem error")
	}

	is, err := c.SIsMember(k, v)
	raise(t, err)
	if !is {
		t.Fatal("sismember error")
	}

	svs, err := c.SMembers(k)
	raise(t, err)
	fmt.Println("smembers", svs)

	length, err := c.SCard(k)
	raise(t, err)
	if length != uint64(len(svs)) {
		t.Fatal("scard not equal smembers")
	}

	typ, err := c.Type(k)
	raise(t, err)
	if typ != "set" {
		t.Fatal("set type error")
	}
}

func TestHash(t *testing.T) {
	getConn(t)

	n := "hash"
	k := "test"
	v := "v4"

	k4 := "hello"
	v4 := "world"

	_, err := c.Del(n)
	raise(t, err)

	oknum, err := c.HSet(n, k, v)
	raise(t, err)
	if oknum != 1 {
		t.Fatal("hset error")
	}

	keysLen, err := c.HLen(n)
	raise(t, err)
	if keysLen != 1 {
		t.Fatal("hlen error")
	}

	_v, err := c.HGet(n, k)
	raise(t, err)
	if _v != v {
		t.Fatal("hget error")
	}

	data := map[string]string{k: v, k4: v4}
	ok, err := c.HMSet(n, data)
	raise(t, err)
	if !ok {
		t.Fatal("hmset error")
	}

	vals, err := c.HVals(n)
	raise(t, err)
	if len(vals) != 2 {
		t.Fatal("hvals error")
	}
	if !inSlice(v, vals) {
		t.Fatal("hvals values error")
	}

	keys, err := c.HKeys(n)
	raise(t, err)
	if len(keys) != 2 {
		t.Fatal("hkeys error")
	}
	if !inSlice(k, keys) {
		t.Fatal("hkeys values error")
	}

	_, err = c.HDel(n, k)
	raise(t, err)
	keys2, _ := c.HKeys(n)
	if len(keys2) != 1 {
		t.Fatal("hdel error")
	}

	typ, err := c.Type(n)
	raise(t, err)
	if typ != "hash" {
		t.Fatal("hash type error")
	}
}
*/
func TestTransaction(t *testing.T) {
	getConn(t)

	name := "pipeline:"

	k0 := "test"
	v0 := "v0"

	k1 := "astring"
	v1 := "v1"

	k2 := "list"
	v2 := "v2"

	k3 := "set"
	v3 := "v3"

	k4 := "deleted-set"
	v4 := "v4"

	k5n := "hash"
	k5k := "field"
	k5v := "v5"

	// 添加数据，测试管道删除操作（未免冲突，添加 KEY 前缀，也是测试）
	c.Prefix = name
	_, err := c.Set(k0, v0)
	raise(t, err)

	_, err = c.SAdd(k4, v4)
	raise(t, err)

	// 管道开始事务
	tc := c.Pipeline()

	err = tc.Del(k0)
	raise(t, err)

	err = tc.Set(k1, v1)
	raise(t, err)

	err = tc.RPush(k2, v2)
	raise(t, err)

	err = tc.SAdd(k3, v3)
	raise(t, err)

	err = tc.SRem(k4, v4)
	raise(t, err)

	err = tc.HSet(k5n, k5k, k5v)
	raise(t, err)

	// 结束管道，执行命令
	_, err = tc.Execute()
	raise(t, err)

	// 测试判定：k0被删除 k4删除v4 k1值为v1 k2长度大于0 v3在k3中 k5n有值
	if has, _ := c.Exsits(k0); has {
		t.Fatal("key should be not found for k0")
	}
	if is, _ := c.SIsMember(k4, v4); is {
		t.Fatal("v4 should be not found in k4")
	}

	if has, _ := c.Exsits(k1); !has {
		t.Fatal("key should be found for k1")
	}
	_v1, err := c.Get(k1)
	raise(t, err)
	if _v1 != v1 {
		t.Fatal("pipe set fail")
	}
	num, _ := c.LLen(k2)
	if num <= 0 {
		t.Fatal("k2 should has values")
	}

	if is, _ := c.SIsMember(k3, v3); !is {
		t.Fatal("v3 should in k3")
	}
	size, _ := c.SCard(k3)
	if size != 1 {
		t.Fatal("k3 length should be equal 1")
	}

	_k5v, err := c.HGet(k5n, k5k)
	raise(t, err)
	if _k5v != k5v {
		t.Fatal("k5 value error for pipe hset")
	}

	tc2 := c.Pipeline()
	tc2.HDel(k5n, k5k)
	_, err = tc2.Execute()
	if err != nil {
		t.Fatal("pipeline hdel error")
	}
}
