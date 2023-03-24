``` 
- ST0263, Reto 3 Aplicación Monolítica con Balanceo
  y Datos Distribuidos (BD y archivos)
- David José Cardona Nieves, djcardonan@eafit.edu.co
- Edwin Nelson Montoya Munera, emontoya@eafit.edu.co
```

# Configuración

1. Instalar nginx en la maquina:

``` Shell
sudo apt-get update
sudo apt install nginx
```

2. Copiar el contenido del archivo nginx.conf presente en esta carpeta del directorio

3. Cambiar los dominios especificados en el archivo

3. Pegar el contenido copiado en la ruta etc/nginx/nginx.conf

4. To create the SSL certificate install certbot:

``` Shell
sudo apt install certbot python3-certbot-nginx
```

5. Revisar el archivo nginx.conf para asegurarse que el server_name sea el mismo dominio para generar SSL

6. Para generar SSL correr el comando

``` Shell
sudo certbot --nginx -d yourdomain.com
```