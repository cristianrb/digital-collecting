package com.github.cristianrb.auth

import com.auth0.jwt.JWT
import com.auth0.jwt.algorithms.Algorithm
import com.github.cristianrb.users.User
import org.koin.core.component.KoinComponent
import java.util.*

interface AuthService {
    suspend fun generateToken(user: User): String
}

class DefaultAuthService(
    private val audience: String,
    private val issuer: String,
    private val secret: String
): AuthService, KoinComponent {

    override suspend fun generateToken(user: User): String {
        return JWT.create()
            .withAudience(audience)
            .withIssuer(issuer)
            .withClaim("username", user.username)
            .withClaim("role", user.role.name)
            .withExpiresAt(Date(System.currentTimeMillis() + 3600000))
            .sign(Algorithm.HMAC256(secret))
    }

}