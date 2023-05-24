package com.github.cristianrb.util

import kotlinx.serialization.ExperimentalSerializationApi
import kotlinx.serialization.KSerializer
import kotlinx.serialization.Serializer
import kotlinx.serialization.encoding.Decoder
import kotlinx.serialization.encoding.Encoder
import java.time.LocalDate
import java.time.format.DateTimeFormatter


@OptIn(ExperimentalSerializationApi::class)
@Serializer(forClass = LocalDate::class)
object LocalDateSerializer : KSerializer<LocalDate> {

    private val isoDateFormatter = DateTimeFormatter.ISO_DATE

    override fun deserialize(decoder: Decoder): LocalDate {
        return LocalDate.parse(decoder.decodeString(), isoDateFormatter)
    }

    override fun serialize(encoder: Encoder, value: LocalDate) {
        encoder.encodeString(isoDateFormatter.format(value))
    }

}