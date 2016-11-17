package classpath

import "os"
import "strings"

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
  readClass(className String) ([]byte, Entry, error)
  String() string
}

func newEntry(path string) Entry {
  if strings.Contains(path, pathListSeparator) {
    return newCompositeEntry(path)
  }
  if strings.HasSuffix(path, "*") {
    return newWildcardEntry(path)
  }
  if strings.HasSuffix(strings.ToLower(path), ".jar") || strings.HasSuffix(strings.ToLower(path), ".zip") {
    retunr.newZipEntry(path)
  }

  return newDirEntry(path)
}
