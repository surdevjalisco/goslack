# goslack

Simple aplicación para consolas que consume la API de slack usando Golang

## Cómo usar goslack

Para ejecutar `goslack` es necesario primeramente compilar el código utilizando [Golang](https://go.dev/).
Después de instalar exitosamente Golang, basta con ejecutar `go build` dentro de la carpeta raíz (donde se encuentra `main.go`).
Una vez compilado el código, aparecerá un binario ejecutable con el nombre `goslack` el cual estará listo para ejecutarse.

### Token de autenticación

Para comunicarse con Slack, goslack utiliza un `token` que define el nivel de acceso permitido para la aplicación.
Para configurar un token, sigue los pasos definidos por slack [aquí](https://api.slack.com/authentication/oauth-v2#classic).
Slack solicitará la creación de una aplicación que actuará en tu nombre y la cuál genera el token.

Una vez creado el token, debes establecer una variable de ambiente que será leída por goslack como parte de la comunicación hacia Slack.
El nombre de la variable de ambiente debe ser `GOSLACK_AUTH_TOKEN`.
Por ejemplo, en la terminal de linux, puedes usar `export GOSLACK_AUTH_TOKEN="xoxb-xxxxxx...."` para definir la variable de ambiente.

> Nota: goslack generará un error si no encuentra la variable de ambiente.

### Operaciones

- Lista de correos electrónicos.  Ejecuta `goslack usuarios emails` para obtener todos los correos electrónicos de los miembros de Slack.

> Nota: El token debe tener el permiso de `users.read` para funcionar. Ver mas [aquí](https://api.slack.com/methods/users.list)
