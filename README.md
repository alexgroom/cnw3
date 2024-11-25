# Serverless and Eventing demo  

This is a demo app for Cloud Native applications using 
Red Hat OpenShift Application Runtimes (Spring Boot, Quarkus, Eclipse Vert.x and Node.js) 
utilizing a microservices architecture.


## CoolStore Online Store App

CoolStore is an online store web application built using Spring Boot, WildFly Swarm, Eclipse Vert.x, 
Node.js and AngularJS adopting the microservices architecture.

* **Web**: A Node.js/Angular front-end
* **API Gateway**: aggregates API calls to back-end services and provides a condenses REST API for front-end
* **Catalog**: a REST API for the product catalog and product information
* **Inventory**: a REST API for product's inventory status

```
                    +-------------+
                    |             |
                    |     Web     |
                    |             |
                    |   Node.js   |
                    |  AngularJS  |
                    +------+------+
                          |
                          v
                    +------+------+
                    |             |
                    | API Gateway |
                    |             |
                    |   .NET.     |
                    |             |
                    +------+------+
                          |
                +---------+---------+
                v                   v
          +------+------+     +------+------+
          |             |     |             |
          |   Catalog   |     |  Inventory  |
          |             |     |             |
          | Spring Boot |     |Quarkus App. |
          |             |     |             |
          +-------------+     +-------------+
```
Associated document is being prepared here (Red Hat internal only atm) https://docs.google.com/document/d/15PY3PdaEXavGPhF_6FIVOrQ2ccOrDtC06f6j0Wo1CyE/edit?usp=sharing
