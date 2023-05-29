#!/usr/bin/env bash

./gradlew generateExposedCode
rm build/tables/com/github/cristianrb/generated/Public.flywaySchemaHistory.kt