``` 
- ST0263, Reto 2 Procesos comunicantes por API REST, RPC y MOM
- David José Cardona Nieves, djcardonan@eafit.edu.co
- Edwin Nelson Montoya Munera, emontoya@eafit.edu.co
```

# 1. Breve descripción de la actividad

El sistema ofrece dos micro-servicios a una API Gateway sirviendo como proxy en frente de estos dos servicios. Los micro-servicios hacen uso de dos diferentes middleware, uno de ellos es gRPC y el otro hace uso de MOM con RabbitMQ. Se ofrecen dos servicios que trabajan sobre un directorio de archivos. Primeramente el poder listar todos los archivos presentes en el directorio, y el segundo es el poder verificar si un archivo en específico existe en esa ruta.

## 1.1. Requerimientos

- El sistema se compone de tres servicios: Un microservicio implementado con gRPC, un microservicio implementado con interfaz MoM y un servicio pasarela que sirve de proxy frente entre los dos microservicios previos.
- El primer microservicio se implementó utilizando gRPC (Protocol Buffers).
- El segundo microservicio se implementó utilizando RabbitMQ.
- Ambos microservicios ofrecen las mismas funcionalidades (listar y buscar archivos).
- La API Gateway implementa un REST API bajo la utilización de JSON.
- Hay archivo de configuración para aspectos como IP, puertos a utilizar, etc.
- Hay soporte para manejar concurrencia y balanceo entre ambos microservicios.

Es decir, todos los requerimientos fueron cumplidos.

# 2. Arquitectura.

El sistema hace uso de dos modos de comunicación: gRPC y MOM.

En el modo gRPC, la API Gateway recibe la solicitud directamente y genera un formato binario específico creado por Google (proto-buffer) para enviar al servidor GRPC. El microservicio desserializa el contenido del mensaje y procede a realizar el procesamiento con la información obtenida, para luego enviar la respuesta de vuelta al API Gateway, que a su vez la envía al cliente con el formato JSON. 

Por otra parte, en el modo MOM, la solicitud se envía a una cola de peticiones de RabbitMQ con un ID único generado. El servidor MOM monitorea constantemente la cola en busca de nuevas peticiones. Una vez que procesa la solicitud, el servidor MOM serializa la respuesta y la envía de vuelta a la cola de respuestas. El cliente MOM compara el ID de la respuesta con el ID de la petición original para confirmar que corresponde a la respuesta correcta. Estos ID son generados mediante la utilización de funciones _random_.

Es decir, el sistema podría utilizar diferentes modos de comunicación según las necesidades (aunque en este caso lo hace únicamente basándose en el balance de las peticiones), y cada uno tiene su propio proceso para enviar y recibir solicitudes y respuestas de los microservicios.

# 3. Ambiente de desarrollo
# 3.1 Dir Tree
![alt text](https://github.com/DavidLakeT/djcardonan-st0263/tree/main/Assets/tree.png)
# 3.2 Lenguaje de programacion
EL lenguaje utilizado para el desarrollo de todo el sistema fue Go (Golang) en su versión 1.20.1.

Las librerias usadas son:
```
- protobuf v2.4.0
- yaml v1.28.1
- grpc v1.53.0
- reflect v1.0.2
- amqp091-go v1.7.0
- gonfig v0.0.02021
```
# 3.3 Configuracion de servidores

## API Gateway
`api_gateway.py`
```
{
    "ip":"127.0.0.1",
    "port":"50051",
    "dir":"../Files"
}
```
## Servidor GRPC
`grpc_server.py`
```
{
    "ip":"127.0.0.1",
    "port":"50051",     
    "dir":"../../Files/"
}
```
## Servidor MOM
`rpc_server.py`

```
{
    "ip":"127.0.0.1",
    "port":"5672",
    "dir":"../../Files/"
}
```

# 4. Ejecucion

Como requisitos para la ejecución de los ejecutables generados con ```go build``` se encuentran:

- Tener Golang instalado.
- Tener RabbitMQ instalado.
- Verificar que los puertos de la máquina estén abiertos (8080).

La carpeta Systemd cuenta con los archivos para la automatización de la inicialización de los microservicios en las instancias.

Este proceso se puede realizar mediante la ejecución del archivo ```build.sh```.

# 5 Consulta a la API Gateway

Una vez se encuentran encendidos los tres componentes del sistema. Es posible acceder a estos utilizando la IP pública de la instancia donde se ejecuten (o localhost, si se ejecutan localmente) de la siguiente forma:

``` Go
# Para listar todos los archivos
localhost:8080/list
```

``` Go
# Para realizar una búsqueda
localhost:8080/search/<filename>
```

**Nota:** Los nombres de los archivos (_filename_) deben incluir la extensión dado que podrían existir múltiples archivos con un nombre.


# 5. referencias:
- https://github.com/grpc/grpc-go
- https://grpc.io/docs/languages/go/basics/
- https://www.rabbitmq.com/tutorials/tutorial-six-go.html
- https://github.com/ruybrito106/mom-client-go
- https://levelup.gitconnected.com/writing-an-rpc-server-in-go-eb9afd56d1e1
