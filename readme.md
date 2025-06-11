# 🐹 Go REST API – Proyecto Backend Profesional

![Go](https://img.shields.io/badge/Go-1.21-blue?logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-blue?logo=postgresql)
![Docker](https://img.shields.io/badge/Docker-ready-blue?logo=docker)
![Swagger](https://img.shields.io/badge/Swagger-UI-green?logo=swagger)
![Tests](https://img.shields.io/badge/Tests-passing-brightgreen?logo=pytest)
![Logger](https://img.shields.io/badge/Logger-active-lightgrey?logo=logstash)

> Una API RESTful escrita en Go usando el potente enrutador [Gorilla Mux](https://github.com/gorilla/mux), lista para producción con validaciones, pruebas, documentación interactiva y contenedores para su despliegue.

---

## 📸 Vista Previa

![preview](https://github.com/santigorbe/GO_REST_API.git)
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

📁 /controllers  
├── /tasks.controller.go  
├── /users.controller.go  
📁 /db  
├── /connection.go  
📁 /docs  
├── /docs.go  
📁 /moddlewares  
├── /errorHandler.go  
├── /logger.go  
📁 /models  
├── /Task.go  
├── /User.go  
📁 /repository  
├── /gorm.tasks.repo.go  
├── /gorm.users.repo.go  
├── /tasks.repository.go  
├── /users.repository.go  
📁 /routes  
├── /index.routes.go  
├── /tasks.routes.go  
├── /users.routes.go  
📁 /services  
├── /tasks.services.go  
├── /users.services.go  
📁 /utils  
├── /response.helper.go  
📄 main.go  
📄 .env  
📄 docker-compose.yml  
📄 Dockerfile  
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

http://localhost:3000/docs 🧭

🧪 Autor
Desarrollado con Go por Santiago Gorbea 👨‍💻
