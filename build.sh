# Instalación Go
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt-get update
sudo apt-get install golang-go

# Instalación librerías Go
sudo apt-get install protobuf-compiler grpc
go get google.golang.org/protobuf/cmd/protoc-gen-go

# Instalación y configuración RabbitMQ
yes Y | sudo apt-get install rabbitmq-server
sudo systemctl start rabbitmq-server
sudo systemctl enable rabbitmq-server

# Configuración servicios
sudo mv /home/ubuntu/djcardonan-st0263/Systemd/API.service /etc/systemd/system/
sudo mv /home/ubuntu/djcardonan-st0263/Systemd/MOM.service /etc/systemd/system/
sudo mv /home/ubuntu/djcardonan-st0263/Systemd/RPC.service /etc/systemd/system/
sudo systemctl start API
sudo systemctl enable API
sudo systemctl start MOM
sudo systemctl enable MOM
sudo systemctl start RPC
sudo systemctl enable RPC
sudo systemctl daemon-reload