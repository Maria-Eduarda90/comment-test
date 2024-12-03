#!/bin/bash

# Carregar variáveis do arquivo .env
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
else
    echo "Arquivo .env não encontrado."
    exit 1
fi

# Funções para as tarefas
MIGRATE_BIN=C:\\Users\\meyh_mary\\go\\bin\\migrate

create_migration() {
    $MIGRATE_BIN create -ext=sql -dir=internal/database/migrations -seq "$1"
}

migrate_up() {
    $MIGRATE_BIN -path=internal/database/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up
}

migrate_down() {
    $MIGRATE_BIN -path=internal/database/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose down
}

# Interpretar o comando
case "$1" in
    create_migration)
        if [ -z "$2" ]; then
            echo "Uso: $0 create_migration <nome>"
            exit 1
        fi
        create_migration "$2"
        ;;
    migrate_up)
        migrate_up
        ;;
    migrate_down)
        migrate_down
        ;;
    *)
        echo "Comando inválido. Uso:"
        echo "$0 create_migration <nome>"
        echo "$0 migrate_up"
        echo "$0 migrate_down"
        exit 1
        ;;
esac
