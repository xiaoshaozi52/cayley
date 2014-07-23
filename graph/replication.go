// Copyright 2014 The Cayley Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graph

// Defines the interface for consistent replication of a graph instance.
//
// Separate from the backend, this dictates how individual triples get
// identified and replicated consistently across (potentially) multiple
// instances. The simplest case is to keep an append-only log of triple
// changes.

type Procedure byte

const (
	Add Procedure = iota
	Delete
)

type Transaction struct {
	Id     int64
	Triple *Triple
	Action Procedure
}

type Replication interface {
	// Get a unique range of triple IDs from the Replication strategy.
	// Returns the inclusive set of unique ids.
	AcquireNextIds(size int64) (start int64, end int64)

	// Returns the highest current ID.
	GetLastId() int64

	// Sends the transactions to the replicas.
	Replicate([]*Transaction)

	// Attempt to acquire the given range of triples from other replicated sources.
	RequestTransactionRange(start int64, end int64)

	// Opens the replication interface.
	Open(TripleStore, Options)
}
