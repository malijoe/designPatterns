package main

/*

Creational:
    Builder
        Separate component for when object construction gets too complicated
        Can create mutually cooperating sub-builders
        Often has a fluent interface

    Factories
        Factory functions (constructors) are common
        Factor can be a simple function or dedicated struct

    Prototype
        Creation of an object from an existing object
        Requires either explicit deep copy or copy through serialization

    Singleton
        When you need to ensure just a single instance exists
        Can be made thread-safe and lazy
        Consider extracting interface or using dependency injection

Structural:
    Adapter:
        Converts the interface you get to the interface you need
    Bridge:
        Decouple abstraction from implementation
    Composite:
        Allows clients to treat individual objects and compositions of objects uniformly
    Decorator:
        Attach additional responsibilities to objects
        Can be done through embedding pointers
    Facade:
        Provide a single unified interface over a set of interfaces
    Flyweight:
        Efficiently support very large numbers of similar objects
    Proxy:
        Provide a surrogate object that forwards calls to the real object while performing additional functions
        e.g., access control, communication, logging, etc.

Behavioral:
    Chain of Responsibility:
        Allow components to process information/events in a chain
        Each element in the chain refers to next element; or
        Make a list and go through it
    Command:
        Encapsulate a request into a separate object
        Good for audit, replay, undo/redo
        Part of CQS/CQRS
    Interpreter:
        Transform textual input into structures (e.g. ASTs)
        Used by interpreters, compilers, static analysis tools, etc.
        Compiler Theory is a separate branch of Computer Science
    Iterator:
        Provides an interface for accessing elements of an aggregate object
    Mediator:
        Provides mediation services between several objects
    Memento:
        Yields tokens representing system states
        Tokens do not allow direct manipulation, but can be used in appropriate APIs
    Observers:
        Allows notifications of changes/happenings in a component
    State:
        We model systems by having one of possible states and transitions between these states
        Such a system is called a state machine
        Special frameworks exists to orchestrate state machines
    Strategy & Template Method
        Both define a skeleton algorithm with details filled in by implementer
        Strategy uses composition; Template Method doesn't
    Visitor:
        Allows non-intrusive addition of functionality to hierarchies
*/
