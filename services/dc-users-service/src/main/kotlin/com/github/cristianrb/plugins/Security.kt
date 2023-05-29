package com.github.cristianrb.plugins

import com.auth0.jwt.JWT
import com.auth0.jwt.algorithms.Algorithm
import io.ktor.server.application.Application
import io.ktor.server.auth.authentication
import io.ktor.server.auth.jwt.JWTPrincipal
import io.ktor.server.auth.jwt.jwt

fun Application.configureSecurity() {
    
    authentication {
            jwt {
                val audience = this@configureSecurity.environment.config.property("jwt.audience").getString()
                val secret = this@configureSecurity.environment.config.property("jwt.secret").getString()
                val issuer = this@configureSecurity.environment.config.property("jwt.issuer").getString()
                realm = this@configureSecurity.environment.config.property("jwt.realm").getString()

                verifier(
                    JWT
                        .require(Algorithm.HMAC256(secret))
                        .withAudience(audience)
                        .withIssuer(issuer)
                        .build()
                )
                validate { credential ->
                    if (credential.payload.audience.contains(audience)) JWTPrincipal(credential.payload) else null
                }
            }
        }
}
