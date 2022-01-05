package bridge

/*
    Overview: Bridge prevents a 'Cartesian product' complexity explosion
    Example:
        - Common type ThreadScheduler
        - Can be preemptive or cooperative
        - Can run on Windows or Unix
        - End up with a 2x2 scenario: WindowsPTS, UnixPTS, WindowsCTS, UnixCTS
    Bridge pattern avoids the entity explosion

    Bridge: A mechanism that decouples an interface (hierarchy) from an implementation (hierarchy)
 */

/*
    Summary:
        Decouple the abstraction from the implementation (Creates a parallel hierarchy)
        Both can exist as hierarchies
        A stronger form of encapsulation
 */