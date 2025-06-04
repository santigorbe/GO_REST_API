# 🐹 Go REST API – Proyecto Backend Profesional

![Go](https://img.shields.io/badge/Go-1.21-blue?logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-blue?logo=postgresql)
![Docker](https://img.shields.io/badge/Docker-ready-blue?logo=docker)
![Swagger](https://img.shields.io/badge/Swagger-UI-green?logo=swagger)
![Tests](https://img.shields.io/badge/Tests-✔️-green)
![Logger](https://img.shields.io/badge/Logging-zap-blue)

> Una API RESTful escrita en Go usando el potente enrutador [Gorilla Mux](https://github.com/gorilla/mux), lista para producción con validaciones, pruebas, documentación interactiva y contenedores.

---

## 📸 Vista Previa

![preview](https://user-images.githubusercontent.com/your-image-path/swagger-ui-preview.png)
*Swagger UI integrado para una documentación interactiva*

---

## 🚀 Características Principales

| 💡 Tecnología       | ✅ Implementado          |
|---------------------|--------------------------|
| 🛠 **Gorilla Mux**   | Enrutador HTTP robusto   |
| 🐘 **PostgreSQL**    | Base de datos relacional |
| 🐳 **Docker**        | Contenerización          |
| 📜 **Swagger**       | Documentación automática |
| ✅ **Validación**     | Validaciones con structs |
| 🧪 **Testing**        | Tests unitarios y mocks  |
| 📝 **Logger**         | Logging profesional  |

---

## 🧰 Estructura del Proyecto

📁 /cmd
📁 /internal
├── /handler
├── /service
├── /repository
├── /middleware
📁 /pkg
├── /logger
├── /validator
├── /swagger
📄 go.mod


---

## 📦 Ejecutar el proyecto

```bash
# Clonar el repositorio
git clone https://github.com/tu-usuario/tu-repo.git

# Levantar con Docker
docker-compose up --build
```

📚 Documentación Swagger
Una vez que el servidor esté en marcha, visita:

http://localhost:8080/swagger 🧭

🧪 Autor
Desarrollado con Go por Santiago Gorbea 👨‍💻
