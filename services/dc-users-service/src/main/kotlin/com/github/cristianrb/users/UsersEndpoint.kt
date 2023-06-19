package com.github.cristianrb.users

import arrow.core.Either
import arrow.core.raise.either
import com.github.cristianrb.util.clientError
import io.ktor.http.HttpStatusCode
import io.ktor.server.application.call
import io.ktor.server.auth.authenticate
import io.ktor.server.auth.jwt.JWTPrincipal
import io.ktor.server.auth.principal
import io.ktor.server.request.receive
import io.ktor.server.response.respond
import io.ktor.server.routing.Route
import io.ktor.server.routing.get
import io.ktor.server.routing.post
import io.ktor.server.routing.route
import org.koin.ktor.ext.inject

fun Route.usersRoutes() {

    val usersService by inject<UsersService>()

    route("/users") {
        post {
            val registrationUser = call.receive<RegistrationUser>()
            val result = either {
                usersService.addUser(registrationUser).bind()
            }

            when (result) {
                is Either.Left -> call.clientError(result.value)
                is Either.Right -> call.respond(HttpStatusCode.Accepted, result.value)
            }
        }

        authenticate {
            get("/me") {
                val user = call.principal<JWTPrincipal>()
                val username = user?.get("username")

                val result = either {
                    usersService.getUserByUsername(username?: "").bind()
                }

                when (result) {
                    is Either.Left -> call.clientError(result.value)
                    is Either.Right -> call.respond(HttpStatusCode.Accepted, result.value)
                }
            }

            post("/retrieve") {
                val user = call.principal<JWTPrincipal>()
                val coins = call.receive<Coins>()
                val userId = user?.getClaim("user_id", Integer::class)?.toInt() ?: -1

                val result = either {
                    usersService.retrieveCoins(userId, coins.coins).bind()
                }

                when (result) {
                    is Either.Left -> call.clientError(result.value)
                    is Either.Right -> call.respond(HttpStatusCode.OK, result.value)
                }
            }
        }

    }

    route("/session") {
        post {
            val loginUser = call.receive<LoginUser>()
            val result = either {
                usersService.authUser(loginUser).bind()
            }

            when (result) {
                is Either.Left -> call.clientError(result.value)
                is Either.Right -> call.respond(HttpStatusCode.OK, result.value)
            }
        }
    }

}