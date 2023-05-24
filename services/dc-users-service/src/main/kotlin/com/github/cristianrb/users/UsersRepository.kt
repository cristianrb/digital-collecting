package com.github.cristianrb.users

import dev.bettercoding.generated.Users
import org.jetbrains.exposed.sql.insert
import org.jetbrains.exposed.sql.transactions.transaction
import org.koin.core.component.KoinComponent

interface UsersRepository {
    fun addUser(userWithHash: RegistrationUser): User
}
class DefaultUsersRepository(): UsersRepository, KoinComponent {

    init {}
    override fun addUser(userWithHash: RegistrationUser): User {
        println(userWithHash)
        println(User(userWithHash.username, UserRole.NORMAL))

        transaction {
            Users.insert {
                it[username] = userWithHash.username
                it[password] = userWithHash.password
                it[role] = UserRole.NORMAL.name
                it[digitalCoins] = 100
            }
        }


        return User(userWithHash.username, UserRole.NORMAL)
    }

}