package com.github.cristianrb.plugins

import com.github.cristianrb.users.DefaultUsersRepository
import com.github.cristianrb.users.DefaultUsersService
import com.github.cristianrb.users.UsersRepository
import com.github.cristianrb.users.UsersService
import io.ktor.server.application.*
import io.ktor.server.application.Application
import org.koin.dsl.module
import org.koin.ktor.plugin.Koin

fun Application.configureKoin() {
    install(Koin) {
        modules(
            module {
                single<UsersService> { DefaultUsersService() }
                single<UsersRepository> { DefaultUsersRepository() }
            }
        )
    }
}