# Documentación del Proyecto PRUEBA TÉCNICA BACKEND BIA

## Descripción del Proyecto

El proyecto PRUEBA TÉCNICA BACKEND BIA es un microservicio desarrollado en el lenguaje de programación Go (Golang) que está conectado a una base de datos MySQL. El microservicio se encarga de almacenar los consumos de energía de los medidores de los clientes en una tabla de la base de datos.

## Requerimientos

El proyecto cumple con los siguientes requerimientos:

1. Crear el microservicio en Golang.
2. Crear una base de datos e importar el archivo CSV de ejemplos de consumos.
3. Crear los endpoints de consulta con sus respectivos contratos.
4. Implementar test unitarios.
5. Subir el repositorio como privado a GitHub y dar permisos al usuario "sebas-assa".

## Configuración de la Base de Datos

La base de datos utilizada en el proyecto es MySQL y está ejecutándose en un contenedor de Docker. Para ejecutar la base de datos, se deben seguir los siguientes pasos:

1. Ejecutar el comando `docker-compose build` para construir la imagen del contenedor.
2. Ejecutar el comando `docker-compose up -d` para levantar el contenedor en segundo plano.

## Migración de la Tabla

Al ejecutar el proyecto con el comando `go run main.go`, se realizará automáticamente la migración de la tabla de consumos en la base de datos. Esto creará la tabla y sus respectivas columnas en la base de datos.

## Carga Masiva de Datos

Para realizar una carga masiva de datos desde el archivo CSV, se debe ejecutar el comando `go run src/importer/seeds/main.go`. Este comando leerá el archivo CSV y cargará los datos en la tabla de consumos en la base de datos.

## Endpoints de Consulta

El microservicio cuenta con los siguientes endpoints de consulta:

1. `/consumption?meters_ids=1&start_date=2023-06-01&end_date=2023-07-10&kind_period=monthly`: Retorna el consumo acumulado en los meses dados por `start_date` y `end_date`.
2. `/consumption?meters_ids=1&start_date=2023-06-01&end_date=2023-06-26&kind_period=weekly`: Retorna los acumulados de cada semana del periodo dado por `start_date` y `end_date`.
3. `/consumption?meters_ids=1&start_date=2023-06-01&end_date=2023-06-10&kind_period=daily`: Retorna el acumulado de consumo diario en el periodo dado por `start_date` y `end_date`.

## Test Unitarios

El proyecto incluye test unitarios para asegurar el correcto funcionamiento de las funciones y endpoints. Los test se pueden ejecutar con el comando `go test`.

## Repositorio en GitHub

El repositorio del proyecto se encuentra en GitHub y es privado. Se han otorgado permisos al usuario "sebas-assa" para acceder al repositorio.

## Versiones Utilizadas

El proyecto ha sido desarrollado utilizando Go version 1.20 y MySQL como base de datos.

## Ejecución del Proyecto

Para ejecutar el proyecto, se deben seguir los siguientes pasos:

1. Ejecutar la base de datos en Docker utilizando los comandos `docker-compose build` y `docker-compose up -d`.
2. Ejecutar el comando `go run main.go` para realizar la migración de la tabla en la base de datos.
3. Ejecutar el comando `go run src/importer/seeds/main.go` para realizar una carga masiva de datos desde el archivo CSV.
4. Ejecutar el comando `go run main.go` para iniciar el microservicio y acceder a los endpoints de consulta.

Con estos pasos, el proyecto estará funcionando correctamente y se podrán realizar las consultas de consumo de energía de los medidores de los clientes.

## Conclusiones

El proyecto PRUEBA TÉCNICA BACKEND BIA es un microservicio eficiente y bien documentado, que cumple con los requerimientos establecidos y utiliza las mejores prácticas de programación en Golang. Con la base de datos ejecutándose en un contenedor de Docker, se facilita la configuración y despliegue del proyecto en diferentes entornos. El uso de test unitarios garantiza la calidad y fiabilidad del microservicio. Además, la carga masiva de datos desde el archivo CSV permite almacenar grandes cantidades de información en la base de datos de manera rápida y eficiente. En resumen, el proyecto es un ejemplo de desarrollo de un microservicio en Golang con una base de datos MySQL en Docker, que puede ser utilizado como base para futuros proyectos similares.