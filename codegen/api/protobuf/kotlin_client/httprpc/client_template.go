package httprpc

import (
	"strings"
	"text/template"

	"github.com/hori-ryota/go-codegen/codegen/api/protobuf/go_server/httprpc"
	"github.com/hori-ryota/go-strcase"
)

type TemplateParam struct {
	KotlinPackage    string
	KotlinClassName  string
	StructdefPackage string
	Services         []Service
}

var ClientTemplate = template.Must(template.New("").Funcs(map[string]interface{}{
	"ToLowerCamel": strcase.ToLowerCamel,
	"ToUpperCamel": strcase.ToUpperCamel,
	"ToURLPath":    httprpc.URLPath,
}).Parse(strings.TrimSpace(`
// Code generated by go-codegen api protobuf kotlin_client httprpc; DO NOT EDIT

{{if .StructdefPackage}}import {{.StructdefPackage}}.*
{{end -}}
import io.ktor.client.HttpClient
import io.ktor.client.request.post
import io.ktor.client.request.url
import io.ktor.client.response.HttpResponse
import io.ktor.client.response.readBytes
import io.ktor.http.ContentType
import io.ktor.http.contentType
import kotlinx.coroutines.*
import kotlinx.serialization.protobuf.ProtoBuf

class {{.KotlinClassName}}(private val urlBase: String, private val client: HttpClient, private val defaultCoroutineScope: CoroutineScope) {

{{- range .Services}}
{{$service := .}}
{{- range .RPCs}}
    fun {{ToLowerCamel .Name}}(input: {{ToUpperCamel .InputType.Obj.Name}}, coroutineScope: CoroutineScope = defaultCoroutineScope): {{if .OutputType}}Deferred<{{ToUpperCamel .OutputType.Obj.Name}}>{{else}}Job{{end}} {
        return coroutineScope.{{if .OutputType}}async{{else}}launch{{end}} {
            val protoData = ProtoBuf.dump({{ToUpperCamel .InputType.Obj.Name}}.serializer(), input)
            val url = "$urlBase/{{ToURLPath $service.Obj.Name .Name}}"
            val contentType = ContentType("application", "protobuf")
            {{- if .OutputType}}
            val response = client.post<HttpResponse> {
                url(url)
                contentType(contentType)
                body = protoData
            }
            // TODO handle error
            ProtoBuf.load({{ToUpperCamel .OutputType.Obj.Name}}.serializer(), response.readBytes())
            {{- else}}
            client.post<HttpResponse> {
                url(url)
                contentType(contentType)
                body = protoData
            }
            // TODO handle error
            {{- end}}
        }
    }
{{- end}}
{{- end}}
}
`)))
