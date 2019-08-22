package jccclient

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/ankadb"
	jccclientdbpb "github.com/zhs007/jccclient/dbpb"
)

// DB -
type DB struct {
	ankaDB ankadb.AnkaDB
}

func makeHostInfoKey(servaddr string, hostname string) string {
	return "hi:" + servaddr + ":" + hostname
}

// NewDB - new DB
func NewDB(dbpath string, httpAddr string, engine string) (*DB, error) {
	cfg := ankadb.NewConfig()

	cfg.AddrHTTP = httpAddr
	cfg.PathDBRoot = dbpath
	cfg.ListDB = append(cfg.ListDB, ankadb.DBConfig{
		Name:   DBName,
		Engine: engine,
		PathDB: DBName,
	})

	ankaDB, err := ankadb.NewAnkaDB(cfg, nil)
	if ankaDB == nil {
		return nil, err
	}

	db := &DB{
		ankaDB: ankaDB,
	}

	return db, err
}

// UpdHostInfo - update hostinfo
func (db *DB) UpdHostInfo(ctx context.Context, servaddr string, hostname string, hi *jccclientdbpb.HostInfo) error {
	if servaddr == "" {
		return ErrNoServAddr
	}

	if hostname == "" {
		return ErrNoHostname
	}

	if hi == nil {
		return ErrNoHostInfo
	}

	buf, err := proto.Marshal(hi)
	if err != nil {
		return err
	}

	err = db.ankaDB.Set(ctx, DBName, makeHostInfoKey(servaddr, hostname), buf)
	if err != nil {
		return err
	}

	return nil
}

// GetHostInfo - get a resource
func (db *DB) GetHostInfo(ctx context.Context, servaddr string, hostname string) (*jccclientdbpb.HostInfo, error) {
	buf, err := db.ankaDB.Get(ctx, DBName, makeHostInfoKey(servaddr, hostname))
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return nil, nil
		}

		return nil, err
	}

	hi := &jccclientdbpb.HostInfo{}

	err = proto.Unmarshal(buf, hi)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
