@file:UseSerializers(
    LocalDateSerializer::class,
)

package com.github.cristianrb.users

import com.github.cristianrb.util.LocalDateSerializer
import kotlinx.serialization.Serializable
import kotlinx.serialization.UseSerializers
import java.time.LocalDate

@Serializable
enum class UserRole {
    SYSTEM,
    ADMIN,
    NORMAL
}

@Serializable
data class RegistrationUser(
    val username: String,
    val password: String,
    val email: String,
    val birthDate: LocalDate
)

fun RegistrationUser.toUserWithHash(hashedPassword: String): RegistrationUser {
    return this.copy(password = hashedPassword)
}

@Serializable
data class LoginUser(
    val username: String,
    val password: String
)

@Serializable
data class User(
    val username: String,
    val role: UserRole
)