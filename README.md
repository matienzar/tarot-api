# tarot-api
Aplicación para poder testear los conocimientos de Go haciendo uso de un modelo de AI como experto en el tarot.

Se ha mirado de hacer uso de una arquitectura hexagonal e incorporar una base de datos para persistir las preguntas realizadas.

# Pasos
## Utilidades a instalar previas a la ejecución
### Docker
- https://docs.docker.com/engine/install/
### GoLang
- https://go.dev/doc/install
## Instalar dependencias

1) ORM

    Para la comunicación con la base de datos haremos uso de un ORM:

    ```go
    go get -u gorm.io/gorm
    ```

2) ORM. Conector con la base de datos SQLLite

    Para conectar el ORM con una base de datos, en este caso SQLLite

    ```go
    go get -u gorm.io/driver/sqlite
    ```
3) Rutas REST

    Para manejar las rutas que haremos uso en nuestra aplicación a través de llamadas API

    ```go
    go get -u github.com/gorilla/mux
    ```
## Ejecutar la aplicación
Ves al apartado [Ejecutar el proyecto](#ejecutar-el-proyecto)
## Probar la aplicación con Postman
Importa la [colección](tarot-api.postman_collection.json) que hay en el proyecto en Postman y mira de ejecutarla. Hay dos métodos:
* Realizar una pregunta
* Mostrar todas las preguntas realizadas


# Comamdos útiles
```go
# Ejecutar el proyecto
go run .
```

```go
# Instalar todas las dependencias que el proyecto necesita
go mod tidy
```

```go
# Formatear el código
go fmt ./...
```

```go
# Verificar posibles errores estáticos en el código
go vet ./...
```

```go
# Ejecutar pruebas unitarias
go test ./...
```

```bash
# Verificar los contenedores en ejecución
docker compose ps
```

```bash
# Ver los logs de los contenedores
docker compose logs
```

# Aplicación de una arquitectura hexagonal

- **Core Domain** (internal/core/domain): Define las entidades y las reglas de negocio.
- **Service Layer** (internal/core/service): Contiene la lógica de negocio.
- **Application Layer** (internal/application): Define los puertos/interfaces que los adaptadores (como repositorios) deben implementar.
- **Infrastructure Layer** (internal/infrastructure): Implementa los adaptadores, como el repositorio de SQLite.

# Ejecutar el proyecto

Acuérdate de configurar antes en el fichero [.env](.env) tu API_KEY 
```bash
docker compose --env-file .env up --build -d
```

Para **parar** el servicio
```bash
docker compose stop app
```

Para **arrancar** el servicio
```bash
docker compose up app
```

El comando que ejecuta docker compose hace esto:
- **apk add --no-cache sqlite**: Esto instala sqlite3 en la imagen basada en Alpine Linux. apk es el gestor de paquetes de Alpine, y --no-cache asegura que no se almacenen archivos innecesarios después de la instalación.
- **sh /app/scripts/init_db.sh**: Ejecuta el script de inicialización de la base de datos.
- **go run main.go**: Inicia tu aplicación Go después de que la base de datos haya sido inicializada
- **export CGO_ENABLED=1**: Necesario para SQLLite