# Go-DesignPattern

1. Builder.
    Why we need to use builder ?
        - It might be necessary to have business rules to validate some parameters or
        derive some added attributes.
        - We might need some code to bring in efficiency.
        - It might be necessary to have idempotency and thread safety in object creation.
        That is, multiple requests for object creation with the same parameters should
        give the same object.
        - The objects might have multiple constructor arguments (typically called
        telescopic constructors), and it is difficult to remember the order of parameters
        for the clients. Some of these parameters might be optional. Such constructors
        frequently lead to bugs in client code.
        
    1.1. Builder Facets.
    1.2. Builder Parameter.
    1.3. Functional Builder.
2. Factories.
    2.1. Factory Function.
    2.2. Interface Factory.
    2.3. Prototye Factory.
3. Prototype.
   
4. Singleton.