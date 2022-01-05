package template
/*
    Overview:
        Similar to Strategy pattern
        Algorithms can be decomposed into common parts + specifics
        Strategy pattern does this through composition
        Template Method performs a similar operation, but
            It's typically just a function, not a struct with a reference to the implementation
            Can still use interfaces; or
            Can be functional (take several functions as parameters)

    Template Method: A skeleton algorithm defined in a function. Function can either use an interface (like Strategy pattern) or can take several functions as arguments.

    Summary:
        Very similar to Strategy
        Typical Implementation:
            Define an interface with common operations
            Make use of those operations inside a function
        Alternative functional approach:
            Make a function that takes several functions
            Can pass in functions that capture local state
            No need for either structs or interfaces
 */
