FROM registry.redhat.io/ubi8/dotnet-60-runtime AS build
# This image provides a .NET Core 6.0 environment you can use to run your .NET
# applications.

ENV PATH=/opt/app-root/src/.local/bin:/opt/app-root/src/bin:/opt/app-root/node_modules/.bin:/opt/app-root/.dotnet/tools/:${PATH} 

LABEL io.k8s.description="Platform for building and running .NET Core 6.0 applications" \
      io.openshift.tags="builder,.net,dotnet,dotnetcore,dotnet-60"

# Labels consumed by Red Hat build service
LABEL name="ubi8/dotnet-60" \
      com.redhat.component="dotnet-60-container" \
      version="6.0" \
      release="1" \
      architecture="x86_64"

# Labels consumed by Eclipse JBoss OpenShift plugin
LABEL com.redhat.dev-mode="DEV_MODE:false" \
      com.redhat.deployments-dir="/opt/app-root/src"

# Switch to root for package installs
USER 0

RUN INSTALL_PKGS="npm nodejs-nodemon dotnet-sdk-60 rsync procps-ng findutils" && \
    yum install -y --setopt=tsflags=nodocs $INSTALL_PKGS && \
    rpm -V $INSTALL_PKGS && \
    yum clean all -y && \
# yum cache files may still exist (and quite large in size)
    rm -rf /var/cache/yum/*

# Directory with the sources is set as the working directory.
RUN mkdir /opt/app-root/src
WORKDIR /opt/app-root/src

# Trigger first time actions.
RUN dotnet help

# copy csproj and restore as distinct layers
COPY *.csproj .
RUN dotnet restore

# copy and publish app and libraries
COPY . .
RUN dotnet publish -c release -o /opt/app-root --no-restore


# Since $HOME is set to /opt/app-root, the yum install may have created config
# directories (such as ~/.pki/nssdb) there. These will be owned by root and can
# cause actions that work on all of /opt/app-root to fail. So we need to fix
# the permissions on those too.
RUN chown -R 1001:0 /opt/app-root && fix-permissions /opt/app-root

# Needed for the `dotnet watch` to detect changes in a container
ENV  DOTNET_USE_POLLING_FILE_WATCHER=true

# Run container by default as user with id 1001 (default)
USER 1001

# Set the default CMD to print the usage of the language image.
FROM registry.redhat.io/ubi8/dotnet-31-runtime
WORKDIR /opt/app-root
COPY --from=build /opt/app-root .
ENTRYPOINT ["dotnet", "gateway.dll"]

