package com.github.cristianrb.plugins

import com.github.cristianrb.auth.AuthService
import com.github.cristianrb.auth.DefaultAuthService
import com.github.cristianrb.users.DefaultUsersRepository
import com.github.cristianrb.users.DefaultUsersService
import com.github.cristianrb.users.UsersRepository
import com.github.cristianrb.users.UsersService
import io.ktor.server.application.*
import io.ktor.server.application.Application
import org.koin.dsl.module
import org.koin.ktor.plugin.Koin

fun Application.configureKoin() {
    val audience = environment.config.property("jwt.audience").getString()
    val issuer = environment.config.property("jwt.issuer").getString()
    val secret = environment.config.property("jwt.secret").getString()


    install(Koin) {
        modules(
            module {
                single<AuthService> { DefaultAuthService(audience, issuer, secret) }
                single<UsersService> { DefaultUsersService() }
                single<UsersRepository> { DefaultUsersRepository() }
            }
        )
    }
}