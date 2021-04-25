package indexdb

import (
	"fmt"

	"github.com/realPy/jswasm"
	"github.com/realPy/jswasm/js"
)

type IndexDB struct {
	dbObject js.Value
}

func OpenIndexDB(name string) (IndexDB, error) {
	var indexdb IndexDB

	if window, err := js.Global().GetWithErr("window"); err == nil {

		if indexedDB, err := window.GetWithErr("indexedDB"); err == nil {

			if waitableOpen, err := indexedDB.CallWithErr("open", js.ValueOf("test")); err == nil {
				ch := jswasm.OnSuccessFailure(waitableOpen)
				results := <-ch
				if results.Success {
					indexdb.dbObject = results.Payload[0]
					return indexdb, nil
				} else {
					return IndexDB{}, fmt.Errorf("Open IndexDB failed")
				}

			} else {
				return IndexDB{}, err
			}
		} else {
			return IndexDB{}, err
		}

	} else {
		return IndexDB{}, err
	}

}
