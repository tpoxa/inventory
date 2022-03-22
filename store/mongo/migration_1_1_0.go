// Copyright 2022 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package mongo

import (
	"context"

	"github.com/mendersoftware/go-lib-micro/mongo/migrate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	mopts "go.mongodb.org/mongo-driver/mongo/options"

	mstore "github.com/mendersoftware/go-lib-micro/store"
)

type migration_1_1_0 struct {
	ms  *DataStoreMongo
	ctx context.Context
}

func (m *migration_1_1_0) Up(from migrate.Version) error {
	databaseName := mstore.DbFromContext(m.ctx, DbName)
	coll := m.ms.client.Database(databaseName).Collection(DbDevicesColl)
	indexView := coll.Indexes()
	keys := bson.D{
		{Key: DbDevAttributesText, Value: "text"},
	}
	name := DbDevAttributesText
	_, err := indexView.CreateOne(m.ctx, mongo.IndexModel{Keys: keys, Options: &mopts.IndexOptions{
		Name: &name,
	}})
	return err
}

func (m *migration_1_1_0) Version() migrate.Version {
	return migrate.MakeVersion(1, 1, 0)
}
