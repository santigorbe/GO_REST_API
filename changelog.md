# Changelog

Todos los cambios notables a este proyecto se documentarán en este archivo.

El formato sigue [Keep a Changelog](https://keepachangelog.com/es-ES/1.0.0/)
y este proyecto sigue [SemVer](https://semver.org/lang/es/).

---

## [1.2.4] - 2025-05-30
### Agregado
- Soporte para paginación y filtros en el endpoint `/tasks`.
- Middleware de recuperación para capturar panic.
- Dockerfile y docker-compose para contenerizar app y base de datos.
- Endpoints de tasks con documentacion para Swagger
- Se agrego el manejo de variables de entorno
- Se agrega la funcion para poder dockerizar la aplicacion completa facilmente

### Cambiado
- Refactor de estructura de repositorios para usar interfaces limpias.

---

## [1.2.3] - 2025-05-29
### Agregado
- Respuestas estandarizadas con helpers JSON.
- Validación de ID en endpoints.
- Se separo el codigo en routes, controllers, services y repository
  
---

## [1.0.0] - 2025-05-29
### Inicial
- CRUD básico de tareas y usuarios con GORM y Gorilla Mux.
