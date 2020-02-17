package httprpc

import (
	"errors"
	"go/types"
	"strings"
	"text/template"

	"github.com/hori-ryota/go-codegen/codegen/api/protobuf/go_server/httprpc"
	"github.com/hori-ryota/go-codegen/util/typeutil"
	"github.com/hori-ryota/go-strcase"
)

type TemplateParam struct {
	PackageName       string
	ImportPackages    string
	Services          []Service
	TypePrinter       typeutil.Printer
	SerializerPackage string
}

var ClientTemplate = template.Must(template.New("").Funcs(map[string]interface{}{
	"ToLowerCamel": strcase.ToLowerCamel,
	"ToUpperCamel": strcase.ToUpperCamel,
	"ToURLPath":    httprpc.URLPath,
	"ContentType": func(serializerPackage string) string {
		switch strings.ToLower(serializerPackage) {
		case "proto":
			return "application/protobuf"
		case "json":
			return "application/json"
		default:
			panic(errors.New("unknown serializer package: " + serializerPackage))
		}
	},
}).Parse(strings.TrimSpace(`
// Code generated by go-codegen api protobuf go_client httprpc; DO NOT EDIT

package {{.PackageName}}

{{.ImportPackages}}

{{$rootParam := .}}

func NewClient(
	httpClient *http.Client,
	urlBase url.URL,
	handleErrorResponse func(resp *http.Response) error,
) Client {
	return Client{
		httpClient: httpClient,
		urlBase: urlBase,
		handleErrorResponse: handleErrorResponse,
	}
}

type Client struct {
	httpClient *http.Client
	urlBase url.URL
	handleErrorResponse func(resp *http.Response) error
}

{{- range .Services}}
{{$service := .}}
{{- range .RPCs}}

{{$returnError := "return err"}}
{{if .OutputClientType}}
{{$returnError = "return output, err"}}
{{end}}

func (c Client){{.Name}}(ctx context.Context, input {{.InputClientType.Obj.Name}}) ({{if .OutputClientType}}output {{.OutputClientType.Obj.Name}}, err {{end}}error) {
	u := c.urlBase
	u.Path = path.Join(u.Path, "{{ToURLPath $service.Obj.Name .Name}}")

	inputProtoType := {{$rootParam.TypePrinter.PrintConverterWitoutErrorCheck "input" .InputClientType .InputProtoType}}
	b, err := {{$rootParam.SerializerPackage}}.Marshal(&inputProtoType)
	if err != nil {
		{{$returnError}}
	}
	r, err := http.NewRequest("POST", u.String(), bytes.NewReader(b))
	if err != nil {
		{{$returnError}}
	}
	r = r.WithContext(ctx)
	r.Header.Add("Content-Type", "{{ContentType $rootParam.SerializerPackage}}")

	resp, err := c.httpClient.Do(r)
	if err != nil {
		{{$returnError}}
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	if resp.StatusCode >= 400 {
		err := c.handleErrorResponse(resp)
		{{$returnError}}
	}

	{{- if .OutputClientType}}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = zaperr.Wrap(err, "failed to read response body", zap.Int("statusCode", resp.StatusCode))
		{{$returnError}}
	}
	outputProtoType := {{$rootParam.TypePrinter.PrintRelativeType .OutputProtoType}}{}
	if err := {{$rootParam.SerializerPackage}}.Unmarshal(body, &outputProtoType); err != nil {
		err = zaperr.Wrap(err, "failed to parse response body", zap.String("body", string(body)))
		{{$returnError}}
	}
	return {{$rootParam.TypePrinter.PrintConverterWitoutErrorCheck "outputProtoType" .OutputProtoType .OutputClientType}}, nil
	{{- else}}
	return nil
	{{- end}}
}
{{- end}}
{{- end}}
`)))

var ClientTemplateUsedPackages = []*types.Package{
	types.NewPackage("net/http", "http"),
	types.NewPackage("net/url", "url"),
	types.NewPackage("bytes", "bytes"),
	types.NewPackage("context", "context"),
	types.NewPackage("io", "io"),
	types.NewPackage("io/ioutil", "ioutil"),
	types.NewPackage("path", "path"),
	types.NewPackage("go.uber.org/zap", "zap"),
	types.NewPackage("github.com/hori-ryota/zaperr", "zaperr"),
}
