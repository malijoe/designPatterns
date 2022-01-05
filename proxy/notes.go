package proxy

/*
    Proxy: A type that functions as an interface to a particular resource. That resource may be remote, expensive to construct, or may require logging or some other added functionality.

    Proxy vs. Decorator
        Proxy tries to provide an identical interface
        Decorator provides an enhanced interface
        Decorator typically aggregates (or has a pointer to) what it is decorating; proxy doesnt have to
        Proxy might not even be working with a materialized object

    Summary:
        A proxy has the same interface as the underlying object
        To create a proxy, simply replicate the existing interface of an object
        Add relevant functionality to the redefined methods
        Different proxies (communication, logging, caching, etc.) have completely different behaviors
 */
