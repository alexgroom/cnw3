
Using jKube with OpenShift 4 to build Coolstore components
===
## Introduction
The familiar Coolstore components can all be built via a variety of build operations. 
1. Inventory - Quarkus app
2. Catalog - Spring Boot app
3. Gateway - Vert.x app

In this exmaple we explore how to build them via Maven using the jKube plugin.

## jKube

[jKube Documentation](https://www.eclipse.org/jkube/docs/openshift-maven-plugin#jkube:deploy)

jKube replaces the defunct fabric8 plugin to build and deploy applications into Kubernetes. it has a special flavour just for OpenShift that takes adavnatge of additional build services it offers - mainly S2I

Maven will typically be used to build a jar or in the case of Quarkus/Graalvm/Mandrel a Linux executable. This then needfs to be packaged up into a container image using jKube. it will invoke local Docker if pushing to k8s, but for OpenShift it will push the application and then invoke S2I to build the image on the server. 

For development systems that lack Docker this is a big benefit. Doubly so for Graalvm based applications since again building Linux native binaries on say Windows/Mac also requires a VM.

In the case of OpenShift, jkube drop replaces fabric8 in terms of configuration style.

## POM files
The POM files for the 3 components have been modified to reference the jKube build plugin.

In the most complex example we see the use of injecting labels and environment variables into the deployment eg:-

```
            <plugin>
                <groupId>org.eclipse.jkube</groupId>
                <artifactId>openshift-maven-plugin</artifactId>
                <version>1.0.0</version>
                <configuration>
                <source>11</source>
                <target>11</target>
                <resources>
                    <labels>
                        <all>
                            <property>
                                <name>app</name>
                                <value>gateway</value>
                            </property>
                            <property>
                                <name> app.kubernetes.io/part-of</name>
                                <value>coolstore</value>
                            </property>
                            <property>
                                <name>app.openshift.io/runtime</name>
                                <value>java</value>
                            </property>
                        </all>
                    </labels>
                    <env>
                        <COMPONENT_CATALOG_HOST>catalog</COMPONENT_CATALOG_HOST>
                        <COMPONENT_INVENTORY_HOST>inventory</COMPONENT_INVENTORY_HOST>
                        <COMPONENT_CATALOG_PORT>8080</COMPONENT_CATALOG_PORT>
                        <COMPONENT_INVENTORY_PORT>8080</COMPONENT_INVENTORY_PORT>
                    </env>
                </resources>
                </configuration>
            </plugin>            

```

Note that jKube also supports configuring PVC and Configmaps requests.

## Invoking a build and deploy
Maven drives the build and deploy process over several build goals or stages. A full build and deployment would look like this, having first logged in to OpenShift and selected a project to use.
```
mvn clean package oc:build oc:resource oc:apply
```
Maven can be used just to execute a single goal in response to say a configuration change.

Example build goals

* oc:build - build the image form the archive using S2I
* oc:resource - creates the resource configuration for a DeploymentConfig such as labels, environment and annotations
* oc:apply - applies these changes

## Example Build Script for all 3 Components
After first cloning this git repo, to build all 3 Coolstore components from a command like would look like this:
```
oc login <host> -u <user> -p <password>
oc new-project coolstore
cd inventory-quarkus
mvn clean package oc:build oc:resource oc:apply -DskipTests
cd ../catalog-springboot
mvn clean package oc:build oc:resource oc:apply -DskipTests
cd ../gateway-vertx
mvn clean package oc:build oc:resource oc:apply -DskipTests
```


