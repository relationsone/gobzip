package gobzip

import (
	"fmt"
	"hash/fnv"
	"path/filepath"

	"github.com/satori/go.uuid"
)

// use uuid to generate file name
func tempFile(dir, prefix string) (string, error) {
	fnv1a := fnv.New64a()
	id, err := uuid.NewV1()
	if err != nil {
		return "", err
	}
	fnv1a.Write(id.Bytes())
	return filepath.Join(fmt.Sprintf("%s", dir), fmt.Sprintf("%s-%d", prefix, uint64(fnv1a.Sum64()))), nil
}
