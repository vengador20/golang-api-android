# Aplicación

## Tecnologías que se utilizaran:

## Lenguaje de programación:

    - Golang

## FrameWork o librerías:

### BackEnd:

    - Fiber (Servidor Web)
    - Jwt (Json Web Token)
    - Mongo driver (Base de datos)
    - Validator (go-playground/validator)
    - Validator Translate (go-playground/universal-translator)

## Arquitectura limpia (Arquitectura hexagonal)
![Clean Architecture](./assets/CleanArchitecture.jpg)

    - Infraestructure (logica de la aplicacion)
    - Application o use cases (casos de uso se utiliza la interface)
    - Domain o entities (entidades de la aplicacion)