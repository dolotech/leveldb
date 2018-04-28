package main

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/golang/glog"
	"flag"
)

func main() {
	flag.Parse()
	db, err := leveldb.OpenFile("./db", nil)

	if err != nil {
		glog.Fatal(err)
	}

	defer db.Close()

	err = db.Put([]byte("key"), []byte("value"), nil)

	if err != nil {
		glog.Error(err)
	}

	data, err := db.Get([]byte("key"), nil)

	if err != nil {
		glog.Error(err)
	}
	glog.Info(string(data))
	err = db.Delete([]byte("key"), nil)

	if err != nil {
		glog.Error(err)
	}

	data, err = db.Get([]byte("key"), nil)

	if err != nil {
		glog.Error(err)
	}
	glog.Info(string(data))

	select {}
}
