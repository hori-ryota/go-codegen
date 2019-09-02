// Code generated by go-codegen api protobuf kotlin_client httprpc; DO NOT EDIT

import com.github.horiryota.gocodegen.api.example.codegen.*
import io.ktor.client.HttpClient
import io.ktor.client.request.post
import io.ktor.client.request.url
import io.ktor.client.response.HttpResponse
import io.ktor.client.response.readBytes
import io.ktor.http.ContentType
import io.ktor.http.contentType
import kotlinx.coroutines.*
import kotlinx.serialization.protobuf.ProtoBuf

class CodegenExampleApi(private val urlBase: String, private val client: HttpClient, private val defaultCoroutineScope: CoroutineScope) {

    fun doSomethingWithOutputAndActor(input: DoingSomethingWithOutputAndActorUsecaseInput, coroutineScope: CoroutineScope = defaultCoroutineScope): Deferred<DoingSomethingWithOutputAndActorUsecaseOutput> {
        return coroutineScope.async {
            val protoData = ProtoBuf.dump(DoingSomethingWithOutputAndActorUsecaseInput.serializer(), input)
            val url = "$urlBase/DoingSomethingWithOutputAndActorUsecase/DoSomethingWithOutputAndActor"
            val contentType = ContentType("application", "protobuf")
            val response = client.post<HttpResponse> {
                url(url)
                contentType(contentType)
                body = protoData
            }
            // TODO handle error
            ProtoBuf.load(DoingSomethingWithOutputAndActorUsecaseOutput.serializer(), response.readBytes())
        }
    }

    fun doSomethingWithOutputWithoutActor(input: DoingSomethingWithOutputWithoutActorUsecaseInput, coroutineScope: CoroutineScope = defaultCoroutineScope): Deferred<DoingSomethingWithOutputAndActorUsecaseOutput> {
        return coroutineScope.async {
            val protoData = ProtoBuf.dump(DoingSomethingWithOutputWithoutActorUsecaseInput.serializer(), input)
            val url = "$urlBase/DoingSomethingWithOutputWithoutActorUsecase/DoSomethingWithOutputWithoutActor"
            val contentType = ContentType("application", "protobuf")
            val response = client.post<HttpResponse> {
                url(url)
                contentType(contentType)
                body = protoData
            }
            // TODO handle error
            ProtoBuf.load(DoingSomethingWithOutputAndActorUsecaseOutput.serializer(), response.readBytes())
        }
    }

    fun doSomethingWithoutOutputAndActor(input: DoingSomethingWithoutOutputAndActorUsecaseInput, coroutineScope: CoroutineScope = defaultCoroutineScope): Job {
        return coroutineScope.launch {
            val protoData = ProtoBuf.dump(DoingSomethingWithoutOutputAndActorUsecaseInput.serializer(), input)
            val url = "$urlBase/DoingSomethingWithoutOutputAndActorUsecase/DoSomethingWithoutOutputAndActor"
            val contentType = ContentType("application", "protobuf")
            client.post<HttpResponse> {
                url(url)
                contentType(contentType)
                body = protoData
            }
            // TODO handle error
        }
    }

    fun doSomethingWithoutOutputWithActor(input: DoingSomethingWithoutOutputWithActorUsecaseInput, coroutineScope: CoroutineScope = defaultCoroutineScope): Job {
        return coroutineScope.launch {
            val protoData = ProtoBuf.dump(DoingSomethingWithoutOutputWithActorUsecaseInput.serializer(), input)
            val url = "$urlBase/DoingSomethingWithoutOutputWithActorUsecase/DoSomethingWithoutOutputWithActor"
            val contentType = ContentType("application", "protobuf")
            client.post<HttpResponse> {
                url(url)
                contentType(contentType)
                body = protoData
            }
            // TODO handle error
        }
    }
}