package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	"github.com/badvassal/wllib/gen/wlerr"
	"github.com/badvassal/wllib/msq"
)

// A backup block has the following structure:
// EncSection:
//     BackupHdr (JSON-encoded)
// PlainSection
//	   gzipped GAME file

const BackupMagic = "wlrand-backup-v1"

type BackupHdr struct {
	Magic   string
	GameIdx int
}

type BackupRecord struct {
	Hdr  BackupHdr
	Data []byte
}

func EncodeBackupBlock(gameIdx int, data []byte) (*msq.Block, error) {
	hdr := &BackupHdr{
		Magic:   BackupMagic,
		GameIdx: gameIdx,
	}

	j, err := json.Marshal(hdr)
	if err != nil {
		return nil, wlerr.Wrapf(err, "failed to marshal backup header")
	}

	b := &bytes.Buffer{}
	gw := gzip.NewWriter(b)
	if _, err := gw.Write(data); err != nil {
		return nil, wlerr.Wrapf(err, "failed to compress backup data")
	}
	if err := gw.Close(); err != nil {
		return nil, wlerr.Wrapf(err, "failed to compress backup data")
	}

	return &msq.Block{
		EncSection:   []byte(j),
		PlainSection: b.Bytes(),
	}, nil
}

func DecodeBackupBlock(b msq.Block) *BackupRecord {
	hdr := &BackupHdr{}
	if err := json.Unmarshal(b.EncSection, hdr); err != nil {
		return nil
	}

	if hdr.Magic != BackupMagic {
		return nil
	}

	buf := bytes.NewBuffer(b.PlainSection)
	gr, err := gzip.NewReader(buf)
	if err != nil {
		log.Errorf("failed to decompress backup data: %s", err.Error())
		return nil
	}
	data, err := ioutil.ReadAll(gr)
	if err != nil {
		log.Errorf("failed to decompress backup data: %s", err.Error())
		return nil
	}

	return &BackupRecord{
		Hdr:  *hdr,
		Data: data,
	}
}

func FindAndDecodeBackupRecords(blocks []msq.Block) (*BackupRecord, *BackupRecord) {
	var r0 *BackupRecord
	var r1 *BackupRecord

	onDup := func(idx int) {
		log.Errorf("discarding backup record with duplicate idx: %d", idx)
	}

	for _, b := range blocks {
		r := DecodeBackupBlock(b)
		if r != nil {
			switch r.Hdr.GameIdx {
			case 0:
				if r0 != nil {
					onDup(0)
				}
				r0 = r

			case 1:
				if r1 != nil {
					onDup(1)
				}
				r1 = r

			default:
				log.Errorf("discarding backup record with bad game idx: "+
					"have=%d want==0||1", r.Hdr.GameIdx)
			}
		}
	}

	if r0 != nil && r1 == nil || r0 == nil && r1 != nil {
		log.Errorf("discarding backup record: wrong count: have=1 want=2")
		return nil, nil
	}

	return r0, r1
}
