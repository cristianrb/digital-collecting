package com.github.cristianrb.util

import arrow.core.Either
import io.ktor.http.HttpStatusCode
import io.ktor.server.application.ApplicationCall
import io.ktor.server.response.respond
import kotlinx.serialization.Serializable

typealias AppResult<T> = Either<AppError, T>

const val UNKNOWN_ERROR = "Unknown Error"

sealed class AppError(open val msg: String, open val httpErrorCode: Int = 400) {

    data class BadRequestError(val brqMsg: String): AppError(brqMsg, 400)
    data class AuthenticationError(val msg2: String): AppError(msg2, 401)

    data class SQLException(val t: Throwable): AppError(t.message ?: UNKNOWN_ERROR, 500)
}

@Serializable
data class HttpAppError(val message: String, val errorCode: Int)
fun AppError.toHttpAppError(): HttpAppError {
    return HttpAppError(this.msg, this.httpErrorCode)
}

suspend fun ApplicationCall.clientError(appError: AppError) {
    respond(
        status = HttpStatusCode.fromValue(appError.httpErrorCode),
        message = appError.toHttpAppError(),
    )
}