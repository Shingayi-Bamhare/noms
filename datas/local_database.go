// Copyright 2016 The Noms Authors. All rights reserved.
// Licensed under the Apache License, version 2.0:
// http://www.apache.org/licenses/LICENSE-2.0

package datas

import (
	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/hash"
	"github.com/attic-labs/noms/types"
)

// Database provides versioned storage for noms values. Each Database instance represents one moment in history. Heads() returns the Commit from each active fork at that moment. The Commit() method returns a new Database, representing a new moment in history.
type LocalDatabase struct {
	cch *cachingChunkHaver
	databaseCommon
}

func newLocalDatabase(cs chunks.ChunkStore) *LocalDatabase {
	return &LocalDatabase{
		newCachingChunkHaver(cs),
		newDatabaseCommon(types.NewValueStore(types.NewBatchStoreAdaptor(cs)), cs),
	}
}

func (lds *LocalDatabase) Commit(datasetID string, commit types.Struct) (Database, error) {
	err := lds.commit(datasetID, commit)
	lds.vs.Flush()
	return &LocalDatabase{
		lds.cch,
		newDatabaseCommon(lds.vs, lds.rt),
	}, err
}

func (lds *LocalDatabase) Delete(datasetID string) (Database, error) {
	err := lds.doDelete(datasetID)
	lds.vs.Flush()
	return &LocalDatabase{
		lds.cch,
		newDatabaseCommon(lds.vs, lds.rt),
	}, err
}

func (lds *LocalDatabase) has(r hash.Hash) bool {
	return lds.cch.Has(r)
}

func (lds *LocalDatabase) batchSink() batchSink {
	return lds.vs.BatchStore()
}

func (lds *LocalDatabase) batchStore() types.BatchStore {
	return lds.vs.BatchStore()
}
