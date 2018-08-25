package schema

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _query_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\x2c\x4d\x2d\xaa\xac\xe6\x52\x50\x50\x50\xc8\x48\xcd\xc9\xc9\xb7\x52\x08\x2e\x29\xca\xcc\x4b\x57\xd4\x01\x8b\x15\x67\x67\x16\x24\x16\x17\x6b\x94\x64\x26\x67\xa7\x96\xf8\x95\xe6\x26\xa5\x16\xc1\x94\x68\x5a\x29\x04\x43\xa4\xb9\x6a\xb9\x00\x01\x00\x00\xff\xff\x44\xdf\x1c\xa8\x4d\x00\x00\x00")

func query_graphql() ([]byte, error) {
	return bindata_read(
		_query_graphql,
		"query.graphql",
	)
}

var _schema_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\xce\x48\xcd\x4d\x54\xa8\xe6\x52\x50\x50\x50\x28\x2c\x4d\x2d\xaa\xb4\x52\x08\x04\x51\x5c\xb5\x80\x00\x00\x00\xff\xff\x54\xe0\x78\x3a\x1b\x00\x00\x00")

func schema_graphql() ([]byte, error) {
	return bindata_read(
		_schema_graphql,
		"schema.graphql",
	)
}

var _skipass_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\xce\xce\x2c\x48\x2c\x2e\x56\xa8\xe6\x52\x50\x50\x50\x28\xc8\x49\xcc\xb3\x52\x08\x2e\x29\xca\xcc\x4b\x57\x04\x8b\x24\x27\x16\xa5\xf8\x95\xe6\x26\xa5\x16\xa1\x8a\x97\x64\x26\x67\xa7\x96\x60\x93\x29\x28\x2d\x4a\xce\x48\x2c\x4e\x75\x49\x2c\x49\xb5\x52\x08\xc9\xcc\x4d\x55\xe4\xaa\x05\x04\x00\x00\xff\xff\xfe\x2e\xd9\x2c\x6c\x00\x00\x00")

func skipass_graphql() ([]byte, error) {
	return bindata_read(
		_skipass_graphql,
		"skipass.graphql",
	)
}

var _time_graphqls = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x08\xc9\xcc\x4d\x55\xc8\x2c\x56\x48\xcc\x53\x08\x72\x73\x36\x36\x36\xb6\x54\x28\xc9\xcc\x4d\x2d\x2e\x49\xcc\x2d\xd0\xe3\x2a\x4e\x4e\xcc\x49\x2c\x02\x2b\x02\x04\x00\x00\xff\xff\xb0\x23\xec\x86\x2b\x00\x00\x00")

func time_graphqls() ([]byte, error) {
	return bindata_read(
		_time_graphqls,
		"time.graphqls",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"query.graphql": query_graphql,
	"schema.graphql": schema_graphql,
	"skipass.graphql": skipass_graphql,
	"time.graphqls": time_graphqls,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"query.graphql": &_bintree_t{query_graphql, map[string]*_bintree_t{
	}},
	"schema.graphql": &_bintree_t{schema_graphql, map[string]*_bintree_t{
	}},
	"skipass.graphql": &_bintree_t{skipass_graphql, map[string]*_bintree_t{
	}},
	"time.graphqls": &_bintree_t{time_graphqls, map[string]*_bintree_t{
	}},
}}
