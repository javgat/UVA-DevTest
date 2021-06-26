# DevTest

DevTest es una aplicación web que permite la creación de preguntas y exámenes de forma colaborativa,
al estilo de una red social.

Para su diseño e implementación se ha utilizado Swagger (OpenApi 2.0), Go (versión 1.15.8), MariaDB
(versión 15.1) y Angular (versión 12.0.1).

DevTest es el Trabajo de Fin de Grado (TFG) de Javier Gatón Herguedas ([javgat](https://github.com/javgat))
para la carrera de Ingeniería Informática en la Universidad de Valladolid.

El TFG se desarrolló al amparo de un convenio de la Escuela de Ingeniería Informática de Valladolid con
la empresa [HP SCDS](https://hpscds.com/), mediante el cual los TFG se realizarán bajo la dirección
conjunta de profesionales de HP y profesores de la Escuela.

La tutorización del TFG se llevó a cabo por Rubén López Fernández, profesional de la empresa, y Valentín
Cardeñoso Payo, profesor de la Escuela de Ingeniería Informática.

## Administración e Instalación

La instalación de la aplicación se explicará para un entorno Linux, poniendo de ejemplo Ubuntu 20.04.

Lo primero de todo es clonar el repositorio de git.

```shell
    git clone https://gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest .
```

Una vez descargado, se puede dividir la instalación en tres partes: Base de datos, BackEnd y FrontEnd.

### Base de Datos

El software de gestión de bases de datos utilizado ha sido MariaDB, en la versión 15.1. Se recomienda
utilizar esa versión o una posterior para el funcionamiento correcto de la aplicación.

Una vez instalada, se deberá ejecutar el fichero de inicialización de la base de datos, en **/Backend/init.sql**.

Finalmente, habrá que crear un directorio llamado *config* dentro de /Backend, que contendrá un fichero
*dbinfo.json* con la información necesaria para conectarse a la base de datos.

```shell
    cd BackEnd
    mkdir config
    cd config
    vi dbinfo.json
```

Dentro de ese fichero habrá un json con la información de la base de datos, siguiendo la siguiente
estructura, poniendo de ejemplo una base de datos alojada en localhost con el puerto 3306, y con el
nombre de usuario "admin" y contraseña "admin":

```json
{
  "Username" : "admin",
  "Pass" : "admin",
  "Host" : "localhost",
  "Port" : "3306",
  "Name" : "uva_devtest"
}
```

### BackEnd

#### Ejecución de pruebas sobre las respuestas

Para la ejecución de pruebas sobre las respuestas enviadas por los usuarios se ha utilizado Docker,
concretamente la versión 20.10.7.

Una vez se finalice la instalación de Docker en el sistema, se recomienda añadir al usuario que ejecute
los procesos del servidor al grupo de usuarios de Docker, para así evitar tener que ejecutar los
procesos con permisos de superusuario.

A continuación hay que construir la imagen de Docker encargada de ejecutar los procesos, para eso
habrá que entrar en el directorio **/BackEnd/docker/** y ejecutar el script **build_container.sh**.

Una vez hecho esto ya sé podrá ejecutar el servidor. Desde el servidor se llamará a los diferentes
scripts que hay en el directorio principalmente mencionado, y los directorios que se compartirán
con el contenedor de Docker estarán en el directorio **/tmp**, dentro de un subdirectorio llamado
**pruebas** que se creará automáticamente, por lo que hay que asegurarse de que el usuario tiene
los permisos necesarios sobre el directorio **/tmp**.

#### Servidor

El lenguaje de programación utilizado en el BackEnd es Go, o Golang, y la versión utilizada es la 1.15.8.
Se recomienda utilizar esa versión o una posterior compatible para compilar el software del BackEnd.

Una vez instalado go, se necesitará compilar el código fuente. Para ello basta con ejecutar el script
**/BackEnd/compile.sh**

```shell
    cd BackEnd
    mkdir bin
    ./compile.sh
```

Habrá que añadir unos archivos extra al directorio **/Backend/config**. Dos relacionados a la conexión HTTPS,
el certificado TLS, que se denominará "cert.pem" y la clave (key) "key.pem". Un tercer archivo estará relacionado
con la configuración del correo electrónico del sistema, que contendrá la dirección de correo, la contraseña, el
servidor de correo, el puerto utilizado, y la dirección de la página web que se enlazará en los correos.

```json
{
 "from": "UVaDevTest@gmail.com",
 "password": "password",
 "serverhost": "smtp.gmail.com",
 "serverport": "587",
 "frontendurl": "https://localhost:4200"
}
```

Finalmente, para activar el servidor habrá que ejecutar el fichero **/Backend/serve.sh**, el cual se puede
configurar para modificar los puertos donde se ejecutarán el servicio HTTP y HTTPS. Para eliminar el servicio
HTTP y quedarse solo con HTTPS, habrá que modificar la línea correspondiente a HTTP en el fichero **/swagger.yml**,
ejecutar el fichero **/BackEnd/generate.sh**, y modificar el fichero **/BackEnd/serve.sh**, eliminando la opción
"--port:".

### FrontEnd

Para ejecutar el FrontEnd se necesita la aplicación compilada, la cual sirve la misma en cualquier sistema
operativo, por lo que la compilación no es necesario que se haga en la misma maquina en la que se despliega el
sistema. Para compilar se necesita angular-cli (ng) versión 12.0.1.

Hay dos tipos de compilación posible, modo producción o modo inseguro, el cual es la aplicación utilizando HTTP
en vez de HTTPS, pues durante el desarrollo del proyecto no había un certificado SSL real, por lo que se usaba
uno autofirmado, el cual daba problemas en algunos navegadores.

Antes de compilarlo, hay que modificar los archivos de entorno y el index. En los de entorno existe un campo
denominado *API_BASE_PATH* que contiene la dirección de la API que se va a llamar. Se trata de los archivos
en el directorio **/FrontEnd/devtest-client/src/environments/**, concretamente **environment.prod.ts** y
**environment.unsecure.ts**.

Los archivos de ``index'' se encuentran en el directorio **/FrontEnd/devtest-client/src/**, concretamente
**index.prod.html** y **index.unsecure.html**. En ambos hay que modificar la etiqueta relativa al Content-Security-Policy,
cambiando la dirección de despliegue del BackEnd por defecto a la que se ha escrito en los archivos de entorno.

Una vez hecho esto, se ha de compilar, siguiendo la configuración que se necesite, o ambas, con el comando ng build.
Por ejemplo se pone el modo producción:

```shell
    ng build --configuration production
```

Se generarán los archivos necesarios para el despliegue en un directorio dentro de **/FrontEnd/devtest-client/dist/**.

Es importante mencionar que, tal y como se explica en la sección 7.3.1 de la memoria del TFG, el
FrontEnd y el BackEnd tienen que estar desplegados bajo el mismo nombre de dominio.
