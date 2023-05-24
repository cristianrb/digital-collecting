package com.github.cristianrb.users

import io.ktor.http.HttpStatusCode
import io.ktor.server.application.call
import io.ktor.server.request.receive
import io.ktor.server.response.respond
import io.ktor.server.routing.Route
import io.ktor.server.routing.post
import io.ktor.server.routing.route
import org.koin.ktor.ext.inject

fun Route.usersRoutes() {

    val usersService by inject<UsersService>()

    route("/users") {
        post {
            val registrationUser = call.receive<RegistrationUser>()
            val user = usersService.addUser(registrationUser)

            call.respond(HttpStatusCode.Accepted, user)
        }
    }

}