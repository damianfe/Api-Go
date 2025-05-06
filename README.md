# 🐹 API-Go: API REST en Go con Gin y PostgreSQL

Este proyecto es una API REST básica desarrollada en Go, con arquitectura tipo **MVC**, utilizando:

- **Gin** como framework web
- **GORM** como ORM
- **PostgreSQL** como base de datos
- **godotenv** para la gestión de variables de entorno

## 📁 Estructura del Proyecto

```
api_go/
├── config/             # Configuración de la base de datos
│   └── database.go
├── controllers/        # Lógica de negocio (handlers)
│   └── user_controller.go
├── models/             # Modelos de datos
│   └── user.go
├── routes/             # Definición de rutas
│   └── user_routes.go
├── .env                # Variables de entorno (no se sube al repo)
├── go.mod              # Módulo y dependencias
└── main.go             # Punto de entrada
```

## ⚙️ Instalación

1. Clona este repositorio:
   ```bash
   git clone https://github.com/tu-usuario/api_go.git
   cd api_go
   ```

2. Instala las dependencias:
   ```bash
   go mod tidy
   ```

3. Crea un archivo `.env` con tu URL de PostgreSQL:
   ```env
   DB_URL=postgresql://usuario:contraseña@host/nombre_db?sslmode=require
   ```

4. Corre la API:
   ```bash
   go run main.go
   ```

## 🧪 Endpoints

### Obtener todos los usuarios

```http
GET /users
```

### Crear un nuevo usuario

```http
POST /users
Content-Type: application/json

{
  "name": "Juan",
  "email": "juan@mail.com"
}
```

## 🧰 Tecnologías utilizadas

- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [godotenv](https://github.com/joho/godotenv)

## 📌 Notas

- Se utiliza `AutoMigrate()` para crear automáticamente la tabla `users` si no existe.
- Esta API es un punto de partida para proyectos más complejos.
- Ideal para aprender Go y practicar buenas prácticas con estructura MVC.

## 🧑‍💻 Autor

Hecho con ❤️ por [Damian].