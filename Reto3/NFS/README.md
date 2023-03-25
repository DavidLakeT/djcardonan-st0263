``` 
- ST0263, Reto 3 Aplicación Monolítica con Balanceo
  y Datos Distribuidos (BD y archivos)
- David José Cardona Nieves, djcardonan@eafit.edu.co
- Edwin Nelson Montoya Munera, emontoya@eafit.edu.co
```

# Configuración

A continuación se plantean unos pasos a seguir para realizar la correcta configuración de la instancia NFS y el respectivo acceso por parte de las instancias WordPress.

1. Crear una nueva VM en Google GCP para alojar el servidor de archivos NFS.
2. Conectar la VM del servidor NFS a la misma red y subred que las VM de WordPress y MySQL.
3. Instalar el paquete de servidor NFS en la VM del servidor NFS utilizando los siguiente comandos en la terminal:

``` Shell
sudo apt-get update
sudo apt-get install nfs-kernel-server
```
4. Configurar el servidor NFS editando el archivo de configuración /etc/exports. Este archivo define los directorios que se compartirán a través de NFS y los permisos de acceso. Por ejemplo, si se desea compartir el directorio /var/nfs con permisos de lectura y escritura para las dos instancias de WordPress, se puede agregar la siguiente línea al archivo /etc/exports:

``` Shell
/var/nfs/general  ip_del_host1(rw,sync,no_root_squash,no_subtree_check)
/var/nfs/general  ip_del_host2(rw,sync,no_root_squash,no_subtree_check)
``` 

5. Reiniciar el servicio de servidor NFS para que se apliquen los cambios de configuración mediante el siguiente comando:

``` Shell
sudo systemctl restart nfs-kernel-server
```

6. Por otra parte, puedes editar el archivo /etc/fstab para agregar la siguiente línea que permite montar el directorio cada que se inicia la instancia:

``` Shell
host_ip:/var/nfs/general    /nfs/general   nfs auto,nofail,noatime,nolock,intr,tcp,actimeo=1800 0 0
```

Aún no has terminado, ahora deberás seguir las instrucciones presentes en el archivo README.md disponible en la carpeta WordPress del repositorio.