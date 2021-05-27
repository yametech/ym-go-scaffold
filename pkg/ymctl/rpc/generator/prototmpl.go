package generator

import (
	"github.com/yametech/ym-go-scaffold/pkg/util"
	"github.com/yametech/ym-go-scaffold/pkg/util/stringx"
	"path/filepath"
	"strings"
)

const rpcTemplateText = `syntax = "proto3";

package {{.package}};

option go_package = "./pb";

message Request {
  string ping = 1;
}

message Response {
  string pong = 1;
}

service {{.serviceName}} {
  rpc Ping(Request) returns(Response);
}
`

// ProtoTmpl returns an sample of a proto file
func ProtoTmpl(out string) error {
	protoFilename := filepath.Base(out)
	serviceName := stringx.From(strings.TrimSuffix(protoFilename, filepath.Ext(protoFilename)))
	//text, err := util.LoadTemplate(category, rpcTemplateFile, rpcTemplateText)
	//if err != nil {
	//	return err
	//}

	dir := filepath.Dir(out)
	err := util.MkdirIfNotExist(dir)
	if err != nil {
		return err
	}

	err = util.With("t").Parse(rpcTemplateText).SaveTo(map[string]string{
		"package":     serviceName.Untitle(),
		"serviceName": serviceName.Title(),
	}, out, false)
	return err
}
