package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/badvassal/wllib/gen/wlerr"
	"github.com/badvassal/wllib/msq"
)

// A backup block body has the following structure:
// SecSection:
//     BackupHdr (JSON-encoded)
// PlainSection
//	   gzipped GAME file

const BackupMagic = "wlrand-backup-v1"

type BackupHdr struct {
	Magic    string
	Filename string
	GameIdx  int `json:",omitempty"` // Deprecated; only used if Filename=="".
}

type BackupRecord struct {
	Hdr  BackupHdr
	Data []byte
}

func EncodeBackupBlock(filename string, data []byte) (*msq.Body, error) {
	hdr := &BackupHdr{
		Magic:    BackupMagic,
		Filename: filename,
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

	return &msq.Body{
		SecSection:   []byte(j),
		PlainSection: b.Bytes(),
	}, nil
}

func DecodeBackupBlock(b msq.Body) *BackupRecord {
	hdr := &BackupHdr{}
	if err := json.Unmarshal(b.SecSection, hdr); err != nil {
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

func FindAndDecodeBackupRecords(bodies []msq.Body) map[string]*BackupRecord {
	rmap := map[string]*BackupRecord{}

	for _, b := range bodies {
		r := DecodeBackupBlock(b)
		if r != nil {
			var filename string
			if r.Hdr.Filename != "" {
				filename = r.Hdr.Filename
			} else {
				// Backwards compatibilty.
				switch r.Hdr.GameIdx {
				case 0:
					filename = "GAME1"

				case 1:
					filename = "GAME2"

				default:
					log.Errorf("discarding backup record with no filename and "+
						"bad game idx: have=%d want==0||1", r.Hdr.GameIdx)
				}
			}

			// Only allow files in the same directory.
			if dir, _ := filepath.Split(filename); dir != "" {
				log.Errorf("discarding backup record with bad filename: %s",
					filename)
			} else if rmap[filename] != nil {
				log.Errorf(
					"discarding backup record with duplicate filename: %s",
					filename)
			} else {
				rmap[filename] = r
			}
		}
	}

	return rmap
}
