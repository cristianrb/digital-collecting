package com.github.cristianrb.users

import arrow.core.Either
import arrow.core.flatMap
import com.github.cristianrb.generated.Users
import com.github.cristianrb.util.AppError
import com.github.cristianrb.util.AppResult
import org.jetbrains.exposed.sql.Column
import org.jetbrains.exposed.sql.insert
import org.jetbrains.exposed.sql.select
import org.jetbrains.exposed.sql.transactions.transaction
import org.jetbrains.exposed.sql.update
import org.koin.core.component.KoinComponent

interface UsersRepository {
    fun addUser(userWithHash: RegistrationUser): AppResult<User>
    fun getUserByUsername(username: String): AppResult<User>
    fun getUserById(userId: Int): AppResult<User>
    fun retrieveCoins(user: User, coins: Int): AppResult<User>
}
class DefaultUsersRepository: UsersRepository, KoinComponent {

    override fun addUser(userWithHash: RegistrationUser): AppResult<User> {
        return Either.catch {
            val row = transaction {
                Users.insert {
                    it[username] = userWithHash.username
                    it[password] = userWithHash.password
                    it[role] = UserRole.NORMAL.name
                    it[digitalCoins] = 100
                }.resultedValues!!.first()
            }


            User(row[Users.id].value, userWithHash.username, null, UserRole.NORMAL, 100)
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
                row[Users.id].value,
                row[Users.username],
                row[Users.password],
                row[Users.role].toUserRole(),
                row[Users.digitalCoins]
            )
        }.mapLeft { AppError.SQLException(it) }
    }

    override fun getUserById(userId: Int): AppResult<User> {
        return Either.catch {
            val row = transaction {
                Users.select {
                    Users.id eq userId
                }.first()
            }

            User(
                row[Users.id].value,
                row[Users.username],
                row[Users.password],
                row[Users.role].toUserRole(),
                row[Users.digitalCoins]
            )
        }.mapLeft { AppError.SQLException(it) }
    }

    override fun retrieveCoins(user: User, coins: Int): AppResult<User> {
        return Either.catch {
            transaction {
                Users.update({Users.id eq user.id}) {
                    it[digitalCoins] = coins
                }
            }
        }
            .mapLeft { AppError.SQLException(it) }
            .flatMap { getUserById(user.id) }
    }

}