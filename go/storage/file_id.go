package storage

import (
	"code.google.com/p/weed-fs/go/util"
	"encoding/hex"
	"strings"
  "code.google.com/p/weed-fs/go/glog"
)

type FileId struct {
	VolumeId VolumeId
	Key      uint64
	Hashcode uint32
}

func NewFileId(VolumeId VolumeId, Key uint64, Hashcode uint32) *FileId {
	return &FileId{VolumeId: VolumeId, Key: Key, Hashcode: Hashcode}
}
func ParseFileId(fid string) *FileId {
	a := strings.Split(fid, ",")
	if len(a) != 2 {
		glog.V(1).Infoln("Invalid fid", fid, ", split length", len(a))
		return nil
	}
	vid_string, key_hash_string := a[0], a[1]
	volumeId, _ := NewVolumeId(vid_string)
	key, hash := ParseKeyHash(key_hash_string)
	return &FileId{VolumeId: volumeId, Key: key, Hashcode: hash}
}
func (n *FileId) String() string {
	bytes := make([]byte, 12)
	util.Uint64toBytes(bytes[0:8], n.Key)
	util.Uint32toBytes(bytes[8:12], n.Hashcode)
	nonzero_index := 0
	for ; bytes[nonzero_index] == 0; nonzero_index++ {
	}
	return n.VolumeId.String() + "," + hex.EncodeToString(bytes[nonzero_index:])
}
