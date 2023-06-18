package com.github.cristianrb.users

import arrow.core.Either
import arrow.core.raise.either
import com.github.cristianrb.auth.AuthService
import com.github.cristianrb.util.AppError
import com.github.cristianrb.util.AppResult
import org.koin.core.component.KoinComponent
import org.koin.core.component.inject
import org.mindrot.jbcrypt.BCrypt

interface UsersService {
    fun addUser(user: RegistrationUser): AppResult<User>
    fun getUserByUsername(username: String): AppResult<User>
    suspend fun authUser(loginUser: LoginUser): AppResult<UserWithToken>
    fun retrieveCoins(userId: Int, coins: Int): AppResult<User>
}
class DefaultUsersService: UsersService, KoinComponent {

    private val usersRepository by inject<UsersRepository>()
    private val authService by inject<AuthService>()

    override fun addUser(user: RegistrationUser): AppResult<User> {
        return either {
            val salt = BCrypt.gensalt()
            val hashedPassword = BCrypt.hashpw(user.password, salt)
            usersRepository.addUser(user.toUserWithHash(hashedPassword)).bind()
        }

    }

    override fun getUserByUsername(username: String): AppResult<User> {
        return either {
            usersRepository.getUserByUsername(username).bind()
        }
    }

    override suspend fun authUser(loginUser: LoginUser): AppResult<UserWithToken> {
        return either {
            val user = usersRepository.getUserByUsername(loginUser.username).bind()
            val valid = BCrypt.checkpw(loginUser.password, user.password)

            return if (valid) {
                val user = UserWithToken(
                    user.username,
                    user.role,
                    authService.generateToken(user)
                )
                Either.Right(user)
            } else {
                Either.Left(AppError.AuthenticationError("Invalid auth"))
            }

        }
    }

    override fun retrieveCoins(userId: Int, coins: Int): AppResult<User> {
        return either {
            val user = usersRepository.getUserById(userId).bind()
            return if (user.coins >= coins) {
                val updatedUser = usersRepository.retrieveCoins(user, user.coins - coins).bind()
                Either.Right(updatedUser)
            } else {
                Either.Left(AppError.BadRequestError("Insufficient coins"))
            }
        }

    }
}