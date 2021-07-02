To build on the CLI or IDE use the following command. The project contains a gradle wrapper so a version will get downloaded to execute the build. 
If a recent gradle is installed and available locally then use the gradle command directly rather than the script version.

` ./gradlew quarkusBuild`

With the addtion of the .s2i assemble and run files combined with this ubi-quarkus image `quay.io/quarkus/ubi-quarkus-native-s2i` you can create 
OpenShift build and deployments using S2I 

To create an OpenShift instance use this to build and deploy via Gradle

```
oc new-app quay.io/quarkus/ubi-quarkus-native-s2i:21.0-java11~https://github.com/alexgroom/cnw3.git
    --context-dir=inventory-quarkus-gradle --build-env BUILD_TYPE=jar --name inventory-gradle
```
    
To create a native image remove the BUILD_TYPE flag
