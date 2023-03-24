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