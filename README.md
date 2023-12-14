# GoKeygenSecure CLI

## Descripción

GoKeygenSecure es una herramienta de línea de comandos para generar contraseñas seguras y evaluar la seguridad de contraseñas existentes.

## Instalación

Para utilizar la herramienta, necesitarás tener Go instalado en tu sistema. Una vez que tengas Go, puedes clonar el repositorio y construir el binario con los siguientes comandos:

```
git clone https://tu-repositorio/GoKeygenSecure.git
cd GoKeygenSecure
go build -o GoKeygenSecure
```

## Ejecución

Después de compilar el programa, puedes ejecutarlo directamente desde la línea de comandos:

```
./GoKeygenSecure
```

## Comandos disponibles

### Generar una contraseña

Para generar una nueva contraseña:

```
./GoKeygenSecure generate
```

Puedes personalizar la generación de la contraseña con los siguientes flags:

- `-l, --length`: Especifica la longitud de la contraseña.
- `-n, --nums`: Incluye números en la contraseña.
- `-s, --syms`: Incluye símbolos en la contraseña.
- `-e, --easy-to-say`: Evita caracteres confusos para hacer la contraseña fácil de decir.
- `-r, --easy-to-read`: Evita caracteres confusos para hacer la contraseña fácil de leer.

Ejemplo de uso con flags:

```
./GoKeygenSecure generate --length 16 --nums --syms
```

### Evaluar una contraseña

Para evaluar la seguridad de una contraseña existente:

```
./GoKeygenSecure evaluate
```

Te pedirá que ingreses la contraseña que deseas evaluar y luego te mostrará el nivel de seguridad de esa contraseña.

## Salir del programa

Para salir del programa, puedes usar el comando `exit` en cualquier momento.

```
exit
```

Esperamos que encuentres útil esta herramienta para gestionar la seguridad de tus contraseñas.
