package com.github.cristianrb.users

import org.koin.core.component.KoinComponent
import org.koin.core.component.inject
import org.mindrot.jbcrypt.BCrypt

interface UsersService {
    fun addUser(user: RegistrationUser): User
}
class DefaultUsersService: UsersService, KoinComponent {

    private val usersRepository by inject<UsersRepository>()
    override fun addUser(user: RegistrationUser): User {
        val salt = BCrypt.gensalt()
        val hashedPassword = BCrypt.hashpw(user.password, salt)
        return usersRepository.addUser(user.toUserWithHash(hashedPassword))
    }
}