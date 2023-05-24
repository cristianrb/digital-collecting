package com.github.cristianrb

import com.github.cristianrb.plugins.configureKoin
import com.github.cristianrb.plugins.configureMonitoring
import com.github.cristianrb.plugins.configureSecurity
import com.github.cristianrb.plugins.configureSerialization
import com.github.cristianrb.users.usersRoutes
import com.zaxxer.hikari.HikariConfig
import com.zaxxer.hikari.HikariDataSource
import io.ktor.server.application.Application
import io.ktor.server.config.ApplicationConfig
import io.ktor.server.routing.routing
import org.flywaydb.core.Flyway
import org.jetbrains.exposed.sql.Database

fun main(args: Array<String>): Unit =
    io.ktor.server.netty.EngineMain.main(args)

@Suppress("unused") // application.conf references the main function. This annotation prevents the IDE from marking it as unused.
fun Application.module() {
    configureKoin()
    configureSerialization()
    configureMonitoring()
    configureSecurity()

    val dataSourceConfig = createDataSourceConfig(environment.config)
    val dataSource = createDataSource(dataSourceConfig)

    val flyway = Flyway.configure().dataSource(dataSource).load()
    flyway.migrate()

    Database.connect(dataSource)

    routing {
        usersRoutes()
    }
}

data class DataSourceConfig(
    val username: String,
    val password: String,
    val host: String,
    val port: String,
    val database: String,
    val schema: String
)

private fun createDataSourceConfig(applicationConfig: ApplicationConfig) = DataSourceConfig(
    applicationConfig.property("ktor.datasource.username").getString(),
    applicationConfig.property("ktor.datasource.password").getString(),
    applicationConfig.property("ktor.datasource.host").getString(),
    applicationConfig.property("ktor.datasource.port").getString(),
    applicationConfig.property("ktor.datasource.database").getString(),
    applicationConfig.property("ktor.datasource.schema").getString()
)

private fun createDataSource(dataSourceConfig: DataSourceConfig): HikariDataSource {
    val hikariConfig = HikariConfig()
    hikariConfig.username = dataSourceConfig.username
    hikariConfig.password = dataSourceConfig.password
    hikariConfig.jdbcUrl = "jdbc:postgresql://${dataSourceConfig.host}:${dataSourceConfig.port}/${dataSourceConfig.database}"
    hikariConfig.schema = dataSourceConfig.schema
    hikariConfig.maximumPoolSize = 10

    return HikariDataSource(hikariConfig)
}
