#!/bin/sh

# Asignamos los permisos del directorio
chmod 777 /tmp

# Crear la base de datos si no existe
if [ ! -f /tmp/tarot.db ]; then
    echo "Creando la base de datos vacía con VACUUM..."
    sqlite3 /tmp/tarot.db "VACUUM;"
fi

# Ejecutar el script SQL para crear las tablas
sqlite3 /tmp/tarot.db < /app/scripts/preguntas.sql

# Verificar que la tabla "preguntas" se haya creado correctamente
echo "Verificando que la tabla 'preguntas' existe..."
TABLE_EXISTS=$(sqlite3 /tmp/tarot.db "SELECT name FROM sqlite_master WHERE type='table' AND name='preguntas';")

if [ "$TABLE_EXISTS" = "preguntas" ]; then
    echo "La tabla 'preguntas' ha sido creada exitosamente."
else
    echo "Error: La tabla 'preguntas' no se ha creado."
    exit 1
fi