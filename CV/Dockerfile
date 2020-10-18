FROM python:3.7.5

ADD grpc_api $HOME/grpc_api

RUN sudo apt-get install -y libsm6 libxext6 libxrender-dev
RUN pip install -r $HOME/grpc_api/requirements.txt

CMD [ "python", ".$HOME/grpc_api/ml_server/server.py"]
