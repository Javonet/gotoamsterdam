FROM mcr.microsoft.com/dotnet/sdk:8.0
WORKDIR /usr/local/app

COPY Binaries ./


COPY jcg.exe ./
COPY bin/Debug/net8.0/RobotConnector.dll ./

CMD jcg RobotConnector.dll
EXPOSE 8080:8080