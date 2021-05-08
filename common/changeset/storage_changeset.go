package changeset

import (
	"bytes"
	"encoding/binary"
	"errors"
	"sort"

	"github.com/ledgerwatch/turbo-geth/common"
	"github.com/ledgerwatch/turbo-geth/common/dbutils"
	"github.com/ledgerwatch/turbo-geth/common/etl"
	"github.com/ledgerwatch/turbo-geth/ethdb"
)

const (
	DefaultIncarnation = uint64(1)
)

var (
	ErrNotFound  = errors.New("not found")
	ErrFindValue = errors.New("find value error")
)

func NewStorageChangeSet() *ChangeSet {
	return &ChangeSet{
		Changes: make([]Change, 0),
		keyLen:  common.AddressLength + common.HashLength + common.IncarnationLength,
	}
}

func EncodeStorage(blockN uint64, s *ChangeSet, f func(k, v []byte) error) error {
	sort.Sort(s)
	keyPart := common.AddressLength + common.IncarnationLength
	for _, cs := range s.Changes {
		newK := make([]byte, common.BlockNumberLength+keyPart)
		binary.BigEndian.PutUint64(newK, blockN)
		copy(newK[8:], cs.Key[:keyPart])
		newV := make([]byte, 0, common.HashLength+len(cs.Value))
		newV = append(append(newV, cs.Key[keyPart:]...), cs.Value...)
		if err := f(newK, newV); err != nil {
			return err
		}
	}
	return nil
}

func DecodeStorage(dbKey, dbValue []byte) (uint64, []byte, []byte) {
	blockN := binary.BigEndian.Uint64(dbKey)
	k := make([]byte, common.AddressLength+common.IncarnationLength+common.HashLength)
	dbKey = dbKey[common.BlockNumberLength:] // remove BlockN bytes
	copy(k, dbKey)
	copy(k[len(dbKey):], dbValue[:common.HashLength])
	v := dbValue[common.HashLength:]
	if len(v) == 0 {
		v = nil
	}

	return blockN, k, v
}

type StorageChangeSet struct{ c ethdb.CursorDupSort }

func (b StorageChangeSet) Find(blockNumber uint64, k []byte) ([]byte, error) {
	return findWithoutIncarnationInStorageChangeSet2(b.c, blockNumber, common.AddressLength, k[:common.AddressLength], k[common.AddressLength:])
}

func (b StorageChangeSet) FindWithIncarnation(blockNumber uint64, k []byte) ([]byte, error) {
	return doSearch2(
		b.c, blockNumber,
		k[:common.AddressLength],
		k[common.AddressLength+common.IncarnationLength:common.AddressLength+common.HashLength+common.IncarnationLength],
		binary.BigEndian.Uint64(k[common.AddressLength:]), /* incarnation */
	)
}

func (b StorageChangeSet) FindWithoutIncarnation(blockNumber uint64, addressToFind []byte, keyToFind []byte) ([]byte, error) {
	return findWithoutIncarnationInStorageChangeSet2(b.c, blockNumber, common.AddressLength, addressToFind, keyToFind)
}

func findWithoutIncarnationInStorageChangeSet2(c ethdb.CursorDupSort, blockNumber uint64, keyPrefixLen int, addrBytesToFind []byte, keyBytesToFind []byte) ([]byte, error) {
	return doSearch2(
		c, blockNumber,
		addrBytesToFind,
		keyBytesToFind,
		0, /* incarnation */
	)
}

func doSearch2(
	c ethdb.CursorDupSort,
	blockNumber uint64,
	addrBytesToFind []byte,
	keyBytesToFind []byte,
	incarnation uint64,
) ([]byte, error) {
	keyPrefixLen := common.AddressLength
	if incarnation == 0 {
		seek := make([]byte, 8+keyPrefixLen)
		binary.BigEndian.PutUint64(seek, blockNumber)
		copy(seek[8:], addrBytesToFind)
		for k, v, err := c.Seek(seek); k != nil; k, v, err = c.Next() {
			if err != nil {
				return nil, err
			}
			_, k, v = DecodeStorage(k, v)
			if !bytes.HasPrefix(k, addrBytesToFind) {
				return nil, ErrNotFound
			}

			stHash := k[keyPrefixLen+common.IncarnationLength:]
			if bytes.Equal(stHash, keyBytesToFind) {
				return v, nil
			}
		}
		return nil, ErrNotFound
	}

	seek := make([]byte, common.BlockNumberLength+keyPrefixLen+common.IncarnationLength)
	binary.BigEndian.PutUint64(seek, blockNumber)
	copy(seek[8:], addrBytesToFind)
	binary.BigEndian.PutUint64(seek[common.BlockNumberLength+keyPrefixLen:], incarnation)
	k := seek
	v, err := c.SeekBothRange(seek, keyBytesToFind)
	if err != nil {
		return nil, err
	}
	if !bytes.HasPrefix(v, keyBytesToFind) {
		return nil, ErrNotFound
	}
	_, _, v = DecodeStorage(k, v)
	return v, nil
}

// RewindDataPlain generates rewind data for all plain buckets between the timestamp
// timestapSrc is the current timestamp, and timestamp Dst is where we rewind
func RewindData(db ethdb.Tx, timestampSrc, timestampDst uint64, tmpdir string, quit <-chan struct{}) (*etl.Collector, error) {
	// Collect list of buckets and keys that need to be considered

	changes := etl.NewCollector(tmpdir, etl.NewOldestEntryBuffer(etl.BufferOptimalSize))

	if err := walkAndCollect(
		changes.Collect,
		db, dbutils.AccountChangeSetBucket,
		timestampDst+1, timestampSrc,
		quit,
	); err != nil {
		return nil, err
	}

	if err := walkAndCollect(
		changes.Collect,
		db, dbutils.StorageChangeSetBucket,
		timestampDst+1, timestampSrc,
		quit,
	); err != nil {
		return nil, err
	}

	return changes, nil
}

func walkAndCollect(collectorFunc func([]byte, []byte) error, db ethdb.Tx, bucket string, timestampDst, timestampSrc uint64, quit <-chan struct{}) error {
	c, err := db.Cursor(bucket)
	if err != nil {
		return err
	}
	defer c.Close()
	return ethdb.Walk(c, dbutils.EncodeBlockNumber(timestampDst), 0, func(dbKey, dbValue []byte) (bool, error) {
		if err := common.Stopped(quit); err != nil {
			return false, err
		}
		timestamp, k, v := Mapper[bucket].Decode(dbKey, dbValue)
		if timestamp > timestampSrc {
			return false, nil
		}
		if innerErr := collectorFunc(common.CopyBytes(k), common.CopyBytes(v)); innerErr != nil {
			return false, innerErr
		}
		return true, nil
	})
}