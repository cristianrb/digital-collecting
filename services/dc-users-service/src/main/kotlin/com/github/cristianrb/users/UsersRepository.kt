package com.github.cristianrb.users

import arrow.core.Either
import com.github.cristianrb.generated.Users
import com.github.cristianrb.util.AppError
import com.github.cristianrb.util.AppResult
import org.jetbrains.exposed.sql.insert
import org.jetbrains.exposed.sql.select
import org.jetbrains.exposed.sql.transactions.transaction
import org.koin.core.component.KoinComponent

interface UsersRepository {
    fun addUser(userWithHash: RegistrationUser): AppResult<User>
    fun getUserByUsername(username: String): AppResult<User>
}
class DefaultUsersRepository: UsersRepository, KoinComponent {

    override fun addUser(userWithHash: RegistrationUser): AppResult<User> {
        return Either.catch {
            transaction {
                Users.insert {
                    it[username] = userWithHash.username
                    it[password] = userWithHash.password
                    it[role] = UserRole.NORMAL.name
                    it[digitalCoins] = 100
                }
            }


            User(userWithHash.username, null, UserRole.NORMAL)
        }.mapLeft { AppError.SQLException(it) }

    }

    override fun getUserByUsername(username: String): AppResult<User> {
        return Either.catch {
            val row = transaction {
                Users.select {
                    Users.username eq username
                }.first()
            }

            User(
                row[Users.username],
                row[Users.password],
                row[Users.role].toUserRole(),
            )
        }.mapLeft { AppError.SQLException(it) }
    }

}