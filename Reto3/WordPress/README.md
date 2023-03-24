``` 
- ST0263, Reto 3 Aplicación Monolítica con Balanceo
  y Datos Distribuidos (BD y archivos)
- David José Cardona Nieves, djcardonan@eafit.edu.co
- Edwin Nelson Montoya Munera, emontoya@eafit.edu.co
```

# Requisitos

Ten en cuenta que para la ejecución de los componentes que se ejecutan en docker debes tener instalado los componentes para la ejecución de este

La instalación de Docker y Git la puedes realizar así:

``` Shell
sudo apt update
sudo apt install docker.io -y
sudo apt install docker-compose -y
sudo apt install git -y
```

Y la configuración para que docker se inicialice cada que se enciende la instancia así:

``` Shell
sudo systemctl enable docker
sudo systemctl start docker
```

Así mismo, el encendido manual lo realizas así:

``` Shell
sudo docker-compose -f docker-compose.yaml up
```

# Configuración

Realiza estos pasos una vez que hayas configurado la instancia NFS siguiendo los pasos disponibles en el archivo README.md de la carpeta NFS del repositorio. 

1. Tendrás que instalar el paquete de cliente NFS en las VM de WordPress mediante la utilización de los siguientes comandos en la terminal:

``` Shell
sudo apt-get update
sudo apt-get install nfs-common
```

2. Luego deberás montar el directorio compartido de NFS en las VM de WordPress mediante el siguiente comando en la terminal:

``` Shell
sudo mount <dirección_IP_del_servidor_NFS>:/var/nfs /mnt/nfs
```

3. Y finalmente solo queda verificar que el directorio compartido de NFS esté montado correctamente en las VM de WordPress mediante el siguiente comando en la terminal:

``` Shell
df -h
```

4. Por otra parte, puedes editar el archivo /etc/fstab para agregar la siguiente línea que permite montar el directorio cada que se inicia la instancia:

``` Shell
host_ip:/var/nfs/general    /nfs/general   nfs auto,nofail,noatime,nolock,intr,tcp,actimeo=1800 0 0
```