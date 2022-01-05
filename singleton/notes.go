package singleton
/*
    Overview:
        For some components it only makes sense to have one in the system
            Database repo
            Object factory
        The construction call is expensive
            only do it once
            we give everyone the same instance
        Want to prevent anyone creating additional copies
        Need to take care of lazy instantiation
    Singleton: A component that is instantiated only once
 */

/*
    Summary:
        Lazy one-time initialization using sync.Once (thread-safe)
        Adhere to DIP: depend on interfaces, not concrete types
 */
