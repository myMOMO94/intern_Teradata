To run, please type the following command on your terminal:


docker build -t tcp_client_server .

docker run -e NUMBYTES="512" -e NUMRUNS="2" tcp_client_server


You can change the numbytes and numruns if you want.