``` 
- ST0263, Reto 3 Aplicación Monolítica con Balanceo
  y Datos Distribuidos (BD y archivos)
- David José Cardona Nieves, djcardonan@eafit.edu.co
- Edwin Nelson Montoya Munera, emontoya@eafit.edu.co
```

# 1. Breve descripción de la actividad

## 1.1. ¿Que aspectos cumplió o desarrolló de la actividad propuesta por el profesor? (requerimientos funcionales y no funcionales)

- Una instancia de GCP con NGINX que equilibra la carga de las peticiones a las instancias de WordPress.
- Dos instancias de WordPress conectadas a la VM de NFS y referidas por la instancia de NGINX.
- Una instancia de GCP que actúe como NFS para compartir información entre las instancias de WordPress.
- Una instancia de GCP con una base de datos MySQL.
Un dominio en el que se puede acceder al sitio web: https://davidlake.tech
- Certificados SSL creados a través de certbot.

# 2. información general de diseño de alto nivel, arquitectura, patrones, mejores prácticas utilizadas.

La arquitectura de este proyecto se basa en la implementación de cinco nodos, cada uno de ellos representando una máquina virtual en GCP.

Cuando el cliente ingresa el dominio en su navegador, el browser busca la IP pública correspondiente a través de Hostinger, la plataforma utilizada para gestionar el DNS. Una vez obtenida la IP, Nginx entra en acción y realiza un balanceo de carga sobre los dos servidores de WordPress, los cuales dependen de un NFS server para el manejo de los archivos y un servidor de base de datos ubicados en la última instancia.

# 3. Descripción del ambiente de desarrollo y técnico: lenguaje de programación, librerias, paquetes, etc, con sus numeros de versiones.

## 3.1) Sistema Operativo

Las cinco instancias están configuradas para trabajar con Ubuntu 22.04 x86.

## 3.2) Dependencias

## como se compila y ejecuta.
## detalles del desarrollo.
## detalles técnicos
## descripción y como se configura los parámetros del proyecto (ej: ip, puertos, conexión a bases de datos, variables de ambiente, parámetros, etc)
## opcional - detalles de la organización del código por carpetas o descripción de algún archivo. (ESTRUCTURA DE DIRECTORIOS Y ARCHIVOS IMPORTANTE DEL PROYECTO, comando 'tree' de linux)
## 
## opcionalmente - si quiere mostrar resultados o pantallazos 

# 4. Descripción del ambiente de EJECUCIÓN (en producción) lenguaje de programación, librerias, paquetes, etc, con sus numeros de versiones.

# IP o nombres de dominio en nube o en la máquina servidor.

## descripción y como se configura los parámetros del proyecto (ej: ip, puertos, conexión a bases de datos, variables de ambiente, parámetros, etc)

## como se lanza el servidor.

## una mini guia de como un usuario utilizaría el software o la aplicación

## opcionalmente - si quiere mostrar resultados o pantallazos 

# 5. otra información que considere relevante para esta actividad.

# referencias:
<debemos siempre reconocer los créditos de partes del código que reutilizaremos, así como referencias a youtube, o referencias bibliográficas utilizadas para desarrollar el proyecto o la actividad>
## sitio1-url 
## sitio2-url
## url de donde tomo info para desarrollar este proyecto

#### versión README.md -> 1.0 (2022-agosto)