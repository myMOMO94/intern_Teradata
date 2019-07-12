To run, please type the following command on your terminal:

docker build -t http_client_server .

docker run -e NUMBYTES="512" -e NUMRUNS="4" http_client_server -verbose=true

You can change the NUMBYTES, NUMRUNS and verbose flag if you want.
